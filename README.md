# Straw

A simple-to-use, and pretty-to-look-at command-line interface tool to help enter project directories using Tmux very easily.

Uses [Bubbletea](https://github.com/charmbracelet/bubbletea) for the heavy TUI stuff

## Features: 
- [X] View and open projects
- [X] Automatically open Tmux in the chosen project
- [ ] Add projects
  - [X] From the command line (project directory)
  - [ ] From the command line (any directory)
  - [ ] From the TUI
- [ ] Delete projects
  - [X] From the command line
  - [ ] From the TUI 

## Support 
| Linux | MacOS | Windows |
| --- | --- |--- |
| ✅ | ✅ | ❌ |


## Installing Guide

#### Dependencies:
- Go
- Tmux

### Arch Linux
Since Straw is on the [AUR](https://aur.archlinux.org/packages/straw) you can use yay or any other aur helper
```
yay -S straw 
```

### General Linux
```
git clone https://github.com/siddhantmadhur/straw
cd straw
go build -ldflags="-X 'main.version=$(git describe --tags --abbrev=0)'" -o ./bin/straw
mv ./bin/straw /bin/straw
```

### General MacOS
```
git clone https://github.com/siddhantmadhur/straw
cd straw
go build -ldflags="-X 'main.version=$(git describe --tags --abbrev=0)'" -o ./bin/straw
mv ./bin/straw /usr/local/bin/straw
```
