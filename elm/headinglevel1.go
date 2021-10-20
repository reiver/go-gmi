package elm

import (
	"fmt"
	"strings"
)

// HeadingLevel1 represents a Gemtext heading-level-1.
//
// If a line in a Gemtext file starts with a "# ", then it is a Gemtext heading-level-1.
type HeadingLevel1 struct {
	value string
	eol   string
}

func (receiver HeadingLevel1) GoString() string {
	return fmt.Sprintf("elm.Create(%q, %q, %q)", receiver.MagicPrefix(), receiver.value, receiver.eol)
}

func (receiver HeadingLevel1) MagicPrefix() string {
	return "# "
}

func (receiver *HeadingLevel1) Set(value string, eol string) {
	if nil == receiver {
		return
	}

	receiver.value = value
	receiver.eol   = eol
}

func (receiver HeadingLevel1) String() string {

	var storage strings.Builder

	storage.WriteString(receiver.MagicPrefix())
	storage.WriteString(receiver.value)
	storage.WriteString(receiver.eol)

	return storage.String()
}
