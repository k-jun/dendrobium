# dendrobium

## test

```sh
$ go test -failfast -v ./
```

## release

```sh
$ go install github.com/mitchellh/gox@latest
$ gox -arch="amd64 arm64" -os="linux darwin windows" \
    -ldflags="\
        -X main.version=`git describe --tag --abbrev=0` \
        -X main.revision=`git rev-list -1 HEAD` \
        -X main.build=`git describe --tags`"\
    -output=./dist/{{.Dir}}_{{.OS}}_{{.Arch}}
```
