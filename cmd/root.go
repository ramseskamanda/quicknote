package cmd

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/ramseskamanda/quicknote/internal/storage"
	"github.com/ramseskamanda/quicknote/internal/tui"
	"github.com/spf13/cobra"
	"strings"
)

var filename = "/home/ramsesk/.local/share/quicknote/"

var (
	listFlag bool
)

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List existing notes")

	return rootCmd.ExecuteContext(context.Background())
}

// CLI args become timestamped notes
var rootCmd = &cobra.Command{
	Use: "quicknote",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.Open(filename)
		if err != nil {
			return err
		}

		switch {
		case listFlag:
			savedNotes, err := db.ListNotes()
			if err != nil {
				return err
			}

			var item []list.Item
			for _, note := range savedNotes {
				item = append(item, note)
			}

			return tui.List(item)
		case len(args) > 0:
			return db.WriteNote(strings.Join(args, " "))
		default:
			return tui.Editor(db.WriteNote)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Everything's saved! :)")
	},
}
