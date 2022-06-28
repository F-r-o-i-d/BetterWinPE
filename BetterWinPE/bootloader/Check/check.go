package Check

import (
	"os/exec"
)

func IsNetwork() bool {
	_, err := exec.Command("cmd", "/c", `ping AEM-GHOST`).Output()
	return err == nil
}
