// Prints all information of an MKV file
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/5rahim/gomkv"
)

type MyParser struct {
}

func (p *MyParser) HandleMasterBegin(id gomkv.ElementID, info gomkv.ElementInfo) (bool, error) {
	switch id {
	default:
		fmt.Printf("%s- %s:\n", indent(info.Level), gomkv.NameForElementID(id))
		return true, nil
	}
}

func (p *MyParser) HandleMasterEnd(id gomkv.ElementID, info gomkv.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleString(id gomkv.ElementID, value string, info gomkv.ElementInfo) error {
	fmt.Printf("%s- %v: %q\n", indent(info.Level), gomkv.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleInteger(id gomkv.ElementID, value int64, info gomkv.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), gomkv.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleFloat(id gomkv.ElementID, value float64, info gomkv.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), gomkv.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleDate(id gomkv.ElementID, value time.Time, info gomkv.ElementInfo) error {
	fmt.Printf("%s- %v: %v\n", indent(info.Level), gomkv.NameForElementID(id), value)
	return nil
}

func (p *MyParser) HandleBinary(id gomkv.ElementID, value []byte, info gomkv.ElementInfo) error {
	switch id {
	case gomkv.SeekIDElement:
		fmt.Printf("%s- %v: %x\n", indent(info.Level), gomkv.NameForElementID(id), value)
	default:
		fmt.Printf("%s- %v: <binary> (%d)\n", indent(info.Level), gomkv.NameForElementID(id), info.Size)
	}
	return nil
}

func main() {
	handler := MyParser{}
	err := gomkv.ParsePath(os.Args[1], &handler)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}

func indent(n int) string {
	return strings.Repeat("  ", n)
}
