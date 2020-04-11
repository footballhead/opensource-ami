#!/bin/bash

# === [ Towners ] ==============================================================

mkdir -p _assets_/towners

# --- [ Griswold the Blacksmith ] ----------------------------------------------

# Actions:
#
#    stand: towners/smith/smithn/*.png
#    walk: towners/smith/smithw/smithw_{0,1,2,3,4,5,6,7}/*.png

# Animation order in spritesheet:
#
#    * stand (SW): 16 frames
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
#    S  (smithw_0)
#    SW (smithw_1)
#    W  (smithw_2)
#    NW (smithw_3)
#    N  (smithw_4)
#    NE (smithw_5)
#    E  (smithw_6)
#    SE (smithw_7)

# width: 96
# height: 96
# nframes (stand): 16
# nframes (walk): 8

# total frames: 16 + 8*8 = 80

echo 'Extracting graphics for "Griswold the Blacksmith"'
montage \
	_dump_/towners/smith/smithn/*.png \
	_dump_/towners/smith/smithw/smithw_0/*.png \
	_dump_/towners/smith/smithw/smithw_1/*.png \
	_dump_/towners/smith/smithw/smithw_2/*.png \
	_dump_/towners/smith/smithw/smithw_3/*.png \
	_dump_/towners/smith/smithw/smithw_4/*.png \
	_dump_/towners/smith/smithw/smithw_5/*.png \
	_dump_/towners/smith/smithw/smithw_6/*.png \
	_dump_/towners/smith/smithw/smithw_7/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 8x \
	-background none \
	_assets_/towners/griswold.png

# --- [ Pepin the Healer ] -----------------------------------------------------

# --- [ Wounded Townsman ] -----------------------------------------------------

# Actions:
#
#    stand: towners/butch/deadguy/*.png

# Animation order in spritesheet:
#
#    * stand (N): 8 frames

# width: 96
# height: 96
# nframes (stand): 8
# TODO: split stand and death anim.

# total frames: 8

echo 'Extracting graphics for "Wounded Townsman"'
montage \
	_dump_/towners/butch/deadguy/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 8x \
	-background none \
	_assets_/towners/wounded_townsman.png

# --- [ Ogden the Tavern owner ] -----------------------------------------------

# Actions:
#
#    stand: towners/twnf/twnfn/*.png
#    walk: towners/twnf/twnfw/twnfw_{0,1,2,3,4,5,6,7}/*.png

# Animation order in spritesheet:
#
#    * stand (SW): 16 frames
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
#    S  (twnfw_0)
#    SW (twnfw_1)
#    W  (twnfw_2)
#    NW (twnfw_3)
#    N  (twnfw_4)
#    NE (twnfw_5)
#    E  (twnfw_6)
#    SE (twnfw_7)

# width: 96
# height: 96
# nframes (stand): 16
# nframes (walk): 8

# total frames: 16 + 8*8 = 80

echo 'Extracting graphics for "Ogden the Tavern owner"'
montage \
	_dump_/towners/twnf/twnfn/*.png \
	_dump_/towners/twnf/twnfw/twnfw_0/*.png \
	_dump_/towners/twnf/twnfw/twnfw_1/*.png \
	_dump_/towners/twnf/twnfw/twnfw_2/*.png \
	_dump_/towners/twnf/twnfw/twnfw_3/*.png \
	_dump_/towners/twnf/twnfw/twnfw_4/*.png \
	_dump_/towners/twnf/twnfw/twnfw_5/*.png \
	_dump_/towners/twnf/twnfw/twnfw_6/*.png \
	_dump_/towners/twnf/twnfw/twnfw_7/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 8x \
	-background none \
	_assets_/towners/ogden.png

# --- [ Cain the Elder ] -------------------------------------------------------

# --- [ Farnham the Drunk ] ----------------------------------------------------

# --- [ Adria the Witch ] ------------------------------------------------------

# Actions:
#
#    stand: towners/townwmn1/witch/*.png

# Animation order in spritesheet:
#
#    * stand (S): 19 frames

# width: 96
# height: 96
# nframes (stand): 19

# total frames: 19

echo 'Extracting graphics for "Adria the Witch"'
montage \
	_dump_/towners/townwmn1/witch/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 19x \
	-background none \
	_assets_/towners/adria.png

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
# height: 96
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

# Actions:
#
#    stand: towners/townboy/pegkid1/*.png

# Animation order in spritesheet:
#
#    * stand (S): 20 frames

# width: 96
# height: 64
# nframes (stand): 20

# total frames: 20

echo 'Extracting graphics for "Wirt the Peg-legged boy"'
montage \
	_dump_/towners/townboy/pegkid1/*.png \
	-gravity south -geometry 96x+0+0 \
	-tile 5x \
	-background none \
	_assets_/towners/wirt.png

# --- [ Cow ] ------------------------------------------------------------------

# Actions:
#
#    stand: towners/animals/cow/cow_{0,1,2,3,4,5,6,7}/*.png

# Animation order in spritesheet:
#
#    * stand (S):  12 frames
#    * stand (SW): 12 frames
#    * stand (W):  12 frames
#    * stand (NW): 12 frames
#    * stand (N):  12 frames
#    * stand (NE): 12 frames
#    * stand (E):  12 frames
#    * stand (SE): 12 frames

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
# height: 128
# nframes (stand): 12

# total frames: 8*12 = 96

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
