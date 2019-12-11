package strdiff

import "testing"

type TestLehvenstein struct {
	S1 string
	S2 string
	D int
}

func TestLevenshteinDistance(t *testing.T) {
	testData := make([]*TestLehvenstein, 0)
	testData = append(testData, &TestLehvenstein{
		S1: "abc",
		S2: "abd",
		D:  1,
	},&TestLehvenstein{
		S1: "abc",
		S2: "ade",
		D:  2,
	},&TestLehvenstein{
		S1: "abc",
		S2: "def",
		D:  3,
	},&TestLehvenstein{
		S1: "abc",
		S2: "abca",
		D:  1,
	},&TestLehvenstein{
		S1: "abc",
		S2: "abcabc",
		D:  3,
	},&TestLehvenstein{
		S1: "abc",
		S2: "ab",
		D:  1,
	},&TestLehvenstein{
		S1: "abc",
		S2: "",
		D:  3,
	})

	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.LevenshteinDistance() != td.D {
			t.Error("Distance between", td.S1, "and", td.S2, "expected to", td.D, "but",sd.LevenshteinDistance())
		}
	}
}

type TestTrigram struct {
	S1 string
	S2 string
	D float32
}

func TestTrigramCompare(t *testing.T) {
	testData := make([]*TestTrigram, 0)
	testData = append(testData, &TestTrigram{
		S1: "Twitter v1",
		S2: "Twitter v2",
		D:  0.6666667,
	},&TestTrigram{
		S1: "Twitter v1",
		S2: "Twitter v1",
		D:  1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.TrigramCompare() != td.D {
			t.Error("trigram Compare between", td.S1, "and", td.S2, "expected to", td.D, "but",sd.TrigramCompare())
		}
	}
}

type TestDjaroDistancce struct {
	S1 string
	S2 string
	DJ float32
}

func TestDjaroDistance(t *testing.T) {
	testData := make([]*TestDjaroDistancce, 0)
	testData = append(testData, &TestDjaroDistancce{
		S1: "martha",
		S2: "marhta",
		DJ:  0.9444444,
	},&TestDjaroDistancce{
		S1: "martha",
		S2: "martha",
		DJ:  1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.DjaroDistance() != td.DJ {
			t.Error("Djaro Distance between", td.S1, "and", td.S2, "expected to", td.DJ, "but", sd.DjaroDistance())
		}
	}
}

func TestDjaroWinklerDistance(t *testing.T) {
	testData := make([]*TestDjaroDistancce, 0)
	testData = append(testData, &TestDjaroDistancce{
		S1: "martha",
		S2: "marhta",
		DJ:  0.96111107,
	},&TestDjaroDistancce{
		S1: "martha",
		S2: "martha",
		DJ:  1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.DjaroWinklerDistance( 0.1) != td.DJ {
			t.Error("Djaro Distance between", td.S1, "and", td.S2, "expected to", td.DJ, "but",sd.DjaroWinklerDistance( 0.1))
		}
	}
}