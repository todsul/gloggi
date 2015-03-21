package main

import (
    "bytes"
    "path"
    "strings"
    "github.com/russross/blackfriday"
)

type Post struct {
    date string
    file string
    html []byte
    name string
    path string
    slug string
}

// Extract the post date from the pathname, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetDateFromFilename(pathname string) {
    info := strings.Split(pathname, ".md")[0]
    date := info[:4] + "-" + info[4:6] +  "-" + info[6:8]

    p.date = date
}

// Extract the post name from the pathname, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetNameFromFilename(pathname string) {
    info := strings.Split(pathname, ".md")[0]
    name := strings.Replace(info[9:], "-", " ", -1)

    p.name = name
}

// Extract the post slug from the pathname, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetSlugFromFilename(pathname string) {
    info := strings.Split(pathname, ".md")[0]
    slug := strings.ToLower(info[9:])

    p.slug = slug
    p.path = "/" + p.slug
}


// Set the post html by first converting the markdown to html and then
// decorating the result with the base post template.
func (p *Post) SetHtmlFromMarkdown(html []byte, md []byte) {
    text := blackfriday.MarkdownBasic(md)

    html = bytes.Replace(html, []byte("{{ date }}"), []byte(p.date), -1)
    html = bytes.Replace(html, []byte("{{ name }}"), []byte(p.name), -1)
    html = bytes.Replace(html, []byte("{{ text }}"), []byte(text), -1)

    p.html = html
}

// Loop through source directories recursively to create category directories,
// post directories, and index html files. Also works for uncategorized posts.
func ProcessPosts(base []byte, in string, out string, posts Posts) Posts {
    for _, info := range GetDirectoryListing(in) {
        pathname := info.Name()

        if info.IsDir() {
            // Create category directory
            CreateDirectory(path.Join(out, pathname))

            // Process category posts recursively
            posts = ProcessPosts(base, path.Join(in, pathname), path.Join(out, pathname), posts)
        } else {
            // Get the markdown
            md := GetFile(path.Join(in, pathname))

            // Build the post
            post := new(Post)
            post.SetDateFromFilename(pathname)
            post.SetNameFromFilename(pathname)
            post.SetSlugFromFilename(pathname)
            post.SetHtmlFromMarkdown(base, md)

            // Create the output directory
            CreateDirectory(path.Join(out,post.slug))

            // Create the output file
            CreateFile(path.Join(out,post.slug, "index.html"), post.html)

            // Append to Post collection
            posts = append(posts, *post)
        }
    }

    return posts
}

type Posts []Post

func (slice Posts) Len() int {
    return len(slice)
}

func (slice Posts) Less(i, j int) bool {
    return slice[i].date > slice[j].date;
}

func (slice Posts) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}
