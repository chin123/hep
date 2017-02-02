// Copyright 2017 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rootio

import (
	"fmt"
	"reflect"
)

type tbranch struct {
	named          tnamed
	attfill        attfill
	compress       int       // compression level and algorithm
	basketSize     int       // initial size of Basket buffer
	entryOffsetLen int       // initial length of entryOffset table in the basket buffers
	writeBasket    int       // last basket number written
	entryNumber    int64     // current entry number (last one filled in this branch)
	offset         int       // offset of this branch
	maxBaskets     int       // maximum number of baskets so far
	splitLevel     int       // branch split level
	entries        int64     // number of entries
	firstEntry     int64     // number of the first entry in this branch
	totBytes       int64     // total number of bytes in all leaves before compression
	zipBytes       int64     // total number of bytes in all leaves after compression
	branches       []tbranch // list of branches of this branch
	leaves         []Leaf    // list of leaves of this branch
	baskets        []Basket  // list of baskets of this branch

	basketBytes []int32 // length of baskets on file
	basketEntry []int64 // table of first entry in each basket
	basketSeek  []int64 // addresses of baskets on file

	fname string // named of file where buffers are stored (empty if in same file as Tree header)

	readbasket  int     // current basket number when reading
	readentry   int64   // current entry number when reading
	firstbasket int64   // first entry in the current basket
	nextbasket  int64   // next entry that will reaquire us to go to the next basket
	currbasket  *Basket // pointer to the current basket

	tree   Tree        // tree header
	mother *tbranch    // top-level parent branch in the tree
	parent *tbranch    // parent branch
	dir    *tdirectory // directory where this branch's buffers are stored
}

func (b *tbranch) Name() string {
	return b.named.Name()
}

func (b *tbranch) Title() string {
	return b.named.Title()
}

func (b *tbranch) Class() string {
	return "TBranch"
}

func (b *tbranch) SetTree(t Tree) {
	b.tree = t
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (b *tbranch) UnmarshalROOT(r *RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	beg := r.Pos()
	vers, pos, bcnt := r.ReadVersion()

	b.tree = nil
	b.currbasket = nil
	b.firstbasket = -1
	b.nextbasket = -1

	if vers < 12 {
		panic(fmt.Errorf("rootio: too old TBanch version (%d<12)", vers))
	}

	if err := b.named.UnmarshalROOT(r); err != nil {
		return err
	}

	if err := b.attfill.UnmarshalROOT(r); err != nil {
		return err
	}

	b.compress = int(r.ReadI32())
	b.basketSize = int(r.ReadI32())
	b.entryOffsetLen = int(r.ReadI32())
	b.writeBasket = int(r.ReadI32())
	b.entryNumber = r.ReadI64()
	b.offset = int(r.ReadI32())
	b.maxBaskets = int(r.ReadI32())
	b.splitLevel = int(r.ReadI32())
	b.entries = r.ReadI64()
	b.firstEntry = r.ReadI64()
	b.totBytes = r.ReadI64()
	b.zipBytes = r.ReadI64()

	{
		var branches objarray
		if err := branches.UnmarshalROOT(r); err != nil {
			r.err = err
			return r.err
		}
		b.branches = make([]tbranch, branches.last+1)
		for i := range b.branches {
			br := branches.At(i).(*tbranch)
			b.branches[i] = *br
		}
	}

	{
		var leaves objarray
		if err := leaves.UnmarshalROOT(r); err != nil {
			r.err = err
			return r.err
		}
		b.leaves = make([]Leaf, leaves.last+1)
		for i := range b.leaves {
			leaf := leaves.At(i).(Leaf)
			b.leaves[i] = leaf
		}
	}
	{
		var baskets objarray
		if err := baskets.UnmarshalROOT(r); err != nil {
			r.err = err
			return r.err
		}
		b.baskets = make([]Basket, baskets.last+1)
		for i := range b.baskets {
			bk := baskets.At(i).(*Basket)
			b.baskets[i] = *bk
		}
	}

	b.basketBytes = make([]int32, b.maxBaskets)
	b.basketEntry = make([]int64, b.maxBaskets)
	b.basketSeek = make([]int64, b.maxBaskets)

	/*isArray*/ _ = r.ReadI8()
	copy(b.basketBytes, r.ReadFastArrayI32(b.maxBaskets))

	/*isArray*/ _ = r.ReadI8()
	_ = r.ReadFastArrayI64(b.maxBaskets)

	/*isArray*/ _ = r.ReadI8()
	copy(b.basketSeek, r.ReadFastArrayI64(b.maxBaskets))

	b.fname = r.ReadString()

	r.CheckByteCount(pos, bcnt, beg, "TBranch")

	if b.splitLevel == 0 && len(b.branches) > 0 {
		b.splitLevel = 1
	}

	return r.Err()
}

func init() {
	f := func() reflect.Value {
		o := &tbranch{}
		return reflect.ValueOf(o)
	}
	Factory.add("TBranch", f)
	Factory.add("*rootio.tbranch", f)
}

var _ Object = (*tbranch)(nil)
var _ Named = (*tbranch)(nil)
var _ Branch = (*tbranch)(nil)
var _ ROOTUnmarshaler = (*tbranch)(nil)
