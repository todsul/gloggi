package main

import (
    "path"
)

var pnme string

func CleanPostOutput(dir string) {
    for _, info := range GetDirectoryListing(dir) {

        if (info.IsDir() && info.Name() != "src" && info.Name() != ".git") {

            pnme = path.Join(dir,info.Name())
            RemoveDirectory(pnme)
        }
    }
}

func CleanHomeOutput(dir string) {

    pnme = path.Join(dir, "index.html")
    RemoveFile(pnme)
}
