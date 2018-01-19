package backend

import (
	"context"
	"math/rand"
	"strings"

        // "google.golang.org/grpc"

        // "google.golang.org/grpc/reflection"
	
	cw "github.com/vatine/codewords"
)

var generated map[int64]bool

var adjectives = []string {
	"tall",
	"short",
	"old",
	"young",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"indigo",
	"violet",
	"putrid",
	"fresh",
	"purple",
}
var nouns = []string {
	"aardvark",
	"wombat",
	"boat",
	"apple",
	"banana",
}

var rndN func(int) int

type server struct{}

func GetAdjective() (string, int) {
	n := rndN(len(adjectives))

	return adjectives[n], n
}

func GetNoun() (string, int) {
	n := rndN(len(nouns))

	return nouns[n], n
}

func init() {
	reset()
}

func reset() {
	rndN = rand.Intn
	generated = make(map[int64]bool)
}

func composite(a, b int) int64 {
	n1 := int64(a & 0xFFFF)
	n2 := int64(b & 0xFFFF)
	return (n1 << 32) | n2
}

func GeneratePair() string {
	for true {
		adj, adjN := GetAdjective()
		noun, nounN := GetNoun()
		comp := composite(adjN, nounN)
		if !generated[comp] {
			generated[comp] = true
			return strings.Join([]string{adj, noun}, " ")
		}
	}
	return ""
}

func GenerateNPairs(n int) []string {
	var rv []string

	for i := 0; i < n; i++ {
		rv = append(rv, GeneratePair())
	}

	return rv
}

func (s *server) GetCodewords(ctx context.Context, in *cw.CodewordsRequest) (*cw.CodewordsResponse, error) {
	return nil, nil
}
