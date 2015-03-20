package main

import (
    "path"
)

// Clean all output directories except src and .git. It has to be this brutal
// because to work as a static blog, the main posts need to be in the root
// directory and not an output directory, which would affect the urls.
func CleanPostOutput(dir string) {
    for _, info := range GetDirectoryListing(dir) {
        if (info.IsDir() && info.Name() != "src" && info.Name() != ".git") {
            pnme := path.Join(dir,info.Name())
            RemoveDirectory(pnme)
        }
    }
}

// Clean the home output index page
func CleanHomeOutput(dir string) {
    pnme := path.Join(dir, "index.html")
    RemoveFile(pnme)
}
