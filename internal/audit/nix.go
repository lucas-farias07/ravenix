package audit

import (
	"os/exec"
	"strings"
)

func parse(out string) (parsed string, count int) {
	sliced := strings.Split(out, "\n")
	for _, line := range sliced {
		if strings.HasPrefix(line, "checking") {
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		count++
		parsed += line + "\n"
	}
	parsed = strings.TrimSpace(parsed)

	return
}

func nixCheck() Finding {
	// check store for corrupted, untrusted or unverified following: https://nix.dev/manual/nix/2.19/command-ref/new-cli/nix3-store-verify
	out, err := exec.Command("nix", "--extra-experimental-features", "nix-command", "store", "verify", "--all").CombinedOutput()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return Finding{severity: 3, description: "could not run nix verify store: " + err.Error()}
		}

		strOut := string(out)
		parsedOut, count := parse(strOut)
		code := exitErr.ExitCode()
		var severityVal int
		switch code {
		case 1, 3, 5, 7: // sum with 1: corrupted
			severityVal = 1
		case 2, 6: // untrusted, not corrupted
			severityVal = 3
		case 4: // only couldn't verify
			severityVal = 3
		default: // unexpected code
			severityVal = 4
		}
		if count > 3 {
			severityVal = max(1, severityVal-1)
		}

		return Finding{exitCode: exitErr.ExitCode(), severity: severityVal, description: string(parsedOut)}
	}

	return Finding{exitCode: 1, severity: 0, description: string(out)}
}
