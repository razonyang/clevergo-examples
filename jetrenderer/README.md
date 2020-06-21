# Jet Renderer Example

Install `packr` for embedding the templates in the binary.

```shell
$ go get -u github.com/gobuffalo/packr/v2/packr2
```

```shell
$ cd jetrenderer && packr2 install
```

The templates is now embedded in the binary.

```shell
$ cd ~ && jetrenderer
```

Then open http://localhost:8080.
