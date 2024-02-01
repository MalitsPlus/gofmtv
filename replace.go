package main

import (
	"bytes"
	"io"
)

func replaceTab2Space(srcBuf *bytes.Buffer) bytes.Buffer {
	var desBuf bytes.Buffer
	s := make([]byte, 256)
	for {
		n, err := srcBuf.Read(s)
		r := bytes.ReplaceAll(s[:n], []byte{'\t'}, []byte("  "))
		desBuf.Write(r)
		if err == io.EOF {
			break
		}
	}
	return desBuf
}

func replaceTab2SpaceSlice(srcSlice []byte) []byte {
	return bytes.ReplaceAll(srcSlice, []byte{'\t'}, []byte("  "))
}
