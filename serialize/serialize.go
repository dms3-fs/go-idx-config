package idxrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dms3-fs/go-idx-config"

	"github.com/facebookgo/atomicfile"
	"github.com/dms3-fs/go-fs-util"
)

// ReadConfigFile reads the idxconfig from `filename` into `cfg`.
func ReadConfigFile(filename string, cfg interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("failure to decode idxconfig: %s", err)
	}
	return nil
}

// WriteConfigFile writes the idxconfig from `cfg` into `filename`.
func WriteConfigFile(filename string, cfg interface{}) error {
	err := os.MkdirAll(filepath.Dir(filename), 0775)
	if err != nil {
		return err
	}

	f, err := atomicfile.New(filename, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	return encode(f, cfg)
}

// encode configuration with JSON
func encode(w io.Writer, value interface{}) error {
	// need to prettyprint, hence MarshalIndent, instead of Encoder
	buf, err := idxconfig.Marshal(value)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}

// Load reads given file and returns the read idxconfig, or error.
func Load(filename string) (*idxconfig.IdxConfig, error) {
	// if nothing is there, fail. User must run 'dms3fs init'
	if !util.FileExists(filename) {
		return nil, errors.New("dms3fs not initialized, please run 'dms3fs init'")
	}

	var cfg idxconfig.IdxConfig
	err := ReadConfigFile(filename, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
