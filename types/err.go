package types

import "errors"

var ErrInvalidCookie = errors.New("invalid cookie in the response header")

var ErrReadBody = errors.New("got unknown error, unable to read body")

var ErrDeserializeBody = errors.New("got error deserializing body")
