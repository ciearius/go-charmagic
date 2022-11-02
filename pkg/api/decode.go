package api

import (
	"errors"
	"fmt"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
)

var errNoMatch = errors.New("no match in index")
var errNotSupported = errors.New("not supported")

// GetDecoder tries to get a matching encoding or returns nil with an error.
func GetDecoder(name string) (encoding.Encoding, error) {
	enc, err := query_all_indexes(name)

	if err != nil {
		return nil, fmt.Errorf("no encoder found for %s: %s", name, err)
	}

	return enc, nil
}

// MustGetDecoder tries to get a matching encoding or panics.
//
// Use discouraged since errors should always be checked.
func MustGetDecoder(name string) encoding.Encoding {
	enc, err := query_all_indexes(name)

	if err != nil {
		panic(err)
	}

	return enc
}

// query_all_indexes looks up a name in all three known encoding indexes and returns the first non-nil encoding.
// In case no non-nil encoding is found an error is returned
func query_all_indexes(name string) (encoding.Encoding, error) {
	var enc encoding.Encoding
	var err error

	// TODO: #7 make sure there are no weird edgecases in encoding indexes

	enc, err = get_encoding(ianaindex.IANA, name)

	if enc != nil {
		return enc, err
	}

	enc, err = get_encoding(ianaindex.MIME, name)

	if enc != nil {
		return enc, err
	}

	return get_encoding(ianaindex.MIB, name)
}

// get_encoding looks up a name in one specific index and checks for the two known errors.
func get_encoding(idx *ianaindex.Index, name string) (encoding.Encoding, error) {
	enc, err := idx.Encoding(name)

	if enc == nil && err != nil {
		return nil, errNoMatch
	}

	if enc == nil && err == nil {
		return nil, errNotSupported
	}

	return enc, nil
}
