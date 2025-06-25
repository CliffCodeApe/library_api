package migrations

import "database/sql"

type createLendingTable struct{}

func (m *createLendingTable) SkipProd() bool {
	return false
}

func getCreateLendingTable() migration {
	return &createLendingTable{}
}

func (m *createLendingTable) Name() string {
	return "create-lending"
}

func (m *createLendingTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE lendings (
			id SERIAL PRIMARY KEY,
			book_id BIGINT NOT NULL,
			user_id BIGINT NOT NULL,
			status VARCHAR(50) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)
	if err != nil {
		return err
	}
	return err
}

func (m *createLendingTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE lendings")
	return err
}
