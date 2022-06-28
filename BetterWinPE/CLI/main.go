package main

import (
	"BetterWinPE/Custom"
	"BetterWinPE/setupEnv"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	workingDirectories string = ""
)

func menu() {
	cmd := exec.Command("cmd.exe", "/k", "C:\\Program Files (x86)\\Windows Kits\\10\\Assessment and Deployment Kit\\Deployment Tools\\DandISetEnv.bat")
	stdin, err := cmd.StdinPipe()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("travail actuellement dans " + workingDirectories)
		fmt.Println("1. definire le clavier (http://www.lingoes.net/en/translator/langcode.htm)")
		fmt.Println("2. monter l'image")
		fmt.Println("3. demonter l'image")
		fmt.Println("4. crée une clée USB bootable")
		reader := bufio.NewScanner(os.Stdin)
		choice := ""
		for reader.Scan() {
			choice = reader.Text()
			break
		}
		switch choice {
		case "1":
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("code :")
			code := ""
			for scanner.Scan() {
				code = scanner.Text()
				break
			}
			Custom.SetLanguage(code, workingDirectories)
		case "2":
			setupEnv.Mount(stdin, workingDirectories)

		case "3":
			setupEnv.UnMount(stdin, workingDirectories)
		}
	}
}

func main() {
	fmt.Println("1. crée un environement de travail")
	fmt.Println("2. specifiée un environement de travail existant")
	choice := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choice = scanner.Text()
		break
	}
	fmt.Println(choice)
	if strings.Contains(choice, "1") {
		fmt.Println("création de l'environement de travail...")
		setupEnv.CreateEnv()
		fmt.Println("Terminer")
		workingDirectories = "C:\\WinPE_amd64_PS"
	}
	if strings.Contains(choice, "2") {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("environement de travail :")
		for scanner.Scan() {
			workingDirectories = scanner.Text()
			break
		}
	}
	menu()
}
