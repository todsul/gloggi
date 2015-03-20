package main

import (
    "fmt"
    "path"
    "time"
)

func main() {

    fmt.Println("*** START ***")

    var (
        base []byte
        posts []*Post
        pfmt string
        snip []byte
    )

    dir := GetArgRootPath()
    start := time.Now()

    // STEP 1: Clean the workspace (see clean.go)

    CleanPostOutput(dir)
    CleanHomeOutput(dir)

    // STEP 2: Create the post output (see post.go)

    // Get the post template
    pfmt = path.Join(dir,"/src/templates/post.html")
    base = GetFile(pfmt)

    // Loop the source post files
    pfmt = path.Join(dir,"/src/posts")
    for _, info := range GetDirectoryListing(pfmt) {

        // Get the markdown
        pfmt = path.Join(dir,"/src/posts",info.Name())
        md := GetFile(pfmt)

        // Build the post
        post := new(Post)
        post.date = post.GetDateFromFilename(info.Name())
        post.name = post.GetNameFromFilename(info.Name())
        post.slug = post.GetSlugFromFilename(info.Name())
        post.html = post.GetHtmlFromMarkdown(base, md)
        post.path = path.Join("/",post.slug)

        // Create the output directory
        pfmt = path.Join(dir,post.slug)
        CreateDirectory(pfmt)

        // Create the output file
        pfmt = path.Join(dir,post.slug,"index.html")
        CreateFile(pfmt, post.html)

        // Collect for home page
        posts = append(posts, post)

        fmt.Println("gloggi.write", pfmt)
    }

    // STEP 3: Create the home index page

    // Get the home template
    pfmt = path.Join(dir,"/src/templates/home.html")
    base = GetFile(pfmt)

    // Get the summary snippet
    pfmt = path.Join(dir,"/src/templates/_summary.html")
    snip = GetFile(pfmt)

    // Build the home page
    home := new(Home)
    home.html = home.GetHtmlFromPosts(base, snip, posts)

    // Create the index.html
    pfmt = path.Join(dir,"index.html")
    CreateFile(pfmt, home.html)

    fmt.Println("gloggi.write", pfmt)
    fmt.Println("*** END", time.Now().Sub(start), "***")
}
