## Logs Cursor

### Implemented with class `DataCursor`

#### `status` - Atomic, no mutex needed

- `EMPTY`
- `REQUESTED`
- `READY`
- `DELIVERED`
- `ERROR`

#### `number_requested` - No mutex required

#### `pending_total_count` - Must acquire reorg_mutex

#### `current_total_key` - Must acquire reorg_mutex

#### `data` - Must acquire reorg_mutex

#### `deleted_data` - Must acquire reorg_mutex

#### `error_string` - Must acquire reorg_mutex

### Different states of Logs Cursor

#### Status starts out as `EMPTY`

#### Status `EMPTY`

- Go code calls `logsCursorRequest`
- Only called when `status` is `EMPTY`
- Set `number_requested` to a non-zero value
- Set `status` to `REQUESTED`

#### Status `REQUESTED`

- Arbcore thread calls `handleLogsCursorRequested`
- Only called when `status` is `REQUESTED`
- Acquires mutex
- If no new messages but deleted_messages is not empty, set `status` to `READY` and return
- If no new messages, don't do anything
- Reads requested logs from database and puts it into cursor
- Set `status` to `ERROR` if error occurred
- Set `status` to `READY` if no errors

#### Status `READY`

- Go code calls `logsCursorGetLogs`
- Acquires mutex
- logsCursorGetLogs returns `TryAgain` if `status` not `REQUESTED`
- Reads current_count from database
- Set cursor pending_total_count to current_count + number of new logs
- Set `status` to `ERROR` if error occurred
- Set `status` to `DELIVERED` if no errors
- Return list of logs and list of deleted logs, remove lists from cursor

#### Status `DELIVERED`

- Go code calls `logsCursorConfirmReceived`
- Acquires mutex
- If deleted logs were added since GetLogs called, set `status` to `READY` and return false
- Writes database current_count to cursor pending_total_count
- Set `status` to `ERROR` if error occurred
- Set `status` to `EMPTY` if no errors

### Handling Reorgs

- Arbcore thread calls `handleLogsCursorReorg` when appropriate regardless of what `status` is set to
- Acquires mutex
- Reads current_count from database
- If current_count > cursor pending_total_count, set cursor pending_total_count to current_count
- if new log_count < cursor pending_total_count:
  - Save any logs that will be deleted into cursor
  - Set cursor pending_total_count to new log_count
  - Only update database current_count if > new log_count (if less, probably cursor probably contains pending logs)
- Remove outdated pending logs in cursor if needed
- If `status` is `READY` but no pending logs or pending deleted logs left, change `status` to `REQUESTED`
