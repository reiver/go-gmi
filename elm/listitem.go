package elm

import (
	"fmt"
	"strings"
)

// ListItem represents a Gemtext list-item.
//
// If a line in a Gemtext file starts with a "* ", then it is a Gemtext list-item.
type ListItem struct {
	value string
	eol   string
}

func (receiver ListItem) GoString() string {
	return fmt.Sprintf("elm.Create(%q, %q, %q)", receiver.MagicPrefix(), receiver.value, receiver.eol)
}

func (receiver ListItem) MagicPrefix() string {
		return "* "
}

func (receiver *ListItem) Set(value string, eol string) {
	if nil == receiver {
		return
	}

	receiver.value = value
	receiver.eol   = eol
}

func (receiver ListItem) String() string {

	var storage strings.Builder

	storage.WriteString(receiver.MagicPrefix())
	storage.WriteString(receiver.value)
	storage.WriteString(receiver.eol)

	return storage.String()
}
