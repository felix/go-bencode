package bencode

import (
	"bytes"
	"io/ioutil"
	"testing"

	alt2 "github.com/jackpal/bencode-go"
	alt1 "github.com/marksamman/bencode"
)

var data []byte

func BenchmarkFelix(b *testing.B) {
	data, _ = ioutil.ReadFile("testdata/torrent.bin")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Decode(data)
	}
}

func BenchmarkMarkSamman(b *testing.B) {
	for n := 0; n < b.N; n++ {
		alt1.Decode(bytes.NewBuffer(data))
	}
}

func BenchmarkJackpal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		alt2.Decode(bytes.NewBuffer(data))
	}
}
