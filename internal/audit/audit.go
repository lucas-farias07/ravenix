package audit

import (
	"fmt"
	"os/exec"
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

func nixCheck() Finding {
	// first: check store
	out, err := exec.Command("nix", "store", "verify", "--all").CombinedOutput()
	if err != nil {
		return Finding{severity: 3, description: string(out)}
	}

	return Finding{description: string(out)}
}
