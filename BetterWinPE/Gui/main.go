package main

import (
	"BetterWinPE/Custom"
	"BetterWinPE/setupEnv"
	"os/exec"

	ui "github.com/VladimirMarkelov/clui"
)

func Prompt(text string) string {
	// mainFrame := ui.AddWindow(0, 20, 60, 10, "input")
	// mainFrame.SetPack(ui.Vertical)
	// mainFrame.SetVisible(false)
	// ui.CreateLabel(mainFrame, 5, 2, text, 1)
	// btnLangue := ui.CreateButton(mainFrame, 40, 1, "ok", 1)
	// input := ui.CreateEditField(mainFrame, 40, "", 1)
	// btnQuit.OnClick(func(ev ui.Event) {
	// 	go ui.Stop()
	// })
	// mainFrame.SetVisible(true)
	// btnLangue.OnClick(func(e ui.Event) {
	// 	return input.Title()
	// })
	return ""
}

func menu(workingDirectories string) *ui.Window {

	work := workingDirectories
	mainFrame := ui.AddWindow(0, 15, 60, 10, "Global Menu working on "+workingDirectories)
	mainFrame.SetPack(ui.Vertical)
	mainFrame.SetVisible(false)
	if setupEnv.IsMounted(workingDirectories) {
		ui.CreateLabel(mainFrame, 5, 5, "Environment est monter", 1)

	} else {
		ui.CreateLabel(mainFrame, 5, 5, "Environment est monter", 1)

	}
	// btnLangue := ui.CreateButton(mainFrame, 40, 1, *workingDirectories, 1)
	btnMount := ui.CreateButton(mainFrame, 40, 1, "Monter l'image", 1)
	btnUnMount := ui.CreateButton(mainFrame, 40, 1, "Demonter l'image", 1)
	BtnProd := ui.CreateButton(mainFrame, 40, 1, "envoyer l'image dans une clée", 1)
	btnLayout := ui.CreateButton(mainFrame, 40, 1, "change le layout du clavier en fr", 1)
	btnBackground := ui.CreateButton(mainFrame, 40, 1, "Change le fond d'écran", 1)
	btnStartup := ui.CreateButton(mainFrame, 40, 1, "Change le startnet", 1)
	btnFolder := ui.CreateButton(mainFrame, 40, 1, workingDirectories, 1)

	btnMount.OnClick(func(e ui.Event) {
		setupEnv.Mount(work)
	})

	btnUnMount.OnClick(func(e ui.Event) {
		setupEnv.UnMount(work)
	})

	btnLayout.OnClick(func(e ui.Event) {
		Custom.SetLanguage("fr-FR", work)
	})

	BtnProd.OnClick(func(e ui.Event) {
		setupEnv.ProdutionOnUsb(work)
	})
	btnFolder.OnClick(func(e ui.Event) {
		cmd, _ := exec.Command("cmd", "/c", "start "+work).Output()
		if cmd != nil {

		}
	})
	btnStartup.OnClick(func(e ui.Event) {
		go func() {
			exec.Command("cmd", "/c", "start notepad C:\\WinPE_amd64_PS\\mount\\Windows\\System32\\startnet.cmd").Output()
		}()
	})

	btnBackground.OnClick(func(ev ui.Event) {
		imgPath := ""

		dlg := ui.CreateFileSelectDialog("Select your background image", "*.jpg", "C:\\Users\\", false, true)
		dlg.OnClose(func() {
			if !dlg.Selected {
				// a user canceled the dialog
				return
			}
			imgPath = dlg.FilePath
			Custom.ChangeBackground(work, imgPath)

		})
	})
	mainFrame.SetVisible(true)
	return mainFrame
}
func main() {
	// round := 0
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	workingDirectories := "a"

	wnd := ui.AddWindow(0, 0, 60, ui.AutoSize, "Scrollable frame")
	wnd.SetSizable(false)

	frm := ui.CreateFrame(wnd, 50, 12, ui.BorderNone, ui.Fixed)
	frm.SetPack(ui.Vertical)
	frm.SetScrollable(true)

	btnQuit := ui.CreateButton(frm, 40, 1, "quit", 1)
	btn := ui.CreateButton(frm, 40, ui.AutoSize, "Crée un environement", 1)

	btnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	btn.OnClick(func(ev ui.Event) {
		setupEnv.CreateEnv()
		menu("C:\\WinPE_amd64_PS")

	})

	btn2 := ui.CreateButton(frm, 40, ui.AutoSize, "Selectione un environement", 1)

	inputFrame := ui.AddWindow(0, 0, 60, ui.AutoSize, "Scrollable frame")
	inputFrame.SetPack(ui.Vertical)
	inputFrame.SetVisible(false)

	inputFrame.SetVisible(false)
	btn2.OnClick(func(ev ui.Event) {
		dlg := ui.CreateFileSelectDialog("Select your environelent", "", "C:\\", true, true)
		dlg.OnClose(func() {
			if !dlg.Selected {
				// a user canceled the dialog
				return
			}
			workingDirectories = dlg.FilePath
			menu(workingDirectories)

		})

	})
	ui.MainLoop()
}
