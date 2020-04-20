all: opensource-ami

opensource-ami: diablo.exe diabdat.mpq _assets_

# Check that tools exist in PATH.

TOOLS=mpq mpqfix cel_dump min_dump til_dump dun_dump dun_merge miniset_dump

# ref: https://stackoverflow.com/a/25668869
EVAL_RHS := $(foreach TOOL, $(TOOLS), \
	$(if $(shell which $(TOOL)), \
		nop, \
		$(error "Unable to locate '$(TOOL)' command in $$PATH. Please run './get_tools.sh'. Also, remember to set $$GOPATH and add $$GOPATH/bin to $$PATH.") \
	) \
)

_assets_: _dump_
	@echo "===> Copying game assets to \"$@\" directory."
	./copy_assets.sh

# Raw dump of game assets.

_dump_: dump_cels dump_mins dump_tils dump_duns dump_minisets post_process

# CEL images.

dump_cels: _dump_/towners/twnf/twnfw/twnfw_7/twnfw_0008.png

_dump_/towners/twnf/twnfw/twnfw_7/twnfw_0008.png: diabdat
	@echo "===> Extracting CEL images."
	cel_dump -a

# MIN files (miniture tiles).

dump_mins: _dump_/_dpieces_/town/town.pal/dpiece_1258.png

_dump_/_dpieces_/town/town.pal/dpiece_1258.png: diabdat
	@echo "===> Extracting MIN files."
	min_dump -a

# TIL files (tiles).

dump_tils: _dump_/_tiles_/town/town.pal/tile_0342.png _dump_/_tiles_special_/town/town.pal/tile_0342.png

_dump_/_tiles_/town/town.pal/tile_0342.png: diabdat
	@echo "===> Extracting TIL files."
	til_dump -a

_dump_/_tiles_special_/town/town.pal/tile_0342.png: diabdat
	@echo "===> Extracting TIL files with special tileset."
	til_dump -a -special

# DUN files (set level dungeon maps).

dump_duns: _dump_/_dungeons_/tristram/sector4s.tmx

_dump_/_dungeons_/tristram/sector4s.tmx: diabdat
	@echo "===> Extracting DUN files."
	dun_dump -a

# Minisets (miniature set levels).

dump_minisets: _dump_/_minisets_/hell/l4ustairs.tmx

_dump_/_minisets_/hell/l4ustairs.tmx: diablo.exe
	@echo "===> Extracting minisets."
	miniset_dump diablo.exe

# Post-processed raw game assets.

post_process: dump_duns _dump_/_dungeons_/tristram/tristram.tmx

_dump_/levels/towndata/tristram.dun:
	@echo "===> Post-processing raw game assets."
	dun_merge -o _dump_/levels/towndata/tristram.dun diabdat/levels/towndata/sector1s.dun diabdat/levels/towndata/sector2s.dun diabdat/levels/towndata/sector3s.dun diabdat/levels/towndata/sector4s.dun

_dump_/_dungeons_/tristram/tristram.tmx: _dump_/levels/towndata/tristram.dun
	dun_dump _dump_/levels/towndata/tristram.dun

# Extracted DIABDAT.MPQ archive.

diabdat: diabdat.mpq
	@echo "===> Extracting MPQ archive."
	mpq -dir $@ -m $<
	mpqfix -mpqdump $@

clean:
	$(RM) -v -r _dump_ _assets_

.PHONY: all clean dump_cels dump_mins dump_tils dump_duns dump_minisets post_process

# DIABDAT.MPQ archive.

diabdat.mpq:
	@if [ ! -f $@ ]; then \
		echo "Unable to locate \"$@\". Please copy \"$@\" to this directory or provide a symlink."; \
		echo; \
		echo "   ln -s /path/to/$@ ."; \
		echo; \
		exit 1; \
	fi;

# DIABLO.EXE executable.

diablo.exe:
	@if [ ! -f $@ ]; then \
		echo "Unable to locate \"$@\" (v1.09b). Please copy \"$@\" to this directory or provide a symlink."; \
		echo; \
		echo "   ln -s /path/to/$@ ."; \
		echo; \
		exit 1; \
	fi;
