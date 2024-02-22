# quicknote

---

`quicknote` is a simple CLI notes app, like jrnl but in Go!

Sometimes you just need to write down some notes when you're deep into
a sensitive directory and don't want to open a new terminal or use vim.
This is is exactly what this is useful for! :)

No markdown, no syntax highlighting, just simple notes notes without leaving the command line.

You can use it to easily write, search, and view notes. Notes are stored as human-readable plain text.

## Features

- [ ] CLI args become timestamped notes
- [ ] No args starts an editor that becomes a timestamped note when closed
- [ ] save notes to single file
- [ ] list notes from cli
- [ ] search notes by time
- [ ] Note titles
- [ ] Edit previous notes
- [ ] https://jrnl.sh/en/stable/usage/
- [ ] https://github.com/charmbracelet/bubbletea/tree/master/examples/split-editors

## Installation

### From source

!TODO

### From AUR or other package managers

!TODO

## Basic Usage

Simply use `quicknote` (personally aliased to `qn`) and an editor will open for you.

You can also run `quicknote my root partition is 320MB because of the /opt/ directory!!` and this note will be saved!

All helpers should be displayed by the application but in case they are not, simply run `quicknote --help`.

## License

[MIT](/LICENSE)
