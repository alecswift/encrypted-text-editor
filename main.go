package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const MicroservicePath = "password_microservice/save_user_password.py"
const MSCommPath = "password_microservice/user_password.txt"
const UserPasswordText = `
	Please enter a username and password. You will have to reenter the username
	and password to regain access to the file
	`
const UserGuideText = `
Purpose: A Desktop application for securely writing, deleting, and saving text files.
Home page User Interface
	Text box: Below the toolbar is a large text box for writing and deleting text.
	Toolbar: On the top of the application is a toolbar for navigating the application.
			The contents are detailed below from left to right.
		Actions: A button that opens a pop down menu with the primary application actions.
			User guide: A button that opens a pop up text box that details the user guide
				for the application.
			New: Opens a new, blank text file.
			Open: Opens the file directory, the user can select any text file from the
				local computer to open into the application. If a password has been set
				for the requested file, then a pop up menu opens requesting the user to
				enter their password.
			Save: Opens the file directory, the user can select any folder in the local
				computer to save the currently worked on text file.
			Delete: Opens a pop up menu that confirms whether or not the user would like
				to delete the current text file. If the user presses yes, then the current
				file is deleted. If the user presses no, then the user returns to the
				application.
			Undo: Undoes the last action. For example, if a character was written last,
				then delete that character. Or if a character was deleted last, then
				rewrite that character.
			Copy: Copies the selected text into the clipboard for temporary storage.
				The copied text can be used later with the paste command.
			Paste: Paste the previously copied text stored in the clipboard at the spot
				selected by the cursor.
			Set Password for file: Opens a pop up box that requests the user to enter and
				confirm a password. Once the user has entered and confirmed a password,
				the user can confirm and set the password for the file. Otherwise the user
				can exit the pop up box and return to the application without setting a
				password.
	Exit button: Appears as an X. This button opens a pop up text box that asks the user to
		confirm that they want to exit the application. If the user presses no, then the user
		is returned to the application. If the user presses yes, then the application closes.
`

func main() {
	myApp := app.New()
	setUpMainWindow(myApp)
	myApp.Run()
	tidyUp()
}

func setUpMainWindow(myApp fyne.App) {
	mainWindow := myApp.NewWindow("Encrypted Text Editor")
	mainWindow.Resize(fyne.NewSize(1200, 700))

	mainWindow.SetMainMenu(actionBar(myApp))
	mainWindowCloseIntercept(myApp, mainWindow)

	textBox := widget.NewMultiLineEntry()
	textBox.AcceptsTab()
	mainWindow.SetContent(textBox)
	mainWindow.Show()
}

func mainWindowCloseIntercept(myApp fyne.App, mainWindow fyne.Window) {
	mainWindow.SetCloseIntercept(func() {
		closeWindow := myApp.NewWindow("Confirm Exit")
		closeWindowContent := makeCloseWindowContent(closeWindow, mainWindow)
		closeWindow.SetContent(closeWindowContent)
		closeWindow.Show()
	})
}

func makeCloseWindowContent(closeWindow, mainWindow fyne.Window) *fyne.Container {
	closeWindowContent := container.NewVBox(
		widget.NewLabel("Please confirm that you would like to exit the application"),
		widget.NewButton("Confirm", func() {
			mainWindow.Close()
			closeWindow.Close()
		}),
	)
	return closeWindowContent
}



func actionBar(myApp fyne.App) *fyne.MainMenu {
	buttons := []*fyne.MenuItem{
		makeUserGuideButton(myApp), makeNewButton(myApp), makeOpenButton(myApp),
		makeSaveButton(myApp), makeDeleteButton(myApp), makeUndoButton(myApp),
		makeCopyButton(myApp), makePasteButton(myApp), makeSetPasswordButton(myApp),
	}
	actionPopDown := fyne.NewMenu("Actions", buttons...)
	actions := fyne.NewMainMenu(actionPopDown)
	return actions
}

func makeActionBarButtons() {

}

func makeUserGuideButton(myApp fyne.App) *fyne.MenuItem {
	userGuideButton := fyne.NewMenuItem("UserGuide", func() {
		userGuideWindow := myApp.NewWindow("User Guide")
		userGuideWindow.Resize(fyne.NewSize(400, 240))
		userGuideText := widget.NewLabel(UserGuideText)
		userGuideWindow.SetContent(userGuideText)
		userGuideWindow.Show()
	})
	return userGuideButton
}

func makeNewButton(myApp fyne.App) *fyne.MenuItem {
	newButton := fyne.NewMenuItem("New File", func() {fmt.Println("Pressed")})
	return newButton
}

func makeOpenButton(myApp fyne.App) *fyne.MenuItem {
	openButton := fyne.NewMenuItem("Open a File", func() {fmt.Println("Pressed")})
	return openButton
}

func makeSaveButton(myApp fyne.App) *fyne.MenuItem {
	saveButton := fyne.NewMenuItem("Save the File", func() {fmt.Println("Pressed")})
	return saveButton
}

func makeDeleteButton(myApp fyne.App) *fyne.MenuItem {
	deleteButton := fyne.NewMenuItem("Delete the File", func() {
		deleteWindow := myApp.NewWindow("Delete")
		deleteWindow.Resize(fyne.NewSize(400, 240))
		deleteWidgets := makeDeleteWidgets()
		deleteWindow.SetContent(deleteWidgets)
		deleteWindow.Show()
	})
	return deleteButton
}

func makeDeleteWidgets() *fyne.Container {
	deleteWidgets := container.NewVBox(
		widget.NewButton("Confirm", func() {}),
		widget.NewLabel("Please confirm that you would like to delete the file"),
		)
	return deleteWidgets
}

func makeUndoButton(myApp fyne.App) *fyne.MenuItem {
	undoButton := fyne.NewMenuItem("Undo the Last action", func() {fmt.Println("Pressed")})
	return undoButton
}

func makeCopyButton(myApp fyne.App) *fyne.MenuItem {
	copyButton := fyne.NewMenuItem("Copy Selected Text", func() {fmt.Println("Pressed")})
	return copyButton
}

func makePasteButton(myApp fyne.App) *fyne.MenuItem {
	pasteButton := fyne.NewMenuItem("Paste Copied Text", func() {})
	return pasteButton
}

func makeSetPasswordButton(myApp fyne.App) *fyne.MenuItem {
	setPasswordButton := fyne.NewMenuItem("Set a Password", func() {
		passwordWindow := myApp.NewWindow("Set a Password")
		passwordWindow.Resize(fyne.NewSize(400, 240))

		passwordWidgets := makePasswordWidgets(passwordWindow)
		passwordWindow.SetContent(passwordWidgets)
		passwordWindow.Show()
	})
	return setPasswordButton
}

func makePasswordWidgets(passwordWindow fyne.Window) *fyne.Container {
	userName, password := widget.NewEntry(), widget.NewEntry()
	userName.SetPlaceHolder("Enter username here")
	password.SetPlaceHolder("Enter password here")
	passwordWidgets := container.NewVBox(userName, password,
		widget.NewButton("Confirm", func() {save(userName.Text, password.Text)
		passwordWindow.Close()
		}),
		widget.NewLabel(UserPasswordText),
	)
	return passwordWidgets
}

func save(userName, password string) int {

	file, err := os.Create(MSCommPath)
	checkFor(err)
	n, err := file.WriteString(userName + "\n" + password)
	checkFor(err)

	file.Sync()
	runMicroservice()
	return n
}

func runMicroservice() {
	cmd := exec.Command("/usr/bin/python3", MicroservicePath)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	go func(){
		if cmd.Run()!= nil{
			panic("Error")
		}
	}()
}

func checkFor(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func tidyUp() {
	fmt.Println("Exited")
}
