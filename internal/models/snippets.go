package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// func para insertar un snippet en la db
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	// statement query para el insert en la tabla snippets
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type before returning.
	return int(id), nil
}

// func que retorna un snippet en base al id en la db
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// func que retorna los 10 ultimos snippets de la db
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
