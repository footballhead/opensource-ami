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

// MinisetInfo records the information related to an embedded miniset.
type MinisetInfo struct {
	addr    uint32
	tmxName string
	dtype   enum.DType
}

// dump retrieves information stored within the diablo.exe executable.
func dump(mpqDir, exePath, outputDir string) error {
	exe, err := d1exe.ParseFile(exePath)
	if err != nil {
		return errors.WithStack(err)
	}
	// Miniset files.
	var minisetInfos []MinisetInfo
	switch exe.Version {
	case d1exe.DiabloVersion109:
		minisetInfos = []MinisetInfo{
			// l1 minisets.
			{addr: d1exe.L1StairsUp1Addr, tmxName: "stairsup.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1StairsUp2Addr, tmxName: "l5stairsup.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1StairsDownAddr, tmxName: "stairsdown.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1CandlestickAddr, tmxName: "lamps.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1StairsDownPoisonAddr, tmxName: "pwaterin.tmx", dtype: enum.DTypeCathedral},
			// l2 minisets.
			{addr: d1exe.L2VertArch1Addr, tmxName: "varch1.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch2Addr, tmxName: "varch2.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch3Addr, tmxName: "varch3.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch4Addr, tmxName: "varch4.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch5Addr, tmxName: "varch5.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch6Addr, tmxName: "varch6.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch7Addr, tmxName: "varch7.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch8Addr, tmxName: "varch8.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch9Addr, tmxName: "varch9.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch10Addr, tmxName: "varch10.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch11Addr, tmxName: "varch11.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch12Addr, tmxName: "varch12.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch13Addr, tmxName: "varch13.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch14Addr, tmxName: "varch14.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch15Addr, tmxName: "varch15.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch16Addr, tmxName: "varch16.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch17Addr, tmxName: "varch17.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch18Addr, tmxName: "varch18.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch19Addr, tmxName: "varch19.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch20Addr, tmxName: "varch20.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch21Addr, tmxName: "varch21.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch22Addr, tmxName: "varch22.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch23Addr, tmxName: "varch23.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch24Addr, tmxName: "varch24.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch25Addr, tmxName: "varch25.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch26Addr, tmxName: "varch26.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch27Addr, tmxName: "varch27.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch28Addr, tmxName: "varch28.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch29Addr, tmxName: "varch29.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch30Addr, tmxName: "varch30.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch31Addr, tmxName: "varch31.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch32Addr, tmxName: "varch32.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch33Addr, tmxName: "varch33.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch34Addr, tmxName: "varch34.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch35Addr, tmxName: "varch35.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch36Addr, tmxName: "varch36.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch37Addr, tmxName: "varch37.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch38Addr, tmxName: "varch38.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch39Addr, tmxName: "varch39.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2VertArch40Addr, tmxName: "varch40.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch1Addr, tmxName: "harch1.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch2Addr, tmxName: "harch2.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch3Addr, tmxName: "harch3.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch4Addr, tmxName: "harch4.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch5Addr, tmxName: "harch5.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch6Addr, tmxName: "harch6.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch7Addr, tmxName: "harch7.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch8Addr, tmxName: "harch8.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch9Addr, tmxName: "harch9.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch10Addr, tmxName: "harch10.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch11Addr, tmxName: "harch11.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch12Addr, tmxName: "harch12.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch13Addr, tmxName: "harch13.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch14Addr, tmxName: "harch14.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch15Addr, tmxName: "harch15.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch16Addr, tmxName: "harch16.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch17Addr, tmxName: "harch17.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch18Addr, tmxName: "harch18.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch19Addr, tmxName: "harch19.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch20Addr, tmxName: "harch20.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch21Addr, tmxName: "harch21.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch22Addr, tmxName: "harch22.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch23Addr, tmxName: "harch23.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch24Addr, tmxName: "harch24.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch25Addr, tmxName: "harch25.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch26Addr, tmxName: "harch26.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch27Addr, tmxName: "harch27.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch28Addr, tmxName: "harch28.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch29Addr, tmxName: "harch29.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch30Addr, tmxName: "harch30.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch31Addr, tmxName: "harch31.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch32Addr, tmxName: "harch32.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch33Addr, tmxName: "harch33.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch34Addr, tmxName: "harch34.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch35Addr, tmxName: "harch35.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch36Addr, tmxName: "harch36.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch37Addr, tmxName: "harch37.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch38Addr, tmxName: "harch38.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch39Addr, tmxName: "harch39.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2HorizArch40Addr, tmxName: "harch40.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2StairsUpAddr, tmxName: "ustairs.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2StairsDownAddr, tmxName: "dstairs.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2StairsToTownAddr, tmxName: "warpstairs.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CrushedColumnAddr, tmxName: "crushcol.tmx", dtype: enum.DTypeCatacombs},
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
			{addr: d1exe.L2CenterDoor1Addr, tmxName: "ctrdoor1.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor2Addr, tmxName: "ctrdoor2.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor3Addr, tmxName: "ctrdoor3.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor4Addr, tmxName: "ctrdoor4.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor5Addr, tmxName: "ctrdoor5.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor6Addr, tmxName: "ctrdoor6.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor7Addr, tmxName: "ctrdoor7.tmx", dtype: enum.DTypeCatacombs},
			{addr: d1exe.L2CenterDoor8Addr, tmxName: "ctrdoor8.tmx", dtype: enum.DTypeCatacombs},
			// l3 minisets.
			{addr: d1exe.L3StairsUpAddr, tmxName: "l3up.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3StairsDownAddr, tmxName: "l3down.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3StairsToTownAddr, tmxName: "l3holdwarp.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite1Addr, tmxName: "l3tite1.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite2Addr, tmxName: "l3tite2.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite3Addr, tmxName: "l3tite3.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite6Addr, tmxName: "l3tite6.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite7Addr, tmxName: "l3tite7.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite8Addr, tmxName: "l3tite8.tmx", dtype: enum.DTypeCaves},   // -> stalacmite (grows from floor up)
			{addr: d1exe.L3Stalactite9Addr, tmxName: "l3tite9.tmx", dtype: enum.DTypeCaves},   // -> stalacmite (grows from floor up)
			{addr: d1exe.L3Stalactite10Addr, tmxName: "l3tite10.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			{addr: d1exe.L3Stalactite11Addr, tmxName: "l3tite11.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			{addr: d1exe.L3Stalactite12Addr, tmxName: "l3tite12.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Stalactite13Addr, tmxName: "l3tite13.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice1Addr, tmxName: "l3crev1.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice2Addr, tmxName: "l3crev2.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice3Addr, tmxName: "l3crev3.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice4Addr, tmxName: "l3crev4.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice5Addr, tmxName: "l3crev5.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice6Addr, tmxName: "l3crev6.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice7Addr, tmxName: "l3crev7.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice8Addr, tmxName: "l3crev8.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice9Addr, tmxName: "l3crev9.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice10Addr, tmxName: "l3crev10.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Crevice11Addr, tmxName: "l3crev11.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Isle1Addr, tmxName: "l3isle1.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Isle2Addr, tmxName: "l3isle2.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Isle3Addr, tmxName: "l3isle3.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Isle4Addr, tmxName: "l3isle4.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Isle5Addr, tmxName: "l3isle5.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Extra1Addr, tmxName: "l3xtra1.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Extra2Addr, tmxName: "l3xtra2.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Extra3Addr, tmxName: "l3xtra3.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Extra4Addr, tmxName: "l3xtra4.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3Extra5Addr, tmxName: "l3xtra5.tmx", dtype: enum.DTypeCaves},
			{addr: d1exe.L3AnvilAddr, tmxName: "l3anvil.tmx", dtype: enum.DTypeCaves},
			// l4 minisets.
			{addr: d1exe.L4StairsUpAddr, tmxName: "l4ustairs.tmx", dtype: enum.DTypeHell},
			{addr: d1exe.L4StairsToTownAddr, tmxName: "l4twarp.tmx", dtype: enum.DTypeHell},
			{addr: d1exe.L4StairsDownAddr, tmxName: "l4dstairs.tmx", dtype: enum.DTypeHell},
			{addr: d1exe.L4Penta1Addr, tmxName: "l4penta.tmx", dtype: enum.DTypeHell},
			{addr: d1exe.L4Penta2Addr, tmxName: "l4penta2.tmx", dtype: enum.DTypeHell},
		}
	case d1exe.DiabloVersionAlpha4:
		minisetInfos = []MinisetInfo{
			// l1 minisets (identical, missing PWATERIN).
			//{addr: d1exe.L1Alpha4StairsUp1Addr, tmxName: "alpha4_stairsup.tmx", dtype: enum.DTypeCathedral},
			//{addr: d1exe.L1Alpha4StairsUp2Addr, tmxName: "alpha4_l5stairsup.tmx", dtype: enum.DTypeCathedral},
			//{addr: d1exe.L1Alpha4StairsDownAddr, tmxName: "alpha4_stairsdown.tmx", dtype: enum.DTypeCathedral},
			//{addr: d1exe.L1Alpha4CandlestickAddr, tmxName: "alpha4_lamps.tmx", dtype: enum.DTypeCathedral},
			// l1 minisets (extra).
			{addr: d1exe.L1Alpha4SarcophagusAddr, tmxName: "alpha4_sarc.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1Alpha4RightCorridorAddr, tmxName: "alpha4_rcorridor.tmx", dtype: enum.DTypeCathedral},
			{addr: d1exe.L1Alpha4LeftCorridorAddr, tmxName: "alpha4_lcorridor.tmx", dtype: enum.DTypeCathedral},
			// l2 minisets (identical, missing WARPSTAIRS).
			//{addr: d1exe.L2Alpha4VertArch1Addr, tmxName: "alpha4_varch1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch2Addr, tmxName: "alpha4_varch2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch3Addr, tmxName: "alpha4_varch3.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch4Addr, tmxName: "alpha4_varch4.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch5Addr, tmxName: "alpha4_varch5.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch6Addr, tmxName: "alpha4_varch6.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch7Addr, tmxName: "alpha4_varch7.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch8Addr, tmxName: "alpha4_varch8.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch9Addr, tmxName: "alpha4_varch9.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch10Addr, tmxName: "alpha4_varch10.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch11Addr, tmxName: "alpha4_varch11.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch12Addr, tmxName: "alpha4_varch12.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch13Addr, tmxName: "alpha4_varch13.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch14Addr, tmxName: "alpha4_varch14.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch15Addr, tmxName: "alpha4_varch15.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch16Addr, tmxName: "alpha4_varch16.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch17Addr, tmxName: "alpha4_varch17.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch18Addr, tmxName: "alpha4_varch18.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch19Addr, tmxName: "alpha4_varch19.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch20Addr, tmxName: "alpha4_varch20.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch21Addr, tmxName: "alpha4_varch21.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch22Addr, tmxName: "alpha4_varch22.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch23Addr, tmxName: "alpha4_varch23.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch24Addr, tmxName: "alpha4_varch24.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch25Addr, tmxName: "alpha4_varch25.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch26Addr, tmxName: "alpha4_varch26.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch27Addr, tmxName: "alpha4_varch27.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch28Addr, tmxName: "alpha4_varch28.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch29Addr, tmxName: "alpha4_varch29.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch30Addr, tmxName: "alpha4_varch30.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch31Addr, tmxName: "alpha4_varch31.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch32Addr, tmxName: "alpha4_varch32.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch33Addr, tmxName: "alpha4_varch33.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch34Addr, tmxName: "alpha4_varch34.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch35Addr, tmxName: "alpha4_varch35.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch36Addr, tmxName: "alpha4_varch36.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch37Addr, tmxName: "alpha4_varch37.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch38Addr, tmxName: "alpha4_varch38.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch39Addr, tmxName: "alpha4_varch39.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4VertArch40Addr, tmxName: "alpha4_varch40.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch1Addr, tmxName: "alpha4_harch1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch2Addr, tmxName: "alpha4_harch2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch3Addr, tmxName: "alpha4_harch3.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch4Addr, tmxName: "alpha4_harch4.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch5Addr, tmxName: "alpha4_harch5.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch6Addr, tmxName: "alpha4_harch6.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch7Addr, tmxName: "alpha4_harch7.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch8Addr, tmxName: "alpha4_harch8.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch9Addr, tmxName: "alpha4_harch9.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch10Addr, tmxName: "alpha4_harch10.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch11Addr, tmxName: "alpha4_harch11.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch12Addr, tmxName: "alpha4_harch12.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch13Addr, tmxName: "alpha4_harch13.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch14Addr, tmxName: "alpha4_harch14.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch15Addr, tmxName: "alpha4_harch15.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch16Addr, tmxName: "alpha4_harch16.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch17Addr, tmxName: "alpha4_harch17.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch18Addr, tmxName: "alpha4_harch18.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch19Addr, tmxName: "alpha4_harch19.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch20Addr, tmxName: "alpha4_harch20.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch21Addr, tmxName: "alpha4_harch21.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch22Addr, tmxName: "alpha4_harch22.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch23Addr, tmxName: "alpha4_harch23.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch24Addr, tmxName: "alpha4_harch24.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch25Addr, tmxName: "alpha4_harch25.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch26Addr, tmxName: "alpha4_harch26.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch27Addr, tmxName: "alpha4_harch27.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch28Addr, tmxName: "alpha4_harch28.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch29Addr, tmxName: "alpha4_harch29.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch30Addr, tmxName: "alpha4_harch30.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch31Addr, tmxName: "alpha4_harch31.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch32Addr, tmxName: "alpha4_harch32.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch33Addr, tmxName: "alpha4_harch33.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch34Addr, tmxName: "alpha4_harch34.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch35Addr, tmxName: "alpha4_harch35.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch36Addr, tmxName: "alpha4_harch36.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch37Addr, tmxName: "alpha4_harch37.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch38Addr, tmxName: "alpha4_harch38.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch39Addr, tmxName: "alpha4_harch39.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4HorizArch40Addr, tmxName: "alpha4_harch40.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4StairsUpAddr, tmxName: "alpha4_ustairs.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4StairsDownAddr, tmxName: "alpha4_dstairs.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CrushedColumnAddr, tmxName: "alpha4_crushcol.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big1Addr, tmxName: "alpha4_big1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big2Addr, tmxName: "alpha4_big2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big3Addr, tmxName: "alpha4_big3.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big4Addr, tmxName: "alpha4_big4.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big5Addr, tmxName: "alpha4_big5.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big6Addr, tmxName: "alpha4_big6.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big7Addr, tmxName: "alpha4_big7.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big8Addr, tmxName: "alpha4_big8.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big9Addr, tmxName: "alpha4_big9.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Big10Addr, tmxName: "alpha4_big10.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins1Addr, tmxName: "alpha4_ruins1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins2Addr, tmxName: "alpha4_ruins2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins3Addr, tmxName: "alpha4_ruins3.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins4Addr, tmxName: "alpha4_ruins4.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins5Addr, tmxName: "alpha4_ruins5.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins6Addr, tmxName: "alpha4_ruins6.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Ruins7Addr, tmxName: "alpha4_ruins7.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Pancreas1Addr, tmxName: "alpha4_pancreas1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4Pancreas2Addr, tmxName: "alpha4_pancreas2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor1Addr, tmxName: "alpha4_ctrdoor1.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor2Addr, tmxName: "alpha4_ctrdoor2.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor3Addr, tmxName: "alpha4_ctrdoor3.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor4Addr, tmxName: "alpha4_ctrdoor4.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor5Addr, tmxName: "alpha4_ctrdoor5.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor6Addr, tmxName: "alpha4_ctrdoor6.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor7Addr, tmxName: "alpha4_ctrdoor7.tmx", dtype: enum.DTypeCatacombs},
			//{addr: d1exe.L2Alpha4CenterDoor8Addr, tmxName: "alpha4_ctrdoor8.tmx", dtype: enum.DTypeCatacombs},
			// l3 minisets (identical, missing L3HOLDWARP, L3ANVIL).
			//{addr: d1exe.L3Alpha4StairsUpAddr, tmxName: "alpha4_l3up.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4StairsDownAddr, tmxName: "alpha4_l3down.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite1Addr, tmxName: "alpha4_l3tite1.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite2Addr, tmxName: "alpha4_l3tite2.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite3Addr, tmxName: "alpha4_l3tite3.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite6Addr, tmxName: "alpha4_l3tite6.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite7Addr, tmxName: "alpha4_l3tite7.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite8Addr, tmxName: "alpha4_l3tite8.tmx", dtype: enum.DTypeCaves},   // -> stalacmite (grows from floor up)
			//{addr: d1exe.L3Alpha4Stalactite9Addr, tmxName: "alpha4_l3tite9.tmx", dtype: enum.DTypeCaves},   // -> stalacmite (grows from floor up)
			//{addr: d1exe.L3Alpha4Stalactite10Addr, tmxName: "alpha4_l3tite10.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			//{addr: d1exe.L3Alpha4Stalactite11Addr, tmxName: "alpha4_l3tite11.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			//{addr: d1exe.L3Alpha4Stalactite12Addr, tmxName: "alpha4_l3tite12.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Stalactite13Addr, tmxName: "alpha4_l3tite13.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice1Addr, tmxName: "alpha4_l3crev1.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice2Addr, tmxName: "alpha4_l3crev2.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice3Addr, tmxName: "alpha4_l3crev3.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice4Addr, tmxName: "alpha4_l3crev4.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice5Addr, tmxName: "alpha4_l3crev5.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice6Addr, tmxName: "alpha4_l3crev6.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice7Addr, tmxName: "alpha4_l3crev7.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice8Addr, tmxName: "alpha4_l3crev8.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice9Addr, tmxName: "alpha4_l3crev9.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice10Addr, tmxName: "alpha4_l3crev10.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Crevice11Addr, tmxName: "alpha4_l3crev11.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Isle1Addr, tmxName: "alpha4_l3isle1.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Isle2Addr, tmxName: "alpha4_l3isle2.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Isle3Addr, tmxName: "alpha4_l3isle3.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Isle4Addr, tmxName: "alpha4_l3isle4.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Isle5Addr, tmxName: "alpha4_l3isle5.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Extra1Addr, tmxName: "alpha4_l3xtra1.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Extra2Addr, tmxName: "alpha4_l3xtra2.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Extra3Addr, tmxName: "alpha4_l3xtra3.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Extra4Addr, tmxName: "alpha4_l3xtra4.tmx", dtype: enum.DTypeCaves},
			//{addr: d1exe.L3Alpha4Extra5Addr, tmxName: "alpha4_l3xtra5.tmx", dtype: enum.DTypeCaves},
			// l3 minisets (extra).
			{addr: d1exe.L3Alpha4Stalactite4Addr, tmxName: "alpha4_l3tite4.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			{addr: d1exe.L3Alpha4Stalactite5Addr, tmxName: "alpha4_l3tite5.tmx", dtype: enum.DTypeCaves}, // -> stalacmite (grows from floor up)
			// l4 minisets (extra).
			{addr: d1exe.L4Alpha4StairsUpAddr, tmxName: "alpha4_l4ustairs.tmx", dtype: enum.DTypeHell},
			{addr: d1exe.L4Alpha4StairsDownAddr, tmxName: "alpha4_l4dstairs.tmx", dtype: enum.DTypeHell},
		}
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
