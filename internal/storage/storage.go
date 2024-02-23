package storage

import (
	"errors"
	"fmt"
	"github.com/ramseskamanda/quicknote/internal/core"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const FilenameFmt = time.RFC1123

type Dir struct {
	path string
}

func Open(where string) (*Dir, error) {
	if _, err := os.Stat(where); errors.Is(err, os.ErrNotExist) {
		if err := initialize(where); err != nil {
			return nil, fmt.Errorf("uninitialized and failed to initialize: %w", err)
		}
	}

	return &Dir{path: where}, nil
}

func initialize(where string) error {
	if err := os.MkdirAll(where, 0750); err != nil {
		return fmt.Errorf("failed to create parent directories: %w", err)
	}

	filename := path.Join(where, time.Now().Format(FilenameFmt)+".txt")
	if err := os.WriteFile(filename, []byte("Notes initialized!"), 0666); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	fmt.Println("Successfully initialized!")

	return nil
}

func (dir *Dir) WriteNote(text string) error {
	if text == "" {
		return nil
	}

	filename := path.Join(dir.path, time.Now().Format(FilenameFmt)+".txt")

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

func (dir *Dir) ListNotes() ([]*core.Note, error) {
	files, err := os.ReadDir(dir.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read storage directory: %w", err)
	}

	var notes []*core.Note
	for _, file := range files {
		if ext := path.Ext(file.Name()); ext != ".txt" {
			log.Printf("found file without .txt extension: %s", file.Name())

			continue
		}

		filename := strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name()))

		timestamp, err := time.Parse(time.RFC1123, filename)
		if err != nil {
			log.Printf(
				"found .txt file but couldn't parse file name: %s",
				fmt.Errorf("%s : %w", filename, err),
			)

			continue
		}

		content, err := os.ReadFile(path.Join(dir.path, file.Name()))
		if err != nil {
			log.Printf(
				"found .txt file but couldn't open it: %s",
				fmt.Errorf("%s : %w", filename, err),
			)

			continue
		}

		notes = append(notes, &core.Note{Text: string(content), Timestamp: timestamp})
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[j].Timestamp.Before(notes[i].Timestamp)
	})

	return notes, nil
}
