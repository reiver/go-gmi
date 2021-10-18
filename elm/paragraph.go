package elm

import (
	"fmt"
	"strings"
)

// Paragraph represents a Gemtext paragraph.
// I.e., what the sometimes gets called “text”.
//
// If a line in a Gemtext file isn't one of the other Gemtext line types,
// then it defaults to a Gemtext paragraph.
type Paragraph struct {
	value string
	eol   string
}

func (receiver Paragraph) GoString() string {
	return fmt.Sprintf("elm.Create(%q, %q, %q)", receiver.MagicPrefix(), receiver.value, receiver.eol)
}

func (receiver Paragraph) MagicPrefix() string {
	return ""
}

func (receiver *Paragraph) Set(value string, eol string) {
	if nil == receiver {
		return
	}

	receiver.value = value
	receiver.eol   = eol
}

func (receiver Paragraph) String() string {

	var storage strings.Builder

	storage.WriteString(receiver.value)
	storage.WriteString(receiver.eol)

	return storage.String()
}
