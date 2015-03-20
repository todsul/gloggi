package main

import (
    "bytes"
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

// Extract the post date from the filename, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetDateFromFilename(filename string) {
    info := strings.Split(filename, ".md")[0]
    date := info[:4] + "-" + info[4:6] +  "-" + info[6:8]

    p.date = date
}

// Extract the post name from the filename, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetNameFromFilename(filename string) {
    info := strings.Split(filename, ".md")[0]
    name := strings.Replace(info[9:], "-", " ", -1)

    p.name = name
}

// Extract the post slug from the filename, which holds the meta data. This
// avoids including foreign data in the raw post markdown file.
// Filename format is YYYYMMDD-name-of-post.md
func (p *Post) SetSlugFromFilename(filename string) {
    info := strings.Split(filename, ".md")[0]
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
