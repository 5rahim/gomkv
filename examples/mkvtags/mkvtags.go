// Prints tags of an MKV file
package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/5rahim/gomkv"
)

type MyParser struct {
	gomkv.DefaultHandler

	title *string
}

func (p *MyParser) HandleString(id gomkv.ElementID, value string, info gomkv.ElementInfo) error {
	switch id {
	case gomkv.TitleElement:
		p.title = &value
	}
	return nil
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	defer file.Close()
	titleh := MyParser{}
	tagsh := gomkv.NewTagsHandler()
	err = gomkv.ParseSections(file, gomkv.NewHandlerChain(&titleh, tagsh), gomkv.InfoElement, gomkv.TagsElement)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	// Print (sorted) tags
	if titleh.title != nil {
		fmt.Printf("- title: %q\n", *titleh.title)
	}
	tags := tagsh.Tags()
	var tagNames []string
	for tagName := range tags {
		tagNames = append(tagNames, tagName)
	}
	sort.Strings(tagNames)
	for _, tagName := range tagNames {
		fmt.Printf("- %s: %q\n", tagName, tags[tagName])
	}
}
