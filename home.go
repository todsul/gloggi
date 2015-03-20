package main

import (
    "bytes"
)

type Home struct {
    html []byte
}

// Content Functions

func (h *Home) GetHtmlFromPosts(base []byte, snip []byte, posts []*Post) []byte {
    // Loop through posts and concat snippets

    html := base
    summary := snip

    var summaries []byte
    for _, post := range posts {
        summary = bytes.Replace(summary, []byte("{{ path }}"), []byte(post.path), -1)
        summary = bytes.Replace(summary, []byte("{{ name }}"), []byte(post.name), -1)

        summaries = append(summaries, summary...)
    }

    html = bytes.Replace(html, []byte("{{ posts }}"), summaries, -1)

    return html
}
