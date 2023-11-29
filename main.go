package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Encrypted Text Editor")
	mainWindow.Resize(fyne.NewSize(1200, 700))

	mainWindow.SetMainMenu(actionBar(myApp))
	mainWindow.SetCloseIntercept(func() {
		closeWindow := myApp.NewWindow("Confirm Exit")
		closeWindowContent := container.NewVBox(widget.NewLabel("Please confirm that you would like to exit the application"),
			widget.NewButton("Confirm", func() {
				mainWindow.Close()
				closeWindow.Close()
			}),
		)
		closeWindow.SetContent(closeWindowContent)
		closeWindow.Show()
	})

	textBox := widget.NewMultiLineEntry()
	textBox.AcceptsTab()
	mainWindow.SetContent(textBox)

	mainWindow.Show()
	myApp.Run()
	tidyUp()
}



func actionBar(myApp fyne.App) *fyne.MainMenu {
	userGuideButton := fyne.NewMenuItem("UserGuide", func() {
		userGuideWindow := myApp.NewWindow("User Guide")
		userGuideWindow.Resize(fyne.NewSize(400, 240))
		userGuideText := widget.NewLabel(
			`Purpose: A Desktop application for securely writing, deleting, and saving text files.
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
		`)
		userGuideWindow.SetContent(userGuideText)
		userGuideWindow.Show()
	})
	newButton := fyne.NewMenuItem("New File", func() {fmt.Println("Pressed")})
	openButton := fyne.NewMenuItem("Open a File", func() {fmt.Println("Pressed")})
	saveButton := fyne.NewMenuItem("Save the File", func() {fmt.Println("Pressed")})
	deleteButton := fyne.NewMenuItem("Delete the File", func() {
		deleteWindow := myApp.NewWindow("Delete")
		deleteWindow.Resize(fyne.NewSize(400, 240))

		deleteWidgets := container.NewVBox(
			widget.NewButton("Confirm", func() {}),
			widget.NewLabel("Please confirm that you would like to delete the file"),
			)
		deleteWindow.SetContent(deleteWidgets)
		deleteWindow.Show()
	})
	undoButton := fyne.NewMenuItem("Undo the Last action", func() {fmt.Println("Pressed")})
	copyButton := fyne.NewMenuItem("Copy Selected Text", func() {fmt.Println("Pressed")})
	pasteButton := fyne.NewMenuItem("Paste Copied Text", func() {})
	setPasswordButton := fyne.NewMenuItem("Set a Password", func() {
		passwordWindow := myApp.NewWindow("Set a Password")
		passwordWindow.Resize(fyne.NewSize(400, 240))

		userName := widget.NewEntry()
		userName.SetPlaceHolder("Enter username here")
		password := widget.NewEntry()
		password.SetPlaceHolder("Enter password here")

		passwordWidgets := container.NewVBox(userName,
			password,
			widget.NewButton("Confirm", func() {
			save(userName.Text, password.Text)
			passwordWindow.Close()
			}),
			widget.NewLabel("Please enter a username and password. You will have to reenter the username and password to regain access to the file"),
			)
		passwordWindow.SetContent(passwordWidgets)
		passwordWindow.Show()
	})
	actionPopDown := fyne.NewMenu("Actions", userGuideButton, newButton, openButton, saveButton, deleteButton, undoButton, copyButton, pasteButton, setPasswordButton)
	actions := fyne.NewMainMenu(actionPopDown)
	return actions
}

func save(userName, password string) int {
	file, err := os.Create("/home/alec/Desktop/code/osu_projects/encrypted_text_editor/password_microservice/user_password.txt")
	if err != nil {
        log.Fatal(err)
	}
	n, err := file.WriteString(userName + "\n" + password)
	if err != nil {
	 log.Fatal(err)
	}
	file.Sync()
	return n
}

func tidyUp() {
	fmt.Println("Exited")
}
