## astral

Crude go library to use astrald via its appsupport.

### Quick start

### Server
```go
port := astral.Listen("myapp")

for request := range port.Next() {
    if request.Caller() == allowedID {
        conn, err := request.Accept()
        // handle the connection
        conn.Close()
    } else {
    	request.Reject()
    }
}
```

### Client
```go
conn, err := astral.Dial(nodeID, "myapp")
// transfer data
conn.Close()
```