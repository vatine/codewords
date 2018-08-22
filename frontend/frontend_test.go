package frontend

import (
	"fmt"
	"testing"

	cw "github.com/vatine/codewords"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockServer struct {
	calls int
}

func (s *mockServer) GetCodewords(ctx context.Context, r *cw.CodewordsRequest, opts ...grpc.CallOption) (*cw.CodewordsResponse, error) {
	_ = ctx
	s.calls += 1

	rv := &cw.CodewordsResponse{}

	for n := int32(0); n < r.Count; n++ {
		rv.Words = append(rv.Words, fmt.Sprintf("testing%d", n))
	}

	return rv, nil
}

func TestNextCodeword1and5(t *testing.T) {
	td := []struct {
		count int
		e     string
	}{
		{1, "testing0"}, {1, "testing1"}, {1, "testing2"},
		{1, "testing3"}, {2, "testing4"}, {2, "testing0"},
	}
	client = &mockServer{}
	ctx := context.Background()
	SetMinimumCacheSize(1)
	SetMaximumCacheSize(5)

	for ix, d := range td {
		response, err := NextCodeword(ctx)
		if err != nil {
			t.Errorf("Unexpected error in test %d", ix)
		}
		if response != d.e {
			t.Errorf("UNexpected codeword, saw %s, expected %s", response, d.e)
		}
		if client.(*mockServer).calls != d.count {
			t.Errorf("Unexpected number of backend calls at ix %d, saw %d expected %d", ix, client.(*mockServer).calls, d.count)
		}
	}
}

func TestNextCodeword1and6(t *testing.T) {
	td := []struct {
		count int
		e     string
	}{
		{1, "testing0"}, {1, "testing1"}, {1, "testing2"},
		{1, "testing3"}, {1, "testing4"}, {2, "testing5"},
		{2, "testing0"},
	}
	client = &mockServer{}
	ctx := context.Background()
	SetMinimumCacheSize(1)
	SetMaximumCacheSize(6)
	cache = []string{}

	for ix, d := range td {
		response, err := NextCodeword(ctx)
		if err != nil {
			t.Errorf("Unexpected error in test %d", ix)
		}
		if response != d.e {
			t.Errorf("Unexpected codeword, saw %s, expected %s", response, d.e)
		}
		if client.(*mockServer).calls != d.count {
			t.Errorf("Unexpected number of backend calls at ix %d, saw %d expected %d", ix, client.(*mockServer).calls, d.count)
		}
	}
}
