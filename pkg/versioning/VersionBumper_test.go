package versioning

import (
	"fmt"
	"testing"
)

func TestBumpAdditive(t *testing.T) {
	data := map[string]string {
		"" : "0",
		"0" : "1",
		"1" : "2",
		"9" : "10",
		"99" : "910",
		"999" : "9910",
		"9999" : "99910",
		"9999999999" : "99999999910",
		"a": "b",
		"aa": "ab",
		"z": "aa",
		"az": "aaa",
		"A": "B",
		"Y": "Z",
		"Z": "AA",
		"AA": "AB",
		"AZ": "AAA",
		"1a": "1b",
		"1z": "1aa",
	}

	for k,v := range data {
		sv := BumpAdditive(k)
		if sv != v {
			t.Error(fmt.Sprintf("%s is expected tobe %s, but %s", k, v, sv))
			t.Fail()
		}
	}
}

func TestBumpRecursive(t *testing.T) {
	data := map[string]string {
		"+1" : "+2",
		"+09" : "+10",
		"1" : "2",
		"9" : "10",
		"+9" : "+0",
		"99" : "100",
		"999" : "1000",
		"5.5.5" : "5.5.6",
		"5.5.9" : "5.6.0",
		"+a" : "+b",
		"+0z" : "+1a",
		"a" : "b",
		"Z" : "AA",
		"ZZ" : "AAA",
		"ZZZ" : "AAAA",
		"A.B.C" : "A.B.D",
		"A.B.Z" : "A.C.A",
	}

	for k,v := range data {
		sv := BumpRecursive(k)
		if sv != v {
			t.Error(fmt.Sprintf("%s is expected tobe %s, but %s", k, v, sv))
			t.Fail()
		}
	}
}

func TestBumpVersion(t *testing.T) {
	testBumpMinor := map[string]string {
		"" : "0.0.0",
		"0" : "0.0.0",
		"1" : "1.0.0",
		"9" : "9.0.0",
		"99" : "99.0.0",
		"1.9" : "1.9.0",
		"9.9" : "9.9.0",
		"9.9.9" : "9.9.9",
		"9.9.9+abc.234.ab34" : "9.9.9+abc.234.ab34",
		"9.9.9-abc.234.ab34" : "9.9.9-abc.234.ab34",
		"9.9.9-abc.234.ab34+22.44.44" : "9.9.9-abc.234.ab34+22.44.44",
	}

	for k,v := range testBumpMinor {
		sv,err := NewSemanticVersion(k)
		if err != nil {
			t.Error(err)
			t.Fail()
		} else {
			if sv.String() != v {
				t.Error(fmt.Sprintf("%s is expected tobe %s, but %s", k, v, sv.String()))
				t.Fail()
			}
		}
	}
}

type BumpTest struct {
	before string
	after string
	bumpType int
	bumpMode int
}

func TestSemanticVersion_Bump(t *testing.T) {
	data := []*BumpTest {
		&BumpTest{
			before:   "0.0.9",
			after:    "0.0.10",
			bumpType: Patch,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0",
			after:    "0.0.1",
			bumpType: Patch,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0-a.b.c+1.2.3",
			after:    "0.0.1-a.b.c+1.2.3",
			bumpType: Patch,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0-a.b.c+1.2.3",
			after:    "0.1.0-a.b.c+1.2.3",
			bumpType: Minor,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0-a.b.c+1.2.3",
			after:    "1.0.0-a.b.c+1.2.3",
			bumpType: Major,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0-a.b.c+1.2.3",
			after:    "0.0.0-a.b.d+1.2.3",
			bumpType: PreRelease,
			bumpMode: Recursive,
		},
		&BumpTest{
			before:   "0.0.0-a.b.c+1.2.3",
			after:    "0.0.0-a.b.c+1.2.4",
			bumpType: Build,
			bumpMode: Recursive,
		},
	}

	for _,v := range data {
		aft := BumpVersion(v.before, v.bumpType, v.bumpMode)
		if aft != v.after {
			t.Errorf("%s expected to %s, but %s", v.before, v.after, aft)
		}
	}
}