# gloggi

Minimalist static blog generator built with Go

## Structure

In order for gloggi to work, the blog needs a basic structure:

```
|- src
    |- posts
        |- YYYYMMDD-name-of-post.md
    |- templates
        |- _summary.html
        |- home.html
        |- post.html
```

## Posts

The sources files for the blog posts are markdown. The filenames must be of the above format, including the date and name of the post. The filename is split and used as metadata for the blog post.

## Templates

There above templates require the following placeholders

* Post: `date`, `name`, `content`
* Home: `posts`
* Summary: `path`, `name`

The syntax for placesholders is `{{ field }}`

## Command

gloggi /path-to-blog
