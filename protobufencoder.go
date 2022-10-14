package main

import (
	"bytes"
	"encoding/binary"
)

func intToBytes(n int) []byte {
	data := int32(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
func encoder(meg []byte) []byte {
	body := make([]byte, 0)
	body = append(body, 0x39)
	body = append(body, intToBytes(len(meg))...)
	body = append(body, meg...)

	return body
}
