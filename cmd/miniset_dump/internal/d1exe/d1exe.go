// Package d1exe provides access to the Diablo 1 game executable.
package d1exe

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/pe" // register PE decoder
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/level/miniset"
)

var (
	// dbg is a logger with the "d1exe:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("d1exe:")+" ", 0)
	// warn is a logger with the "d1exe:" prefix which logs warning messages to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("d1exe:")+" ", 0)
)

// Executable provides access to information stored within the diablo.exe
// executable.
type Executable struct {
	// Miniset DUN files.
	Minisets map[uint32]*miniset.Miniset

	// Underlying file of the executable.
	file *bin.File
}

// ParseFile parses the given diablo.exe executable.
func ParseFile(exePath string) (*Executable, error) {
	dbg.Printf("parsing %q.", exePath)
	// Read file contents.
	buf, err := ioutil.ReadFile(exePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Sanity check.
	sum := fmt.Sprintf("%040x", sha1.Sum(buf))
	switch sum {
	case "accedfe32775d4a1984451309608c2a2d39ad406":
		// v1.00dbg, not supported.
		return nil, errors.Errorf("support for parsing diablo.exe v1.00dbg not yet implemented (expected version 1.09 or 1.09b)")
	case "9633d9bbf9a2dc3e88571cead6ac09d689c2abcf":
		// v1.00, not supported.
		return nil, errors.Errorf("support for parsing diablo.exe v1.00 not yet implemented (expected version 1.09 or 1.09b)")
	case "958e3503839a0798d544868d5fc3dff05c02e9fa":
		// v1.04, not supported.
		return nil, errors.Errorf("support for parsing diablo.exe v1.04 not yet implemented (expected version 1.09 or 1.09b)")
	case "d8af136407ea1c019c881f31a180a0bceda39226":
		// v1.07, not supported.
		return nil, errors.Errorf("support for parsing diablo.exe v1.07 not yet implemented (expected version 1.09 or 1.09b)")
	case "c2c111a1825c0410eec93f2f0dc872dd49f8c0db":
		// v1.08, not supported.
		return nil, errors.Errorf("support for parsing diablo.exe v1.08 not yet implemented (expected version 1.09 or 1.09b)")
	case "2119e1c8b818c27a06948979560cdeb4bec9ae65":
		// v1.09, supported.
		dbg.Printf("supported version of diablo.exe (v1.09)")
	case "ebaee2acb462a0ae9c895a0e33079c94796cb0b6":
		// v1.09b, supported.
		dbg.Printf("supported version of diablo.exe (v1.09b)")
	case "e59538ac8de87063e5d3e921a0c5d629e8d54c4e":
		// v1.09b (no CD), supported.
		dbg.Printf("supported version of diablo.exe (v1.09b (no CD))")
	case "a024b35350f94959e6c5e2b7e042da1289d9ee79":
		// v1.09b (AJenbo), supported.
		dbg.Printf("supported version of diablo.exe (v1.09b (AJenbo))")
	default:
		return nil, errors.Errorf("support for parsing diablo.exe with SHA1 hashsum `%s` not yet implemented", sum)
	}
	// Parse executable.
	file, err := bin.ParseFile(exePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	exe := &Executable{
		Minisets: make(map[uint32]*miniset.Miniset),
		file:     file,
	}
	// Parse miniset DUN files.
	if err := exe.parseMinisets(); err != nil {
		return nil, errors.WithStack(err)
	}
	return exe, nil
}
