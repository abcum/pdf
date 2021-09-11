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

	"github.com/abcum/pdf/colour"
)

var documentOptions = map[string]bool{
	"inmemory":       true,
	"linearize":      true,
	"masterpassword": true,
	"optimize":       true,
	"permissions":    true,
}

var pageOptions = map[string]bool{
	"artbox":     true,
	"bleedbox":   true,
	"cloneboxes": true,
	"cropbox":    true,
	"mediabox":   true,
	"rotate":     true,
}

var leafOptions = map[string]bool{
	"cloneboxes": true,
}

var textOptions = map[string]bool{
	"adjustmethod":    true,
	"alignment":       true,
	"avoidbreak":      true,
	"avoidemptybegin": true,
	"boxsize":         true,
	"charspacing":     true,
	"embedding":       true,
	"encoding":        true,
	"fallbackfonts":   true,
	"fillcolor":       true,
	"fontname":        true,
	"fontsize":        true,
	"glyphcheck":      true,
	"hortabmethod":    true,
	"hortabsize":      true,
	"leading":         true,
	"matchbox":        true,
	"replacementchar": true,
	"ruler":           true,
	"shrinklimit":     true,
	"stamp":           true,
	"strokecolor":     true,
	"strokewidth":     true,
	"subsetting":      true,
	"tabalignment":    true,
	"textrendering":   true,
	"wordspacing":     true,
}

var flowOptions = map[string]bool{
	"firstlinedist": true,
	"fitmethod":     true,
	"lastlinedist":  true,
	"maxlines":      true,
	"orientate":     true,
	"showborder":    true,
	"verticalalign": true,
}

var loadOptions = map[string]bool{
	"honorclippingpath": true,
	"honoriccprofile":   true,
	"iccprofile":        true,
	"ignoremask":        true,
	"ignoreorientation": true,
	"invert":            true,
	"passthrough":       true,
}

var placeOptions = map[string]bool{
	"boxsize":    true,
	"dpi":        true,
	"margin":     true,
	"matchbox":   true,
	"orientate":  true,
	"position":   true,
	"rotate":     true,
	"showborder": true,
}

var imageOptions = map[string]bool{
	"boxsize":    true,
	"fitmethod":  true,
	"orientate":  true,
	"position":   true,
	"scale":      true,
	"showborder": true,
}

var tcellOptions = map[string]bool{
	"avoidwordsplitting": true,
	"colspan":            true,
	"colwidth":           true,
	"margin":             true,
	"marginbottom":       true,
	"marginleft":         true,
	"marginright":        true,
	"marginto":           true,
	"rowheight":          true,
	"rowspan":            true,
}

var tableOptions = map[string]bool{
	"fill":      true,
	"showcells": true,
	"showgrid":  true,
	"stroke":    true,
}

var stateOptions = map[string]bool{
	"blendmode":     true,
	"borderwidth":   true,
	"fillcolor":     true,
	"fillrule":      true,
	"linecap":       true,
	"linejoin":      true,
	"linewidth":     true,
	"miterlimit":    true,
	"opacityfill":   true,
	"opacitystroke": true,
	"shading":       true,
	"strokecolor":   true,
}

type Options map[string]interface{}

func (opts Options) cull(allowed map[string]bool) (out string) {
	for k, v := range opts {
		if _, ok := allowed[k]; ok {
			switch v := v.(type) {
			case nil:
				break
			default:
				out = fmt.Sprintf("%s %s={%v}", out, k, v)
			case *colour.Colour:
				out = fmt.Sprintf("%s %s={%v}", out, k, v.String())
			}
		}
	}
	return out
}
