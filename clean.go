package main

import (
    "path"
)

// Clean all output directories except src and .git. It has to be this brutal
// because to work as a static blog, the main posts need to be in the root
// directory and not an output directory, which would affect the urls.
func ProcessClean(dir string) {
    // STEP 1 Remove non-source directories
    for _, info := range GetDirectoryListing(dir) {
        if info.IsDir() && info.Name() != "src" && info.Name() != ".git" {
            RemoveDirectory(path.Join(dir,info.Name()))
        }
    }

    // STEP 2 Remove home page
    RemoveFile(path.Join(dir, "index.html"))
}
