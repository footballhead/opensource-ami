#!/bin/bash

# === [ Tilesets ] =============================================================

mkdir -p _assets_/_tilesets_

# * tristram.png  -> (same)
# * cathedral.png -> (same)
# * catacombs.png -> (same)
# * caves.png     -> (same)
# * hell.png      -> (same)

cp _dump_/_tilesets_/tristram.png  _assets_/_tilesets_/
cp _dump_/_tilesets_/cathedral.png _assets_/_tilesets_/
cp _dump_/_tilesets_/catacombs.png _assets_/_tilesets_/
cp _dump_/_tilesets_/caves.png     _assets_/_tilesets_/
cp _dump_/_tilesets_/hell.png      _assets_/_tilesets_/

# === [ TMX maps ] =============================================================

# --- [ Tristram ] -------------------------------------------------------------

mkdir -p _assets_/_maps_/tristram

# * tristram.tmx -> (same)

cp _dump_/_dungeons_/tristram/tristram.tmx _assets_/_maps_/tristram/

# --- [ Cathedral ] ------------------------------------------------------------

mkdir -p _assets_/_maps_/cathedral

# * banner1.tmx  -> ogdens_sign.tmx
# * banner2.tmx  -> ogdens_sign_pre.tmx
# * hero1.tmx    -> hero.tmx                            (unused?)
# * hero2.tmx    -> hero_pre.tmx                        (unused?)
# * lv1mazea.tmx -> maze.tmx
# * lv1mazeb.tmx -> maze_pre.tmx
# * rnd1.tmx     -> (same)                              (unused?)
# * rnd2.tmx     -> (same)                              (unused?)
# * rnd3.tmx     -> (same)                              (unused?)
# * rnd4.tmx     -> (same)                              (unused?)
# * rnd5.tmx     -> the_butchers_chamber_arches.tmx     (unused?)
# * rnd6.tmx     -> the_butchers_chamber.tmx
# * sklkng.tmx   -> skeleton_kings_lair_post.tmx        (unused?)
# * sklkng1.tmx  -> skeleton_kings_lair.tmx
# * sklkng2.tmx  -> skeleton_kings_lair_pre.tmx
# * sklkngdr.tmx -> skeleton_kings_lair_door.tmx        (unused?)
# * skngdc.tmx   -> skeleton_kings_lair_down_closed.tmx (unused?)
# * skngdo.tmx   -> skeleton_kings_lair_down_open.tmx
# * vile1.tmx    -> archbishop_lazarus_lair.tmx
# * vile2.tmx    -> archbishop_lazarus_lair_pre.tmx

cp _dump_/_dungeons_/cathedral/banner1.tmx  _assets_/_maps_/cathedral/ogdens_sign.tmx
cp _dump_/_dungeons_/cathedral/banner2.tmx  _assets_/_maps_/cathedral/ogdens_sign_pre.tmx
cp _dump_/_dungeons_/cathedral/hero1.tmx    _assets_/_maps_/cathedral/hero.tmx
cp _dump_/_dungeons_/cathedral/hero2.tmx    _assets_/_maps_/cathedral/hero_pre.tmx
cp _dump_/_dungeons_/cathedral/lv1mazea.tmx _assets_/_maps_/cathedral/maze.tmx
cp _dump_/_dungeons_/cathedral/lv1mazeb.tmx _assets_/_maps_/cathedral/maze_pre.tmx
cp _dump_/_dungeons_/cathedral/rnd1.tmx     _assets_/_maps_/cathedral/
cp _dump_/_dungeons_/cathedral/rnd2.tmx     _assets_/_maps_/cathedral/
cp _dump_/_dungeons_/cathedral/rnd3.tmx     _assets_/_maps_/cathedral/
cp _dump_/_dungeons_/cathedral/rnd4.tmx     _assets_/_maps_/cathedral/
cp _dump_/_dungeons_/cathedral/rnd5.tmx     _assets_/_maps_/cathedral/the_butchers_chamber_arches.tmx
cp _dump_/_dungeons_/cathedral/rnd6.tmx     _assets_/_maps_/cathedral/the_butchers_chamber.tmx
cp _dump_/_dungeons_/cathedral/sklkng.tmx   _assets_/_maps_/cathedral/skeleton_kings_lair_post.tmx
cp _dump_/_dungeons_/cathedral/sklkng1.tmx  _assets_/_maps_/cathedral/skeleton_kings_lair.tmx
cp _dump_/_dungeons_/cathedral/sklkng2.tmx  _assets_/_maps_/cathedral/skeleton_kings_lair_pre.tmx
cp _dump_/_dungeons_/cathedral/sklkngdr.tmx _assets_/_maps_/cathedral/skeleton_kings_lair_door.tmx
cp _dump_/_dungeons_/cathedral/skngdc.tmx   _assets_/_maps_/cathedral/skeleton_kings_lair_down_closed.tmx
cp _dump_/_dungeons_/cathedral/skngdo.tmx   _assets_/_maps_/cathedral/skeleton_kings_lair_down_open.tmx
cp _dump_/_dungeons_/cathedral/vile1.tmx    _assets_/_maps_/cathedral/archbishop_lazarus_lair.tmx
cp _dump_/_dungeons_/cathedral/vile2.tmx    _assets_/_maps_/cathedral/archbishop_lazarus_lair_pre.tmx

# --- [ Catacombs ] ------------------------------------------------------------

mkdir -p _assets_/_maps_/catacombs

# * blind1.tmx   -> halls_of_the_blind.tmx
# * blind2.tmx   -> halls_of_the_blind_pre.tmx
# * blood1.tmx   -> valor.tmx
# * blood2.tmx   -> valor_pre.tmx
# * blood3.tmx   -> valor_mid.tmx
# * bonecha1.tmx -> bone_chamber_pre.tmx
# * bonecha2.tmx -> bone_chamber.tmx
# * bonestr1.tmx -> bone_chamber_stairs_pre.tmx
# * bonestr2.tmx -> bone_chamber_stairs.tmx

cp _dump_/_dungeons_/catacombs/blind1.tmx   _assets_/_maps_/catacombs/halls_of_the_blind.tmx
cp _dump_/_dungeons_/catacombs/blind2.tmx   _assets_/_maps_/catacombs/halls_of_the_blind_pre.tmx
cp _dump_/_dungeons_/catacombs/blood1.tmx   _assets_/_maps_/catacombs/valor.tmx
cp _dump_/_dungeons_/catacombs/blood2.tmx   _assets_/_maps_/catacombs/valor_pre.tmx
cp _dump_/_dungeons_/catacombs/blood3.tmx   _assets_/_maps_/catacombs/valor_mid.tmx
cp _dump_/_dungeons_/catacombs/bonecha1.tmx _assets_/_maps_/catacombs/bone_chamber_pre.tmx
cp _dump_/_dungeons_/catacombs/bonecha2.tmx _assets_/_maps_/catacombs/bone_chamber.tmx
cp _dump_/_dungeons_/catacombs/bonestr1.tmx _assets_/_maps_/catacombs/bone_chamber_stairs_pre.tmx
cp _dump_/_dungeons_/catacombs/bonestr2.tmx _assets_/_maps_/catacombs/bone_chamber_stairs.tmx

# --- [ Caves ] ----------------------------------------------------------------

mkdir -p _assets_/_maps_/caves

# * anvil.tmx    -> anvil_of_fury.tmx
# * foulwatr.tmx -> poisoned_water_supply.tmx
# * lair.tmx     -> lair.tmx                  (unused?)

cp _dump_/_dungeons_/caves/anvil.tmx    _assets_/_maps_/caves/anvil_of_fury.tmx
cp _dump_/_dungeons_/caves/foulwatr.tmx _assets_/_maps_/caves/poisoned_water_supply.tmx
cp _dump_/_dungeons_/caves/lair.tmx     _assets_/_maps_/caves/lair.tmx

# --- [ Hell ] -----------------------------------------------------------------

mkdir -p _assets_/_maps_/hell

# diab1.tmx    -> diablo_first.tmx
# diab2a.tmx   -> diablo_second_pre.tmx
# diab2b.tmx   -> diablo_second.tmx
# diab3a.tmx   -> diablo_third_pre.tmx
# diab3b.tmx   -> diablo_third.tmx
# diab4a.tmx   -> diablo_final_pre.tmx
# diab4b.tmx   -> diablo_final.tmx
# vile1.tmx    -> archbishop_lazarus_lair_multiplayer.tmx
# vile2.tmx    -> archbishop_lazarus_lair_multiplayer_pre.tmx (unused?)
# vile3.tmx    -> archbishop_lazarus_lair_switch.tmx          (unused?)
# warlord.tmx  -> warlord_of_blood.tmx
# warlord2.tmx -> warlord_of_blood_pre.tmx

cp _dump_/_dungeons_/hell/diab1.tmx    _assets_/_maps_/hell/diablo_first.tmx
cp _dump_/_dungeons_/hell/diab2a.tmx   _assets_/_maps_/hell/diablo_second_pre.tmx
cp _dump_/_dungeons_/hell/diab2b.tmx   _assets_/_maps_/hell/diablo_second.tmx
cp _dump_/_dungeons_/hell/diab3a.tmx   _assets_/_maps_/hell/diablo_third_pre.tmx
cp _dump_/_dungeons_/hell/diab3b.tmx   _assets_/_maps_/hell/diablo_third.tmx
cp _dump_/_dungeons_/hell/diab4a.tmx   _assets_/_maps_/hell/diablo_final_pre.tmx
cp _dump_/_dungeons_/hell/diab4b.tmx   _assets_/_maps_/hell/diablo_final.tmx
cp _dump_/_dungeons_/hell/vile1.tmx    _assets_/_maps_/hell/archbishop_lazarus_lair_multiplayer.tmx
cp _dump_/_dungeons_/hell/vile2.tmx    _assets_/_maps_/hell/archbishop_lazarus_lair_multiplayer_pre.tmx
cp _dump_/_dungeons_/hell/vile3.tmx    _assets_/_maps_/hell/archbishop_lazarus_lair_switch.tmx
cp _dump_/_dungeons_/hell/warlord.tmx  _assets_/_maps_/hell/warlord_of_blood.tmx
cp _dump_/_dungeons_/hell/warlord2.tmx _assets_/_maps_/hell/warlord_of_blood_pre.tmx
