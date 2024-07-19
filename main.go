package main

import "go_web_boilerplate/app"

func che(err error) {
    if err != nil {
        app.LogError(err)
    }
}

func main() {
    _, err := app.DbInterfaceInit("ahoj.db")
    che(err)
}
