package main

import (
	"errors"
	"fmt"

	"github.com/cearius/go-charmagic/pkg/matching"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
)

var errNoMatch = errors.New("no match in index")
var errNotSupported = errors.New("not supported")

// TODO: refactor and define clear interface

func GetDecoder(name string) (enc encoding.Encoding, err error) {
	enc, err = ianaindex.IANA.Encoding(name)

	if err == nil {
		return
	}

	enc, err = ianaindex.MIME.Encoding(name)

	if err == nil {
		return
	}

	enc, err = ianaindex.MIB.Encoding(name)

	if err == nil {
		return
	}

	return nil, fmt.Errorf("no encoder found for %s", name)
}

func GetDecoderFromResult(r *matching.Result) (enc encoding.Encoding) {
	name := r.Charset

	// ianaindex.IANA, ianaindex.MIME, ianaindex.MIB,

	enc = get_encoding(ianaindex.IANA, name)

	if enc != nil {
		return
	}

	enc = get_encoding(ianaindex.MIME, name)

	if enc != nil {
		return
	}

	enc = get_encoding(ianaindex.MIB, name)

	if enc == nil {
		panic(fmt.Errorf("no encoder found for '%s'", name))
	}

	return enc
}

// get_encoding tries to fetch the encoding matching the name from the provided Index
func get_encoding(idx *ianaindex.Index, name string) encoding.Encoding {
	enc, err := queryIndex(idx, name)

	if enc != nil {
		return enc
	}

	if err == errNotSupported {
		panic(fmt.Errorf("encoding '%s' not supported", name))
	}

	if err == errNoMatch {
		return nil
	}

	return enc
}

// queryIndex checks the index and returns an encoding and a bool.
func queryIndex(idx *ianaindex.Index, name string) (enc encoding.Encoding, err error) {

	// enc == nil && err != nil => no match
	// enc == nil && err == nil => not supported
	enc, err = idx.Encoding(name)

	if enc == nil && err != nil {
		return nil, errNoMatch
	}

	if enc == nil && err == nil {
		return nil, errNotSupported
	}

	return
}
