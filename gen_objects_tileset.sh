#!/bin/bash

# Object IDs (in order of occurrence in _assets_/_tilesets_/objects.png):
#
#    * object ID 83: altboy/altboy.png
#    * object ID 24: angel/angel.png
#    * object ID 77: armstand/armstand_0001.png
#    * object ID 89: armstand/armstand_0001.png
#    * object ID 78: armstand/armstand_0002.png
#    * object ID 12: banner/banner_0001.png
#    * object ID 11: banner/banner_0002.png
#    * object ID 13: banner/banner_0003.png
#    * object ID 57: barrel/barrel_0001.png
#    * object ID 58: barrelex/barrelex_0001.png
#    * object ID 55: bcase/bcase_0001.png
#    * object ID 62: bcase/bcase_0003.png
#    * object ID 63: bcase/bcase_0004.png
#    * object ID 86: bkslbrnt/bkslbrnt_0001.png
#    * object ID 66: bloodfnt/bloodfnt_0001.png (animated)
#    * object ID 52: book1/book1_0001.png
#    * object ID 71: book1/book1_0001.png
#    * object ID 72: book1/book1_0004.png
#    * object ID 88: book1/book1_0004.png
#    * object ID 25: book2/book2_0001.png
#    * object ID 64: book2/book2_0001.png
#    * object ID 41: book2/book2_0004.png
#    * object ID 61: book2/book2_0004.png
#    * object ID 26: burncros/burncros_0001.png (animated)
#    * object ID 91: burncros/burncros_0001.png (animated)
#    * object ID 65: candle2/candle2_0001.png (animated)
#    * object ID 87: candle2/candle2_0001.png (animated)
#    * object ID 9:  candle2/candle2_0001.png (animated)
#    * object ID 80: cauldren/cauldren_0001.png
#    * object ID 5:  chest1/chest1_0001.png
#    * object ID 68: chest1/chest1_0001.png
#    * object ID 6:  chest2/chest2_0001.png
#    * object ID 69: chest2/chest2_0001.png
#    * object ID 7:  chest3/chest3_0001.png
#    * object ID 70: chest3/chest3_0001.png
#    * object ID 97: chest3/chest3_0001.png
#    * object ID 20: cruxsk1/cruxsk1_0001.png
#    * object ID 21: cruxsk2/cruxsk2_0001.png
#    * object ID 22: cruxsk3/cruxsk3_0001.png
#    * object ID 67: decap/decap_0001.png
#    * object ID 96: decap/decap_0002.png
#    * object ID 49: flame1/flame1_0001.png
#    * object ID 79: goatshrn/goatshrn_0001.png (animated)
#    * object ID 0:  l1braz/l1braz_0001.png (animated)
#    * object ID 1:  l1doors/l1_1.pal/l1doors_0001.png
#    * object ID 2:  l1doors/l1_1.pal/l1doors_0002.png
#    * object ID 42: l2doors/l2_1.pal/l2doors_0001.png
#    * object ID 43: l2doors/l2_1.pal/l2doors_0002.png
#    * object ID 74: l3doors/l3_1.pal/l3doors_0001.png
#    * object ID 75: l3doors/l3_1.pal/l3doors_0002.png
#    * object ID 4:  lever/lever_0001.png
#    * object ID 50: lever/lever_0001.png
#    * object ID 59: lshrineg/lshrineg_0001.png
#    * object ID 95: lzstand/lzstand_0001.png
#    * object ID 84: mcirl/mcirl_0001.png
#    * object ID 85: mcirl/mcirl_0003.png
#    * object ID 81: mfountn/mfountn_0001.png (animated)
#    * object ID 51: miniwatr/miniwatr_0001.png (animated)
#    * object ID 94: mushptch/mushptch_0001.png
#    * object ID 27: nude2/nude2_0001.png (animated)
#    * object ID 73: pedistl/pedistl_0001.png
#    * object ID 76: pfountn/pfountn_0001.png (animated)
#    * object ID 23: rockstan/rockstan.png
#    * object ID 60: rshrineg/rshrineg_0001.png
#    * object ID 48: sarc/sarc_0001.png
#    * object ID 3:  skulfire/skulfire_0001.png (animated)
#    * object ID 14: skulpile/skulpile.png
#    * object ID 28: switch4/switch4_0001.png
#    * object ID 82: tfountn/tfountn_0001.png (animated)
#    * object ID 29: tnudem/tnudem_0001.png
#    * object ID 30: tnudem/tnudem_0002.png
#    * object ID 31: tnudem/tnudem_0003.png
#    * object ID 32: tnudem/tnudem_0004.png
#    * object ID 33: tnudew/tnudew_0001.png
#    * object ID 34: tnudew/tnudew_0002.png
#    * object ID 35: tnudew/tnudew_0003.png
#    * object ID 53: traphole/traphole_0001.png
#    * object ID 54: traphole/traphole_0002.png
#    * object ID 36: tsoul/tsoul_0001.png
#    * object ID 37: tsoul/tsoul_0002.png
#    * object ID 38: tsoul/tsoul_0003.png
#    * object ID 39: tsoul/tsoul_0004.png
#    * object ID 40: tsoul/tsoul_0005.png
#    * object ID 56: weapstnd/weapstnd_0001.png
#    * object ID 90: weapstnd/weapstnd_0001.png
#    * object ID 92: weapstnd/weapstnd_0001.png
#    * object ID 93: weapstnd/weapstnd_0002.png
#    * object ID 46: wtorch1/wtorch1_0001.png (animated)
#    * object ID 47: wtorch2/wtorch2_0001.png (animated)
#    * object ID 45: wtorch3/wtorch3_0001.png (animated)
#    * object ID 44: wtorch4/wtorch4_0001.png (animated)

# Merge angel.
montage \
	_dump_/objects/angel/angel_0001.png \
	_dump_/objects/angel/angel_0002.png \
	-gravity south -geometry x128+0+0 \
	-tile 2x \
	-background none \
	_dump_/objects/angel/angel_pre_crop.png

# Crop angel to 128x128.
convert _dump_/objects/angel/angel_pre_crop.png -crop 128x128+0+0 +repage _dump_/objects/angel.png
rm -f _dump_/objects/angel/angel_pre_crop.png

echo 'Creating "_assets_/_tilesets_/objects.png'
montage \
	_dump_/objects/altboy/altboy.png \
	_dump_/objects/angel.png \
	_dump_/objects/armstand/armstand_0001.png \
	_dump_/objects/armstand/armstand_0001.png \
	_dump_/objects/armstand/armstand_0002.png \
	_dump_/objects/banner/banner_0001.png \
	_dump_/objects/banner/banner_0002.png \
	_dump_/objects/banner/banner_0003.png \
	_dump_/objects/barrel/barrel_0001.png \
	_dump_/objects/barrelex/barrelex_0001.png \
	_dump_/objects/bcase/bcase_0001.png \
	_dump_/objects/bcase/bcase_0003.png \
	_dump_/objects/bcase/bcase_0004.png \
	_dump_/objects/bkslbrnt/bkslbrnt_0001.png \
	_dump_/objects/bloodfnt/bloodfnt_0001.png \
	_dump_/objects/book1/book1_0001.png \
	_dump_/objects/book1/book1_0001.png \
	_dump_/objects/book1/book1_0004.png \
	_dump_/objects/book1/book1_0004.png \
	_dump_/objects/book2/book2_0001.png \
	_dump_/objects/book2/book2_0001.png \
	_dump_/objects/book2/book2_0004.png \
	_dump_/objects/book2/book2_0004.png \
	_dump_/objects/burncros/burncros_0001.png \
	_dump_/objects/burncros/burncros_0001.png \
	_dump_/objects/candle2/candle2_0001.png \
	_dump_/objects/candle2/candle2_0001.png \
	_dump_/objects/candle2/candle2_0001.png \
	_dump_/objects/cauldren/cauldren_0001.png \
	_dump_/objects/chest1/chest1_0001.png \
	_dump_/objects/chest1/chest1_0001.png \
	_dump_/objects/chest2/chest2_0001.png \
	_dump_/objects/chest2/chest2_0001.png \
	_dump_/objects/chest3/chest3_0001.png \
	_dump_/objects/chest3/chest3_0001.png \
	_dump_/objects/chest3/chest3_0001.png \
	_dump_/objects/cruxsk1/cruxsk1_0001.png \
	_dump_/objects/cruxsk2/cruxsk2_0001.png \
	_dump_/objects/cruxsk3/cruxsk3_0001.png \
	_dump_/objects/decap/decap_0001.png \
	_dump_/objects/decap/decap_0002.png \
	_dump_/objects/flame1/flame1_0001.png \
	_dump_/objects/goatshrn/goatshrn_0001.png \
	_dump_/objects/l1braz/l1braz_0001.png \
	_dump_/objects/l1doors/l1_1.pal/l1doors_0001.png \
	_dump_/objects/l1doors/l1_1.pal/l1doors_0002.png \
	_dump_/objects/l2doors/l2_1.pal/l2doors_0001.png \
	_dump_/objects/l2doors/l2_1.pal/l2doors_0002.png \
	_dump_/objects/l3doors/l3_1.pal/l3doors_0001.png \
	_dump_/objects/l3doors/l3_1.pal/l3doors_0002.png \
	_dump_/objects/lever/lever_0001.png \
	_dump_/objects/lever/lever_0001.png \
	_dump_/objects/lshrineg/lshrineg_0001.png \
	_dump_/objects/lzstand/lzstand_0001.png \
	_dump_/objects/mcirl/mcirl_0001.png \
	_dump_/objects/mcirl/mcirl_0003.png \
	_dump_/objects/mfountn/mfountn_0001.png \
	_dump_/objects/miniwatr/miniwatr_0001.png \
	_dump_/objects/mushptch/mushptch_0001.png \
	_dump_/objects/nude2/nude2_0001.png \
	_dump_/objects/pedistl/pedistl_0001.png \
	_dump_/objects/pfountn/pfountn_0001.png \
	_dump_/objects/rockstan/rockstan.png \
	_dump_/objects/rshrineg/rshrineg_0001.png \
	_dump_/objects/sarc/sarc_0001.png \
	_dump_/objects/skulfire/skulfire_0001.png \
	_dump_/objects/skulpile/skulpile.png \
	_dump_/objects/switch4/switch4_0001.png \
	_dump_/objects/tfountn/tfountn_0001.png \
	_dump_/objects/tnudem/tnudem_0001.png \
	_dump_/objects/tnudem/tnudem_0002.png \
	_dump_/objects/tnudem/tnudem_0003.png \
	_dump_/objects/tnudem/tnudem_0004.png \
	_dump_/objects/tnudew/tnudew_0001.png \
	_dump_/objects/tnudew/tnudew_0002.png \
	_dump_/objects/tnudew/tnudew_0003.png \
	_dump_/objects/traphole/traphole_0001.png \
	_dump_/objects/traphole/traphole_0002.png \
	_dump_/objects/tsoul/tsoul_0001.png \
	_dump_/objects/tsoul/tsoul_0002.png \
	_dump_/objects/tsoul/tsoul_0003.png \
	_dump_/objects/tsoul/tsoul_0004.png \
	_dump_/objects/tsoul/tsoul_0005.png \
	_dump_/objects/weapstnd/weapstnd_0001.png \
	_dump_/objects/weapstnd/weapstnd_0001.png \
	_dump_/objects/weapstnd/weapstnd_0001.png \
	_dump_/objects/weapstnd/weapstnd_0002.png \
	_dump_/objects/wtorch1/wtorch1_0001.png \
	_dump_/objects/wtorch2/wtorch2_0001.png \
	_dump_/objects/wtorch3/wtorch3_0001.png \
	_dump_/objects/wtorch4/wtorch4_0001.png \
	-gravity south -geometry 128x128+0+0 \
	-tile 13x \
	-background none \
	_assets_/_tilesets_/objects.png
