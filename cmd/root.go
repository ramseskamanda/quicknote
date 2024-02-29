package cmd

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/ramseskamanda/quicknote/internal/core"
	"github.com/ramseskamanda/quicknote/internal/storage"
	"github.com/ramseskamanda/quicknote/internal/tui"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var filename = os.Getenv("HOME") + "/.local/share/quicknote/"

var (
	listFlag bool
	Version  string
	Commit   string
)

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List existing notes")

	return rootCmd.ExecuteContext(context.Background())
}

// CLI args become timestamped notes
var rootCmd = &cobra.Command{
	Use:     "quicknote [note (optional)]",
	Version: fmt.Sprintf("%s (%s)", Version, Commit),
	Short:   "Quick notes from your command line",
	Long: `
quicknote is a simple CLI notes app, like jrnl but in Go!

Sometimes you just need to write down some notes when you're deep into a sensitive directory and don't want to open a new terminal or use vim. This is is exactly what this is useful for! :)
No markdown, no syntax highlighting, just simple notes notes without leaving the command line.
You can use it to easily write, search, and view notes. Notes are stored as human-readable plain text.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.Open(filename)
		if err != nil {
			return err
		}

		switch {
		case len(args) > 0:
			return db.WriteNote(time.Now(), strings.Join(args, " "))
		case listFlag:
			savedNotes, err := db.ListNotes()
			if err != nil {
				return err
			}

			var item []list.Item
			for _, note := range savedNotes {
				item = append(item, note)
			}

			onDelete := func(item list.Item) error {
				note := item.(*core.Note)

				return db.DeleteNote(note.Timestamp)
			}

			list, err := tui.List(item, onDelete)
			if err != nil {
				return err
			}

			if list.Selected == nil {
				return nil
			}

			note, ok := list.Selected.(*core.Note)
			if !ok {
				return nil
			}

			onSave := func(text string) error {
				return db.WriteNote(note.Timestamp, text)
			}

			return tui.NewEditor(note.Timestamp.String(), note.Text, onSave)
		default:
			timestamp := time.Now()

			onSave := func(text string) error {
				return db.WriteNote(timestamp, text)
			}

			return tui.NewEditor("New note", "", onSave)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Everything's saved! :)")
	},
}
