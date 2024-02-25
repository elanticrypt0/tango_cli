package cmd

import "fmt"

func appBanner(version string) {

	fmt.Println(" ### ### ### ####")
	fmt.Println("    TANGO_CLI    ")
	fmt.Printf("        V. %s    ", version)
	fmt.Println(" ### ### ### ####")
	fmt.Println("")
}
