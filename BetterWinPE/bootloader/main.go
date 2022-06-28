package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func network() {
	print("Connectée")
	exec.Command("cmd", "/c", `net use Q: \\AEM-GHOST\srv-projet /user:STORE STORE1`).Output()
	exec.Command("cmd", "/c", "start Q:\\Code\\bin\\x64\\Gui.exe").Output()
}

func offline() {

	var target string
	for _, i := range "abcdefghijklmnopqrstuvwxyz" {
		key := string(i) + ":"
		fmt.Println(key)
		_, err := os.Stat(key)
		if err != nil {

		} else {
			_, err := os.Stat(key + "\\Code")
			fmt.Println("Code found")

			if err == nil {
				_, err := os.Stat(key + "\\Image")
				fmt.Println("Image found")

				if err == nil {
					target = key + "\\"

				}

			}
		}

	}
	fmt.Println(target)
	exec.Command(target + "\\Gui.exe").Output()
}

func main() {
	exec.Command("powershell", "Get-Partition -DriveLetter D | Set-Partition -NewDriveLetter Y").Output()
	exec.Command("powershell", "Get-Partition -DriveLetter K | Set-Partition -NewDriveLetter X").Output()
	exec.Command("powershell", "Get-Partition -DriveLetter C | Set-Partition -NewDriveLetter Z").Output()
	var round int = 0
	fmt.Println("checking network ...")
	// networkStat := Check.IsNetwork()
	networkStat := false

	if networkStat {
		fmt.Println("Running in online mode")
		network()
	} else {
		for i := 0; i <= 5; i++ {
			fmt.Println("Running in offline mode")
			time.Sleep(3 * time.Second)
		}

		// networkStat = Check.IsNetwork()
		round += 1
	}

}
