#!/bin/bash

# === [ Towners ] ==============================================================

mkdir -p _assets_/towners

# --- [ Griswold the Blacksmith ] ----------------------------------------------

# --- [ Pepin the Healer ] -----------------------------------------------------

# --- [ Wounded Townsman ] -----------------------------------------------------

# Actions:
#
#    stand

# width: 96
# nframes: 8

echo 'Extracting graphics for "Wounded Townsman"'
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

# Actions:
#
#    stand: towners/townwmn1/wmnn/*.png
#    walk: towners/townwmn1/wmnw/wmnw_{0,1,2,3,4,5,6,7}/*.png

# Animation order in spritesheet:
#
#    * stand (SE): 18 frames
#    * walk  (S):   8 frames
#    * walk  (SW):  8 frames
#    * walk  (W):   8 frames
#    * walk  (NW):  8 frames
#    * walk  (N):   8 frames
#    * walk  (NE):  8 frames
#    * walk  (E):   8 frames
#    * walk  (SE):  8 frames

# Directions:
#
#    S  (wmnw_0)
#    SW (wmnw_1)
#    W  (wmnw_2)
#    NW (wmnw_3)
#    N  (wmnw_4)
#    NE (wmnw_5)
#    E  (wmnw_6)
#    SE (wmnw_7)

# width: 96
# nframes (stand): 18
# nframes (walk): 8

# total frames: 18 + 8*8 = 82

echo 'Extracting graphics for "Gillian the Barmaid"'
montage \
	_dump_/towners/townwmn1/wmnn/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_0/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_1/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_2/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_3/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_4/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_5/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_6/*.png \
	_dump_/towners/townwmn1/wmnw/wmnw_7/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 41x \
	-background none \
	_assets_/towners/gillian.png

# --- [ Wirt the Peg-legged boy ] ----------------------------------------------

# --- [ Cow ] ------------------------------------------------------------------

# Actions:
#
#    stand

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

echo 'Extracting graphics for "Cow"'
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
