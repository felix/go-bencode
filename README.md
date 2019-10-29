# go-bencode

A Go library for handling data in the bencode format. This is the format used
by the Mainline DHT and BitTorrent protocols.

## Usage

All encoded data objects are `[]byte`. This is the format returned by many
network libraries such as `UDPConn.Read()` etc.

```go
import bencode "src.userspace.com.au/go-bencode"

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

Just for fun and comparison (go v1.13):

```sh
go test -bench=.
goos: freebsd
goarch: amd64
pkg: src.userspace.com.au/go-bencode
BenchmarkFelix-4                  140659              8518 ns/op
BenchmarkMarkSamman-4              72080             16857 ns/op
BenchmarkJackpal-4                 70902             17194 ns/op
BenchmarkDecodeWithString-4      6856792               183 ns/op
BenchmarkDecodeString-4         10423938               130 ns/op
BenchmarkDecodeWithInt-4         8644882               146 ns/op
BenchmarkDecodeInt-4            11680635               107 ns/op
BenchmarkDecodeWithList-4        2046406               527 ns/op
BenchmarkDecodeList-4            2494192               464 ns/op
BenchmarkDecodeWithDict-4        2505790               479 ns/op
BenchmarkDecodeDict-4            2674564               480 ns/op
PASS
ok      src.userspace.com.au/go-bencode 17.455s
```

