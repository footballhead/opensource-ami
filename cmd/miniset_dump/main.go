// The miniset_dump tool extracts miniset files from the Diablo 1 game (*.exe ->
// *.tmx).
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Noofbiz/tmx"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/d1/enum"
	"github.com/sanctuary/formats/level/miniset"
	"github.com/sanctuary/formats/level/til"
	"github.com/sanctuary/opensource-ami/cmd/miniset_dump/internal/d1exe"
)

var (
	// dbg is a logger with the "miniset_dump:" prefix which logs debug messages
	// to standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("miniset_dump:")+" ", 0)
	// warn is a logger with the "miniset_dump:" prefix which logs warning
	// messages to standard error.
	warn = log.New(os.Stderr, term.RedBold("miniset_dump:")+" ", 0)
)

// usage prints tool usage information.
func usage() {
	const use = `
Extract miniset files from the Diablo 1 game (*.exe -> *.tmx).

Usage:

	miniset_dump [OPTION]... diablo.exe

Flags:
`
	fmt.Fprint(os.Stderr, use[1:])
	flag.PrintDefaults()
}

const (
	// Minisets subdirectory (i.e. "_dump_/_minisets_"); containing output TMX
	// files.
	minisetsSubdir = "_minisets_"
	// Tilesets subdirectory (i.e. "_dump_/_tilesets_") containing output tileset
	// PNG spritesheets.
	tilesetsSubdir = "_tilesets_"
)

func main() {
	// Parse command line arguments.
	var (
		// mpqDir specifies the path to an extracted "diabdat.mpq".
		mpqDir string
		// Output directory.
		outputDir string
	)
	flag.Usage = usage
	flag.StringVar(&outputDir, "o", "_dump_", "output directory")
	flag.StringVar(&mpqDir, "mpqdir", "diabdat", `path to extracted "diabdat.mpq"`)
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	exePath := flag.Arg(0)
	// Dump information from diablo.exe.
	if err := dump(mpqDir, exePath, outputDir); err != nil {
		warn.Printf("unable to extract miniset TMX files from %q; %+v", exePath, err)
		fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
}

// dump retrieves information stored within the diablo.exe executable.
func dump(mpqDir, exePath, outputDir string) error {
	exe, err := d1exe.ParseFile(exePath)
	if err != nil {
		return errors.WithStack(err)
	}
	// Miniset files.
	minisetInfos := []struct {
		addr    uint32
		tmxName string
		dtype   enum.DType
	}{
		// l1
		{addr: d1exe.L1StairsUp1Addr, tmxName: "stairs_up1.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairsUp2Addr, tmxName: "stairs_up2.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairsDownAddr, tmxName: "stairs_down.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1CandlestickAddr, tmxName: "candlestick.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairsDownPoisonAddr, tmxName: "stairs_down_poison.tmx", dtype: enum.DTypeCathedral},
		// l2.
		{addr: d1exe.L2VertArch1Addr, tmxName: "vert_arch1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch2Addr, tmxName: "vert_arch2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch3Addr, tmxName: "vert_arch3.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch4Addr, tmxName: "vert_arch4.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch5Addr, tmxName: "vert_arch5.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch6Addr, tmxName: "vert_arch6.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch7Addr, tmxName: "vert_arch7.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch8Addr, tmxName: "vert_arch8.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch9Addr, tmxName: "vert_arch9.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch10Addr, tmxName: "vert_arch10.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch11Addr, tmxName: "vert_arch11.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch12Addr, tmxName: "vert_arch12.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch13Addr, tmxName: "vert_arch13.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch14Addr, tmxName: "vert_arch14.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch15Addr, tmxName: "vert_arch15.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch16Addr, tmxName: "vert_arch16.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch17Addr, tmxName: "vert_arch17.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch18Addr, tmxName: "vert_arch18.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch19Addr, tmxName: "vert_arch19.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch20Addr, tmxName: "vert_arch20.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch21Addr, tmxName: "vert_arch21.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch22Addr, tmxName: "vert_arch22.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch23Addr, tmxName: "vert_arch23.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch24Addr, tmxName: "vert_arch24.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch25Addr, tmxName: "vert_arch25.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch26Addr, tmxName: "vert_arch26.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch27Addr, tmxName: "vert_arch27.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch28Addr, tmxName: "vert_arch28.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch29Addr, tmxName: "vert_arch29.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch30Addr, tmxName: "vert_arch30.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch31Addr, tmxName: "vert_arch31.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch32Addr, tmxName: "vert_arch32.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch33Addr, tmxName: "vert_arch33.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch34Addr, tmxName: "vert_arch34.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch35Addr, tmxName: "vert_arch35.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch36Addr, tmxName: "vert_arch36.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch37Addr, tmxName: "vert_arch37.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch38Addr, tmxName: "vert_arch38.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch39Addr, tmxName: "vert_arch39.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2VertArch40Addr, tmxName: "vert_arch40.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch1Addr, tmxName: "horiz_arch1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch2Addr, tmxName: "horiz_arch2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch3Addr, tmxName: "horiz_arch3.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch4Addr, tmxName: "horiz_arch4.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch5Addr, tmxName: "horiz_arch5.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch6Addr, tmxName: "horiz_arch6.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch7Addr, tmxName: "horiz_arch7.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch8Addr, tmxName: "horiz_arch8.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch9Addr, tmxName: "horiz_arch9.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch10Addr, tmxName: "horiz_arch10.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch11Addr, tmxName: "horiz_arch11.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch12Addr, tmxName: "horiz_arch12.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch13Addr, tmxName: "horiz_arch13.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch14Addr, tmxName: "horiz_arch14.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch15Addr, tmxName: "horiz_arch15.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch16Addr, tmxName: "horiz_arch16.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch17Addr, tmxName: "horiz_arch17.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch18Addr, tmxName: "horiz_arch18.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch19Addr, tmxName: "horiz_arch19.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch20Addr, tmxName: "horiz_arch20.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch21Addr, tmxName: "horiz_arch21.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch22Addr, tmxName: "horiz_arch22.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch23Addr, tmxName: "horiz_arch23.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch24Addr, tmxName: "horiz_arch24.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch25Addr, tmxName: "horiz_arch25.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch26Addr, tmxName: "horiz_arch26.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch27Addr, tmxName: "horiz_arch27.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch28Addr, tmxName: "horiz_arch28.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch29Addr, tmxName: "horiz_arch29.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch30Addr, tmxName: "horiz_arch30.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch31Addr, tmxName: "horiz_arch31.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch32Addr, tmxName: "horiz_arch32.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch33Addr, tmxName: "horiz_arch33.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch34Addr, tmxName: "horiz_arch34.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch35Addr, tmxName: "horiz_arch35.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch36Addr, tmxName: "horiz_arch36.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch37Addr, tmxName: "horiz_arch37.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch38Addr, tmxName: "horiz_arch38.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch39Addr, tmxName: "horiz_arch39.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2HorizArch40Addr, tmxName: "horiz_arch40.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2StairsUpAddr, tmxName: "stairs_up.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2StairsDownAddr, tmxName: "stairs_down.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2StairsToTownAddr, tmxName: "stairs_to_town.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CrushedColumnAddr, tmxName: "crushed_column.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big1Addr, tmxName: "big1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big2Addr, tmxName: "big2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big3Addr, tmxName: "big3.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big4Addr, tmxName: "big4.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big5Addr, tmxName: "big5.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big6Addr, tmxName: "big6.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big7Addr, tmxName: "big7.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big8Addr, tmxName: "big8.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big9Addr, tmxName: "big9.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Big10Addr, tmxName: "big10.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins1Addr, tmxName: "ruins1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins2Addr, tmxName: "ruins2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins3Addr, tmxName: "ruins3.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins4Addr, tmxName: "ruins4.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins5Addr, tmxName: "ruins5.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins6Addr, tmxName: "ruins6.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Ruins7Addr, tmxName: "ruins7.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Pancreas1Addr, tmxName: "pancreas1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2Pancreas2Addr, tmxName: "pancreas2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor1Addr, tmxName: "center_door1.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor2Addr, tmxName: "center_door2.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor3Addr, tmxName: "center_door3.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor4Addr, tmxName: "center_door4.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor5Addr, tmxName: "center_door5.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor6Addr, tmxName: "center_door6.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor7Addr, tmxName: "center_door7.tmx", dtype: enum.DTypeCatacombs},
		{addr: d1exe.L2CenterDoor8Addr, tmxName: "center_door8.tmx", dtype: enum.DTypeCatacombs},
		// l3.
		{addr: d1exe.L3StairsUpAddr, tmxName: "stairs_up.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3StairsDownAddr, tmxName: "stairs_down.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3StairsToTownAddr, tmxName: "stairs_to_town.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite1Addr, tmxName: "stalactite1.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite2Addr, tmxName: "stalactite2.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite3Addr, tmxName: "stalactite3.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite6Addr, tmxName: "stalactite6.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite7Addr, tmxName: "stalactite7.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite8Addr, tmxName: "stalactite8.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite9Addr, tmxName: "stalactite9.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite10Addr, tmxName: "stalactite10.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite11Addr, tmxName: "stalactite11.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite12Addr, tmxName: "stalactite12.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Stalactite13Addr, tmxName: "stalactite13.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice1Addr, tmxName: "crevice1.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice2Addr, tmxName: "crevice2.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice3Addr, tmxName: "crevice3.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice4Addr, tmxName: "crevice4.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice5Addr, tmxName: "crevice5.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice6Addr, tmxName: "crevice6.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice7Addr, tmxName: "crevice7.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice8Addr, tmxName: "crevice8.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice9Addr, tmxName: "crevice9.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice10Addr, tmxName: "crevice10.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Crevice11Addr, tmxName: "crevice11.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Isle1Addr, tmxName: "isle1.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Isle2Addr, tmxName: "isle2.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Isle3Addr, tmxName: "isle3.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Isle4Addr, tmxName: "isle4.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Isle5Addr, tmxName: "isle5.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Extra1Addr, tmxName: "extra1.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Extra2Addr, tmxName: "extra2.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Extra3Addr, tmxName: "extra3.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Extra4Addr, tmxName: "extra4.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3Extra5Addr, tmxName: "extra5.tmx", dtype: enum.DTypeCaves},
		{addr: d1exe.L3AnvilAddr, tmxName: "anvil.tmx", dtype: enum.DTypeCaves},
		// l4.
		{addr: d1exe.L4StairsUpAddr, tmxName: "stairs_up.tmx", dtype: enum.DTypeHell},
		{addr: d1exe.L4StairsToTownAddr, tmxName: "stairs_to_town.tmx", dtype: enum.DTypeHell},
		{addr: d1exe.L4StairsDownAddr, tmxName: "stairs_down.tmx", dtype: enum.DTypeHell},
		{addr: d1exe.L4Penta1Addr, tmxName: "penta1.tmx", dtype: enum.DTypeHell},
		{addr: d1exe.L4Penta2Addr, tmxName: "penta2.tmx", dtype: enum.DTypeHell},
	}
	for _, minisetInfo := range minisetInfos {
		mini, ok := exe.Minisets[minisetInfo.addr]
		if !ok {
			warn.Printf("unable to locate miniset at address 0x%08X", minisetInfo.addr)
			continue
		}
		levelInfo := getLevelInfo(minisetInfo.dtype)
		tmxMap, err := convertMinisetToTmx(mpqDir, mini, levelInfo)
		if err != nil {
			return errors.WithStack(err)
		}
		dtypeDir := strings.ToLower(minisetInfo.dtype.String())
		tmxPath := filepath.Join(outputDir, minisetsSubdir, dtypeDir, minisetInfo.tmxName)
		if err := dumpTmx(tmxPath, tmxMap); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

const (
	// Tile width (in pixels) of each tile on the map.
	mapTileWidth = 64
	// Tile height (in pixels) of each tile on the map.
	mapTileHeight = 32
)

// LevelInfo records information about a Diablo 1 tileset.
type LevelInfo struct {
	// Relative CEL path.
	relCelPath string
	// Relative MIN path.
	relMinPath string
	// Relative TIL path.
	relTilPath string
	// Tileset name; "cathedral", "catacombs", "caves", "hell" or "tristram".
	tilesetName string
	// Tileset spritesheet dimensions.
	tilesetWidth, tilesetHeight int
	// Tileset tile dimensions (i.e. dungeon piece dimensions).
	tilesetTileWidth, tilesetTileHeight int
}

// getLevelInfo returns the level information associated with the Diablo 1
// tileset contained within the given level directory.
func getLevelInfo(dtype enum.DType) LevelInfo {
	levelInfo := LevelInfo{}
	switch dtype {
	case enum.DTypeCathedral:
		levelInfo.relCelPath = "levels/l1data/l1.cel"
		levelInfo.relMinPath = "levels/l1data/l1.min"
		levelInfo.relTilPath = "levels/l1data/l1.til"
		levelInfo.tilesetName = "cathedral"
		levelInfo.tilesetWidth = 2048
		levelInfo.tilesetHeight = 2400
		levelInfo.tilesetTileWidth = 64
		levelInfo.tilesetTileHeight = 160
	case enum.DTypeCatacombs:
		levelInfo.relCelPath = "levels/l2data/l2.cel"
		levelInfo.relMinPath = "levels/l2data/l2.min"
		levelInfo.relTilPath = "levels/l2data/l2.til"
		levelInfo.tilesetName = "catacombs"
		levelInfo.tilesetWidth = 2048
		levelInfo.tilesetHeight = 2880
		levelInfo.tilesetTileWidth = 64
		levelInfo.tilesetTileHeight = 160
	case enum.DTypeCaves:
		levelInfo.relCelPath = "levels/l3data/l3.cel"
		levelInfo.relMinPath = "levels/l3data/l3.min"
		levelInfo.relTilPath = "levels/l3data/l3.til"
		levelInfo.tilesetName = "caves"
		levelInfo.tilesetWidth = 2048
		levelInfo.tilesetHeight = 2880
		levelInfo.tilesetTileWidth = 64
		levelInfo.tilesetTileHeight = 160
	case enum.DTypeHell:
		levelInfo.relCelPath = "levels/l4data/l4.cel"
		levelInfo.relMinPath = "levels/l4data/l4.min"
		levelInfo.relTilPath = "levels/l4data/l4.til"
		levelInfo.tilesetName = "hell"
		levelInfo.tilesetWidth = 2048
		levelInfo.tilesetHeight = 3840
		levelInfo.tilesetTileWidth = 64
		levelInfo.tilesetTileHeight = 256
	case enum.DTypeTristram:
		levelInfo.relCelPath = "levels/towndata/town.cel"
		levelInfo.relMinPath = "levels/towndata/town.min"
		levelInfo.relTilPath = "levels/towndata/town.til"
		levelInfo.tilesetName = "tristram"
		levelInfo.tilesetWidth = 4096
		levelInfo.tilesetHeight = 5120
		levelInfo.tilesetTileWidth = 64
		levelInfo.tilesetTileHeight = 256
	default:
		panic(fmt.Errorf(`support for dungeon type %v not yet implemented; ensure that the path is relative to the MPQ directory (e.g. "levels/l1data")`, dtype))
	}
	return levelInfo
}

// convertMinisetToTmx converts the given miniset to TMX format.
func convertMinisetToTmx(mpqDir string, mini *miniset.Miniset, levelInfo LevelInfo) (*tmx.Map, error) {
	// Create TMX map.
	m := &tmx.Map{
		Orientation: "isometric",
		Width:       int(2 * mini.Width),
		Height:      int(2 * mini.Height),
		TileWidth:   mapTileWidth,
		TileHeight:  mapTileHeight,
	}
	// Create TMX tileset.
	tilesetPngName := fmt.Sprintf("%s.png", levelInfo.tilesetName)
	tilesetImg := tmx.Image{
		// Traverse up to the root of the "_dump_" directory.
		Source: filepath.Join("../..", tilesetsSubdir, tilesetPngName),
		Width:  float64(levelInfo.tilesetWidth),
		Height: float64(levelInfo.tilesetHeight),
	}
	tileset := tmx.Tileset{
		FirstGID:   1,
		Name:       levelInfo.tilesetName,
		TileWidth:  levelInfo.tilesetTileWidth,
		TileHeight: levelInfo.tilesetTileHeight,
		Image:      []tmx.Image{tilesetImg},
	}
	m.Tilesets = append(m.Tilesets, tileset)
	// Convert tiles to dungeon pieces. Each tile (2x2) correspond to four
	// dungeon pieces (1x1).
	tilPath := filepath.Join(mpqDir, levelInfo.relTilPath)
	tiles, err := til.ParseFile(tilPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	addLayer(m, "before", mini.Before, tiles)
	addLayer(m, "after", mini.After, tiles)
	return m, nil
}

// addLayer adds a layer with the given name of dungeon peice IDs to the TMX
// map. The given tile IDs are translated to dungeon piece IDS using the
// specified tile definitions.
func addLayer(m *tmx.Map, layerName string, tileIDs [][]uint8, tiles []til.Tile) {
	dpieceIDs := dpiecesFromTiles(tileIDs, tiles)
	// Convert to CSV format.
	inner := &strings.Builder{}
	inner.WriteString("\n")
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			dpieceID := dpieceIDs[x][y]
			inner.WriteString(strconv.Itoa(dpieceID))
			if !(x == m.Width-1 && y == m.Height-1) {
				inner.WriteString(",")
			}
		}
		inner.WriteString("\n")
	}
	data := tmx.Data{
		Encoding: "csv",
		Inner:    inner.String(),
	}
	layer := tmx.Layer{
		Name:    layerName,
		Data:    []tmx.Data{data},
		Width:   m.Width,
		Height:  m.Height,
		Opacity: 1.0,
		Visible: 1,
	}
	m.Layers = append(m.Layers, layer)
}

// dpiecesFromTiles converts the given two-dimensional array of tile IDs to a
// two-dimensional array of dungeon piece IDs. The width and height of the
// output array are twice that of the input array.
func dpiecesFromTiles(tileIDs [][]uint8, tiles []til.Tile) [][]int {
	width := 2 * len(tileIDs)
	height := 2 * len(tileIDs[0])
	dpieceIDs := make([][]int, width)
	for x := range dpieceIDs {
		dpieceIDs[x] = make([]int, height)
	}
	for xx := range tileIDs {
		for yy := range tileIDs[xx] {
			tileID := int(tileIDs[xx][yy])
			if tileID == 0 {
				continue
			}
			x := 2 * xx
			y := 2 * yy
			t := tiles[tileID-1]
			dpieceIDs[x][y] = int(t.Top) + 1
			dpieceIDs[x+1][y] = int(t.Right) + 1
			dpieceIDs[x][y+1] = int(t.Left) + 1
			dpieceIDs[x+1][y+1] = int(t.Bottom) + 1
		}
	}
	return dpieceIDs
}

// dumpTmx stores the given TMX map to the specified TMX path.
func dumpTmx(tmxPath string, m *tmx.Map) error {
	dbg.Printf("creating %q.", tmxPath)
	tmxDir := filepath.Dir(tmxPath)
	if err := os.MkdirAll(tmxDir, 0755); err != nil {
		return errors.WithStack(err)
	}
	buf, err := xml.MarshalIndent(m, "", "\t")
	if err != nil {
		return errors.WithStack(err)
	}
	if err := ioutil.WriteFile(tmxPath, buf, 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
