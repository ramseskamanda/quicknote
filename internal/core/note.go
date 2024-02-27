package core

import (
	"time"
)

const (
	maxTitleLength = 20
)

type Note struct {
	Text      string
	Timestamp time.Time
}

func (note *Note) Title() string {
	if len(note.Text) > maxTitleLength {
		return note.Text[:maxTitleLength] + "..."
	}

	return note.Text
}

func (note *Note) Description() string { return note.Timestamp.String() }

func (note *Note) FilterValue() string { return note.Text }
