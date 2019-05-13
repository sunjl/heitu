package shell

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	hostNameCmd  = "hostname"
	ipCmd        = "ip addr show %s | awk '/inet / {print $2}' | cut -d/ -f1"
	processorCmd = "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	memoryCmd    = "free -b | egrep '^Mem:' | awk '{print substr($2, index($0,$1))}'"
	diskCmd      = "df -B1 | grep '^/dev/%s' | awk '{print $2}'"
	totalDiskCmd = "df -B1 | grep '^/dev/[hs]d' | awk '{s+=$2} END {print s}'"
	platformCmd  = "lsb_release -d | awk '{print substr($0, index($0,$2))}'"
)

func cmdOutput(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return strings.TrimSpace(string(out))
}

func GroupName(groupFlag string) string {
	if groupFlag != "" {
		return groupFlag
	}
	var groupName string
	hostName := HostName()
	re := regexp.MustCompile(`^([a-z]+)([-_]*[0-9]*)$`)
	results := re.FindStringSubmatch(hostName)
	if len(results) > 0 {
		groupName = results[1]
	}
	return groupName
}

func HostName() string {
	return cmdOutput(hostNameCmd)
}

func IP(device string) string {
	cmd := fmt.Sprintf(ipCmd, device)
	return cmdOutput(cmd)
}

func Processor() string {
	return cmdOutput(processorCmd)
}

func Memory() int {
	memorySize, _ := strconv.Atoi(cmdOutput(memoryCmd))
	return memorySize
}

func Disk(device string) int {
	cmd := fmt.Sprintf(diskCmd, device)
	diskSize, _ := strconv.Atoi(cmdOutput(cmd))
	return diskSize
}

func TotalDisk() int {
	diskSize, _ := strconv.Atoi(cmdOutput(totalDiskCmd))
	return diskSize
}

func Platform() string {
	return cmdOutput(platformCmd)
}
