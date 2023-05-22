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
	"testing"
)

// typical czech test sentence ;-)
const baseTestString = "Příliš žluťoučký kůň příšerně úpěl ďábelské ódy"

func TestRemoveDiacritics(t *testing.T) {
	testStr := baseTestString
	expStr := "Prilis zlutoucky kun priserne upel dabelske ody"
	checkRemoveDiacritics(testStr, expStr, t)

	// other accents (non comlete, but all I founded)
	testStr = "áéíóúýčďěňřšťžůåøĉĝĥĵŝŭšžõäöüàâçéèêëîïôùûüÿžščćđáéíóúąęėįųūčšžāēīūčšžļķņģáéíóúöüőűäöüëïąćęłńóśźżáàãâçéêíóõôăâîșțáäčďéíĺľňóôŕšťúýžáéíñóúüåäöâçîşûğăâđêôơưáàãảạ"
	expStr = "aeiouycdenrstzuaocghjsuszoaouaaceeeeiiouuuyzsccdaeiouaeeiuucszaeiucszlkngaeiouououaoueiacelnoszzaaaaceeioooaaistaacdeillnoorstuyzaeinouuaaoacisugaadeoouaaaaa"
	checkRemoveDiacritics(testStr, expStr, t)

	testStr = "ÁÉÍÓÚÝČĎĚŇŘŠŤŽŮÅØĈĜĤĴŜŬŠŽÕÄÖÜÀÂÇÉÈÊËÎÏÔÙÛÜŸŽŠČĆĐÁÉÍÓÚĄĘĖĮŲŪČŠŽĀĒĪŪČŠŽĻĶŅĢÁÉÍÓÚÖÜŐŰÄÖÜËÏĄĆĘŁŃÓŚŹŻÁÀÃÂÇÉÊÍÓÕÔĂÂÎȘȚÁÄČĎÉÍĹĽŇÓÔŔŠŤÚÝŽÁÉÍÑÓÚÜÅÄÖÂÇÎŞÛĞĂÂĐÊÔƠƯÁÀÃẢẠ"
	expStr = "AEIOUYCDENRSTZUAOCGHJSUSZOAOUAACEEEEIIOUUUYZSCCDAEIOUAEEIUUCSZAEIUCSZLKNGAEIOUOUOUAOUEIACELNOSZZAAAACEEIOOOAAISTAACDEILLNOORSTUYZAEINOUUAAOACISUGAADEOOUAAAAA"
	checkRemoveDiacritics(testStr, expStr, t)

	testStr = "áạàảãăắặằẳẵâấậầẩẫéẹèẻẽêếệềểễiíịìỉĩoóọòỏõôốộồổỗơớợờởỡúụùủũưứựừửữyýỵỳỷỹđ"
	expStr = "aaaaaaaaaaaaaaaaaeeeeeeeeeeeiiiiiioooooooooooooooooouuuuuuuuuuuyyyyyyd"
	checkRemoveDiacritics(testStr, expStr, t)

	testStr = "ÁẠÀẢÃĂẮẶẰẲẴÂẤẬẦẨẪÉẸÈẺẼÊẾỆỀỂỄÍỊÌỈĨÓỌÒỎÕÔỐỘỒỔỖƠỚỢỜỞỠÚỤÙỦŨƯỨỰỪỬỮÝỴỲỶỸĐ"
	expStr = "AAAAAAAAAAAAAAAAAEEEEEEEEEEEIIIIIOOOOOOOOOOOOOOOOOUUUUUUUUUUUYYYYYD"
	checkRemoveDiacritics(testStr, expStr, t)
}

func checkRemoveDiacritics(testStr string, expStr string, t *testing.T) {
	resultStr := removeDiacritics(testStr)
	if resultStr != expStr {
		t.Errorf("at input '%v' expected '%v' but got '%v'", testStr, expStr, resultStr)
	}
}

func TestSearchSettings(t *testing.T) {
	runSearch(t, true, false, true, true, "Veřejný", "vere", true)

	runSearch(t, true, false, true, false, baseTestString, "Zlutoucky", true)
	runSearch(t, true, false, true, false, baseTestString, "zlutoucky", true)
	runSearch(t, true, true, true, false, baseTestString, "Zlutoucky", false)
	runSearch(t, true, true, true, true, baseTestString, "zlutoucky", true)

	runSearch(t, false, false, true, false, baseTestString, "žlutoucky", true)
	runSearch(t, false, false, true, false, baseTestString, "Žlutoucky", false)

	runSearch(t, false, false, true, true, baseTestString, "žluťoučký", true)
	runSearch(t, false, false, true, false, baseTestString, "žluťoučký", true)
	runSearch(t, false, false, false, false, baseTestString, "žluťoučký", true)
	runSearch(t, false, false, false, false, baseTestString, "zlutoucky", false)
	runSearch(t, false, false, true, true, baseTestString, "zlutoucky", true)
}

func runSearch(t *testing.T, ignorecase, smartcase, ignorediacritics, smartdiacritics bool, base, pattern string, expected bool) {
	matched, _ := searchMatch(base, pattern, ignorecase, ignorediacritics, smartcase, smartdiacritics, false, false)
	if matched != expected {
		t.Errorf("False search for" +
			" ignorecase = " + strconv.FormatBool(ignorecase) + ", " +
			" smartcase = " + strconv.FormatBool(smartcase) + ", " +
			" ignoredia = " + strconv.FormatBool(ignorediacritics) + ", " +
			" smartdia = " + strconv.FormatBool(smartdiacritics) + ", ")
	}
}
