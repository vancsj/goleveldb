// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This LevelDB Go implementation is based on LevelDB C++ implementation.
// Which contains the following header:
//   Copyright (c) 2011 The LevelDB Authors. All rights reserved.
//   Use of this source code is governed by a BSD-style license that can be
//   found in the LEVELDBCPP_LICENSE file. See the LEVELDBCPP_AUTHORS file
//   for names of contributors.

package db

import (
	"encoding/binary"
	"io"
)

type readByteReader interface {
	io.Reader
	io.ByteReader
}

func sliceBytes(b []byte) []byte {
	z, n := binary.Uvarint(b)
	return b[n : n+int(z)]
}

func sliceBytesTest(b []byte) (valid bool, v, rest []byte) {
	z, n := binary.Uvarint(b)
	m := n + int(z)
	if n <= 0 || m > len(b) {
		return
	}
	valid = true
	v = b[n:m]
	rest = b[m:]
	return
}

func readBytes(r readByteReader) (b []byte, err error) {
	var n uint64
	n, err = binary.ReadUvarint(r)
	if err != nil {
		return
	}
	b = make([]byte, n)
	_, err = io.ReadFull(r, b)
	if err != nil {
		b = nil
	}
	return
}