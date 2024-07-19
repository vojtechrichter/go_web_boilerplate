package app

import (
    "os"
    "io"
    "log"
)

func LogError(err error) {
    file, err := os.OpenFile("/tmp/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Error occured while trying to write to a log file")
    }
    defer file.Close()

    writer := io.MultiWriter(os.Stdout, file)
    log.SetOutput(writer)
    log.Println(err)
    log.SetOutput(os.Stdout)
}
