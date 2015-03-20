# gloggi

Minimalist static blog generator built with Go

![Alt gloggi](https://d30y9cdsu7xlg0.cloudfront.net/svg/bc9b3e47-b69a-4622-a725-7f4429824187.svg?Expires=1426848721&Signature=CxmMZIe~~2AWgQCXwb2kygAkMjMevrQzgrATnALOjPj6dXEDmg5dNqni6yv~DETZ-O-ZxMtx48dDWhyfoDFgtCGadxan5iiFxKzEzcb8J0GRMYhtFlRrdvEus11Wz~4pi6~eakbudZr3AsrB9iDWqyh0lWNdaIe4~VdKN89ie74_&Key-Pair-Id=APKAI5ZVHAXN65CHVU2Q)

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
        |- _summary.html
        |- home.html
        |- post.html
```

The source files for the blog posts are markdown. The filenames must be of the above format, including the date and name of the post. The filename is split and used as metadata for the blog post.

## Templates

There above templates require the following placeholders

* Post: `date`, `name`, `text`
* Home: `posts`
* Summary: `path`, `name`

The syntax for placesholders is `{{ field }}`

## Command

To run gloggi on the example blog:

`gloggi example`

To run gloggi on your blog:

`gloggi /path/to/blog`

## Todo

* Compile LessCSS
* Handle category folders
* Write category index.html files
* Include templates for header, footer, etc
* Use for the prod version of todsul.com
