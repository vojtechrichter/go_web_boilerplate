package app

import (
    "database/sql"
    _ "modernc.org/sqlite"
)

func DbInterfaceInit(dbPath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite", dbPath)
    if err != nil {
        return nil, err
    }

    return db, nil
}
