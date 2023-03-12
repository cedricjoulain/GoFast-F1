// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package zlib implements reading and writing of zlib format compressed data,
as specified in RFC 1950.

The implementation provides filters that uncompress during reading
and compress during writing.  For example, to write compressed data
to a buffer:

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

and to read that data back:

	r, err := zlib.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()
*/
package main

import (
	"bufio"
	"compress/flate"
	"io"
)

const (
	zlibDeflate   = 8
	zlibMaxWindow = 7
)

type reader struct {
	r            flate.Reader
	decompressor io.ReadCloser
	err          error
}

// Resetter resets a ReadCloser returned by NewReader or NewReaderDict
// to switch to a new underlying Reader. This permits reusing a ReadCloser
// instead of allocating a new one.
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader) error
}

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement io.ByteReader, the decompressor may read more
// data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when done.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r io.Reader) (io.ReadCloser, error) {
	z := new(reader)
	err := z.Reset(r)
	if err != nil {
		return nil, err
	}
	return z, nil
}

func (z *reader) Read(p []byte) (int, error) {
	if z.err != nil {
		return 0, z.err
	}

	var n int
	n, z.err = z.decompressor.Read(p)
	if z.err != io.EOF {
		// In the normal case we return here.
		return n, z.err
	}

	// Finished file
	return n, io.EOF
}

// Calling Close does not close the wrapped io.Reader originally passed to NewReader.
// In order for the ZLIB checksum to be verified, the reader must be
// fully consumed until the io.EOF.
func (z *reader) Close() error {
	if z.err != nil && z.err != io.EOF {
		return z.err
	}
	z.err = z.decompressor.Close()
	return z.err
}

func (z *reader) Reset(r io.Reader) error {
	*z = reader{decompressor: z.decompressor}
	if fr, ok := r.(flate.Reader); ok {
		z.r = fr
	} else {
		z.r = bufio.NewReader(r)
	}

	if z.decompressor == nil {
		z.decompressor = flate.NewReader(z.r)
	} else {
		z.decompressor.(flate.Resetter).Reset(z.r, nil)
	}
	return nil
}
