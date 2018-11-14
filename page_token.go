package ripper

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/pkg/errors"
)

func validatePageToken(p *PaginateParams, pageToken string) error {
	tokenParams, err := decodePageToken(pageToken)
	if err != nil {
		return err
	}
	p.Page = tokenParams.Page
	if *p != *tokenParams {
		return errors.Wrap(newInvalidPageTokenError(), "the PaginateParams is not corresponding to the pageToken")
	}
	return nil
}

func encodeToPageToken(p *PaginateParams) (s string, err error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err = enc.Encode(p)
	if err != nil {
		err = errors.Wrap(newInvalidPageTokenError(), err.Error())
		return
	}
	s = base64.URLEncoding.EncodeToString(b.Bytes())
	return
}

func decodePageToken(pageToken string) (p *PaginateParams, err error) {
	if pageToken == "" {
		panic("pageToken can't be empty")
	}
	byte, err := base64.URLEncoding.DecodeString(pageToken)
	if err != nil {
		err = errors.Wrap(newInvalidPageTokenError(), err.Error())
		return
	}
	b := bytes.NewBuffer(byte)
	dec := gob.NewDecoder(b)
	p = &PaginateParams{}
	err = dec.Decode(p)
	if err != nil {
		err = errors.Wrap(newInvalidPageTokenError(), err.Error())
		return
	}
	return
}
