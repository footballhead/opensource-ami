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
		{addr: d1exe.L1StairUp1Addr, tmxName: "stair_up1.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairUp2Addr, tmxName: "stair_up2.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairDownAddr, tmxName: "stair_down.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1CandlestickAddr, tmxName: "candlestick.tmx", dtype: enum.DTypeCathedral},
		{addr: d1exe.L1StairDownPoisonAddr, tmxName: "stair_down_poison.tmx", dtype: enum.DTypeCathedral},
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
