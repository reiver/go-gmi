package elm

import (
	"fmt"
	"strings"
)

// BlockQuotation represents a Gemtext block-quotation.
// I.e., what the sometimes gets called “block-quote”.
//
// If a line in a Gemtext file starts with a ">", then it is a Gemtext block-quotation.
type BlockQuotation struct {
	value string
	eol   string
}

func (receiver BlockQuotation) GoString() string {
	return fmt.Sprintf("elm.Create(%q, %q, %q)", receiver.MagicPrefix(), receiver.value, receiver.eol)
}

func (receiver BlockQuotation) MagicPrefix() string {
	return ">"
}

func (receiver *BlockQuotation) Set(value string, eol string) {
	if nil == receiver {
		return
	}

	receiver.value = value
	receiver.eol   = eol
}

func (receiver BlockQuotation) String() string {

	var storage strings.Builder

	storage.WriteString(receiver.MagicPrefix())
	storage.WriteString(receiver.value)
	storage.WriteString(receiver.eol)

	return storage.String()
}
