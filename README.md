# noqqe.de

This is the source of [noqqe.de](https://noqqe.de).

# Build

[![Build Status](https://travis-ci.org/noqqe/noqqe.de.svg?branch=master)](https://travis-ci.org/noqqe/noqqe.de)

Build an run the blog on `localhost:1313`

``` bash
$ git clone https://github.com/noqqe/noqqe.de
$ cd noqqe.de
$ brew install hugo
$ hugo server -w
```

Just building works quite the same, ends up with the `documentroot` in
`public/`

``` bash
$ cd noqqe.de
$ hugo
```

Or on a custom server (to overwrite items in config.yaml) use

``` bash
hugo server --baseURL=foo.example.com --port=1313 --bind=0.0.0.0 --watch
```

Think about running `markdownlint`

# Contribute

Patches / Pull Requests regarding

* Grammar
* Misspelling
* Comma
* HTML
* CSS

are very welcome.
