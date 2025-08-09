package repository

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	notesTable = "notes"
)

type Note interface {
	CreateNote(userId int, input model.CreateNoteInput) (int, error)
	GetNoteByID(userId, noteId int) (model.Note, error)
	GetAllNotes(userId int) ([]model.Note, error)
	UpdateNote(userId, noteId int, input model.UpdateNoteInput) error
	DeleteNote(userId, noteId int) error
}

type Repository struct {
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note: NewNotePostgres(db),
	}
}