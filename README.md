# dendrobium

## release

```sh
$ go install github.com/mitchellh/gox@latest
$ gox -arch="amd64 arm64" -os="linux darwin windows" -output=./dist/{{.Dir}}_{{.OS}}_{{.Arch}}
```

## test

```sh
$ go test -failfast -v ./
```
