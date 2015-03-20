package main

import (
    "fmt"
    "path"
    "time"
)

// The main gloggi process in a step-by-step format to clean the output,
// process the raw blog post content, and create index pages.
func main() {
    fmt.Println("*** START ***")

    var (
        base []byte
        posts []*Post
        pnme string
    )

    dir := GetArgRootPath()
    start := time.Now()

    // STEP 1: Clean the workspace (see clean.go)

    CleanPostOutput(dir)
    CleanHomeOutput(dir)

    // STEP 2: Create the post output (see post.go)

    // Get the post template
    pnme = path.Join(dir, "/src/templates/post.html")
    base = GetFile(pnme)

    // Loop the source post files
    pnme = path.Join(dir, "/src/posts")
    for _, info := range GetDirectoryListing(pnme) {

        // Get the markdown
        filename := info.Name()
        pnme = path.Join(dir, "/src/posts", filename)
        md := GetFile(pnme)

        // Build the post
        post := new(Post)
        post.SetDateFromFilename(filename)
        post.SetNameFromFilename(filename)
        post.SetSlugFromFilename(filename)
        post.SetHtmlFromMarkdown(base, md)

        // Create the output directory
        pnme = path.Join(dir,post.slug)
        CreateDirectory(pnme)

        // Create the output file
        pnme = path.Join(dir,post.slug, "index.html")
        CreateFile(pnme, post.html)

        // Collect for home page
        posts = append(posts, post)

        fmt.Println("gloggi.write", pnme)
    }

    // STEP 3: Create the home index page

    // Get the home template
    pnme = path.Join(dir, "/src/templates/home.html")
    base = GetFile(pnme)

    // Get the summary template
    pnme = path.Join(dir, "/src/templates/_summary.html")
    summ := GetFile(pnme)

    // Build the home page
    home := new(Home)
    home.html = home.GetHtmlFromPosts(base, summ, posts)

    // Create the index.html
    pnme = path.Join(dir, "index.html")
    CreateFile(pnme, home.html)

    fmt.Println("gloggi.write", pnme)
    fmt.Println("*** END", time.Now().Sub(start), "***")
}
