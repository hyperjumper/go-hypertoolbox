package versioning

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	SemVerRegex = `^([0-9]+)(\.([0-9]+))?(\.([0-9]+))?`
	PreReleaseRegex = `\-[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*`
	BuildRegex = `\+[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*`
	LowCase = "abcdefghijklmnopqrstuvwxyz"
	UpCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers = "01234567890"

	Minor = iota
	Major
	Patch
	Build
	PreRelease
	Additive
	Recursive
)

// NewSemanticVersion creates new instance of SemanticVersion by parsing the provided version string.
func NewSemanticVersion(version string) (*SemanticVersion, error) {
	sv := &SemanticVersion{
		Major:      0,
		Minor:      0,
		Patch:      0,
		Build:      "",
		PreRelease: "",
	}

	if len(version) == 0 {
		return sv, nil
	}


	rexp := regexp.MustCompile(SemVerRegex)
	rexpPreRR := regexp.MustCompile(PreReleaseRegex)
	rexpBuild := regexp.MustCompile(BuildRegex)


	if !rexp.Match([]byte(version)) {
		return nil, errors.New(fmt.Sprintf("version [%s] is not according to Sematic Versioning standard.", version))
	}

	group := rexp.FindAllSubmatch([]byte(version),-1)

	if len(group[0]) > 1 {
		major, _ := strconv.Atoi(string(group[0][1]))
		sv.Major = uint(major)
	}
	if len(group[0]) > 3 {
		minor, _ := strconv.Atoi(string(group[0][3]))
		sv.Minor = uint(minor)
	}
	if len(group[0]) > 5 {
		patch, _ := strconv.Atoi(string(group[0][5]))
		sv.Patch = uint(patch)
	}

	if rexpPreRR.Match([]byte(version)) {
		sv.PreRelease = rexpPreRR.FindString(version)[1:]
	}

	if rexpBuild.Match([]byte(version)) {
		sv.Build = rexpBuild.FindString(version)[1:]
	}

	return sv, nil
}

// SemanticVersion struct store information about Major number, Minor number, Patch number, build and prerelease identifier.
type SemanticVersion struct {
	Major uint
	Minor uint
	Patch uint
	Build string
	PreRelease string
}

// Bump will bump the SemanticVersion, based on which version element to bump.
func (semver *SemanticVersion) Bump(bumpType, bumpMode int) {
	switch bumpType {
	case Minor :
		semver.Minor++
	case Major :
		semver.Major++
	case Patch :
		semver.Patch++
	case Build:
		switch bumpMode {
		case Recursive:
			semver.Build = BumpRecursive(semver.Build)
		case Additive:
			semver.Build = BumpAdditive(semver.Build)
		}
	case PreRelease:
		switch bumpMode {
		case Recursive:
			semver.PreRelease = BumpRecursive(semver.PreRelease)
		case Additive:
			semver.PreRelease = BumpAdditive(semver.PreRelease)
		}
	}
}

// BumpRecursive will recursively bump the version number so the versioning scheme is incremental without adding new digits.
// for example, build number 10009 will be bumped to 10010 naturaly. abcdez bumped to abcdfa. etc.
func BumpRecursive(str string) string {
	return bumpRec(str, len(str)-1)
}

func bumpRec(str string, indice int) string {
	if indice < 0 {
		return str
	}
	char := str[indice:indice+1]
	if char == "Z" {
		if indice == 0 {
			return "A" + str[:indice] + "A" + str[indice+1:]
		}
		return bumpRec(str[:indice] + "A" + str[indice+1:], indice - 1)
	}
	if char == "z" {
		if indice == 0 {
			return "a" + str[:indice] + "a" + str[indice+1:]
		}
		return bumpRec(str[:indice] + "a" + str[indice+1:], indice - 1)
	}
	if char == "9" {
		if indice == 0 {
			return "1" + str[:indice] + "0" + str[indice+1:]
		}
		return bumpRec(str[:indice] + "0" + str[indice+1:], indice - 1)
	}
	if strings.Contains(LowCase, char) {
		cindice := strings.Index(LowCase, char)
		return str[:indice] + LowCase[cindice+1: cindice+2] + str[indice+1:]
	}
	if strings.Contains(UpCase, char) {
		cindice := strings.Index(UpCase, char)
		return str[:indice] + UpCase[cindice+1: cindice+2] + str[indice+1:]
	}
	if strings.Contains(Numbers, char) {
		cindice := strings.Index(Numbers, char)
		return str[:indice] + Numbers[cindice+1: cindice+2] + str[indice+1:]
	}
	return bumpRec(str, indice - 1)
}

// BumpAdditive will  bump the version number by increment the last digit/character element.
// for example, build number 10009 will be bumped to 100010, notice how after 9 is increased 10. abcdez bumped to abcdeaa. etc.
func BumpAdditive(str string) string {
	if len(str) == 0 {
		return "0"
	}
	lastChar := str[len(str)-1:]
	if lastChar == "Z" {
		return str[:len(str)-1]+"AA"
	}
	if lastChar == "z" {
		return str[:len(str)-1]+"aa"
	}
	if lastChar == "9" {
		return str[:len(str)-1]+"10"
	}
	if strings.Contains(LowCase, lastChar) {
		indice := strings.Index(LowCase, lastChar)
		return str[:len(str)-1]+LowCase[indice+1: indice+2]
	}
	if strings.Contains(UpCase, lastChar) {
		indice := strings.Index(UpCase, lastChar)
		return str[:len(str)-1]+UpCase[indice+1: indice+2]
	}
	if strings.Contains(Numbers, lastChar) {
		indice := strings.Index(Numbers, lastChar)
		return str[:len(str)-1]+Numbers[indice+1: indice+2]
	}
	return str
}


func (semver *SemanticVersion) String() string {
	mainVer := fmt.Sprintf("%d.%d.%d", semver.Major, semver.Minor, semver.Patch)
	if len(semver.PreRelease) > 0 {
		mainVer = fmt.Sprintf("%s-%s", mainVer, semver.PreRelease)
	}
	if len(semver.Build) > 0 {
		mainVer = fmt.Sprintf("%s+%s", mainVer, semver.Build)
	}
	return mainVer
}

// BumpVersion will return the bumped version string.
func BumpVersion(ver string, bumpType, bumpMode int) string {
	sv, err := NewSemanticVersion(ver)
	if err != nil {
		return ver
	}
	sv.Bump(bumpType, bumpMode)
	return sv.String()
}
