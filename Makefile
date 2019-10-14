all: opensource-ami

opensource-ami: _assets_

# Final output of game assets.

_assets_: _dump_ _extra_

# Post-processed raw game assets.

_extra_: _dump_/_dungeons_/tristram/tristram.tmx

_dump_/_dungeons_/tristram/tristram.tmx: _extra_/levels/towndata/tristram.dun | dun_dump
	dun_dump _extra_/levels/towndata/tristram.dun

_extra_/levels/towndata/tristram.dun: dump_duns | dun_merge
	@echo "===> Post-processing raw game assets."
	mkdir -p _extra_/levels/towndata
	dun_merge -output _extra_/levels/towndata/tristram.dun diabdat/levels/towndata/sector{1,2,3,4}s.dun

tristram.tmx:

# Raw dump of game assets.

_dump_: dump_cels dump_mins dump_tils dump_duns

# CEL images.

dump_cels: _dump_/towners/twnf/twnfw/twnfw_7/twnfw_0008.png | cel_dump

_dump_/towners/twnf/twnfw/twnfw_7/twnfw_0008.png: diabdat
	@echo "===> Extracting CEL images."
	cel_dump -a

# MIN files (miniture tiles).

dump_mins: _dump_/_dpieces_/town/town.pal/dpiece_1258.png | min_dump

_dump_/_dpieces_/town/town.pal/dpiece_1258.png: diabdat
	@echo "===> Extracting MIN files."
	min_dump -a

# TIL files (tiles).

dump_tils: _dump_/_tiles_/town/town.pal/tile_0342.png | til_dump

_dump_/_tiles_/town/town.pal/tile_0342.png: diabdat
	@echo "===> Extracting TIL files."
	til_dump -a

# DUN files (set level dungeon maps).

dump_duns: _dump_/_dungeons_/tristram/sector4s.tmx | dun_dump

_dump_/_dungeons_/tristram/sector4s.tmx: diabdat
	@echo "===> Extracting DUN files."
	dun_dump -a

# Extracted DIABDAT.MPQ archive.

diabdat: diabdat.mpq | mpq mpqfix
	@echo "===> Extracting MPQ archive."
	mpq -dir $@ -m $<
	mpqfix -mpqdump $@

clean:
	$(RM) -v -r _dump_ _extra_ _assets_

.PHONY: all clean dump_cels dump_mins dump_tils dump_duns mpq mpqfix cel_dump min_dump til_dump dun_dump dun_merge

# DIABDAT.MPQ archive.

diabdat.mpq:
	@if [ ! -f $@ ]; then                                                                          \
		echo "Unable to locate \"$@\". Please copy \"$@\" to this directory or provide a symlink."; \
		echo;                                                                                       \
		echo "   ln -s /path/to/diabdat.mpq .";                                                     \
		echo;                                                                                       \
		exit 1;                                                                                     \
	fi;

# Tools.

mpq:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/mpq";                                           \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

mpqfix:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   github.com/mewrnd/blizzconv/cmd/mpqfix";                                    \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

cel_dump:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/formats/cmd/cel_dump";                          \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

min_dump:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/formats/cmd/min_dump";                          \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

til_dump:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/formats/cmd/til_dump";                          \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

dun_dump:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/formats/cmd/dun_dump";                          \
		echo;                                                                                \
		exit 1;                                                                              \
	fi

dun_merge:
	@if ! which $@ &> /dev/null ; then                                                      \
		echo "Unable to locate the \"$@\" command in PATH. Please install the \"$@\" tool."; \
		echo "Also, remember to add GOPATH/bin to PATH.";                                    \
		echo;                                                                                \
		echo "   go get github.com/sanctuary/formats/cmd/dun_merge";                         \
		echo;                                                                                \
		exit 1;                                                                              \
	fi
