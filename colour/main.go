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

package colour

import (
	"fmt"

	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

type Colour struct {
	col colorful.Color
	c   float64
	m   float64
	y   float64
	k   float64
}

func Hex(h string) *Colour {
	col, err := colorful.Hex(h)
	if err != nil {
		panic(err)
	}
	return &Colour{col: col}
}

func Rgb(r, g, b float64) *Colour {
	return &Colour{
		col: colorful.Color{
			R: r / 255,
			G: g / 255,
			B: b / 255,
		},
	}
}

func Hcl(h, c, l float64) *Colour {
	return &Colour{
		col: colorful.Hcl(h, c/100, l/100),
	}
}

func Hsl(h, s, l float64) *Colour {
	return &Colour{
		col: colorful.Hsl(h, s/100, l/100),
	}
}

func Hsv(h, s, v float64) *Colour {
	return &Colour{
		col: colorful.Hsv(h, s/100, v/100),
	}
}

func Lab(l, a, b float64) *Colour {
	return &Colour{
		col: colorful.Lab(l/100, a/100, b/100),
	}
}

func Luv(l, u, v float64) *Colour {
	return &Colour{
		col: colorful.Luv(l/100, u/100, v/100),
	}
}

func Xyy(x, y, z float64) *Colour {
	return &Colour{
		col: colorful.Xyy(x/100, y/100, z/100),
	}
}

func Xyz(x, y, z float64) *Colour {
	return &Colour{
		col: colorful.Xyz(x/100, y/100, z/100),
	}
}

func Cmyk(c, m, y, k float64) *Colour {
	return &Colour{
		col: colorful.Color{
			R: float64((1 - c/100) * (1 - k/100)),
			G: float64((1 - m/100) * (1 - k/100)),
			B: float64((1 - y/100) * (1 - k/100)),
		},
		c: c,
		m: m,
		y: y,
		k: k,
	}
}

func Load(kind string, vals ...float64) *Colour {

	switch kind {
	default:
		panic("Incorrect type")
	case "rgb":
		switch len(vals) {
		case 3:
			return Rgb(vals[0], vals[1], vals[2])
		default:
			panic("Incorrect number of arguments for RGB colour")
		}
	case "cmyk":
		switch len(vals) {
		case 3:
			return Cmyk(vals[0], vals[1], vals[2], vals[3])
		default:
			panic("Incorrect number of arguments for CMYK colour")
		}
	}

}

func (this *Colour) IsCmyk() bool {
	return this.c != 0 && this.m != 0 && this.y != 0 && this.k != 0
}

func (this *Colour) Hex() string {
	return this.col.Hex()
}

func (this *Colour) Rgb() [3]uint8 {
	r, g, b := this.col.RGB255()
	return [3]uint8{r, g, b}
}

func (this *Colour) Rgba() [4]uint8 {
	r, g, b := this.col.RGB255()
	return [4]uint8{r, g, b, 255}
}

func (this *Colour) Hcl() [3]float64 {
	h, c, l := this.col.Hcl()
	return [3]float64{fixed(h), fixed(c * 100), fixed(l * 100)}
}

func (this *Colour) Hsl() [3]float64 {
	h, s, l := this.col.Hsl()
	return [3]float64{fixed(h), fixed(s * 100), fixed(l * 100)}
}

func (this *Colour) Hsv() [3]float64 {
	h, s, v := this.col.Hsv()
	return [3]float64{fixed(h), fixed(s * 100), fixed(v * 100)}
}

func (this *Colour) Lab() [3]float64 {
	l, a, b := this.col.Lab()
	return [3]float64{fixed(l * 100), fixed(a * 100), fixed(b * 100)}
}

func (this *Colour) Luv() [3]float64 {
	l, u, v := this.col.Luv()
	return [3]float64{fixed(l * 100), fixed(u * 100), fixed(v * 100)}
}

func (this *Colour) Xyy() [3]float64 {
	x, y, z := this.col.Xyy()
	return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
}

func (this *Colour) Xyz() [3]float64 {
	x, y, z := this.col.Xyz()
	return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
}

func (this *Colour) Cmyk() [4]float64 {
	switch this.IsCmyk() {
	case true:
		return [4]float64{this.c, this.m, this.y, this.k}
	default:
		r, g, b := this.col.RGB255()
		c, m, y, k := color.RGBToCMYK(r, g, b)
		return [4]float64{
			fixed(float64(c) / 2.55),
			fixed(float64(m) / 2.55),
			fixed(float64(y) / 2.55),
			fixed(float64(k) / 2.55),
		}
	}
}

func (this *Colour) ToRGBA() color.RGBA {
	r, g, b := this.col.RGB255()
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func (this *Colour) ToCMYK() color.CMYK {
	r, g, b := this.col.RGB255()
	c, m, y, k := color.RGBToCMYK(r, g, b)
	return color.CMYK{C: c, M: m, Y: y, K: k}
}

func (this *Colour) String() string {
	switch this.IsCmyk() {
	default:
		ru, gu, bu := this.col.RGB255()
		rp, gp, bp := float64(ru), float64(gu), float64(bu)
		return fmt.Sprintf("rgb %v %v %v", fixed(rp/255), fixed(gp/255), fixed(bp/255))
	case true:
		cu, mu, yu, ku := this.c, this.m, this.y, this.k
		cp, mp, yp, kp := float64(cu), float64(mu), float64(yu), float64(ku)
		return fmt.Sprintf("cmyk %v %v %v %v", fixed(cp/100), fixed(mp/100), fixed(yp/100), fixed(kp/100))
	}
}

func (this *Colour) Output() (string, float64, float64, float64, float64) {
	switch this.IsCmyk() {
	default:
		ru, gu, bu := this.col.RGB255()
		rp, gp, bp := float64(ru), float64(gu), float64(bu)
		return "rgb", fixed(rp / 255), fixed(gp / 255), fixed(bp / 255), 0
	case true:
		cu, mu, yu, ku := this.c, this.m, this.y, this.k
		cp, mp, yp, kp := float64(cu), float64(mu), float64(yu), float64(ku)
		return "cmyk", fixed(cp / 100), fixed(mp / 100), fixed(yp / 100), fixed(kp / 100)
	}
}
