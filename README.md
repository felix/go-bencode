# go-bencode

A Go library for handling data in the bencode format. This is the format used
by the Mainline DHT and BitTorrent protocols.


## Install

    go get -u github.com/felix/go-bencode


## Usage

All encoded data objects are `[]byte`. This is the format returned by many
network libraries such as `UDPConn.Read()` etc.

    import bencode "github.com/felix/go-bencode"

    var i interface{}
    var packet []byte
    var str string

    i, err = bencode.Decode(packet)

    // Or if you know the type
    str, err = bencode.DecodeString(packet)

    // Encoding a string, int, slice or map
    packet, err = bencode.Encode(str)
