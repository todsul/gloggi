package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

// Create a new directory
func CreateDirectory(path string) {
    os.Mkdir(path, 0750)
}

// Create a new file
func CreateFile(path string, output []byte) {
    ioutil.WriteFile(path, output, 0750)
}

// Get a directory listing and return the file info
func GetDirectoryListing(path string) []os.FileInfo {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println(err)
	}

    return files;
}

// Get a file's contents
func GetFile(path string) []byte {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err)
	}

    return file
}

// Remove a file
func RemoveFile(path string) {
    err := os.Remove(path)
    if err != nil {
        fmt.Println("ERROR", err)
    }
}

// Remove a directory, including all children
func RemoveDirectory(path string) {
    err := os.RemoveAll(path)
    if err != nil {
        fmt.Println("ERROR", err)
    }
}
