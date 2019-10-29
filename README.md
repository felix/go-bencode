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

Just for fun and comparison:

```sh
go test -bench=.
goos: freebsd
goarch: amd64
pkg: src.userspace.com.au/go-bencode
BenchmarkFelix-4                  100000             22520 ns/op
BenchmarkMarkSamman-4              30000             47871 ns/op
BenchmarkJackpal-4                 30000             46966 ns/op
BenchmarkDecodeWithString-4      5000000               379 ns/op
BenchmarkDecodeString-4         10000000               199 ns/op
BenchmarkDecodeWithInt-4        10000000               239 ns/op
BenchmarkDecodeInt-4            10000000               187 ns/op
BenchmarkDecodeWithList-4        1000000              1047 ns/op
BenchmarkDecodeList-4            2000000               869 ns/op
BenchmarkDecodeWithDict-4        2000000               972 ns/op
BenchmarkDecodeDict-4            2000000               959 ns/op
PASS
ok      src.userspace.com.au/go-bencode   24.999s
```

