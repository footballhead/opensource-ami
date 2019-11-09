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

# === [ Minisets ] =============================================================

# --- [ Cathedral ] ------------------------------------------------------------

mkdir -p _assets_/_minisets_/cathedral

# Diablo 1.09b.
#
# * stairsup.tmx   -> stairs_up1.tmx
# * l5stairsup.tmx -> stairs_up2.tmx
# * stairsdown.tmx -> stairs_down.tmx
# * lamps.tmx      -> candlestick.tmx
# * pwaterin.tmx   -> stairs_down_poison.tmx

cp _dump_/_minisets_/cathedral/stairsup.tmx   _assets_/_minisets_/cathedral/stairs_up1.tmx
cp _dump_/_minisets_/cathedral/l5stairsup.tmx _assets_/_minisets_/cathedral/stairs_up2.tmx
cp _dump_/_minisets_/cathedral/stairsdown.tmx _assets_/_minisets_/cathedral/stairs_down.tmx
cp _dump_/_minisets_/cathedral/lamps.tmx      _assets_/_minisets_/cathedral/candlestick.tmx
cp _dump_/_minisets_/cathedral/pwaterin.tmx   _assets_/_minisets_/cathedral/stairs_down_poison.tmx

# Diablo pre-release-demo [dalpha] (1996-08-17).
#
# * alpha4_sarc.tmx      -> sarcophagus.tmx
# * alpha4_rcorridor.tmx -> right_corridor.tmx
# * alpha4_lcorridor.tmx -> left_corridor.tmx

cp _dump_/_minisets_/cathedral/alpha4_sarc.tmx      _assets_/_minisets_/cathedral/sarcophagus.tmx
cp _dump_/_minisets_/cathedral/alpha4_rcorridor.tmx _assets_/_minisets_/cathedral/right_corridor.tmx
cp _dump_/_minisets_/cathedral/alpha4_lcorridor.tmx _assets_/_minisets_/cathedral/left_corridor.tmx

# --- [ Catacombs ] ------------------------------------------------------------

mkdir -p _assets_/_minisets_/catacombs

# * varch1.tmx     -> vert_arch1.tmx
# * varch2.tmx     -> vert_arch2.tmx
# * varch3.tmx     -> vert_arch3.tmx
# * varch4.tmx     -> vert_arch4.tmx
# * varch5.tmx     -> vert_arch5.tmx
# * varch6.tmx     -> vert_arch6.tmx
# * varch7.tmx     -> vert_arch7.tmx
# * varch8.tmx     -> vert_arch8.tmx
# * varch9.tmx     -> vert_arch9.tmx
# * varch10.tmx    -> vert_arch10.tmx
# * varch11.tmx    -> vert_arch11.tmx
# * varch12.tmx    -> vert_arch12.tmx
# * varch13.tmx    -> vert_arch13.tmx
# * varch14.tmx    -> vert_arch14.tmx
# * varch15.tmx    -> vert_arch15.tmx
# * varch16.tmx    -> vert_arch16.tmx
# * varch17.tmx    -> vert_arch17.tmx
# * varch18.tmx    -> vert_arch18.tmx
# * varch19.tmx    -> vert_arch19.tmx
# * varch20.tmx    -> vert_arch20.tmx
# * varch21.tmx    -> vert_arch21.tmx
# * varch22.tmx    -> vert_arch22.tmx
# * varch23.tmx    -> vert_arch23.tmx
# * varch24.tmx    -> vert_arch24.tmx
# * varch25.tmx    -> vert_arch25.tmx
# * varch26.tmx    -> vert_arch26.tmx
# * varch27.tmx    -> vert_arch27.tmx
# * varch28.tmx    -> vert_arch28.tmx
# * varch29.tmx    -> vert_arch29.tmx
# * varch30.tmx    -> vert_arch30.tmx
# * varch31.tmx    -> vert_arch31.tmx
# * varch32.tmx    -> vert_arch32.tmx
# * varch33.tmx    -> vert_arch33.tmx
# * varch34.tmx    -> vert_arch34.tmx
# * varch35.tmx    -> vert_arch35.tmx
# * varch36.tmx    -> vert_arch36.tmx
# * varch37.tmx    -> vert_arch37.tmx
# * varch38.tmx    -> vert_arch38.tmx
# * varch39.tmx    -> vert_arch39.tmx
# * varch40.tmx    -> vert_arch40.tmx
# * harch1.tmx     -> horiz_arch1.tmx
# * harch2.tmx     -> horiz_arch2.tmx
# * harch3.tmx     -> horiz_arch3.tmx
# * harch4.tmx     -> horiz_arch4.tmx
# * harch5.tmx     -> horiz_arch5.tmx
# * harch6.tmx     -> horiz_arch6.tmx
# * harch7.tmx     -> horiz_arch7.tmx
# * harch8.tmx     -> horiz_arch8.tmx
# * harch9.tmx     -> horiz_arch9.tmx
# * harch10.tmx    -> horiz_arch10.tmx
# * harch11.tmx    -> horiz_arch11.tmx
# * harch12.tmx    -> horiz_arch12.tmx
# * harch13.tmx    -> horiz_arch13.tmx
# * harch14.tmx    -> horiz_arch14.tmx
# * harch15.tmx    -> horiz_arch15.tmx
# * harch16.tmx    -> horiz_arch16.tmx
# * harch17.tmx    -> horiz_arch17.tmx
# * harch18.tmx    -> horiz_arch18.tmx
# * harch19.tmx    -> horiz_arch19.tmx
# * harch20.tmx    -> horiz_arch20.tmx
# * harch21.tmx    -> horiz_arch21.tmx
# * harch22.tmx    -> horiz_arch22.tmx
# * harch23.tmx    -> horiz_arch23.tmx
# * harch24.tmx    -> horiz_arch24.tmx
# * harch25.tmx    -> horiz_arch25.tmx
# * harch26.tmx    -> horiz_arch26.tmx
# * harch27.tmx    -> horiz_arch27.tmx
# * harch28.tmx    -> horiz_arch28.tmx
# * harch29.tmx    -> horiz_arch29.tmx
# * harch30.tmx    -> horiz_arch30.tmx
# * harch31.tmx    -> horiz_arch31.tmx
# * harch32.tmx    -> horiz_arch32.tmx
# * harch33.tmx    -> horiz_arch33.tmx
# * harch34.tmx    -> horiz_arch34.tmx
# * harch35.tmx    -> horiz_arch35.tmx
# * harch36.tmx    -> horiz_arch36.tmx
# * harch37.tmx    -> horiz_arch37.tmx
# * harch38.tmx    -> horiz_arch38.tmx
# * harch39.tmx    -> horiz_arch39.tmx
# * harch40.tmx    -> horiz_arch40.tmx
# * ustairs.tmx    -> stairs_up.tmx
# * dstairs.tmx    -> stairs_down.tmx
# * warpstairs.tmx -> stairs_to_town.tmx
# * crushcol.tmx   -> crushed_column.tmx
# * big1.tmx       -> (same)
# * big2.tmx       -> (same)
# * big3.tmx       -> (same)
# * big4.tmx       -> (same)
# * big5.tmx       -> (same)
# * big6.tmx       -> (same)
# * big7.tmx       -> (same)
# * big8.tmx       -> (same)
# * big9.tmx       -> (same)
# * big10.tmx      -> (same)
# * ruins1.tmx     -> (same)
# * ruins2.tmx     -> (same)
# * ruins3.tmx     -> (same)
# * ruins4.tmx     -> (same)
# * ruins5.tmx     -> (same)
# * ruins6.tmx     -> (same)
# * ruins7.tmx     -> (same)
# * pancreas1.tmx  -> (same)
# * pancreas2.tmx  -> (same)
# * ctrdoor1.tmx   -> center_door1.tmx
# * ctrdoor2.tmx   -> center_door2.tmx
# * ctrdoor3.tmx   -> center_door3.tmx
# * ctrdoor4.tmx   -> center_door4.tmx
# * ctrdoor5.tmx   -> center_door5.tmx
# * ctrdoor6.tmx   -> center_door6.tmx
# * ctrdoor7.tmx   -> center_door7.tmx
# * ctrdoor8.tmx   -> center_door8.tmx

cp _dump_/_minisets_/catacombs/varch1.tmx     _assets_/_minisets_/catacombs/vert_arch1.tmx
cp _dump_/_minisets_/catacombs/varch2.tmx     _assets_/_minisets_/catacombs/vert_arch2.tmx
cp _dump_/_minisets_/catacombs/varch3.tmx     _assets_/_minisets_/catacombs/vert_arch3.tmx
cp _dump_/_minisets_/catacombs/varch4.tmx     _assets_/_minisets_/catacombs/vert_arch4.tmx
cp _dump_/_minisets_/catacombs/varch5.tmx     _assets_/_minisets_/catacombs/vert_arch5.tmx
cp _dump_/_minisets_/catacombs/varch6.tmx     _assets_/_minisets_/catacombs/vert_arch6.tmx
cp _dump_/_minisets_/catacombs/varch7.tmx     _assets_/_minisets_/catacombs/vert_arch7.tmx
cp _dump_/_minisets_/catacombs/varch8.tmx     _assets_/_minisets_/catacombs/vert_arch8.tmx
cp _dump_/_minisets_/catacombs/varch9.tmx     _assets_/_minisets_/catacombs/vert_arch9.tmx
cp _dump_/_minisets_/catacombs/varch10.tmx    _assets_/_minisets_/catacombs/vert_arch10.tmx
cp _dump_/_minisets_/catacombs/varch11.tmx    _assets_/_minisets_/catacombs/vert_arch11.tmx
cp _dump_/_minisets_/catacombs/varch12.tmx    _assets_/_minisets_/catacombs/vert_arch12.tmx
cp _dump_/_minisets_/catacombs/varch13.tmx    _assets_/_minisets_/catacombs/vert_arch13.tmx
cp _dump_/_minisets_/catacombs/varch14.tmx    _assets_/_minisets_/catacombs/vert_arch14.tmx
cp _dump_/_minisets_/catacombs/varch15.tmx    _assets_/_minisets_/catacombs/vert_arch15.tmx
cp _dump_/_minisets_/catacombs/varch16.tmx    _assets_/_minisets_/catacombs/vert_arch16.tmx
cp _dump_/_minisets_/catacombs/varch17.tmx    _assets_/_minisets_/catacombs/vert_arch17.tmx
cp _dump_/_minisets_/catacombs/varch18.tmx    _assets_/_minisets_/catacombs/vert_arch18.tmx
cp _dump_/_minisets_/catacombs/varch19.tmx    _assets_/_minisets_/catacombs/vert_arch19.tmx
cp _dump_/_minisets_/catacombs/varch20.tmx    _assets_/_minisets_/catacombs/vert_arch20.tmx
cp _dump_/_minisets_/catacombs/varch21.tmx    _assets_/_minisets_/catacombs/vert_arch21.tmx
cp _dump_/_minisets_/catacombs/varch22.tmx    _assets_/_minisets_/catacombs/vert_arch22.tmx
cp _dump_/_minisets_/catacombs/varch23.tmx    _assets_/_minisets_/catacombs/vert_arch23.tmx
cp _dump_/_minisets_/catacombs/varch24.tmx    _assets_/_minisets_/catacombs/vert_arch24.tmx
cp _dump_/_minisets_/catacombs/varch25.tmx    _assets_/_minisets_/catacombs/vert_arch25.tmx
cp _dump_/_minisets_/catacombs/varch26.tmx    _assets_/_minisets_/catacombs/vert_arch26.tmx
cp _dump_/_minisets_/catacombs/varch27.tmx    _assets_/_minisets_/catacombs/vert_arch27.tmx
cp _dump_/_minisets_/catacombs/varch28.tmx    _assets_/_minisets_/catacombs/vert_arch28.tmx
cp _dump_/_minisets_/catacombs/varch29.tmx    _assets_/_minisets_/catacombs/vert_arch29.tmx
cp _dump_/_minisets_/catacombs/varch30.tmx    _assets_/_minisets_/catacombs/vert_arch30.tmx
cp _dump_/_minisets_/catacombs/varch31.tmx    _assets_/_minisets_/catacombs/vert_arch31.tmx
cp _dump_/_minisets_/catacombs/varch32.tmx    _assets_/_minisets_/catacombs/vert_arch32.tmx
cp _dump_/_minisets_/catacombs/varch33.tmx    _assets_/_minisets_/catacombs/vert_arch33.tmx
cp _dump_/_minisets_/catacombs/varch34.tmx    _assets_/_minisets_/catacombs/vert_arch34.tmx
cp _dump_/_minisets_/catacombs/varch35.tmx    _assets_/_minisets_/catacombs/vert_arch35.tmx
cp _dump_/_minisets_/catacombs/varch36.tmx    _assets_/_minisets_/catacombs/vert_arch36.tmx
cp _dump_/_minisets_/catacombs/varch37.tmx    _assets_/_minisets_/catacombs/vert_arch37.tmx
cp _dump_/_minisets_/catacombs/varch38.tmx    _assets_/_minisets_/catacombs/vert_arch38.tmx
cp _dump_/_minisets_/catacombs/varch39.tmx    _assets_/_minisets_/catacombs/vert_arch39.tmx
cp _dump_/_minisets_/catacombs/varch40.tmx    _assets_/_minisets_/catacombs/vert_arch40.tmx
cp _dump_/_minisets_/catacombs/harch1.tmx     _assets_/_minisets_/catacombs/horiz_arch1.tmx
cp _dump_/_minisets_/catacombs/harch2.tmx     _assets_/_minisets_/catacombs/horiz_arch2.tmx
cp _dump_/_minisets_/catacombs/harch3.tmx     _assets_/_minisets_/catacombs/horiz_arch3.tmx
cp _dump_/_minisets_/catacombs/harch4.tmx     _assets_/_minisets_/catacombs/horiz_arch4.tmx
cp _dump_/_minisets_/catacombs/harch5.tmx     _assets_/_minisets_/catacombs/horiz_arch5.tmx
cp _dump_/_minisets_/catacombs/harch6.tmx     _assets_/_minisets_/catacombs/horiz_arch6.tmx
cp _dump_/_minisets_/catacombs/harch7.tmx     _assets_/_minisets_/catacombs/horiz_arch7.tmx
cp _dump_/_minisets_/catacombs/harch8.tmx     _assets_/_minisets_/catacombs/horiz_arch8.tmx
cp _dump_/_minisets_/catacombs/harch9.tmx     _assets_/_minisets_/catacombs/horiz_arch9.tmx
cp _dump_/_minisets_/catacombs/harch10.tmx    _assets_/_minisets_/catacombs/horiz_arch10.tmx
cp _dump_/_minisets_/catacombs/harch11.tmx    _assets_/_minisets_/catacombs/horiz_arch11.tmx
cp _dump_/_minisets_/catacombs/harch12.tmx    _assets_/_minisets_/catacombs/horiz_arch12.tmx
cp _dump_/_minisets_/catacombs/harch13.tmx    _assets_/_minisets_/catacombs/horiz_arch13.tmx
cp _dump_/_minisets_/catacombs/harch14.tmx    _assets_/_minisets_/catacombs/horiz_arch14.tmx
cp _dump_/_minisets_/catacombs/harch15.tmx    _assets_/_minisets_/catacombs/horiz_arch15.tmx
cp _dump_/_minisets_/catacombs/harch16.tmx    _assets_/_minisets_/catacombs/horiz_arch16.tmx
cp _dump_/_minisets_/catacombs/harch17.tmx    _assets_/_minisets_/catacombs/horiz_arch17.tmx
cp _dump_/_minisets_/catacombs/harch18.tmx    _assets_/_minisets_/catacombs/horiz_arch18.tmx
cp _dump_/_minisets_/catacombs/harch19.tmx    _assets_/_minisets_/catacombs/horiz_arch19.tmx
cp _dump_/_minisets_/catacombs/harch20.tmx    _assets_/_minisets_/catacombs/horiz_arch20.tmx
cp _dump_/_minisets_/catacombs/harch21.tmx    _assets_/_minisets_/catacombs/horiz_arch21.tmx
cp _dump_/_minisets_/catacombs/harch22.tmx    _assets_/_minisets_/catacombs/horiz_arch22.tmx
cp _dump_/_minisets_/catacombs/harch23.tmx    _assets_/_minisets_/catacombs/horiz_arch23.tmx
cp _dump_/_minisets_/catacombs/harch24.tmx    _assets_/_minisets_/catacombs/horiz_arch24.tmx
cp _dump_/_minisets_/catacombs/harch25.tmx    _assets_/_minisets_/catacombs/horiz_arch25.tmx
cp _dump_/_minisets_/catacombs/harch26.tmx    _assets_/_minisets_/catacombs/horiz_arch26.tmx
cp _dump_/_minisets_/catacombs/harch27.tmx    _assets_/_minisets_/catacombs/horiz_arch27.tmx
cp _dump_/_minisets_/catacombs/harch28.tmx    _assets_/_minisets_/catacombs/horiz_arch28.tmx
cp _dump_/_minisets_/catacombs/harch29.tmx    _assets_/_minisets_/catacombs/horiz_arch29.tmx
cp _dump_/_minisets_/catacombs/harch30.tmx    _assets_/_minisets_/catacombs/horiz_arch30.tmx
cp _dump_/_minisets_/catacombs/harch31.tmx    _assets_/_minisets_/catacombs/horiz_arch31.tmx
cp _dump_/_minisets_/catacombs/harch32.tmx    _assets_/_minisets_/catacombs/horiz_arch32.tmx
cp _dump_/_minisets_/catacombs/harch33.tmx    _assets_/_minisets_/catacombs/horiz_arch33.tmx
cp _dump_/_minisets_/catacombs/harch34.tmx    _assets_/_minisets_/catacombs/horiz_arch34.tmx
cp _dump_/_minisets_/catacombs/harch35.tmx    _assets_/_minisets_/catacombs/horiz_arch35.tmx
cp _dump_/_minisets_/catacombs/harch36.tmx    _assets_/_minisets_/catacombs/horiz_arch36.tmx
cp _dump_/_minisets_/catacombs/harch37.tmx    _assets_/_minisets_/catacombs/horiz_arch37.tmx
cp _dump_/_minisets_/catacombs/harch38.tmx    _assets_/_minisets_/catacombs/horiz_arch38.tmx
cp _dump_/_minisets_/catacombs/harch39.tmx    _assets_/_minisets_/catacombs/horiz_arch39.tmx
cp _dump_/_minisets_/catacombs/harch40.tmx    _assets_/_minisets_/catacombs/horiz_arch40.tmx
cp _dump_/_minisets_/catacombs/ustairs.tmx    _assets_/_minisets_/catacombs/stairs_up.tmx
cp _dump_/_minisets_/catacombs/dstairs.tmx    _assets_/_minisets_/catacombs/stairs_down.tmx
cp _dump_/_minisets_/catacombs/warpstairs.tmx _assets_/_minisets_/catacombs/stairs_to_town.tmx
cp _dump_/_minisets_/catacombs/crushcol.tmx   _assets_/_minisets_/catacombs/crushed_column.tmx
cp _dump_/_minisets_/catacombs/big1.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big2.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big3.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big4.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big5.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big6.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big7.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big8.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big9.tmx       _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/big10.tmx      _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins1.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins2.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins3.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins4.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins5.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins6.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ruins7.tmx     _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/pancreas1.tmx  _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/pancreas2.tmx  _assets_/_minisets_/catacombs/
cp _dump_/_minisets_/catacombs/ctrdoor1.tmx   _assets_/_minisets_/catacombs/center_door1.tmx
cp _dump_/_minisets_/catacombs/ctrdoor2.tmx   _assets_/_minisets_/catacombs/center_door2.tmx
cp _dump_/_minisets_/catacombs/ctrdoor3.tmx   _assets_/_minisets_/catacombs/center_door3.tmx
cp _dump_/_minisets_/catacombs/ctrdoor4.tmx   _assets_/_minisets_/catacombs/center_door4.tmx
cp _dump_/_minisets_/catacombs/ctrdoor5.tmx   _assets_/_minisets_/catacombs/center_door5.tmx
cp _dump_/_minisets_/catacombs/ctrdoor6.tmx   _assets_/_minisets_/catacombs/center_door6.tmx
cp _dump_/_minisets_/catacombs/ctrdoor7.tmx   _assets_/_minisets_/catacombs/center_door7.tmx
cp _dump_/_minisets_/catacombs/ctrdoor8.tmx   _assets_/_minisets_/catacombs/center_door8.tmx

# --- [ Caves ] ----------------------------------------------------------------

mkdir -p _assets_/_minisets_/caves

# * l3up.tmx       -> stairs_up.tmx
# * l3down.tmx     -> stairs_down.tmx
# * l3holdwarp.tmx -> stairs_to_town.tmx
# * l3tite1.tmx    -> stalactite1.tmx
# * l3tite2.tmx    -> stalactite2.tmx
# * l3tite3.tmx    -> stalactite3.tmx
# * l3tite6.tmx    -> stalactite6.tmx
# * l3tite7.tmx    -> stalactite7.tmx
# * l3tite8.tmx    -> stalactite8.tmx
# * l3tite9.tmx    -> stalactite9.tmx
# * l3tite10.tmx   -> stalactite10.tmx
# * l3tite11.tmx   -> stalactite11.tmx
# * l3tite12.tmx   -> stalactite12.tmx
# * l3tite13.tmx   -> stalactite13.tmx
# * l3crev1.tmx    -> crevice1.tmx
# * l3crev2.tmx    -> crevice2.tmx
# * l3crev3.tmx    -> crevice3.tmx
# * l3crev4.tmx    -> crevice4.tmx
# * l3crev5.tmx    -> crevice5.tmx
# * l3crev6.tmx    -> crevice6.tmx
# * l3crev7.tmx    -> crevice7.tmx
# * l3crev8.tmx    -> crevice8.tmx
# * l3crev9.tmx    -> crevice9.tmx
# * l3crev10.tmx   -> crevice10.tmx
# * l3crev11.tmx   -> crevice11.tmx
# * l3isle1.tmx    -> isle1.tmx
# * l3isle2.tmx    -> isle2.tmx
# * l3isle3.tmx    -> isle3.tmx
# * l3isle4.tmx    -> isle4.tmx
# * l3isle5.tmx    -> isle5.tmx
# * l3xtra1.tmx    -> extra1.tmx
# * l3xtra2.tmx    -> extra2.tmx
# * l3xtra3.tmx    -> extra3.tmx
# * l3xtra4.tmx    -> extra4.tmx
# * l3xtra5.tmx    -> extra5.tmx
# * l3anvil.tmx    -> anvil.tmx

cp _dump_/_minisets_/caves/l3up.tmx       _assets_/_minisets_/caves/stairs_up.tmx
cp _dump_/_minisets_/caves/l3down.tmx     _assets_/_minisets_/caves/stairs_down.tmx
cp _dump_/_minisets_/caves/l3holdwarp.tmx _assets_/_minisets_/caves/stairs_to_town.tmx
cp _dump_/_minisets_/caves/l3tite1.tmx    _assets_/_minisets_/caves/stalactite1.tmx
cp _dump_/_minisets_/caves/l3tite2.tmx    _assets_/_minisets_/caves/stalactite2.tmx
cp _dump_/_minisets_/caves/l3tite3.tmx    _assets_/_minisets_/caves/stalactite3.tmx
cp _dump_/_minisets_/caves/l3tite6.tmx    _assets_/_minisets_/caves/stalactite6.tmx
cp _dump_/_minisets_/caves/l3tite7.tmx    _assets_/_minisets_/caves/stalactite7.tmx
cp _dump_/_minisets_/caves/l3tite8.tmx    _assets_/_minisets_/caves/stalactite8.tmx
cp _dump_/_minisets_/caves/l3tite9.tmx    _assets_/_minisets_/caves/stalactite9.tmx
cp _dump_/_minisets_/caves/l3tite10.tmx   _assets_/_minisets_/caves/stalactite10.tmx
cp _dump_/_minisets_/caves/l3tite11.tmx   _assets_/_minisets_/caves/stalactite11.tmx
cp _dump_/_minisets_/caves/l3tite12.tmx   _assets_/_minisets_/caves/stalactite12.tmx
cp _dump_/_minisets_/caves/l3tite13.tmx   _assets_/_minisets_/caves/stalactite13.tmx
cp _dump_/_minisets_/caves/l3crev1.tmx    _assets_/_minisets_/caves/crevice1.tmx
cp _dump_/_minisets_/caves/l3crev2.tmx    _assets_/_minisets_/caves/crevice2.tmx
cp _dump_/_minisets_/caves/l3crev3.tmx    _assets_/_minisets_/caves/crevice3.tmx
cp _dump_/_minisets_/caves/l3crev4.tmx    _assets_/_minisets_/caves/crevice4.tmx
cp _dump_/_minisets_/caves/l3crev5.tmx    _assets_/_minisets_/caves/crevice5.tmx
cp _dump_/_minisets_/caves/l3crev6.tmx    _assets_/_minisets_/caves/crevice6.tmx
cp _dump_/_minisets_/caves/l3crev7.tmx    _assets_/_minisets_/caves/crevice7.tmx
cp _dump_/_minisets_/caves/l3crev8.tmx    _assets_/_minisets_/caves/crevice8.tmx
cp _dump_/_minisets_/caves/l3crev9.tmx    _assets_/_minisets_/caves/crevice9.tmx
cp _dump_/_minisets_/caves/l3crev10.tmx   _assets_/_minisets_/caves/crevice10.tmx
cp _dump_/_minisets_/caves/l3crev11.tmx   _assets_/_minisets_/caves/crevice11.tmx
cp _dump_/_minisets_/caves/l3isle1.tmx    _assets_/_minisets_/caves/isle1.tmx
cp _dump_/_minisets_/caves/l3isle2.tmx    _assets_/_minisets_/caves/isle2.tmx
cp _dump_/_minisets_/caves/l3isle3.tmx    _assets_/_minisets_/caves/isle3.tmx
cp _dump_/_minisets_/caves/l3isle4.tmx    _assets_/_minisets_/caves/isle4.tmx
cp _dump_/_minisets_/caves/l3isle5.tmx    _assets_/_minisets_/caves/isle5.tmx
cp _dump_/_minisets_/caves/l3xtra1.tmx    _assets_/_minisets_/caves/extra1.tmx
cp _dump_/_minisets_/caves/l3xtra2.tmx    _assets_/_minisets_/caves/extra2.tmx
cp _dump_/_minisets_/caves/l3xtra3.tmx    _assets_/_minisets_/caves/extra3.tmx
cp _dump_/_minisets_/caves/l3xtra4.tmx    _assets_/_minisets_/caves/extra4.tmx
cp _dump_/_minisets_/caves/l3xtra5.tmx    _assets_/_minisets_/caves/extra5.tmx
cp _dump_/_minisets_/caves/l3anvil.tmx    _assets_/_minisets_/caves/anvil.tmx

# --- [ Hell ] -----------------------------------------------------------------

mkdir -p _assets_/_minisets_/hell

# * l4ustairs.tmx -> stairs_up.tmx
# * l4twarp.tmx   -> stairs_to_town.tmx
# * l4dstairs.tmx -> stairs_down.tmx
# * l4penta.tmx   -> penta1.tmx
# * l4penta2.tmx  -> penta2.tmx

cp _dump_/_minisets_/hell/l4ustairs.tmx _assets_/_minisets_/hell/stairs_up.tmx
cp _dump_/_minisets_/hell/l4twarp.tmx   _assets_/_minisets_/hell/stairs_to_town.tmx
cp _dump_/_minisets_/hell/l4dstairs.tmx _assets_/_minisets_/hell/stairs_down.tmx
cp _dump_/_minisets_/hell/l4penta.tmx   _assets_/_minisets_/hell/penta1.tmx
cp _dump_/_minisets_/hell/l4penta2.tmx  _assets_/_minisets_/hell/penta2.tmx
