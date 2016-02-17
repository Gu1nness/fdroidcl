// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package fdroidcl

import (
	"encoding/hex"
	"strings"
	"time"
)

type CommaList []string

func (cl *CommaList) FromString(s string) {
	*cl = strings.Split(s, ",")
}

func (cl *CommaList) String() string {
	return strings.Join(*cl, ",")
}

func (cl *CommaList) UnmarshalText(text []byte) error {
	cl.FromString(string(text))
	return nil
}

type HexHash struct {
	Type string `xml:"type,attr"`
	Data HexVal `xml:",chardata"`
}

type HexVal []byte

func (hv *HexVal) FromString(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*hv = b
	return nil
}

func (hv *HexVal) String() string {
	return hex.EncodeToString(*hv)
}

func (hv *HexVal) UnmarshalText(text []byte) error {
	return hv.FromString(string(text))
}

type DateVal struct {
	time.Time
}

func (dv *DateVal) FromString(s string) error {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*dv = DateVal{t}
	return nil
}

func (dv *DateVal) String() string {
	return dv.Format("2006-01-02")
}

func (dv *DateVal) UnmarshalText(text []byte) error {
	return dv.FromString(string(text))
}
