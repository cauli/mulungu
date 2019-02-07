package postgres

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// DB is a singleton instance of the Storage
var DB *Storage

// Storage is a struct that holds important
// database info like the connection pool and resource in use used
type Storage struct {
	Pool     *pgx.ConnPool
	Resource string
}

const bootstrapQuery = `CREATE TABLE IF NOT EXISTS tree (
	id varchar(40) NOT NULL,
	account varchar(40) NOT NULL,
	data jsonb,
	PRIMARY KEY (id, account)
);`

// New will generate a Storage containing a
// pool of connections for a given resource type
func New(resource string) (Storage, error) {
	pool, err := pgx.NewConnPool(getConfig())
	if err != nil {
		return Storage{}, err
	}

	_, err = pool.Exec(bootstrapQuery)
	if err != nil {
		return Storage{}, err
	}

	return Storage{pool, resource}, nil
}

// SetMainStorage will a single-instance storage globally
func SetMainStorage(storage *Storage) {
	DB = storage
}

// Load will make a SELECT query to a database given a resource an id
func (storage *Storage) Load(resource, id string) (result string, notFound bool, e error) {
	var response []byte

	if storage.Pool == nil {
		return "", false, fmt.Errorf("No database connection available")
	}

	query := fmt.Sprintf("SELECT data FROM %v WHERE id = $1 and account = 'default'", resource)
	err := storage.Pool.QueryRow(query, id).Scan(&response)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", true, fmt.Errorf("Entry was not found")
		}
		return "", false, fmt.Errorf("An storage error has occurred: %v", err)
	}

	return string(response), false, nil
}

// Delete will make a DELETE query to the database given a resource and an id
func (storage *Storage) Delete(resource, id string) error {
	if storage.Pool == nil {
		return fmt.Errorf("No database connection available")
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE id = $1 AND account = 'default'", resource)
	_, err := storage.Pool.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// Save will make a INSERT query to the database given a resource and an id
func (storage *Storage) Save(resource, id, data string) error {
	if storage.Pool == nil {
		return fmt.Errorf("No database connection available")
	}

	values := fmt.Sprintf(`('%v','default','%v')`, id, data)
	_, err := storage.Pool.Exec(fmt.Sprintf("INSERT INTO %v VALUES %v ON CONFLICT(id,account) DO UPDATE SET data = EXCLUDED.data;", resource, values))
	if err != nil {
		return err
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
	config.Port = 5432

	return config
}
