package Custom

import (
	"fmt"
	"os/exec"
)

func SetLanguage(code string, workingDirectories string) {
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, `Dism /Image:"`+workingDirectories+`\Mount" /Set-InputLocale:`+code)
	}()

	cmd.CombinedOutput()

}
