# noqqe.de

This is the source of [noqqe.de](https://noqqe.de).

# Build

Build an run the blog on `localhost:1313`

```
$ git clone https://github.com/noqqe/noqqe.de
$ cd noqqe.de
$ brew install hugo
$ hugo server -w
```

Just building works quite the same, ends up with the `documentroot` in
`public/`

```
$ cd noqqe.de
$ hugo
```

Or on a custom server (to overwrite items in config.yaml) use

```
hugo server --baseURL=foo.example.com --port=1313 --bind=0.0.0.0 --watch
```

# Contribute

Patches / Pull Requests regarding

* Grammar
* Misspelling
* Comma
* HTML
* CSS

are very welcome. Please do **not** open pull requests for
`content/sammelsurium` posts! They are maintained somewhere else. It would
be a waste of efforts.
