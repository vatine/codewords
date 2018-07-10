package backend

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
)

var generated map[int64]bool
var stateLock sync.Mutex

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
var maxConsecutiveCollisions int = 3

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

func GeneratePair() (string, error) {
	stateLock.Lock()
	defer stateLock.Unlock()
	for c := 0; c < maxConsecutiveCollisions; c++ {
		adj, adjN := GetAdjective()
		noun, nounN := GetNoun()
		comp := composite(adjN, nounN)
		if !generated[comp] {
			generated[comp] = true
			return strings.Join([]string{adj, noun}, " "), nil
		}
	}
	return "", errors.New("Too many collisions")
}

func GenerateNPairs(n int32) ([]string, error) {
	var rv []string

	for i := int32(0); i < n; i++ {
		s, e := GeneratePair()
		if e != nil {
			fmt.Printf("backend error: %s\n", e)
			return rv, e
		} else {
			rv = append(rv, s)
		}
	}

	return rv, nil
}

func PersistGenerated(w io.Writer) {
	for ix, _ := range generated {
		fmt.Fprintf(w, "%x\n", ix)
	}
}

func ReadGenerated(r io.Reader) {
	
}
