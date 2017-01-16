// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crc8 implements the 8-bit cyclic redundancy check, or CRC-8,
// checksum. See http://en.wikipedia.org/wiki/Cyclic_redundancy_check for
// information.
package crc8

// The size of a CRC-8 checksum in bytes.
const Size = 1

// Table is a 256-word table representing the polynomial for efficient processing.
type Table [256]uint8

// MakeTable returns a Table constructed from the specified polynomial.
// The contents of this Table must not be modified.
func MakeTable(poly uint8) *Table {
	return makeTable(poly)
}

// makeTable returns the Table constructed from the specified polynomial.
func makeTable(poly uint8) *Table {
	t := new(Table)
	for i := 0; i < 256; i++ {
		crc := uint8(i)
		for j := 0; j < 8; j++ {
			if crc&0x80 != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
		t[i] = crc
	}
	return t
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	crc uint8
	tab *Table
}

// New creates a new hash.Hash8 computing the CRC-32 checksum
// using the polynomial represented by the Table.
// Its Sum method will lay the value out in big-endian byte order.
func New(tab *Table) Hash8 { return &digest{0, tab} }

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return 1 }

func (d *digest) Reset() { d.crc = 0 }

func update(crc uint8, tab *Table, p []byte) uint8 {
	for _, v := range p {
		crc = tab[byte(crc)^v]
	}
	return crc
}

// Update returns the result of adding the bytes in p to the crc.
func Update(crc uint8, tab *Table, p []byte) uint8 {
	return update(crc, tab, p)
}

func (d *digest) Write(p []byte) (n int, err error) {
	d.crc = Update(d.crc, d.tab, p)
	return len(p), nil
}

func (d *digest) Sum8() uint8 { return d.crc }

func (d *digest) Sum(in []byte) []byte {
	return append(in, d.crc)
}

// Checksum returns the CRC-8 checksum of data
// using the polynomial represented by the Table.
func Checksum(data []byte, tab *Table) uint8 { return Update(0, tab, data) }
