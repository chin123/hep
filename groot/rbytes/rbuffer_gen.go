// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Automatically generated. DO NOT EDIT.

package rbytes

import (
	"encoding/binary"
	"math"
)

func (r *RBuffer) ReadArrayU16(sli []uint16) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 2*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 2
		cur = end
		v := binary.BigEndian.Uint16(sub[beg:end])
		sli[i] = v

	}
	r.r.c = end
}

func (r *RBuffer) ReadU16() uint16 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 2
	v := binary.BigEndian.Uint16(r.r.p[beg:r.r.c])
	return v

}

func (r *RBuffer) ReadArrayU32(sli []uint32) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 4*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 4
		cur = end
		v := binary.BigEndian.Uint32(sub[beg:end])
		sli[i] = v

	}
	r.r.c = end
}

func (r *RBuffer) ReadU32() uint32 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 4
	v := binary.BigEndian.Uint32(r.r.p[beg:r.r.c])
	return v

}

func (r *RBuffer) ReadArrayU64(sli []uint64) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 8*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 8
		cur = end
		v := binary.BigEndian.Uint64(sub[beg:end])
		sli[i] = v

	}
	r.r.c = end
}

func (r *RBuffer) ReadU64() uint64 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 8
	v := binary.BigEndian.Uint64(r.r.p[beg:r.r.c])
	return v

}

func (r *RBuffer) ReadArrayI16(sli []int16) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 2*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 2
		cur = end
		v := binary.BigEndian.Uint16(sub[beg:end])
		sli[i] = int16(v)
	}
	r.r.c = end
}

func (r *RBuffer) ReadI16() int16 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 2
	v := binary.BigEndian.Uint16(r.r.p[beg:r.r.c])
	return int16(v)
}

func (r *RBuffer) ReadArrayI32(sli []int32) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 4*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 4
		cur = end
		v := binary.BigEndian.Uint32(sub[beg:end])
		sli[i] = int32(v)
	}
	r.r.c = end
}

func (r *RBuffer) ReadI32() int32 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 4
	v := binary.BigEndian.Uint32(r.r.p[beg:r.r.c])
	return int32(v)
}

func (r *RBuffer) ReadArrayI64(sli []int64) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 8*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 8
		cur = end
		v := binary.BigEndian.Uint64(sub[beg:end])
		sli[i] = int64(v)
	}
	r.r.c = end
}

func (r *RBuffer) ReadI64() int64 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 8
	v := binary.BigEndian.Uint64(r.r.p[beg:r.r.c])
	return int64(v)
}

func (r *RBuffer) ReadArrayF32(sli []float32) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 4*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 4
		cur = end
		v := binary.BigEndian.Uint32(sub[beg:end])
		sli[i] = math.Float32frombits(v)
	}
	r.r.c = end
}

func (r *RBuffer) ReadF32() float32 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 4
	v := binary.BigEndian.Uint32(r.r.p[beg:r.r.c])
	return math.Float32frombits(v)
}

func (r *RBuffer) ReadArrayF64(sli []float64) {
	if r.err != nil {
		return
	}
	n := len(sli)
	if n <= 0 || int64(n) > r.Len() {
		return
	}

	cur := r.r.c
	end := r.r.c + 8*len(sli)
	sub := r.r.p[cur:end]
	cur = 0
	for i := range sli {
		beg := cur
		end := cur + 8
		cur = end
		v := binary.BigEndian.Uint64(sub[beg:end])
		sli[i] = math.Float64frombits(v)
	}
	r.r.c = end
}

func (r *RBuffer) ReadF64() float64 {
	if r.err != nil {
		return 0
	}
	beg := r.r.c
	r.r.c += 8
	v := binary.BigEndian.Uint64(r.r.p[beg:r.r.c])
	return math.Float64frombits(v)
}
