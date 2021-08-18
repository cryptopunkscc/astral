package astral

import (
	"errors"
	"fmt"
	"github.com/cryptopunkscc/astrald/services/appsupport/proto"
	"github.com/google/uuid"
	"io"
	"net"
	"os"
	"path/filepath"
)

const ctlSocket = "ctl.sock"

type Astral struct {
	rootDir string
}

func NewAstral(rootDir string) *Astral {
	return &Astral{rootDir: rootDir}
}

func (astral *Astral) Listen(port string) (*Port, error) {
	// connect to the daemon
	conn, err := astral.daemon()
	if err != nil {
		return nil, err
	}

	// set up unix socket
	path := filepath.Join(os.TempDir(), ".astral-"+uuid.NewString())

	// register
	res, err := conn.Register(port, path)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("register error: %v", err)
	}
	if res.Status != proto.StatusOK {
		conn.Close()
		return nil, fmt.Errorf("register error: %v", errors.New(res.Error))
	}

	return NewPort(path, conn)
}

func (astral *Astral) Dial(identity string, port string) (io.ReadWriteCloser, error) {
	conn, err := astral.daemon()
	if err != nil {
		return nil, err
	}

	// Send the request
	res, err := conn.Connect(identity, port)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("connect error: %v", err)
	}
	if res.Status != proto.StatusOK {
		conn.Close()
		return nil, fmt.Errorf("connect error: %v", errors.New(res.Error))
	}

	return conn, nil
}

func (astral *Astral) daemon() (*proto.Socket, error) {
	c, err := net.Dial("unix", astral.socketPath(ctlSocket))
	if err != nil {
		return nil, err
	}

	return proto.NewSocket(c), nil
}

func (astral *Astral) socketPath(name string) string {
	return filepath.Join(astral.rootDir, name)
}
