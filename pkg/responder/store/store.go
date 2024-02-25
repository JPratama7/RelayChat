package store

import (
	"database/sql"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func NewStore(db *sql.DB, dialect string, log waLog.Logger) (container *sqlstore.Container, err error) {
	container = sqlstore.NewWithDB(db, dialect, log)
	err = container.Upgrade()
	return
}
