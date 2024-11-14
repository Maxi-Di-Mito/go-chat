# Chat Room

## main loop: listener

Listens for connections and start a new client with it's own connection object and ID in the ronin clients list starts goroutine for the connection,
sends the available channel list

## Joining room

### New room case

Starts a new room with the client initiating, starts the channel and goroutine for the rooms messages

### Existing room

Adds the client to the room's client list, and adds the room's channel to the client object
