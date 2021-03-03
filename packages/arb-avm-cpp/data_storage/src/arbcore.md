# ArbCore

## Thread Categories

Primary thread categories to be aware of:

- ArbCore
  - Primary thread responsible for database interaction and handling reorgs. The thread and all related data is encapsulated in the ArbCore object. Only one ArbCore thread exists in a given application.
- MachineThread
  - Is responsible for executing machine instance. The thread and all related data is encapsulated in the MachineThread object. Multiple independent MachineThreads may be active at the same time.
  - There is a single MachineThread used by ArbCore to provide the canonical state of L2 at any given time.
  - Additional MachineThread instances may be instantiated by other threads for various purposes.
- RocksDB Database
  - The RocksDB database is threadsafe, but race conditions can occur when querying multiple data types.
  - Race conditions currently managed with mutexes, will eventually move to using RocksDB snapshots.
- Other
  - All other threads that interact with ArbCore and/or MachineThread.

## Thread Interactions

### ArbCore Thread

The `ArbCore::core_thread` processes message entries, performs reorgs on database entries, handles `LogsCursor` queries and runs MachineThread in sequential steps, that way reorgs can easily modify all the components without worrying about race conditions.

#### ArbCore Thread Receiving Messages

- The message delivery itself does not require a mutex, but if a reorg occurs, `ArbCore::core_reorg_mutex` needs to be acquired.
- The function `ArbCore::DeliverMessages()` is used by other threads to insert messages into the core thread. The state is managed between threads using the atomic enum `ArbCore::message_data_status`.

  - `ArbCore::MESSAGES_EMPTY`: (Out) Ready to receive messages
  - `ArbCore::MESSAGES_READY`: (In) Messages in vector
  - `ArbCore::MESSAGES_SUCCESS`: (Out) Messages processed successfully
  - `ArbCore::MESSAGES_NEED_OLDER`: (Out) Last message invalid, need older messages
  - `ArbCore::MESSAGES_ERROR`: (Out) Error processing messages

- The core thread only accesses or updates the following variables when `message_data_status` is set to `MESSAGES_READY`. Correspondingly, other threads should only deliver messages when `message_data_status` is set to `MESSAGES_EMPTY`. The states `MESSAGES_ERROR` and `MESSAGES_NEED_OLDER` should be cleared by calling `ArbCore::messagesClearError()`.

  - `ArbCore::message_data_status`
  - `ArbCore::message_data.messages`
  - `ArbCore::message_data.previous_inbox_acc`
  - `ArbCore::message_data.last_block_complete`

- `ArbCore::deliverMessages` must only be called when `messagesEmpty` returns true
- `ArbCore::messagesEmpty` can be called at any time
- `ArbCore::messagesStatus` can be called at any time
- `ArbCore::messagesClearError` must only be called when `messagesStatus` returns `MESSAGES_NEED_OLDER` or `MESSAGES_ERROR`

#### ArbCore Thread and MachineThread

- Machine is started and results collected synchronously by core thread, so no mutex is required.
- See `MachineThread` section below for specifics of MachineThread interaction.
- When reorg occurs and Machine has already consumed obsolete messages, MachineThread will be appropriately destroyed and recreated from last valid checkpoint.
- If an error occurs while executing machine, the core thread is aborted.

#### ArbCore Thread and RocksDB

The core thread can modify multiple tables during reorg. This causes problems with other threads that query data from the same tables, so the mutex `ArbCore::core_reorg_mutex` is acquired whenever a reorg is done so that other threads can use the mutex to avoid getting invalid data from race conditions.

The following data is stored by core thread and appropriately modified whenever a reorg occurs:

- Checkpoints
- Messages
- Logs
- Sends
- LogsCursor (outside of DB)
- MachineThread (outside of DB)

#### ArbCore Thread and LogsCursor

- All database interaction dealing with `LogsCursor` includiung reorgs is handled synchronously by the core thread, so no mutex is needed for database interaction
- The function `logsCursorGetLogs` needs reorg mutex to ensure underlying data is not affected by reorg while retrieving logs.
- The state is communicated between threads using the atomic enum `ArbCore::logs_cursor.status`

  - `DataCursor::EMPTY`: (Out) Ready to receive request for data
  - `DataCursor::REQUESTED`: (In) Data requested
  - `DataCursor::CONFIRMED`: (In) Data count to confirm as received
  - `DataCursor::ERROR`: (Out) Error getting data

- `ArbCore::logsCursorRequest`: Asynchronously request the next X logs
  - Only call after successful call to `logsCursorGetLogs` or `logsCursorClearError`
- `ArbCore::logsCursorGetLogs`: Asynchronously get the logs requested using `logsCursorRequest`
  - Only call after successful call to `logsCursorRequest`
  - If `nullOpt` is returned, call `logsCursorGetDeletedLogs`
- `ArbCore::logsCursorGetDeletedLogs`: Asynchronously get logs deleted by reorg
  - Only call after `logsCursorGetLogs` returns `nullopt`
  - If `nullOpt` is returned, call `logsCursorCheckError`
- `ArbCore::logsCursorCheckError`: Checks if error has occurred
  - Call `logsCursorClearError` if returns true
- `ArbCore::logsCursorClearError`: returns error string and clears error state
  - Only call if `LogsCursorCheckError` returns true
- `ArbCore::logsCursorSetConfirmedCount`: Sets the count of logs that are confirmed as processed
  - Should be called after logs are received and before `logsCursorRequest` is called for next batch of logs

#### ArbCore and ExecutionCursor

- The core thread is not involved with creating or advancing execution cursors. The reorg mutex does need to be required to avoid a reorg causing inconsistent results.
- When an `ExecutionCursor` is initially created or an existing one is advanced, any existing data is checked to see if affected by reorg.

### MachineThread

- The state is communicated between threads using atomic enum `MachineThread::machine_status`.

  - `MachineThread::MACHINE_NONE`: Machine is not currently running and no results are stored
  - `MachineThread::MACHINE_RUNNING`: Thread is currently running machine
  - `MachineThread::MACHINE_ABORTED`: Thread stopped after being aborted, machine will need to be recreated
  - `MachineThread::MACHINE_SUCCESS`: Thread finished successfully, results of run are waiting
  - `MachineThread::MACHINE_ERROR`: Error occurred, machine will need to be recreated

- `MachineThread::status()` can be called at any time.
- `MachineThread::runMachine()` should only be called when `MachineThread::status()` returns `MachineThread::MACHINE_NONE`.
- `MachineThread::abortMachine()` can be called at any time.
- `MachineThread::getErrorString()` can be called at any time, but usually only called after `MachineThread::status()` returns `MachineThread::MACHINE_ERROR`.
- `MachineThread::clearError()` can be called at any time, but usually only called after calling `MachineThread::getErrorString()`.
- `MachineThread::getAssertion()` should only be called when `MachineThread::status()` returns `MachineThread::MACHINE_SUCCESS`.
