package main

import (
	"fmt"
)

// Style is the heavyweight object that has to be referred to by reference to conserve memory and disk space
type Style struct {
	Font     string
	FontSize int
	Colour   string
}

// Section is a portion of the document that has a single style applied to it
type Section struct {
	Style *Style
	Data  string
}

// String is the Section Stringer
func (s Section) String() string {
	return fmt.Sprintf("{Style: %+v Data: %s", *s.Style, s.Data)
}

// Document is the entire document, consisting only of sections to guarantee styles are applied as intended
type Document struct {
	Name     string
	Sections []Section
}

func main() {
	monacoLargeRed := &Style{
		Font:     "monaco",
		FontSize: 18,
		Colour:   "red",
	}

	timesSmallBlack := &Style{
		Font:     "times",
		FontSize: 9,
		Colour:   "black",
	}

	section1 := Section{
		Style: timesSmallBlack,
		Data:  "The times small black section",
	}

	section2 := Section{
		Style: monacoLargeRed,
		Data:  "The monaco large red section",
	}

	document := Document{
		Name: "The document to end all documents",
		Sections: []Section{
			section1,
			section2,
		},
	}

	fmt.Printf("%+v\n", document)
}
