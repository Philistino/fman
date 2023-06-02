package entryinfo

import "testing"

// func TestGetFilePreview(t *testing.T) {

// 	file := "Line 0\nLine 1\nLine 2\nLine 3\nLine 4\nLine 5\nLine 6\nLine 7\nLine 8\nLine 9\nLine 10\nLine 11\nLine 12\nLine 13\nLine 14\nLine 15\nLine 16\nLine 17\nLine 18\nLine 19"

// 	got, err := getFilePreviewFunc(context.Background(), strings.NewReader(file), 10, 20)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println("got", got)
// 	t.Error()

// }

func TestHighlightSyntax(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc     string
		name     string
		preview  string
		expected string
	}{
		{
			desc:     "empty",
			name:     "",
			preview:  "",
			expected: "",
		},
		{
			desc:    "go",
			name:    "go",
			preview: "package main\n\nfunc main()\n{\n}\n",
			expected: `[1m[37mpackage main[0m[1m[37m
[0m[1m[37m
[0m[1m[37mfunc main()[0m[1m[37m
[0m[1m[37m{[0m[1m[37m
[0m[1m[37m}[0m[1m[37m
[0m`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := highlightSyntax(tC.name, tC.preview)
			if got != tC.expected {
				t.Errorf("expecting %s, got %v", tC.expected, got)
			}
		})
	}
}
