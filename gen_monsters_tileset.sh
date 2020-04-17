#!/bin/bash

# Monster IDs (in order of occurrence in _assets_/_tilesets_/monsters.png):
# Notice that monster ID 56-59 are missing since the monsters/worm/ graphic is
# not present.
#
#    // monsters/zombie
#    * monster ID 0: Zombie
#    * monster ID 1: Ghoul
#    * monster ID 2: Rotting Carcass
#    * monster ID 3: Black Death
#    // monsters/falspear
#    * monster ID 4: Fallen One (spear)
#    * monster ID 5: Carver (spear)
#    * monster ID 6: Devil Kin (spear)
#    * monster ID 7: Dark One (spear)
#    // monsters/skelaxe
#    * monster ID 8:  Skeleton (axe)
#    * monster ID 9:  Corpse Axe (axe)
#    * monster ID 10: Burning Dead (axe)
#    * monster ID 11: Horror (axe)
#    // monsters/falsword
#    * monster ID 12: Fallen One (sword)
#    * monster ID 13: Carver (sword)
#    * monster ID 14: Devil Kin (sword)
#    * monster ID 15: Dark One (sword)
#    // monsters/scav
#    * monster ID 16: Scavenger
#    * monster ID 17: Plague Eater
#    * monster ID 18: Shadow Beast
#    * monster ID 19: Bone Gasher
#    // monsters/skelbow
#    * monster ID 20: Skeleton (bow)
#    * monster ID 21: Corpse Bow (bow)
#    * monster ID 22: Burning Dead (bow)
#    * monster ID 23: Horror (bow)
#    // monsters/skelsd
#    * monster ID 24: Skeleton Captain
#    * monster ID 25: Corpse Captain
#    * monster ID 26: Burning Dead Captain
#    * monster ID 27: Horror Captain
#    // monsters/tsneak
#    * monster ID 28: Invisible Lord
#    // monsters/sneak
#    * monster ID 29: Hidden
#    * monster ID 30: Stalker
#    * monster ID 31: Unseen
#    * monster ID 32: Illusion Weaver
#    // monsters/goatlord
#    * monster ID 33: Lord Sayter
#    // monsters/goatmace
#    * monster ID 34: Flesh Clan (mace)
#    * monster ID 35: Stone Clan (mace)
#    * monster ID 36: Fire Clan (mace)
#    * monster ID 37: Night Clan (mace)
#    // monsters/bat
#    * monster ID 38: Fiend
#    * monster ID 39: Blink
#    * monster ID 40: Gloom
#    * monster ID 41: Familiar
#    // monsters/goatbow
#    * monster ID 42: Flesh Clan (bow)
#    * monster ID 43: Stone Clan (bow)
#    * monster ID 44: Fire Clan (bow)
#    * monster ID 45: Night Clan (bow)
#    // monsters/acid
#    * monster ID 46: Acid Beast
#    * monster ID 47: Poison Spitter
#    * monster ID 48: Pit Beast
#    * monster ID 49: Lava Maw
#    // monsters/sking
#    * monster ID 50: Skeleton King
#    // monsters/fatc
#    * monster ID 51: The Butcher
#    // monsters/fat
#    * monster ID 52: Overlord
#    * monster ID 53: Mud Man
#    * monster ID 54: Toad Demon
#    * monster ID 55: Flayed One
#    // monsters/magma
#    * monster ID 60: Magma Demon
#    * monster ID 61: Blood Stone
#    * monster ID 62: Hell Stone
#    * monster ID 63: Lava Lord
#    // monsters/rhino
#    * monster ID 64: Horned Demon
#    * monster ID 65: Mud Runner
#    * monster ID 66: Frost Charger
#    * monster ID 67: Obsidian Lord
#    // monsters/demskel
#    * monster ID 68: Bone Demon
#    // monsters/thin
#    * monster ID 69: Red Death
#    * monster ID 70: Litch Demon
#    * monster ID 71: Undead Balrog
#    // monsters/fireman
#    * monster ID 72: Incinerator
#    * monster ID 73: Flame Lord
#    * monster ID 74: Doom Fire
#    * monster ID 75: Hell Burner
#    // monsters/thin
#    * monster ID 76: Red Storm
#    * monster ID 77: Storm Rider
#    * monster ID 78: Storm Lord
#    * monster ID 79: Maelstorm
#    // monsters/bigfall
#    * monster ID 80: Devil Kin Brute
#    // monsters/gargoyle
#    * monster ID 81: Winged-Demon
#    * monster ID 82: Gargoyle
#    * monster ID 83: Blood Claw
#    * monster ID 84: Death Wing
#    // monsters/mega
#    * monster ID 85: Slayer
#    * monster ID 86: Guardian
#    * monster ID 87: Vortex Lord
#    * monster ID 88: Balrog
#    // monsters/snake
#    * monster ID 89: Cave Viper
#    * monster ID 90: Fire Drake
#    * monster ID 91: Gold Viper
#    * monster ID 92: Azure Drake
#    // monsters/black
#    * monster ID 93: Black Knight
#    * monster ID 94: Doom Guard
#    * monster ID 95: Steel Lord
#    * monster ID 96: Blood Knight
#    // monsters/unrav
#    * monster ID 97:  Unraveler
#    * monster ID 98:  Hollow One
#    * monster ID 99:  Pain Master
#    * monster ID 100: Reality Weaver
#    // monsters/succ
#    * monster ID 101: Succubus
#    * monster ID 102: Snow Witch
#    * monster ID 103: Hell Spawn
#    * monster ID 104: Soul Burner
#    // monsters/mage
#    * monster ID 105: Counselor
#    * monster ID 106: Magistrate
#    * monster ID 107: Cabalist
#    * monster ID 108: Advocate
#    // monsters/golem
#    * monster ID 109: Golem
#    // monsters/diablo
#    * monster ID 110: The Dark Lord
#    // monsters/darkmage
#    * monster ID 111: The Arch-Litch Malignus

# NOTE: Graphics missing (monsters/worm/*.cl2):
#
#   // monsters/worm
#   * monster ID 56: Wyrm
#   * monster ID 57: Cave Slug
#   * monster ID 58: Devil Wyrm
#   * monster ID 59: Devourer

echo 'Creating "_assets_/_tilesets_/monsters.png'
montage \
	_dump_/monsters/zombie/zombiea/zombiea_0/zombiea_0001.png \
	_dump_/monsters/zombie/zombiea/bluered.trn/zombiea_0/zombiea_0001.png \
	_dump_/monsters/zombie/zombiea/grey.trn/zombiea_0/zombiea_0001.png \
	_dump_/monsters/zombie/zombiea/yellow.trn/zombiea_0/zombiea_0001.png \
	_dump_/monsters/falspear/phalla/fallent.trn/phalla_0/phalla_0001.png \
	_dump_/monsters/falspear/phalla/dark.trn/phalla_0/phalla_0001.png \
	_dump_/monsters/falspear/phalla/phalla_0/phalla_0001.png \
	_dump_/monsters/falspear/phalla/blue.trn/phalla_0/phalla_0001.png \
	_dump_/monsters/skelaxe/sklaxa/white.trn/sklaxa_0/sklaxa_0001.png \
	_dump_/monsters/skelaxe/sklaxa/skelt.trn/sklaxa_0/sklaxa_0001.png \
	_dump_/monsters/skelaxe/sklaxa/sklaxa_0/sklaxa_0001.png \
	_dump_/monsters/skelaxe/sklaxa/black.trn/sklaxa_0/sklaxa_0001.png \
	_dump_/monsters/falsword/falla/fallent.trn/falla_0/falla_0001.png \
	_dump_/monsters/falsword/falla/dark.trn/falla_0/falla_0001.png \
	_dump_/monsters/falsword/falla/falla_0/falla_0001.png \
	_dump_/monsters/falsword/falla/blue.trn/falla_0/falla_0001.png \
	_dump_/monsters/scav/scava/scava_0/scava_0001.png \
	_dump_/monsters/scav/scava/scavbr.trn/scava_0/scava_0001.png \
	_dump_/monsters/scav/scava/scavbe.trn/scava_0/scava_0001.png \
	_dump_/monsters/scav/scava/scavw.trn/scava_0/scava_0001.png \
	_dump_/monsters/skelbow/sklbwa/white.trn/sklbwa_0/sklbwa_0001.png \
	_dump_/monsters/skelbow/sklbwa/skelt.trn/sklbwa_0/sklbwa_0001.png \
	_dump_/monsters/skelbow/sklbwa/sklbwa_0/sklbwa_0001.png \
	_dump_/monsters/skelbow/sklbwa/black.trn/sklbwa_0/sklbwa_0001.png \
	_dump_/monsters/skelsd/sklsra/white.trn/sklsra_0/sklsra_0001.png \
	_dump_/monsters/skelsd/sklsra/skelt.trn/sklsra_0/sklsra_0001.png \
	_dump_/monsters/skelsd/sklsra/sklsra_0/sklsra_0001.png \
	_dump_/monsters/skelsd/sklsra/black.trn/sklsra_0/sklsra_0001.png \
	_dump_/monsters/tsneak/tsneaka/tsneaka_0/tsneaka_0001.png \
	_dump_/monsters/sneak/sneaka/sneaka_0/sneaka_0001.png \
	_dump_/monsters/sneak/sneaka/sneakv2.trn/sneaka_0/sneaka_0001.png \
	_dump_/monsters/sneak/sneaka/sneakv3.trn/sneaka_0/sneaka_0001.png \
	_dump_/monsters/sneak/sneaka/sneakv1.trn/sneaka_0/sneaka_0001.png \
	_dump_/monsters/goatlord/goatla/goatla_0/goatla_0001.png \
	_dump_/monsters/goatmace/goata/goata_0/goata_0001.png \
	_dump_/monsters/goatmace/goata/beige.trn/goata_0/goata_0001.png \
	_dump_/monsters/goatmace/goata/red.trn/goata_0/goata_0001.png \
	_dump_/monsters/goatmace/goata/gray.trn/goata_0/goata_0001.png \
	_dump_/monsters/bat/bata/red.trn/bata_0/bata_0001.png \
	_dump_/monsters/bat/bata/bata_0/bata_0001.png \
	_dump_/monsters/bat/bata/grey.trn/bata_0/bata_0001.png \
	_dump_/monsters/bat/bata/orange.trn/bata_0/bata_0001.png \
	_dump_/monsters/goatbow/goatba/goatba_0/goatba_0001.png \
	_dump_/monsters/goatbow/goatba/beige.trn/goatba_0/goatba_0001.png \
	_dump_/monsters/goatbow/goatba/red.trn/goatba_0/goatba_0001.png \
	_dump_/monsters/goatbow/goatba/gray.trn/goatba_0/goatba_0001.png \
	_dump_/monsters/acid/acida/acida_0/acida_0001.png \
	_dump_/monsters/acid/acida/acidblk.trn/acida_0/acida_0001.png \
	_dump_/monsters/acid/acida/acidb.trn/acida_0/acida_0001.png \
	_dump_/monsters/acid/acida/acidr.trn/acida_0/acida_0001.png \
	_dump_/monsters/sking/skinga/white.trn/skinga_0/skinga_0001.png \
	_dump_/monsters/fatc/fatca/fatca_0/fatca_0001.png \
	_dump_/monsters/fat/fata/fata_0/fata_0001.png \
	_dump_/monsters/fat/fata/blue.trn/fata_0/fata_0001.png \
	_dump_/monsters/fat/fata/fatb.trn/fata_0/fata_0001.png \
	_dump_/monsters/fat/fata/fatf.trn/fata_0/fata_0001.png \
	_dump_/monsters/magma/magmaa/magmaa_0/magmaa_0001.png \
	_dump_/monsters/magma/magmaa/yellow.trn/magmaa_0/magmaa_0001.png \
	_dump_/monsters/magma/magmaa/blue.trn/magmaa_0/magmaa_0001.png \
	_dump_/monsters/magma/magmaa/wierd.trn/magmaa_0/magmaa_0001.png \
	_dump_/monsters/rhino/rhinoa/rhinoa_0/rhinoa_0001.png \
	_dump_/monsters/rhino/rhinoa/orange.trn/rhinoa_0/rhinoa_0001.png \
	_dump_/monsters/rhino/rhinoa/blue.trn/rhinoa_0/rhinoa_0001.png \
	_dump_/monsters/rhino/rhinoa/rhinob.trn/rhinoa_0/rhinoa_0001.png \
	_dump_/monsters/demskel/demskla/thinv3.trn/demskla_0/demskla_0001.png \
	_dump_/monsters/thin/thina/thinv3.trn/thina_0/thina_0001.png \
	_dump_/monsters/thin/thina/thinv3.trn/thina_0/thina_0001.png \
	_dump_/monsters/thin/thina/thinv3.trn/thina_0/thina_0001.png \
	_dump_/monsters/fireman/firema/firema_0/firema_0001.png \
	_dump_/monsters/fireman/firema/firema_0/firema_0001.png \
	_dump_/monsters/fireman/firema/firema_0/firema_0001.png \
	_dump_/monsters/fireman/firema/firema_0/firema_0001.png \
	_dump_/monsters/thin/thina/thinv3.trn/thina_0/thina_0001.png \
	_dump_/monsters/thin/thina/thina_0/thina_0001.png \
	_dump_/monsters/thin/thina/thinv2.trn/thina_0/thina_0001.png \
	_dump_/monsters/thin/thina/thinv1.trn/thina_0/thina_0001.png \
	_dump_/monsters/bigfall/fallga/fallga_0/fallga_0001.png \
	_dump_/monsters/gargoyle/gargoa/gargoa_0/gargoa_0001.png \
	_dump_/monsters/gargoyle/gargoa/gare.trn/gargoa_0/gargoa_0001.png \
	_dump_/monsters/gargoyle/gargoa/gargbr.trn/gargoa_0/gargoa_0001.png \
	_dump_/monsters/gargoyle/gargoa/gargb.trn/gargoa_0/gargoa_0001.png \
	_dump_/monsters/mega/megaa/megaa_0/megaa_0001.png \
	_dump_/monsters/mega/megaa/guard.trn/megaa_0/megaa_0001.png \
	_dump_/monsters/mega/megaa/vtexl.trn/megaa_0/megaa_0001.png \
	_dump_/monsters/mega/megaa/balr.trn/megaa_0/megaa_0001.png \
	_dump_/monsters/snake/snakea/snakea_0/snakea_0001.png \
	_dump_/monsters/snake/snakea/snakr.trn/snakea_0/snakea_0001.png \
	_dump_/monsters/snake/snakea/snakg.trn/snakea_0/snakea_0001.png \
	_dump_/monsters/snake/snakea/snakb.trn/snakea_0/snakea_0001.png \
	_dump_/monsters/black/blacka/blacka_0/blacka_0001.png \
	_dump_/monsters/black/blacka/blkkntrt.trn/blacka_0/blacka_0001.png \
	_dump_/monsters/black/blacka/blkkntbt.trn/blacka_0/blacka_0001.png \
	_dump_/monsters/black/blacka/blkkntbe.trn/blacka_0/blacka_0001.png \
	_dump_/monsters/unrav/unrava/unrava_0/unrava_0001.png \
	_dump_/monsters/unrav/unrava/unrava_0/unrava_0001.png \
	_dump_/monsters/unrav/unrava/unrava_0/unrava_0001.png \
	_dump_/monsters/unrav/unrava/unrava_0/unrava_0001.png \
	_dump_/monsters/succ/scbsa/scbsa_0/scbsa_0001.png \
	_dump_/monsters/succ/scbsa/succb.trn/scbsa_0/scbsa_0001.png \
	_dump_/monsters/succ/scbsa/succrw.trn/scbsa_0/scbsa_0001.png \
	_dump_/monsters/succ/scbsa/succbw.trn/scbsa_0/scbsa_0001.png \
	_dump_/monsters/mage/magea/magea_0/magea_0001.png \
	_dump_/monsters/mage/magea/cnselg.trn/magea_0/magea_0001.png \
	_dump_/monsters/mage/magea/cnselgd.trn/magea_0/magea_0001.png \
	_dump_/monsters/mage/magea/cnselbk.trn/magea_0/magea_0001.png \
	_dump_/monsters/golem/golema/golema_0/golema_0001.png \
	_dump_/monsters/diablo/diabloa/diabloa_0/diabloa_0001.png \
	_dump_/monsters/darkmage/dmagea/dmagea_0/dmagea_0001.png \
	-gravity south -geometry 160x160+0+0 \
	-tile 12x \
	-background none \
	_assets_/_tilesets_/monsters.png
