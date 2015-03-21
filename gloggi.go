package main

import (
    "fmt"
    "path"
    "time"
)

var posts Posts

// The main gloggi process in a step-by-step format to clean the output,
// process the raw blog post content, and create index pages.
func main() {
    fmt.Println("*** START ***")
    start := time.Now()

    root := GetArgRootPath()

    paths := map[string]string {
        "in": path.Join(root, "/src/posts"),
        "out": path.Join(root),
    }

    templates := map[string][]byte {
        "home": GetFile(path.Join(root, "/src/templates/home.html")),
        "summ": GetFile(path.Join(root, "/src/templates/summ.html")),
        "post": GetFile(path.Join(root, "/src/templates/post.html")),
    }

    // STEP 1: Clean the workspace (see clean.go)

    ProcessClean(paths["out"])

    // STEP 2: Create post output (see post.go)

    posts := ProcessPosts(templates["post"], paths["in"], paths["out"], posts)

    // STEP 3: Create home output (see home.go)

    ProcessHome(templates["home"], templates["summ"], paths["out"], posts)

    end := time.Now()
    fmt.Println("*** END", end.Sub(start), "***")
}
