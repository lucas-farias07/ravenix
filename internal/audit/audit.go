package audit

import (
	"os/exec"
)

type Finding struct {
	exitCode    int
	severity    int // 5-1: 5 is least severe, 1 is critical
	description string
}

func Run() (findings []Finding) {
	findings = append(findings, nixCheck())

	vulnixPath, err := exec.LookPath("vulnix")
	if err != nil {
		findings = append(findings, Finding{description: "Vulnix is not in path. Should be packaged with the flake." +
			" If you have built the project from source, please install Vulnix!"})
	} else {
		findings = append(findings, vulnixCheck(vulnixPath))
	}

	return
}
