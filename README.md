# go-bencode

[![GoDoc](https://godoc.org/github.com/felix/go-bencode?status.svg)](http://godoc.org/github.com/felix/go-bencode)
[![Build Status](https://cloud.drone.io/api/badges/felix/go-bencode/status.svg)](https://cloud.drone.io/felix/go-bencode)
[![Coverage Status](https://coveralls.io/repos/github/felix/go-bencode/badge.svg?branch=master)](https://coveralls.io/github/felix/go-bencode?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/felix/go-bencode)](https://goreportcard.com/report/github.com/felix/go-bencode)

A Go library for handling data in the bencode format. This is the format used
by the Mainline DHT and BitTorrent protocols.

## Usage

All encoded data objects are `[]byte`. This is the format returned by many
network libraries such as `UDPConn.Read()` etc.

```go
import bencode "src.userspace.com.au/felix/go-bencode"
// or import bencode "github.com/felix/go-bencode"

var i interface{}
var packet []byte
var str string

i, err = bencode.Decode(packet)

// Or if you know the type
str, err = bencode.DecodeString(packet)

// Encoding a string, int, slice or map
packet, err = bencode.Encode(str)
```

## Benchmarks

Just for fun, and a comparison:

```sh
go test -bench=.

BenchmarkFelix-4                  100000             16544 ns/op
BenchmarkMarkSamman-4              50000             32664 ns/op
BenchmarkJackpal-4                 50000             36755 ns/op
```

