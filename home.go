package main

import (
    "bytes"
    "sort"
)

type Home struct {
    html []byte
}

// Set the home html by looping through the posts, using the summary template
// to create post summaries, then decorating the result with the base template.
func (h *Home) SetHtmlFromPosts(html []byte, summ []byte, posts Posts) {
    var summaries []byte

    // Sort the Posts
    sort.Sort(posts)

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

func ProcessHome(html []byte, summ []byte, out string, posts Posts) {
    home := new(Home)
    home.SetHtmlFromPosts(html, summ, posts)

    CreateFile(out + "/index.html", home.html)
}
