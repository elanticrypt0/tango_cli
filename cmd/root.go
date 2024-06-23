/*
Copyright © 2024 NAME HERE elanticrypt0@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"tango_cli/pkg/cmdrunner"
	"tango_cli/pkg/filemaker"
	"tango_cli/pkg/parser"

	"github.com/spf13/cobra"
)

// the API folder
const APIPATH = "./api/"
const FRONTENDROOTPATH = "/frontend"

var cmdRunner = cmdrunner.New()

// func init() {
// cmdRunner.AppendToRootPath(APIPATH)
// }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tango_cli",
	Short: "CLI to create CRUD or make a build for Tango",
	Long:  `CLI to create CRUD or make a build for Tango`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creación de archivos individuales de features, models y routes",
	Long:  `Crear features, models, views, Api`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
		}
	},
}

var createPackApiCmd = &cobra.Command{
	Use:   "createapicrud",
	Short: "Creación de archivos paquetes de features, models y routes",
	Long:  `Creación de archivos paquetes de features, models y routes`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker
		var namespace string
		var templateSelected string = "api"

		if len(args) > 0 {
			namespace = args[0]
		}

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.AppendToRootPath("/api"), "app", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var createHttpClient = &cobra.Command{
	Use:   "httpclient",
	Short: "Crea una clase para realizar las peticiones ajax al servidor.",
	Long:  `Crea una clase para realizar las peticiones ajax al servidor.`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker

		namespace := "_tangoclient"
		templateSelected := "httpclient"

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.GetRootPath(), "frontend", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", namespace)
		fmt.Println("Mode: ", "forcedMode = true")
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var createModelCmd = &cobra.Command{
	Use:   "createmodel",
	Short: "Creación de archivo de modelo",
	Long:  `Creación de archivo de modelo`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker
		var namespace string
		var templateSelected string = "model"

		if len(args) > 0 {
			namespace = args[0]
		}

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.AppendToRootPath("/api"), "app", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var showAppConfig = &cobra.Command{
	Use:   "appconfig",
	Short: "Muestra la configuracion de la app",
	Long:  `Muestra la configuracion de la app`,
	Run: func(cmd *cobra.Command, args []string) {

		appconfig := cmdRunner.LoadAppConfig()
		fmt.Printf("%+v\n", appconfig)

	},
}

var makeBuild = &cobra.Command{
	Use:   "build",
	Short: "Crea el build de la app",
	Long:  `Crea el build de la app`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: agregar la opcion de compilar para diferentes plataformas.
		appconfig := cmdRunner.LoadAppConfig()

		// fmt.Printf("%+v\n", appconfig)

		rootPath := cmdRunner.GetRootPath()
		appnamePlusVersion := parseAppNameAndVersion(appconfig.Name, appconfig.Version)

		fmt.Printf("%+v\n", appnamePlusVersion)
		// tiene que crear la carpeta build
		buildpath := rootPath + "/build/"
		cmdRunner.Mkdir(buildpath)
		// luego crear una carpeta con el nombre de app + version (app_version)
		buildpath += appnamePlusVersion
		cmdRunner.Mkdir(buildpath)
		// logs, configuracion, public,cookies,uploads, _db
		cmdRunner.Mkdir(buildpath + "/logs")
		fmt.Printf("	> Creando la carpeta %s\n", "/logs")
		cmdRunner.Mkdir(buildpath + "/_db")
		fmt.Printf("	> Creando la carpeta %s\n", "/_db")
		cmdRunner.Mkdir(buildpath + "/cookies")
		fmt.Printf("	> Creando la carpeta %s\n", "/cookies")
		cmdRunner.Mkdir(buildpath + "/uploads")
		fmt.Printf("	> Creando la carpeta %s\n", "/uploads")
		// err := cmdRunner.CopyAll()
		err := cmdrunner.CopyDirectory(rootPath+"/api/config", buildpath+"/config")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("	> Copiando la carpeta %s\n", "/config")
		// err = cmdRunner.CopyAll(rootPath+"/api/public", buildpath+"/public")
		err = cmdrunner.CopyDirectory(rootPath+"/api/public", buildpath+"/public")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("	> Copiando la carpeta %s\n", "/public")
		// crear el ejecutable
		// dentro el ejecutable con el mismo nombre

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	appBanner("0.9.3")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tango_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createPackApiCmd)
	rootCmd.AddCommand(createHttpClient)
	rootCmd.AddCommand(createModelCmd)
	rootCmd.AddCommand(showAppConfig)
	rootCmd.AddCommand(makeBuild)
}
