package astral

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const defaultAstralDir = "/var/run/astrald"

var defaultInstance = NewAstral(astralDir())

func instance() *Astral {
	return defaultInstance
}

func Listen(port string) (*Port, error) {
	return instance().Listen(port)
}

func Dial(identity string, port string) (io.ReadWriteCloser, error) {
	return instance().Dial(identity, port)
}

func astralDir() string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return defaultAstralDir
	}

	dir := filepath.Join(cfgDir, "astrald")
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		fmt.Println("astrald dir error:", err)
		return defaultAstralDir
	}

	return dir
}
