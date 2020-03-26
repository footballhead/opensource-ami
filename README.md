# Open source-ami

**Note**, Open source-ami requires an original copy of `diabdat.mpq`. None of the Diablo 1 game assets are provided by this project. To get a legitimate copy of the game assets, please refer to the [GoG release of Diablo 1](https://www.gog.com/game/diablo).

The aim of this project is to automate the conversion of Diablo 1 game assets to file formats with open specifications.

## Usage

```bash
git clone https://github.com/sanctuary/opensource-ami
cd opensource-ami
./get_tools.sh
go install ./...
ln -s /path/to/diablo-v1.09b.exe diablo.exe
ln -s /path/to/diabdat.mpq diabdat.mpq
make
```
