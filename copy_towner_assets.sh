#!/bin/bash

# === [ Towners ] ==============================================================

mkdir -p _assets_/towners

# --- [ Griswold the Blacksmith ] ----------------------------------------------

# --- [ Pepin the Healer ] -----------------------------------------------------

# --- [ Wounded Townsman ] -----------------------------------------------------

# width: 96
# nframes: 8

echo 'Extracting graphics for "Wounded Townsman" towner'
montage \
	_dump_/towners/butch/deadguy/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 8x \
	-background none \
	_assets_/towners/wounded_townsman.png

# --- [ Ogden the Tavern owner ] -----------------------------------------------

# --- [ Cain the Elder ] -------------------------------------------------------

# --- [ Farnham the Drunk ] ----------------------------------------------------

# --- [ Adria the Witch ] ------------------------------------------------------

# --- [ Gillian the Barmaid ] --------------------------------------------------

# --- [ Wirt the Peg-legged boy ] ----------------------------------------------

# --- [ Cow ] ------------------------------------------------------------------

# Directions:
#
#    S  (cow_0)
#    SW (cow_1)
#    W  (cow_2)
#    NW (cow_3)
#    N  (cow_4)
#    NE (cow_5)
#    E  (cow_6)
#    SE (cow_7)

# width: 128
# nframes: 12

echo 'Extracting graphics for "Cow" towner'
montage \
	_dump_/towners/animals/cow/cow_0/*.png \
	_dump_/towners/animals/cow/cow_1/*.png \
	_dump_/towners/animals/cow/cow_2/*.png \
	_dump_/towners/animals/cow/cow_3/*.png \
	_dump_/towners/animals/cow/cow_4/*.png \
	_dump_/towners/animals/cow/cow_5/*.png \
	_dump_/towners/animals/cow/cow_6/*.png \
	_dump_/towners/animals/cow/cow_7/*.png \
	-gravity south -geometry 128x+0+0 \
	-tile 12x \
	-background none \
	_assets_/towners/cow.png
