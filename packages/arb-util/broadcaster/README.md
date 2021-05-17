# broadcaster
This package defines a web socket server that allows for clients to receive a data feed of inboxMessages sequence numbers prior to availability on the main chain. The corresponding broacastclient package contains a package that connects to this server and receives messages.

## Start a server
```
	broadcasterSettings := BroadcasterSettings{
		Addr:      ":9428", // the port to listen on
		Workers:   128,
		Queue:     1,
		IoTimeout: 2 * time.Second,
	}
	broadcaster := NewBroadcaster(broadcasterSettings)

	broadcaster.startBroadcaster()
```
## send a message server
```
	broadcaster.Broadcast(messages)
```

See also:

`arb-util/broadcastclient`

`arb-util/inbox/inboxMessage.go`