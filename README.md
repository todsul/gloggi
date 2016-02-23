# gloggi

An ultra-minimalist static blog generator built with Go

gloggi is a minimalist blog generator on purpose; it's designed to let you focus
on the content and not the typical window dressing.

gloggi has the following minimal features:

1. Categories created from source directory names
2. Metadate extracted from post filename (YYYYMMDD-name-of-post.md)
3. Raw markdown in post files (no foreign metadata)


## Install

To install gloggi, run the following:

`go get github.com/todsul/gloggi`

## Structure

In order for gloggi to work, the blog needs a basic structure:

```
|- src
    |- posts
        |- YYYYMMDD-name-of-post.md
    |- templates
        |- home.html
        |- post.html
        |- summ.html
```

The source files for the blog posts are markdown. The filenames must be of the above format, including the date and name of the post. The filename is split and used as metadata for the blog post.

## Templates

There above templates require the following placeholders

* Post: `date`, `name`, `text`
* Home: `posts`
* Summ: `path`, `name`

The syntax for placesholders is `{{ field }}`

## Command

To run gloggi on the example blog:

`gloggi example`

To run gloggi on your blog:

`gloggi /path/to/blog`
