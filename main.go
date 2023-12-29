package main

import (
	"os"
	"tango_cli/filemaker"
	"tango_cli/parser"
)

func main() {

	paramsSetted := 0
	//todo

	//1. leer el comando

	//2. leer el nombre del paquete o elemento

	//3. Este si tiene CRUD te tira un crud. si tiene FULL te tira un full. Sino se pone que se quiere crear (feature, model,route,view)

	//3. leer el modo (forced)

	p := parser.New()
	fm := filemaker.New()

	if len(os.Args) > 1 {

		fm.SetRootPath(os.Args[0])
		p.Read(os.Args[1])

		paramsSetted = paramsSetted + 2

	}

	if len(os.Args) >= 2 {

		// aca define que va a crear

		paramsSetted = paramsSetted + 1
	}

	if len(os.Args) >= 3 {

		// aca define el modo SI es forzado o no

		fm.SetForcedMode(true)

		paramsSetted = paramsSetted + 1
	}

	// aca ejecuta todo

	if paramsSetted > 3 {

	}

}
