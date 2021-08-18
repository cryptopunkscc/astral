package astral

import (
	"github.com/cryptopunkscc/astrald/services/appsupport/proto"
	"io"
)

type Request struct {
	caller string
	query  string
	socket *proto.Socket
}

func (request Request) Caller() string {
	return request.caller
}

func (request Request) Query() string {
	return request.query
}

func (request *Request) Accept() (io.ReadWriteCloser, error) {
	return request.socket, request.socket.OK()
}

func (request *Request) Reject() {
	request.socket.Error("rejected")
}
