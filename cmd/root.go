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
	version  string
	commit   string
)

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List existing notes")

	return rootCmd.ExecuteContext(context.Background())
}

// CLI args become timestamped notes
var rootCmd = &cobra.Command{
	Use:     "quicknote",
	Version: fmt.Sprintf("%s (%s)", version, commit),
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
