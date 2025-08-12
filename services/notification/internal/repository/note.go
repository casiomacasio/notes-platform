package repository

import (
	"errors"
	// "fmt"
	// "github.com/casiomacasio/notes-platform/services/notification/internal/model"
)

var (
	ErrNoteNotFound = errors.New("note not found")
)

// type NotePostgres struct {
// 	db *sqlx.DB
// }

// func NewNotePostgres(db *sqlx.DB) *NotePostgres {
// 	return &NotePostgres{db: db}
// }

// func (r *NotePostgres) CreateNote(userId int, input model.CreateNoteInput) (int, error) {
// 	var id int
// 	query := fmt.Sprintf(
// 		"INSERT INTO %s (user_id, title, content) VALUES ($1, $2, $3) RETURNING id",
// 		notesTable,
// 	)
// 	err := r.db.QueryRow(query, userId, input.Title, input.Content).Scan(&id)
// 	return id, err
// }

// func (r *NotePostgres) GetNoteByID(userId, noteId int) (model.Note, error) {
// 	var note model.Note
// 	query := fmt.Sprintf("SELECT id, user_id, title, content FROM %s WHERE id = $1 AND user_id = $2", notesTable)
// 	err := r.db.Get(&note, query, noteId, userId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return model.Note{}, ErrNoteNotFound
// 		}
// 		return model.Note{}, err
// 	}
// 	return note, nil
// }

// func (r *NotePostgres) GetAllNotes(userId int) ([]model.Note, error) {
// 	var notes []model.Note
// 	query := fmt.Sprintf("SELECT id, user_id, title, content FROM %s WHERE user_id = $1", notesTable)
// 	err := r.db.Select(&notes, query, userId)
// 	return notes, err
// }

// func (r *NotePostgres) UpdateNote(userId, noteId int, input model.UpdateNoteInput) error {
// 	query := fmt.Sprintf("UPDATE %s SET", notesTable)
// 	args := []interface{}{}
// 	argIdx := 1

// 	if input.Title != nil {
// 		query += fmt.Sprintf(" title = $%d,", argIdx)
// 		args = append(args, *input.Title)
// 		argIdx++
// 	}
// 	if input.Content != nil {
// 		query += fmt.Sprintf(" content = $%d,", argIdx)
// 		args = append(args, *input.Content)
// 		argIdx++
// 	}

// 	if len(args) == 0 {
// 		return nil
// 	}

// 	query += " updated_at = NOW()"
// 	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
// 	args = append(args, noteId)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }

// func (r *NotePostgres) DeleteNote(userId, noteId int) error {
// 	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND user_id = $2", notesTable)
// 	result, err := r.db.Exec(query, noteId, userId)
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rows == 0 {
// 		return ErrNoteNotFound
// 	}
// 	return nil
// }
