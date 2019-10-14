package d1exe

import (
	"bytes"

	"github.com/decomp/exp/bin"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/level/miniset"
)

// ref: v1.09b.
const (
	// l1 minisets.
	L1StairUp1Addr        = 0x479EC8 // STAIRSUP
	L1StairUp2Addr        = 0x479EEC // L5STAIRSUP
	L1StairDownAddr       = 0x479F10 // STAIRSDOWN
	L1CandlestickAddr     = 0x479F2C // LAMPS
	L1StairDownPoisonAddr = 0x479F38 // PWATERIN
)

// parseMinisets parses the miniset DUN files contained within the executable.
func (exe *Executable) parseMinisets() error {
	l1StairUp1, err := exe.parseMiniset(L1StairUp1Addr)
	if err != nil {
		return errors.WithStack(err)
	}
	exe.Minisets[L1StairUp1Addr] = l1StairUp1
	l1StairUp2, err := exe.parseMiniset(L1StairUp2Addr)
	if err != nil {
		return errors.WithStack(err)
	}
	exe.Minisets[L1StairUp2Addr] = l1StairUp2
	l1StairDown, err := exe.parseMiniset(L1StairDownAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	exe.Minisets[L1StairDownAddr] = l1StairDown
	l1Candlestick, err := exe.parseMiniset(L1CandlestickAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	exe.Minisets[L1CandlestickAddr] = l1Candlestick
	l1StairDownPoison, err := exe.parseMiniset(L1StairDownPoisonAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	exe.Minisets[L1StairDownPoisonAddr] = l1StairDownPoison
	return nil
}

// parseMinisets parses the miniset DUN file at the given address contained
// within the executable.
func (exe *Executable) parseMiniset(minisetAddr uint32) (*miniset.Miniset, error) {
	dbg.Printf("parsing miniset at 0x%08X", minisetAddr)
	data := exe.file.Data(bin.Address(minisetAddr))
	r := bytes.NewReader(data)
	m, err := miniset.Parse(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return m, nil
}
