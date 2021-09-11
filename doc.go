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
)

type Document struct {
	val int
	fnt map[string]*Font
	sch func(string) io.Reader
	trk struct {
		fonts []*Font // Fonts
		pages []*Page // Pages
		flows []*Flow // Flows
		holds []*Hold // Holds
	}
}

func New(o Options) *Document {

	doc := new(Document)

	opt := o.cull(documentOptions)

	doc.fnt = make(map[string]*Font)

	if _, err := PDFLib.BeginDocument("", opt); err != nil {
		panic(err)
	}

	return doc

}

func (this *Document) Load(n string) *Font {

	if font, ok := this.fnt[n]; ok {
		return font
	}

	return this.Font(n, fetcher(n))

}

func (this *Document) Font(n string, v io.Reader) *Font {

	if font, ok := this.fnt[n]; ok {
		return font
	}

	this.fnt[n] = NewFont(this, n, v)

	this.trk.fonts = append(this.trk.fonts, this.fnt[n])

	return this.fnt[n]

}

func (this *Document) Page(w, h float64, o Options) *Page {

	for i := len(this.trk.pages) - 1; i >= 0; i-- {
		this.trk.pages[i].Close()
		this.trk.pages = this.trk.pages[:len(this.trk.pages)-1]
	}

	page := NewPage(this, w, h, o)

	this.trk.pages = append(this.trk.pages, page)

	return page

}

func (this *Document) Flow(t string, o Options) *Flow {

	flow := NewFlow(this, t, o)

	this.trk.flows = append(this.trk.flows, flow)

	return flow

}

func (this *Document) Text(t string, x, y float64, o Options) *Document {

	opts := o.cull(textOptions)

	if val, ok := o["fontname"].(string); ok {
		this.Load(val)
	}

	if err := PDFLib.FitTextline(t, x, y, opts); err != nil {
		panic(err)
	}

	return this

}

func (this *Document) Image(r io.Reader, x, y, w, h float64, o Options) *Document {

	hold := NewHold(this, r, o)

	hold.Place(x, y, w, h, o)

	hold.Close()

	return this

}

func (this *Document) Embed(r io.Reader, o Options) *Hold {

	hold := NewHold(this, r, o)

	this.trk.holds = append(this.trk.holds, hold)

	return hold

}

func (this *Document) Place(l *Leaf, x, y float64, o Options) *Document {

	opts := o.cull(pageOptions)

	if err := PDFLib.FitPdiPage(l.ref, x, y, opts); err != nil {
		panic(err)
	}

	return this

}

func (this *Document) Pipe(w io.Writer) {

	var err error
	var out []byte

	// Close pages

	for i := len(this.trk.pages) - 1; i >= 0; i-- {
		this.trk.pages[i].Close()
		this.trk.pages = this.trk.pages[:len(this.trk.pages)-1]
	}

	// Close flows

	for i := len(this.trk.flows) - 1; i >= 0; i-- {
		this.trk.flows[i].Close()
		this.trk.flows = this.trk.flows[:len(this.trk.flows)-1]
	}

	// Close loads

	for i := len(this.trk.holds) - 1; i >= 0; i-- {
		this.trk.holds[i].Close()
		this.trk.holds = this.trk.holds[:len(this.trk.holds)-1]
	}

	// Close fonts

	for i := len(this.trk.fonts) - 1; i >= 0; i-- {
		this.trk.fonts[i].Close()
		this.trk.fonts = this.trk.fonts[:len(this.trk.fonts)-1]
	}

	// End document

	if err = PDFLib.EndDocument(""); err != nil {
		panic(err)
	}

	// Get document

	if out, _, err = PDFLib.GetBuffer(); err != nil {
		panic(err)
	}

	// Write output

	w.Write(out)

}
