xp-runners-go
=============
[![Build statuson travis-ci](https://travis-ci.org/kiesel/xp-runners-go.svg)](https://travis-ci.org/kiesel/xp-runners-go)

This project is an effort in implementing the [XP Runners](https://github.com/xp-framework/xp-runners) (which are currently implemented in *bash* / *C#* on Windows) in golang.

Complilation
------------

This assumes, you have a working Go setup available, and `$GOPATH` declared validly.

* Get code and dependencies:

```sh
$ go get github.com/kiesel/xp-runners-go
```

* Compile binaries

```sh
$ go install github.com/kiesel/xp-runners-go/xp
$ go install github.com/kiesel/xp-runners-go/xpcli
```

* Run test suite

```sh
$ go test -v github.com/kiesel/xp-runners-go/xp
$ go test -v github.com/kiesel/xp-runners-go/xpcli
$ go test -v github.com/kiesel/xp-runners-go/runner
```

The compiled binaries reside in `$GOPATH/bin`.


Completeness
------------

Implementation completeness in comparison to the original runners is as follows:

* [ ] cgen
* [ ] doclet
* [ ] unittest
* [ ] xar
* [ ] xcc
* [ ] xpi
* [ ] xpws
* [x] xp
* [x] xpcli

