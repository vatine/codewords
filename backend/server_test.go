package backend

import (
	"golang.org/x/net/context"
	"testing"
	cw "github.com/vatine/codewords"
)

func TestAssignment(t *testing.T) {
	var s cw.CodewordsServiceServer
	s1 := Server{}
	s = &s1
	_ = s
	if false {
		t.Errorf("Not expected to see this")
	}
}

func TestServer(t *testing.T) {
	GenRandSeq(0, 0, 0, 0, 0, 1, 1, 0)
	req := &cw.CodewordsRequest{Count: 3}
	s := &Server{}
	resp, _ := s.GetCodewords(context.Background(), req)
	expected := []string{"tall aardvark", "tall wombat", "short aardvark"}

	for ix, seen := range resp.Words {
		exp := expected[ix]
		if exp != seen {
			t.Errorf("Expected %s, saw %s, at index %d", exp, seen, ix)
		}
	}
}
