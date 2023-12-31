package main

import (
	"fmt"
	"os"
	"tango_cli/filemaker"
)

func main() {

	paramsSetted := 0

	//Example: tango_cli photo basic f

	var fm *filemaker.FileMaker

	if len(os.Args) > 1 {

		fm = filemaker.New(os.Args[1])
		fmt.Println(os.Args[0])
		fm.SetRootPath("")
		fm.SetAppDir("app")

		paramsSetted = paramsSetted + 1

	}

	if len(os.Args) >= 2 {

		// aca define que va a crear
		fm.SetMode(os.Args[2])
		paramsSetted = paramsSetted + 1
	}

	if len(os.Args) >= 4 {

		// aca define el modo SI es forzado o no
		fmt.Println("Forced Mode")
		fm.SetForcedMode(true)

		paramsSetted = paramsSetted + 1
	}

	// aca ejecuta todo

	if paramsSetted >= 2 {

		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Making files...")
		fm.MakeIt()
		fmt.Println("-----------------")
		fmt.Println("- Check the model to add backticks to tags")
		fmt.Println("- Now just add the routes call to the app/routes/setupapproutes.go")
	}

}
