package audit

import (
	"fmt"
	"os/exec"
	"strings"
)

type Finding struct {
	severity    int // 5-1: 5 is least severe, 1 is critical
	description string
}

func Run() {
	find1 := nixCheck()
	fmt.Println(find1.severity)
	fmt.Println(find1.description)
}

func parse(out string) (parsed string) {
	sliced := strings.Split(out, "\n")
	for _, line := range sliced {
		if strings.HasPrefix(line, "checking") {
			continue
		}
		parsed += line
	}

	return
}

func nixCheck() Finding {
	// first: check store
	out, err := exec.Command("nix", "store", "verify", "--all").CombinedOutput()
	if err != nil {
		strOut := string(out)
		parsedOut := parse(strOut)
		return Finding{severity: 3, description: string(parsedOut)}
	}

	return Finding{description: string(out)}
}
