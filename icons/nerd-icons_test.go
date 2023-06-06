package icons

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/Philistino/fman/icons/nerdicons"
	"github.com/mattn/go-runewidth"
)

func TestGen(t *testing.T) {

	getGlyph := func(name string) string {
		glyph, ok := nerdicons.Icons[name]
		if !ok {
			log.Println(name)
			return ""
		}
		return glyph
	}
	output := make([]iconInfo2, 0, len(iconsByExt))
	for _, icon := range iconsByExt {
		icon.Glyph = getGlyph(icon.Name)
		_, ok := iconExt[icon.Id]
		if !ok {
			// log.Println(icon.Id)
			output = append(output, icon)
		}
		// log.Println(icon.Glyph, icon.Id, icon.Name)
		w := runewidth.StringWidth(icon.Glyph)
		if w > 1 {
			t.Error(icon.Glyph, icon.Name, w)
		}
	}
	// log.Println(len(output))
	// err := WriteStructsToJSONFile("icons.json", output)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Error()
}

func WriteStructsToJSONFile(filename string, structs interface{}) error {
	// Create a new file writer.
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Marshal the structs to JSON.
	data, err := json.MarshalIndent(structs, "", "\t")
	if err != nil {
		return err
	}

	// Write the JSON data to the file.
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
