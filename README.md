# go-makefile-example

Example project for the Makefile explained in [A Good Makefile for Go](http://azer.bike/journal/a-good-makefile-for-go/) blog post.

## Setup

Run following commands to try it out:

* git clone https://github.com/azer/go-makefile-example.git
* cd go-makefile-example
* make install
* make start

It should now be running at :7777 (defined in `.env` file). You can try making a change and
see how the server will get recompiled and restarted automatically.

## Manual

Run `make help` to list available commands:

```
  λ  make help

 Choose a command run in go-makefile-example:

  install   Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
  start     Start in development mode. Auto-starts when code changes.
  stop      Stop development mode.
  watch     Run given command when code changes. e.g; make watch run="echo 'hey'"
  compile   Compile the binary.
  exec      Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
  clean     Clean build files. Runs `go clean` internally.
```
