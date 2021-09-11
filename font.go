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
	"fmt"
	"io"
	"io/ioutil"
)

type Font struct {
	doc *Document
}

func NewFont(doc *Document, n string, r io.Reader) *Font {

	var err error
	var res []byte

	font := &Font{
		doc: doc,
	}

	name := fmt.Sprintf("%s.ttf", n)

	if res, err = ioutil.ReadAll(r); err != nil {
		panic(err)
	}

	if err = PDFLib.CreatePvf(name, res, "copy"); err != nil {
		panic(err)
	}

	if _, err = PDFLib.LoadFont(n, "unicode", "embedding={true} subsetting={true}"); err != nil {
		panic(err)
	}

	if err = PDFLib.DeletePvf(name); err != nil {
		panic(err)
	}

	return font

}

func (this *Font) Close() {

	// Don't do anything as the font will be deleted automatically

}
