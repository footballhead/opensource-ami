#!/bin/bash

# === [ Towners ] ==============================================================

mkdir -p _assets_/towners

echo 'Extracting graphics for "Wounded Townsman" towner'
# width: 96
montage \
	_dump_/towners/butch/deadguy/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 8x \
	-background none \
	_assets_/towners/wounded_townsman.png
