package cmdrunner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// This run OS commands

type CmdRunner struct {
	Script_rootpath string
}

func newCmdRunner(script_rootpath string) CmdRunner {
	return CmdRunner{
		Script_rootpath: script_rootpath,
	}
}

func (cr *CmdRunner) RunLines(lines []string) {
	// chequea primero si el comando está dentro de los comandos espaciales
	// sino está lo ejecuta
	for _, command := range lines {
		command_splitted := strings.Split(command, " ")
		command_name := command_splitted[0]
		command_args := command_splitted[1:]
		if !cr.isSpecialCmdAndExecute(command_name, command_args) {
			cr.RunSliceArgs(command_name, command_args)
		}
	}

}

func (cr *CmdRunner) isSpecialCmdAndExecute(cmd string, args []string) bool {
	isSpecial := true

	switch cmd {
	case "mkdir":
		cr.Mkdir(args[0])
	case "cd":
		cr.Cd(args[0])
	case "echo":
		cr.Echo(strings.Join(args, " "))
	case "pwd":
		cr.Run("pwd")
	default:
		isSpecial = false
	}
	return isSpecial

}

func (cr *CmdRunner) RunSliceArgs(name string, arg []string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("> %s\n", cmd.String())
	return cmd.Run()
}

func (cr *CmdRunner) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("> %s\n", cmd.String())
	return cmd.Run()
}

func (cr *CmdRunner) Mkdir(newDir string) error {
	// checks is a directory exists
	// if is not then create
	newDir = cr.Script_rootpath + "/" + newDir
	_, err := os.Stat(newDir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.Run("mkdir", newDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CmdRunner) Cd(dirpath string) error {
	// checks is a directory exists
	// if is not then create
	dirpath = cr.Script_rootpath + "/" + dirpath
	_, err := os.Stat(dirpath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.Run("cd", dirpath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CmdRunner) Echo(msg string) {
	fmt.Println(msg)
}
