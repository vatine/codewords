package backend

import (
	cw "github.com/vatine/codewords"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) GetCodewords(ctx context.Context, in *cw.CodewordsRequest) (*cw.CodewordsResponse, error) {
	pairs, e := GenerateNPairs(in.Count)
	rv := &cw.CodewordsResponse{
		Words: pairs,
	}

	return rv, e
}
