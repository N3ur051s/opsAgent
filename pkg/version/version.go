package version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Major  int64
	Minor  int64
	Patch  int64
	Pre    string
	Meta   string
	Commit string
}

var versionRx = regexp.MustCompile(`(\d+\.\d+\.\d+)(\-[^\+]+)*(\+.+)*`)

func Agent() (Version, error) {
	return New(AgentVersion, Commit)
}

func New(version, commit string) (Version, error) {
	toks := versionRx.FindStringSubmatch(version)
	if len(toks) == 0 || toks[0] != version {
		return Version{}, fmt.Errorf("Version string has wrong format")
	}

	parts := strings.Split(toks[1], ".")
	major, _ := strconv.ParseInt(parts[0], 10, 64)
	minor, _ := strconv.ParseInt(parts[1], 10, 64)
	patch, _ := strconv.ParseInt(parts[2], 10, 64)

	pre := strings.Replace(toks[2], "-", "", 1)

	meta := strings.Replace(toks[3], "+", "", 1)

	av := Version{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Pre:    pre,
		Meta:   meta,
		Commit: commit,
	}

	return av, nil
}

func (v *Version) String() string {
	ver := v.GetNumber()
	if v.Pre != "" {
		ver = fmt.Sprintf("%s-%s", ver, v.Pre)
	}
	if v.Meta != "" {
		ver = fmt.Sprintf("%s+%s", ver, v.Meta)
	}
	if v.Commit != "" {
		if v.Meta != "" {
			ver = fmt.Sprintf("%s.commit.%s", ver, v.Commit)
		} else {
			ver = fmt.Sprintf("%s+commit.%s", ver, v.Commit)
		}
	}

	return ver
}

func (v *Version) GetNumber() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) GetNumberAndPre() string {
	version := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Pre != "" {
		version = fmt.Sprintf("%s-%s", version, v.Pre)
	}
	return version
}
