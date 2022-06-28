package Custom

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ChangeBackground(workingDirectories string, backgroundPath string) {
	// takeown /f
	userdomain := os.Getenv("USERDOMAIN")
	// usernames := os.Getenv("username")
	go func() {
		cmd := exec.Command("powershell")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			fmt.Fprintln(stdin, `$NewAcl = Get-Acl -Path `+workingDirectories+`\mount\Windows\System32\winpe.jpg`)
			fmt.Fprintln(stdin, `$identity = "`+userdomain+`\Administrators"`)
			fmt.Fprintln(stdin, `$fileSystemRights = "FullControl"`)
			fmt.Fprintln(stdin, `$type = "Allow"`)
			fmt.Fprintln(stdin, `$fileSystemAccessRuleArgumentList = $identity, $fileSystemRights, $type"`)
			fmt.Fprintln(stdin, `$fileSystemAccessRule = New-Object -TypeName System.Security.AccessControl.FileSystemAccessRule -ArgumentList $fileSystemAccessRuleArgumentList`)
			fmt.Fprintln(stdin, `$NewAcl.SetAccessRule($fileSystemAccessRule)`)
			fmt.Fprintln(stdin, `Set-Acl -Path "`+workingDirectories+`\mount\Windows\System32\winpe.jpg" -AclObject $NewAcl`)

		}()
		cmd.CombinedOutput()

		err = os.Rename(backgroundPath, workingDirectories+`\mount\Windows\System32\winpe.jpg`)
		if err != nil {
			panic(err)
		}
	}()

}
