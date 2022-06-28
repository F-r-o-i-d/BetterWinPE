package setupEnv

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var ListComponets = []string{
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-WMI.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-WMI_en-us.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-NetFX.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-NetFX_en-us.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-Scripting.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-Scripting_en-us.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-PowerShell.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-PowerShell_en-us.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-StorageWMI.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-StorageWMI_en-us.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\WinPE-DismCmdlets.cab"`,
	`Dism /Add-Package /Image:"C:\WinPE_amd64_PS\mount" /PackagePath:"C:\Program Files (x86)\Windows Kits\10\Assessment and Deployment Kit\Windows Preinstallation Environment\amd64\WinPE_OCs\en-us\WinPE-DismCmdlets_en-us.cab"`,
}
var running bool

func copyDir(stdin io.WriteCloser) {
	running = true
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, "copype amd64 C:\\WinPE_amd64_PS")
		running = false
	}()
}

func basicComponent(stdin io.WriteCloser) {
	running = true

	go func() {
		defer stdin.Close()
		for _, v := range ListComponets {
			fmt.Fprintln(stdin, v)

		}
		running = false

	}()
}
func Mount(stdin io.WriteCloser, workingDirectories string) {
	running = true
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, `Dism /Mount-Image /ImageFile:"`+workingDirectories+`\media\sources\boot.wim" /Index:1 /MountDir:"`+workingDirectories+`\mount"`)
		running = false

	}()
}

func UnMount(stdin io.WriteCloser, workingDirectories string) {
	go func() {
		defer stdin.Close()

		fmt.Fprintln(stdin, `Dism /Unmount-Image /MountDir:"`+workingDirectories+`\mount" /Commit`)
		fmt.Fprintln(stdin, `dism /Cleanup-Wim`)
		err := os.RemoveAll(workingDirectories + "/mount")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func CreateEnv() string {
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	workingDirectories := `C:\WinPE_amd64_PS`
	copyDir(stdin)
	time.Sleep(1 * time.Second)

	Mount(stdin, workingDirectories)
	time.Sleep(1 * time.Second)

	basicComponent(stdin)
	time.Sleep(1 * time.Second)

	return workingDirectories
}
