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

package pdf

import (
	"io"
	"io/ioutil"
)

type File struct {
	ref int
	trk struct {
		leafs []*Leaf
	}
}

func Open(r io.Reader) *File {

	var err error
	var res []byte

	this := new(File)

	name := uniq()

	if res, err = ioutil.ReadAll(r); err != nil {
		panic(err)
	}

	if err = PDFLib.CreatePvf(name, res, "copy"); err != nil {
		panic(err)
	}

	if this.ref, err = PDFLib.OpenPdiDocument(name, ""); err != nil {
		panic(err)
	}

	if err = PDFLib.DeletePvf(name); err != nil {
		panic(err)
	}

	return this

}

func (this *File) Close() {

	for i := len(this.trk.leafs) - 1; i >= 0; i-- {
		this.trk.leafs[i].Close()
		this.trk.leafs = this.trk.leafs[:len(this.trk.leafs)-1]
	}

	if err := PDFLib.ClosePdiDocument(this.ref); err != nil {
		panic(err)
	}

}

func (this *File) Page(p int, o Options) *Leaf {

	leaf := NewLeaf(this, p, o)

	this.trk.leafs = append(this.trk.leafs, leaf)

	return leaf

}

func (this *File) Pages() int {

	val, err := PDFLib.PcosGetNumber(this.ref, "length:pages")
	if err != nil {
		panic(err)
	}

	return int(val)

}
