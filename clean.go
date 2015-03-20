package main

import (
    "path"
)

var pfmt string

func CleanPostOutput(dir string) {
    for _, info := range GetDirectoryListing(dir) {

        if (info.IsDir() && info.Name() != "src" && info.Name() != ".git") {

            pfmt = path.Join(dir,info.Name())
            RemoveDirectory(pfmt)
        }
    }
}

func CleanHomeOutput(dir string) {

    pfmt = path.Join(dir,"index.html")
    RemoveFile(pfmt)
}
