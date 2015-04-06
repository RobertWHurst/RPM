package main

import "flag"
import "fmt"
import "os/exec"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		runCommand(args[0], args[1:])
	} else {
		fmt.Println("rpm <command> [args...]")
	}
}

func runCommand(command string, args []string) {
	cmd := exec.Command("rpm-"+command, args...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("rpm-" + command + " does not exist")
	}
	fmt.Print(string(stdout))
}
