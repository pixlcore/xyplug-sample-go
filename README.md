# Overview

This is an example Go module which can be downloaded and executed by a single `go run` command. It is designed to read JSON from STDIN, and output JSON to STDOUT, as part of a bare-bones simple xyOps Event Plugin. 

Here is the Plugin command for this package:

```sh
go run github.com/pixlcore/xyplug-sample-go@v1.0.0
```

Here is an example invocation with some test data piped in:

```sh
echo '{"xy":1,"test":[2,3]}' | go run github.com/pixlcore/xyplug-sample-go@v1.0.0
```

Expected output:

```
Read JSON from STDIN:
{
  "test": [
    2,
    3
  ],
  "xy": 1
}

{ "xy":1, "code":0 }
```

Note: The pretty-printed JSON key ordering may vary. xyOps will ignore and pass through the pretty-printed output, and only consume the compact message at the end.

# Instructions

- Ensure you have Go installed and configured (`go version`).
- Make sure your repository root contains a proper `go.mod` with `module github.com/pixlcore/xyplug-sample-go`.
- The `main.go` at the repo root provides the `main` package entrypoint required by `go run`.
- Tag the repo to match the version in the docs (currently `v1.0.0`):

```sh
git tag v1.0.0
git push origin v1.0.0
```

# Implementation notes

- Uses Go's standard library `encoding/json` to parse and pretty-print.
- Reads all input from STDIN; if no input is provided, only the final JSON line is printed.
- Always prints the final line: `{ "xy":1, "code":0 }` before exiting.

