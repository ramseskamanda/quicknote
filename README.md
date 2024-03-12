# Quicknote CLI

[![Latest Release](https://img.shields.io/github/release/ramseskamanda/quicknote.svg)](https://github.com/ramseskamanda/quicknote/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramseskamanda/quicknote)](https://goreportcard.com/report/github.com/ramseskamanda/quicknote)
[![License](https://img.shields.io/github/license/ramseskamanda/quicknote)](https://github.com/ramseskamanda/quicknote/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/ramseskamanda/quicknote/graph/badge.svg?token=MNDQOUZMEJ)](https://codecov.io/gh/ramseskamanda/quicknote)

## About

`quicknote` is a simple CLI notes app, like [jrnl](https://jrnl.sh/en/stable/usage/) but simpler and in Go!

| ![Menu screenshot](/docs/screenshots/menu.png) | ![Editor screenshot](/docs/screenshots/editor.png) |
|:----------------------------------------------:|:--------------------------------------------------:|
|                List your notes                 |                     Edit notes                     |

Sometimes you just need to write down some notes when you're deep into
a sensitive directory and don't want to open a new terminal or use vim.
This is is exactly what this is useful for! :)

No markdown, no syntax highlighting, just simple notes notes without leaving the command line.

You can use it to easily write, search, and view notes. Notes are stored as human-readable plain text.

## Installation

### From source

Make sure you have go>=1.22 installed, then run:

```shell
go install github.com/ramseskamanda/quicknote
```

### From package managers

#### Homebrew

```shell
brew tap ramseskamanda/tap
brew install ramseskamanda/tap/quicknote
```

#### AUR

```shell
pacman -S quicknote
```

## Basic Usage

To create a new note without opening an editor, simply use:

```shell
quicknote i really need to remember this!!
```

If you prefer opening an editor and writing a note over time in a simple, no-nonsense editor, you can do that by using:

```shell
quicknote
```

To see all the notes you've written so far, use:

```shell
quicknote --list
```

> Note: you can search (<kbd>/</kbd>), delete (<kbd>ctrl + d</kbd>), and edit (<kbd>e</kbd>/<kbd>⏎</kbd>) any note in the list. For more commands, use the help (<kbd>?</kbd>).

### The editor

You don't have to use the provided editor! It's made for simplicity with minimal key binds and/or distractions.
All the notes are saved as `.txt` files in the `$HOME/.local/share/quicknote/` directory so feel free to use any other editor of your choice.

All helpers should be displayed by the application but in case they are not, simply run `quicknote --help`.

## License

[MIT](/LICENSE)
