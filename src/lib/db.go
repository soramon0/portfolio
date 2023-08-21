package lib

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/soramon0/portfolio/src/internal/database"
)

func NewDB(url string, l *AppLogger) *database.Queries {
	db, err := sql.Open("postgres", url)
	if err != nil {
		l.ErrorFatal(err)
	}
	return database.New(db)
}
