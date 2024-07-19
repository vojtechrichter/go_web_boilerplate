package app

import (
    "github.com/gin-gonic/gin"
    "database/sql"
)

func RoutesInit(r *gin.Engine, db *sql.DB) {
    protected := r.Group("/")
    protected.Use(AuthRequiredMW)
    {
        protected.GET("/", func(c *gin.Context) {
        })
    }
}

func GinInit(db *sql.DB) *gin.Engine {
    r := gin.Default()
    r.Static("/static", "./static")
    r.Static("/node_modules", "./")

    RoutesInit(r, db)

    return r
}
