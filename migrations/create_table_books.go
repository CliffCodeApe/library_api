package migrations

import "database/sql"

type createBookTable struct{}

func (m *createBookTable) SkipProd() bool {
	return false
}

func getCreateBookTable() migration {
	return &createBookTable{}
}

func (m *createBookTable) Name() string {
	return "create-book"
}

func (m *createBookTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			publisher VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL,
			year INT NOT NULL,
			pages INT NOT NULL,
			stock INT NOT NULL,
			ISBN VARCHAR(20) NOT NULL UNIQUE,
			language VARCHAR(50) NOT NULL,
			thumbnail TEXT NOT NULL,
			file_path TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)

	if err != nil {
		return err
	}

	return err
}

func (m *createBookTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE books")

	return err
}
