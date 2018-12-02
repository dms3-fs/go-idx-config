package idxrepo

import (
	"os"
	"runtime"
	"testing"

	idxconfig "github.com/dms3-fs/go-idx-config"
)

func TestConfig(t *testing.T) {
	const filename = ".dms3config"
	cfgWritten := new(idxconfig.IdxConfig)
	cfgWritten.Identity.PeerID = "faketest"

	err := WriteConfigFile(filename, cfgWritten)
	if err != nil {
		t.Fatal(err)
	}
	cfgRead, err := Load(filename)
	if err != nil {
		t.Fatal(err)
	}
	if cfgWritten.Identity.PeerID != cfgRead.Identity.PeerID {
		t.Fatal()
	}
	st, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("cannot stat idxconfig file: %v", err)
	}

	if runtime.GOOS != "windows" { // see https://golang.org/src/os/types_windows.go
		if g := st.Mode().Perm(); g&0117 != 0 {
			t.Fatalf("idxconfig file should not be executable or accessible to world: %v", g)
		}
	}
}
