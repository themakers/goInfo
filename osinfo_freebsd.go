package osinfo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func GetInfo() *OSInfo {
	out := _getInfo()
	for strings.Index(out, "broken pipe") != -1 {
		out = _getInfo()
		time.Sleep(500 * time.Millisecond)
	}
	osStr := strings.Replace(out, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	gio := &OSInfo{Kernel: osInfo[0], Core: osInfo[1], Platform: runtime.GOARCH, OS: osInfo[2], CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()
	return gio
}

func _getInfo() string {
	cmd := exec.Command("uname", "-sri")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("getInfo:", err)
	}
	return out.String()
}
