package main

import (
    "go-api-catalog/database"
    "go-api-catalog/database/migrations"
)

func main() {
    db := database.ConnectDB()
    defer db.Close()

    migrations.RunMigrations(db)
}