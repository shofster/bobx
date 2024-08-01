package net

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"testing"
)

/*
  Description: test GO Binaries - GOB encoding / decoding
*/

func TestGOBmap(t *testing.T) {

	var bufe bytes.Buffer
	enc := gob.NewEncoder(&bufe)

	me := make(map[string]string)
	me["foo"] = "bar"

	if err := enc.Encode(me); err != nil {
		t.Fatal(err)
	}

	t.Log(bufe.Bytes())

	// io.Reader
	bufd := bytes.NewBuffer(bufe.Bytes())
	dec := gob.NewDecoder(bufd)

	md := make(map[string]string)

	if err := dec.Decode(&md); err != nil {
		log.Fatal(err)
	}

	t.Log("[foo] = ", md["foo"]) // "bar"

}

type Gid int
type Gop int

const (
	Gcreate = iota
	Gopen
	Gclose
	Gread
	Gwrite
)

func (g Gop) String() string {
	switch g {
	case Gcreate:
		return "create"
	case Gopen:
		return "open"
	case Gclose:
		return "close"
	case Gread:
		return "read"
	case Gwrite:
		return "write"
	default:
		return fmt.Sprintf("%d", int(g))
	}
}

// Structs encode and decode only exported fields
type GobS struct {
	ID      Gid
	OP      Gop
	Comment string
	Tags    map[string]string
}

func (s GobS) String() string {
	str := fmt.Sprintf("GobS: ID=%d, OP=%s", s.ID, s.OP)
	str += fmt.Sprintf("\n\"%s\"", s.Comment)
	for k, v := range s.Tags {
		str += fmt.Sprintf("\n %s -> %s", k, v)
	}
	return str
}
func TestGOBstruct(t *testing.T) {
	g := GobS{
		ID:      5545,
		OP:      Gcreate,
		Comment: "I'm a comment",
		Tags:    make(map[string]string),
	}
	g.Tags["tag1"] = "value1"
	g.Tags["tag2"] = "value2"
	t.Logf("original:\n%s", g)

	var bufe bytes.Buffer
	enc := gob.NewEncoder(&bufe)

	if err := enc.Encode(g); err != nil {
		t.Fatal(err)
	}

	t.Logf("length %d\n%v\n", len(bufe.Bytes()), bufe.Bytes())

	// io.Reader
	bufd := bytes.NewBuffer(bufe.Bytes())
	dec := gob.NewDecoder(bufd)

	gd := &GobS{}
	if err := dec.Decode(gd); err != nil {
		log.Fatal(err)
	}
	t.Logf("gobbed:\n%s", gd)

}
