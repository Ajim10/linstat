package sysos

import (
	"bufio"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type Info struct {
	Architecture  string
	OS            string
	Kernel        string
	Distro        string
	DistroVersion string
	KernelVersion string
	BuildDate     string
	BuildNumber   string
	OSRelease     OSRelease
}

type OSRelease struct {
	Name            string
	VersioID        string
	Version         string
	VersionCodename string
	ID              string
	HomeURL         string
	SupportURL      string
	BugReportURL    string
}

func Stat() (Info, error) {
	osRelease, err := parseOSRelease()
	if err != nil {
		return Info{}, err
	}
	version, err := parseVersion()
	if err != nil {
		return Info{}, err
	}
	return Info{
		Architecture:  runtime.GOARCH,
		OS:            runtime.GOOS,
		OSRelease:     osRelease,
		KernelVersion: version.KernelVersion,
		BuildNumber:   version.BuildNumber,
		BuildDate:     version.BuildDate,
	}, nil
}

func parseOSRelease() (OSRelease, error) {
	content, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return OSRelease{}, err
	}
	values := make(map[string]string)
	for line := range strings.SplitSeq(string(content), "\n") {
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		values[key] = strings.Trim(value, `"`)
	}
	return OSRelease{
		Name:            values["NAME"],
		Version:         values["VERSION"],
		VersionCodename: values["VERSION_CODENAME"],
		ID:              values["ID"],
		HomeURL:         values["HOME_URL"],
		SupportURL:      values["SUPPORT_URL"],
		BugReportURL:    values["BUG_REPORT_URL"],
	}, nil
}

func parseVersion() (Info, error) {
	file, err := os.OpenFile("/proc/version", os.O_RDONLY, 0444)
	if err != nil {
		return Info{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return Info{}, err
	}

	info := Info{}

	reKernel, err := regexp.Compile(`Linux version (\S+)`)
	if err != nil {
		return Info{}, err
	}
	if match := reKernel.FindStringSubmatch(line); len(match) > 1 {
		info.KernelVersion = match[1]
	}

	reBuild, err := regexp.Compile(`\(([^)]+)\)`)
	if err != nil {
		return Info{}, err
	}
	if match := reBuild.FindStringSubmatch(line); len(match) > 1 {
		info.BuildNumber = match[1]
	}

	reDate, err := regexp.Compile(`#\d+\s+\S+\s+(.+)$`)
	if err != nil {
		return Info{}, err
	}
	if match := reDate.FindStringSubmatch(line); len(match) > 1 {
		info.BuildDate = match[1]
	}
	return info, nil
}
