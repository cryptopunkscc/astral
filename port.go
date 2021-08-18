package astral

import (
	"github.com/cryptopunkscc/astrald/services/appsupport/proto"
	"io"
	"net"
)

type Port struct {
	path     string
	closer   io.Closer
	requests chan *Request
}

func NewPort(path string, closer io.Closer) (*Port, error) {
	var err error
	var port = &Port{
		path:     path,
		closer:   closer,
		requests: make(chan *Request),
	}

	l, err := net.Listen("unix", port.path)
	if err != nil {
		return nil, err
	}

	go func() {
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}

			socket := proto.NewSocket(conn)
			r, err := socket.ReadRequest()
			if err != nil {
				socket.Close()
				continue
			}

			port.requests <- &Request{
				caller: r.Identity,
				query:  r.Port,
				socket: socket,
			}
		}
	}()

	return port, nil
}

func (port *Port) Next() <-chan *Request {
	return port.requests
}

func (port *Port) Close() error {
	return port.closer.Close()
}
