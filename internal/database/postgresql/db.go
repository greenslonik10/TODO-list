package postgresql

import (
	"os"

	"github.com/jackc/pgx"
)

func NewPostgresConn() (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConnectionString(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return nil, err
	}

	return conn, err

}
