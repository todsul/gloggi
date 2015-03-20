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

// Filename Functions

func (p *Post) SetDateFromFilename(filename string) {
    // Filename format is YYYYMMDD-name-of-post.md

    info := strings.Split(filename, ".md")[0]
    date := info[:4] + "-" + info[4:6] +  "-" + info[6:8]

    p.date = date
}

func (p *Post) SetNameFromFilename(filename string) {
    // Filename format is YYYYMMDD-name-of-post.md

    info := strings.Split(filename, ".md")[0]
    name := strings.Replace(info[9:], "-", " ", -1)

    p.name = name
}

func (p *Post) SetSlugFromFilename(filename string) {
    // Filename format is YYYYMMDD-name-of-post.md

    info := strings.Split(filename, ".md")[0]
    slug := strings.ToLower(info[9:])

    p.slug = slug
    p.path = "/" + p.slug
}

// Content Functions

func (p *Post) SetHtmlFromMarkdown(base []byte, md []byte) {
    // Using github.com/russross/blackfriday to translate markdown

    html := base
    text := blackfriday.MarkdownBasic(md)

    html = bytes.Replace(html, []byte("{{ date }}"), []byte(p.date), -1)
    html = bytes.Replace(html, []byte("{{ name }}"), []byte(p.name), -1)
    html = bytes.Replace(html, []byte("{{ text }}"), []byte(text), -1)

    p.html = html
}
