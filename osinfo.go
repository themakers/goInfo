package osinfo

import (
	"io/ioutil"
	"regexp"
	"runtime"
	"strings"
)

type OSInfo struct {
	Kernel   string
	Core     string
	Platform string
	OS       string

	LinuxRelease *LinuxOSRelease
}

func (gi *OSInfo) IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

func (gi *OSInfo) IsWindows() bool {
	return runtime.GOOS == "windows"
}

func (gi *OSInfo) IsLinux() bool {
	return runtime.GOOS == "linux"
}

func (gi *OSInfo) IsFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}

type LinuxOSRelease struct {
	Name             string
	Version          string
	ID               string
	IDLike           []string
	VersionCodename  string
	VersionID        string
	PrettyName       string
	ANSIColor        string
	CPEName          string
	HomeURL          string
	SupportURL       string
	BugReportURL     string
	PrivacyPolicyURL string
	BuildID          string
	Variant          string
	VariantID        string

	ExtendedFields map[string]string
}

// https://www.freedesktop.org/software/systemd/man/os-release.html
func newLinuxOSRelease() (*LinuxOSRelease, error) {
	osrData, err := (func(paths ...string) (data []byte, err error) {
		for _, path := range paths {
			data, err = ioutil.ReadFile(path)
			if err == nil {
				break
			}
		}
		return
	})(
		"/etc/os-release",
		"/usr/lib/os-release",
	)
	if err != nil {
		return nil, err
	}

	osr := &LinuxOSRelease{}

	for _, mch := range regexp.MustCompile(`([A-Za-z0-9\_]+)=(.*)\n`).FindAllStringSubmatch(string(osrData), -1) {
		name, val := mch[1], strings.Trim(mch[2], " \"")
		switch name {
		case "NAME":
			osr.Name = val
		case "VERSION":
			osr.Version = val
		case "ID":
			osr.ID = val
		case "ID_LIKE":
			osr.IDLike = strings.Split(val, " ")
		case "VERSION_CODENAME":
			osr.VersionCodename = val
		case "VERSION_ID":
			osr.VersionID = val
		case "PRETTY_NAME":
			osr.PrettyName = val
		case "ANSI_COLOR":
			osr.ANSIColor = val
		case "CPE_NAME":
			osr.CPEName = val
		case "HOME_URL":
			osr.HomeURL = val
		case "SUPPORT_URL":
			osr.SupportURL = val
		case "BUG_REPORT_URL":
			osr.BugReportURL = val
		case "PRIVACY_POLICY_URL":
			osr.PrivacyPolicyURL = val
		case "BUILD_ID":
			osr.BuildID = val
		case "VARIANT":
			osr.Variant = val
		case "VARIANT_ID":
			osr.VariantID = val
		default:
			if osr.ExtendedFields == nil {
				osr.ExtendedFields = make(map[string]string)
			}
			osr.ExtendedFields[name] = val
		}
	}

	return osr, nil
}
