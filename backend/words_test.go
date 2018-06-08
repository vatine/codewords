package backend

import (
	"testing"
)

func fakeRand(n int) int {
	_ = n
	return 1
}

func GenRandSeq(n ...int) {
	ix := 0
	l := len(n)
	rndN = func(m int) int {
		base := n[ix % l]
		ix = ix + 1
		return base % m
	}
}

func TestAdjectives(t *testing.T) {
	eWord := "short"
	eInt := 1

	rndN = fakeRand
	sWord, sInt := GetAdjective()

	if sWord != eWord {
		t.Errorf("Saw adjective %s, expected %s", sWord, eWord)
	}
	if sInt != eInt {
		t.Errorf("Saw adjective #%d, expected %d", sInt, eInt)
	}
}

func TestNouns(t *testing.T) {
	eWord := "wombat"
	eInt := 1

	rndN = fakeRand
	sWord, sInt := GetNoun()

	if sWord != eWord {
		t.Errorf("Saw noun %s, expected %s", sWord, eWord)
	}
	if sInt != eInt {
		t.Errorf("Saw noun #%d, expected %d", sInt, eInt)
	}
}

func TestGenerated(t *testing.T) {
	reset()
	GenRandSeq(0, 0, 0, 0, 0, 1)
	w1, _ := GeneratePair()
	if w1 != "tall aardvark" {
		t.Errorf("Expected a tall aardvark, saw %s", w1)
	}
	w2, _ := GeneratePair()
	if w2 != "tall wombat" {
		t.Errorf("Expected a tall wombat, saw %s", w2)
	}
	
	
}

func TestMultiGenerated(t *testing.T) {
	reset()
	GenRandSeq(0, 0, 0, 0, 0, 1, 1, 0)
	seen, _ := GenerateNPairs(3)
	expected := []string{"tall aardvark", "tall wombat", "short aardvark"}

	for ix, s := range seen {
		e := expected[ix]
		if s != e {
			t.Errorf("At index %d, saw %s, expected %s", ix, s, e)
		}
	}
}
