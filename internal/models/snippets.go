package models

import (
	"database/sql"
	"errors"
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
	// statement query para el insert en la tabla snippets
	stmt := `SELECT * FROM snippets 
				WHERE expires > UTC_TIMESTAMP() AND ID = ?`

	row := m.DB.QueryRow(stmt, id)

	// Creo una estructura de snippet vacía y la comparo con la row que recibí,
	// ya que el método QueryRow() no devuelve un error en sí
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil

}

// func que retorna los 10 ultimos snippets de la db
func (m *SnippetModel) Latest() ([]*Snippet, error) {

	stmt := `SELECT * FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	// --------------------------------- CRITICO --------------------------------- //
	// defer siempre después de chequear el error q devuelve el query (Y SIEMPRE VA!)
	defer rows.Close()

	// Creo un slice para manejar los snippets
	snippets := []*Snippet{}

	// los rows funcionan como una lista (con iteradores)
	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
