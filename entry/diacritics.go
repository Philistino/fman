// MIT License

// Copyright (c) 2016 Gökçehan Kara

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package entry

import (
	"strconv"
	"unicode"
)

var normMap map[rune]rune

func init() {
	normMap = make(map[rune]rune)

	// (not only) european
	appendTransliterate(
		"ěřůøĉĝĥĵŝŭèùÿėįųāēīūļķņģőűëïąćęłńśźżõșțčďĺľňŕšťýžéíñóúüåäöçîşûğăâđêôơưáàãảạ",
		"eruocghjsueuyeiuaeiulkngoueiacelnszzostcdllnrstyzeinouuaaocisugaadeoouaaaaa",
	)

	// Vietnamese
	appendTransliterate(
		"áạàảãăắặằẳẵâấậầẩẫéẹèẻẽêếệềểễiíịìỉĩoóọòỏõôốộồổỗơớợờởỡúụùủũưứựừửữyýỵỳỷỹđ",
		"aaaaaaaaaaaaaaaaaeeeeeeeeeeeiiiiiioooooooooooooooooouuuuuuuuuuuyyyyyyd",
	)
}

func appendTransliterate(base, norm string) {
	normRunes := []rune(norm)
	baseRunes := []rune(base)

	lenNorm := len(normRunes)
	lenBase := len(baseRunes)
	if lenNorm != lenBase {
		panic("Base and normalized strings have differend length: base=" + strconv.Itoa(lenBase) + ", norm=" + strconv.Itoa(lenNorm)) // programmer error in constant length
	}

	for i := 0; i < lenBase; i++ {
		normMap[baseRunes[i]] = normRunes[i]

		baseUpper := unicode.ToUpper(baseRunes[i])
		normUpper := unicode.ToUpper(normRunes[i])

		normMap[baseUpper] = normUpper
	}
}

// Remove diacritics and make lowercase.
func removeDiacritics(baseString string) string {
	var normalizedRunes []rune
	for _, baseRune := range baseString {
		if normRune, ok := normMap[baseRune]; ok {
			normalizedRunes = append(normalizedRunes, normRune)
		} else {
			normalizedRunes = append(normalizedRunes, baseRune)
		}
	}
	return string(normalizedRunes)
}
