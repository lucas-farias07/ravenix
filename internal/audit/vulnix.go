package audit

import (
	"os/exec"
)

type vulnixReturn struct {
	Name        string             `json:"name"`
	Pname       string             `json:"pname"`
	Version     string             `json:"version"`
	Derivation  string             `json:"derivation"`
	AffectedBy  []string           `json:"affected_by"`
	Whitelisted []string           `json:"whitelisted"`
	CVSSv3      map[string]float64 `json:"cvssv3_basescore"`
	Description map[string]string  `json:"description"`
}

func vulnixCheck(path string) Finding {
	out, err := exec.Command(path, "--system", "--json").Output() // catch the vulnerabilities in json formatting so there is no parsing yay
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return Finding{severity: 5,
				description: "Vulnix system check could not run:" + err.Error()}
		}
		return Finding{exitCode: exitErr.ExitCode(), severity: 5,
			description: "Vulnix system check could not run:" + err.Error()}
	}

	return Finding{}
}
