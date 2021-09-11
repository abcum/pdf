// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// go:build cgo

package pdf

import (
	"fmt"
)

type Leaf struct {
	doc *File
	pge int
	ref int
}

func NewLeaf(doc *File, p int, o Options) *Leaf {

	var err error

	this := &Leaf{
		doc: doc,
		pge: p,
	}

	opts := o.cull(leafOptions)

	if this.ref, err = PDFLib.OpenPdiPage(doc.ref, p, opts); err != nil {
		panic(err)
	}

	return this

}

func (this *Leaf) Close() {

	if err := PDFLib.ClosePdiPage(this.ref); err != nil {
		panic(err)
	}

}

func (this *Leaf) Width() float64 {

	val, err := PDFLib.PcosGetNumber(this.doc.ref, fmt.Sprintf("pages[%d]/width", this.pge-1))
	if err != nil {
		panic(err)
	}

	return val

}

func (this *Leaf) Height() float64 {

	val, err := PDFLib.PcosGetNumber(this.doc.ref, fmt.Sprintf("pages[%d]/height", this.pge-1))
	if err != nil {
		panic(err)
	}

	return val

}
