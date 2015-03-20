package main

import (
    "fmt"
    "os"
)

// Get the root blog directory path as an argument to the main program command
// line. The format is `gloggi /path/to/blog` and the path must exist.
func GetArgRootPath() string {
    var dir string

    // Check that an argument exists
    if len(os.Args) > 1 {
        dir = os.Args[1]
    } else {
        fmt.Println("ERROR Must pass blog directory as argument")
        os.Exit(1)
    }

    // Check that the directory exists
    _, err := os.Stat(dir)
    if err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }

    return dir
}
