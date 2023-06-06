//go:build ignore
// +build ignore

package main

import "testing"

func TestGenerate(t *testing.T) {
	css, err := FetchCss("https://raw.githubusercontent.com/ryanoasis/nerd-fonts/master/css/nerd-fonts-generated.css")
	if err != nil {
		t.Fatal(err)
	}
	icons := ParseCss(css)
	for k, v := range icons {
		if len(v) != 7 {
			t.Fatal(k, v)
		}
	}
}
