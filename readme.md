# ffw - Friendly Files for the Web
Say that you received 200 pictures with crappy names from the designer of the project you're on.
Oh boy, you can already feel your motivation coming down as you're renaming your files one by one manually.

**Wait** no more, with fftw - the easiest tool you can find over the globe to rename all your files to be web friendly
web.
# Installation

- on macOs
```bash
brew tap neonima/ffw
brew install ffw
```

- with golang
```bash
go get github.com/neonima/fftw
```

- 👉 for other platform go the [release](https://github.com/neonima/fftw/releases/) page and download the binaries you'll need

# Usage

## Problem

Say that you have these crappy png named files in `./src/images`
```text
- src/images/Le fichier galère à renommer 1.png
- src/images/Le fichier galère à renommer 2.png
- src/images/Le fichier galère à renommer 3.png
- src/images/Le fichier galère à renommer 4.png
```

## Solution

simply run `cd src/images && ffw` or with the source (-s) flag `ffw -s ./src/images`
```text
- src/images/le_fichier_galee_a_renommer_1.png
- src/images/le_fichier_galee_a_renommer_2.png
- src/images/le_fichier_galee_a_renommer_3.png
- src/images/le_fichier_galee_a_renommer_4.png
```

## Problem

Say that you received crappy png named files alongside files you do not wish to rename
```text
- src/images/mon fichier custom.jeanclaudevandamn
- src/images/Le fichier galère à renommer 1.png
- src/images/Le fichier galère à renommer 2.png
- src/images/Le fichier galère à renommer 3.png
- src/images/Le fichier galère à renommer 4.png
```

## Solution
Add `-e=[.your_extension1,.your_extension2]`

simply run `cd src/images && ffw -e=.png` or `ffw -s ./src/images -e=.png`
```text
- src/images/mon fichier custom.jeanclaudevandamn
- src/images/le_fichier_galee_a_renommer_1.png
- src/images/le_fichier_galee_a_renommer_2.png
- src/images/le_fichier_galee_a_renommer_3.png
- src/images/le_fichier_galee_a_renommer_4.png
```

# How can I check how the files will be changed?
For dry run use `-d` flag

# Help page

```text
NAME:
   fftw - A dead simple tool to rename you file for smooth web access!

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --recursive, -r             (default: Use this flag is you want to rename files recursively)
   --extensions .png, -e .png  -e=.png,.pdf | rename only .png et `.pdf` file
   --source value, -s value    (default: source path where to rename files)
   --dry, -d                   (default: shows all the file that will be modified with their new name)
   --help, -h                  show help (default: false)
```