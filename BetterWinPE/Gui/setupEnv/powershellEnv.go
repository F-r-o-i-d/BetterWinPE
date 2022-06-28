package setupEnv

import (
	"fmt"
	"io"
	"io/ioutil"
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

	for _, v := range ListComponets {
		fmt.Fprintln(stdin, v)
		exec.Command("powershell", v).Output()
	}

}
func Mount(workingDirectories string) {
	running = true
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, err := cmd.StdinPipe()
	if err != nil {

	}
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, `Dism /Mount-Image /ImageFile:"`+workingDirectories+`\media\sources\boot.wim" /Index:1 /MountDir:"`+workingDirectories+`\mount"`)
		running = false

	}()
	cmd.CombinedOutput()
}

func IsMounted(workingDirectories string) bool {
	list, _ := ioutil.ReadDir(workingDirectories + "\\mount")
	return len(list) > 1
}

func UnMount(workingDirectories string) {
	Copy("bin\\bootloader.exe", workingDirectories+`\mount\Windows\System32\bootloader.exe`)

	running = true
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()

		fmt.Fprintln(stdin, `Dism /Unmount-Image /MountDir:"`+workingDirectories+`\mount" /Commit`)
		fmt.Fprintln(stdin, `dism /Cleanup-Wim`)

	}()
	cmd.CombinedOutput()
	err = os.RemoveAll(workingDirectories + "/mount/")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(workingDirectories+"/mount", 0755)
	if err != nil {
		log.Fatal(err)
	}

}
func Copy(from string, dest string) {
	bytesRead, err := ioutil.ReadFile(from)
	if err != nil {
		log.Fatal(err)
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(dest, bytesRead, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
func ProdutionOnUsb(workingDirectories string) {
	exec.Command("cmd", "/c", "taskkill /f /im explorer.exe").Output()
	exec.Command("cmd", "/c", "diskpart /s diskpart.txt").Output()
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()

		fmt.Fprintln(stdin, `MakeWinPEMedia /UFD /F `+workingDirectories+` F:`)

	}()
	cmd.CombinedOutput()
	os.Mkdir("O:\\Code", 0755)
	os.Mkdir("O:\\Image", 0755)
	os.Mkdir("O:\\Image\\Image", 0755)
	os.Mkdir("O:\\Code\\widget", 0755)
	os.Mkdir("O:\\Code\\Diskpart", 0755)
	os.Create("O:\\Code\\widget\\button.aem")
	Copy("bin\\Gui.exe", "O:\\Gui.exe")
	exec.Command("cmd", "/c", "start explorer.exe").Output()
	exec.Command("cmd", "/c", "start o:").Output()

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

	Mount(workingDirectories)
	time.Sleep(1 * time.Second)

	basicComponent(stdin)
	time.Sleep(1 * time.Second)

	return workingDirectories
}
