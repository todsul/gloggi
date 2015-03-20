package main

import (
    "fmt"
    "os"
)

func GetArgRootPath() string {
    var dir string

    if len(os.Args) > 1 {
        dir = os.Args[1]
    } else {
        fmt.Println("ERROR Must pass blog directory as argument")
        os.Exit(1)
    }

    // Check that directory exists
    _, err := os.Stat(dir)
    if err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }

    return dir
}
