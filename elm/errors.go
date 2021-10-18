package elm

import (
	"errors"
)

var (
	errInternalError  = errors.New("elm: internal error")
	errNilRuneScanner = errors.New("elm: nil rune scanner")
)
