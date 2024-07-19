package app

import (
    "os"
    "github.com/lpernett/godotenv"
    "database/sql"
    "github.com/gin-gonic/gin"
)

type App struct {
    gin *gin.Engine
    db *sql.DB
}

func Init() (*App, error) {
    app := new(App)

    err := godotenv.Load()
    if err != nil {
        return nil, err
    }
    dbInt, err := DbInterfaceInit(os.Getenv("DB_NAME"))
    if err != nil {
        return nil, err
    }
    app.db = dbInt

    app.gin = GinInit(app.db)

    return app, nil
}
