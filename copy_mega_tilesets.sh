#!/bin/bash

# === [ Mega tilesets ] ========================================================

mkdir -p _assets_/_mega_tilesets_

# --- [ Tristram ] -------------------------------------------------------------

echo 'Creating "_assets_/_mega_tilesets_/tristram.png"'
montage \
	_dump_/_tiles_/town/town.pal/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 19x \
	-background none \
	_assets_/_mega_tilesets_/tristram.png

# --- [ Cathedral ] ------------------------------------------------------------

echo 'Creating "_assets_/_mega_tilesets_/cathedral.png"'
montage \
	_dump_/_tiles_/l1/l1_1.pal/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 23x \
	-background none \
	_assets_/_mega_tilesets_/cathedral.png

# --- [ Catacombs ] ------------------------------------------------------------

echo 'Creating "_assets_/_mega_tilesets_/catacombs.png"'
montage \
	_dump_/_tiles_/l2/l2_1.pal/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 16x \
	-background none \
	_assets_/_mega_tilesets_/catacombs.png

# --- [ Caves ] ----------------------------------------------------------------

echo 'Creating "_assets_/_mega_tilesets_/caves.png"'
montage \
	_dump_/_tiles_/l3/l3_1.pal/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 13x \
	-background none \
	_assets_/_mega_tilesets_/caves.png

# --- [ Hell ] -----------------------------------------------------------------

echo 'Creating "_assets_/_mega_tilesets_/hell.png"'
montage \
	_dump_/_tiles_/l4/l4_1.pal/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 23x \
	-background none \
	_assets_/_mega_tilesets_/hell.png
