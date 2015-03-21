package main

import (
    "bytes"
)

type Home struct {
    html []byte
}

// Set the home html by looping through the posts, using the summary template
// to create post summaries, then decorating the result with the base template.
func (h *Home) SetHtmlFromPosts(html []byte, summ []byte, posts []*Post) {
    var summaries []byte

    // Loop through Post collection
    for _, post := range posts {
        html := summ

        // Replace summary template variables
        html = bytes.Replace(html, []byte("{{ path }}"), []byte(post.path), -1)
        html = bytes.Replace(html, []byte("{{ name }}"), []byte(post.name), -1)

        summaries = append(summaries, html...)
    }

    // Replace base template variables
    html = bytes.Replace(html, []byte("{{ posts }}"), summaries, -1)

    h.html = html
}

func ProcessHome(html []byte, summ []byte, out string, posts []*Post) {
    home := new(Home)
    home.SetHtmlFromPosts(html, summ, posts)

    CreateFile(out + "/index.html", home.html)
}
