package postgres

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

type Storage struct {
	conn     *pgx.ConnPool
	Resource string
}

const bootstrapQuery = `CREATE TABLE IF NOT EXISTS fixed_rules (
	id varchar(40) NOT NULL,
	account varchar(40) NOT NULL,
	data jsonb,
	PRIMARY KEY (id, account)
);`

func New(resource string) (Storage, error) {
	pgx.DefaultTypeFormats["jsonb"] = pgx.TextFormatCode

	conn, err := pgx.NewConnPool(getConfig())
	if err != nil {
		return Storage{}, err
	}

	_, err = conn.Exec(bootstrapQuery)
	if err != nil {
		return Storage{}, err
	}

	return Storage{conn, resource}, nil
}

func (storage *Storage) Load(id string, target interface{}) error {
	var response []byte

	if storage.conn == nil {
		return fmt.Errorf("No database connection available")
	}

	query := fmt.Sprintf("SELECT data FROM %v WHERE id = $1 and account = 'default'", storage.Resource)

	err := storage.conn.QueryRow(query).Scan(&response)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("Entry was not found")
		}
		return fmt.Errorf("An storage error has occurred")
	}

	err = json.Unmarshal(response, target)
	if err != nil {
		return fmt.Errorf("An error occurred while loading id `%s` from storage", id)
	}

	return nil
}

func getConfig() pgx.ConnPoolConfig {
	var config pgx.ConnPoolConfig

	config.Host = os.Getenv("POSTGRES_HOST")
	config.User = os.Getenv("POSTGRES_USER")
	config.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Database = os.Getenv("POSTGRES_DB")
	config.MaxConnections = 25

	return config
}
