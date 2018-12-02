// package idxconfig implements the dms3fs idxconfig file datastructures and utilities.
package idxconfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// IdxConfig is used to load dms3fs index parameters idxconfig files.
type IdxConfig struct {
        Indexer   Indexer
        Metadata  Metadata
        Retriever Retriever
		RepoSet	  RepoSet
}

const (
	// DefaultPathName is the default idxconfig dir name
	DefaultPathName = ".dms3-fs/index"
	// DefaultPathRoot is the path to the default idxconfig dir location.
	DefaultPathRoot = "~/" + DefaultPathName
	// DefaultIdxConfigFile is the filename of the configuration file
	DefaultIdxConfigFile = "idxconfig"
	// EnvDir is the environment variable used to change the path root.
	EnvDir = "DMS3_PATH"	// must be relative
)

// PathRoot returns the default configuration root directory
func PathRoot() (string, error) {
	dir := os.Getenv(EnvDir)
	var err error
	if len(dir) == 0 {
		dir, err = homedir.Expand(DefaultPathRoot)
	}
	return dir, err
}

// Path returns the path `extension` relative to the configuration root. If an
// empty string is provided for `configroot`, the default root is used.
func Path(configroot, extension string) (string, error) {
	if len(configroot) == 0 {
		dir, err := PathRoot()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, extension), nil

	}
	return filepath.Join(configroot, extension), nil
}

// Filename returns the configuration file path given a configuration root
// directory. If the configuration root directory is empty, use the default one
func Filename(configroot string) (string, error) {
	return Path(configroot, DefaultIdxConfigFile)
}

// HumanOutput gets a idxconfig value ready for printing
func HumanOutput(value interface{}) ([]byte, error) {
	s, ok := value.(string)
	if ok {
		return []byte(strings.Trim(s, "\n")), nil
	}
	return Marshal(value)
}

// Marshal configuration with JSON
func Marshal(value interface{}) ([]byte, error) {
	// need to prettyprint, hence MarshalIndent, instead of Encoder
	return json.MarshalIndent(value, "", "  ")
}

func FromMap(v map[string]interface{}) (*IdxConfig, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	var conf IdxConfig
	if err := json.NewDecoder(buf).Decode(&conf); err != nil {
		return nil, fmt.Errorf("failure to decode idxconfig: %s", err)
	}
	return &conf, nil
}

func ToMap(conf *IdxConfig) (map[string]interface{}, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(conf); err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.NewDecoder(buf).Decode(&m); err != nil {
		return nil, fmt.Errorf("failure to decode idxconfig: %s", err)
	}
	return m, nil
}
