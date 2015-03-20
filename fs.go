package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func CreateDirectory(path string) {
    os.Mkdir(path, 0750)
}

func CreateFile(path string, output []byte) {
    ioutil.WriteFile(path, output, 0750)
}

func GetDirectoryListing(path string) []os.FileInfo {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println(err)
	}

    return files;
}

func GetFile(path string) []byte {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err)
	}

    return file
}

func RemoveFile(path string) {
    err := os.Remove(path)
    if err != nil {
        fmt.Println("ERROR", err)
    }
}

func RemoveDirectory(path string) {
    err := os.RemoveAll(path)
    if err != nil {
        fmt.Println("ERROR", err)
    }
}
