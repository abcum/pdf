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

import "github.com/abcum/pdf/colour"

type Flow struct {
	doc *Document
	opt Options
	txt string
	ref int
}

func NewFlow(doc *Document, t string, o Options) *Flow {

	this := &Flow{
		doc: doc,
		opt: o,
		txt: t,
	}

	opts := o.cull(textOptions)

	if val, ok := o["fontname"].(string); ok {
		this.doc.Load(val)
	}

	this.ref, _ = PDFLib.CreateTextflow(t, opts)

	return this

}

func (this *Flow) Close() {

	if err := PDFLib.DeleteTextflow(this.ref); err != nil {
		panic(err)
	}

}

func (this *Flow) Box(x, y, r, t float64, o Options) *Flow {

	opts := o.cull(flowOptions)

	if val, ok := o["backgroundcolor"]; ok {
		if col, ok := val.(*colour.Colour); ok {
			k, v1, v2, v3, v4 := col.Output()
			PDFLib.SetColor("fill", k, v1, v2, v3, v4)
			PDFLib.Rect(x, y, r-x, t-y)
			PDFLib.Fill()
		}
	}

	if _, err := PDFLib.FitTextflow(this.ref, x, y, r, t, opts); err != nil {
		panic(err)
	}

	return this

}
