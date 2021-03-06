package cpu

import (
	"os/exec"
	"strconv"
	"strings"
)

var cpuMap = map[string]string{
	"machdep.cpu.vendor":       "vendor_id",
	"machdep.cpu.brand_string": "model_name",
	"hw.physicalcpu":           "cpu_cores",
	"hw.cpufrequency":          "mhz",
	"machdep.cpu.family":       "family",
	"machdep.cpu.model":        "model",
	"machdep.cpu.stepping":     "stepping",
}

func getCpuInfo() (cpuInfo map[string]string, err error) {

	cpuInfo = make(map[string]string)

	for option, key := range cpuMap {
		out, err := exec.Command("sysctl", "-n", option).Output()
		if err != nil {
			return cpuInfo, err
		}
		cpuInfo[key] = strings.Trim(string(out), "\n")
	}

	mhz, err := strconv.Atoi(cpuInfo["mhz"])
	if err != nil {
		return cpuInfo, err
	}
	cpuInfo["mhz"] = strconv.Itoa(mhz / 1000000)

	return
}
