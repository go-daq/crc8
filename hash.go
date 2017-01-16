// Copyright 2017 The go-daq Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc8

import (
	"hash"
)

// Hash8 is the common interface implemented by all 8-bit hash functions.
type Hash8 interface {
	hash.Hash
	Sum8() uint8
}
