package service

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/casiomacasio/notes-platform/services/note/internal/repository"
)

type Note interface {
	CreateNote(userId int, input model.CreateNoteInput) (int, error)
	GetNoteByID(userId, noteId int) (model.Note, error)
	GetAllNotes(userId int) ([]model.Note, error)
	UpdateNote(userId, noteId int, input model.UpdateNoteInput) error
	DeleteNote(userId, noteId int) error
}

type Authorization interface {
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Note
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Note: NewNoteService(repos.Note),
		Authorization: NewAuthService(),
	}
}
