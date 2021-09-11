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
	"io"
	"io/ioutil"

	"github.com/abcum/pdf/colour"
)

type Hold struct {
	doc *Document
	opt Options
	ref int
}

func NewHold(doc *Document, r io.Reader, o Options) *Hold {

	var err error
	var res []byte

	this := &Hold{
		doc: doc,
		opt: o,
	}

	name := uniq()

	opts := o.cull(loadOptions)

	if res, err = ioutil.ReadAll(r); err != nil {
		panic(err)
	}

	if err = PDFLib.CreatePvf(name, res, "copy"); err != nil {
		panic(err)
	}

	if this.ref, err = PDFLib.LoadImage("auto", name, opts); err != nil {
		panic(err)
	}

	if err = PDFLib.DeletePvf(name); err != nil {
		panic(err)
	}

	return this

}

func (this *Hold) Close() {

	if err := PDFLib.CloseImage(this.ref); err != nil {
		panic(err)
	}

}

func (this *Hold) Place(x, y, w, h float64, o Options) *Hold {

	opts := o.cull(imageOptions)

	if val, ok := o["backgroundcolor"]; ok {
		if col, ok := val.(*colour.Colour); ok {
			k, v1, v2, v3, v4 := col.Output()
			PDFLib.SetColor("fill", k, v1, v2, v3, v4)
			PDFLib.Rect(x, y, w, h)
			PDFLib.Fill()
		}
	}

	if err := PDFLib.FitImage(this.ref, x, y, opts); err != nil {
		panic(err)
	}

	return this

}
