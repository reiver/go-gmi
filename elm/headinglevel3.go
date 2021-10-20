package elm

import (
	"fmt"
	"strings"
)

// HeadingLevel3 represents a Gemtext heading-level-3.
//
// If a line in a Gemtext file starts with a "## ", then it is a Gemtext heading-level-3.
type HeadingLevel3 struct {
	value string
	eol   string
}

func (receiver HeadingLevel3) GoString() string {
	return fmt.Sprintf("elm.Create(%q, %q, %q)", receiver.MagicPrefix(), receiver.value, receiver.eol)
}

func (receiver HeadingLevel3) MagicPrefix() string {
	return "### "
}

func (receiver *HeadingLevel3) Set(value string, eol string) {
	if nil == receiver {
		return
	}

	receiver.value = value
	receiver.eol   = eol
}

func (receiver HeadingLevel3) String() string {

	var storage strings.Builder

	storage.WriteString(receiver.MagicPrefix())
	storage.WriteString(receiver.value)
	storage.WriteString(receiver.eol)

	return storage.String()
}
