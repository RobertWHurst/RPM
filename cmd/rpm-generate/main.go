package main

import "flag"

// import "fmt"
import "github.com/RobertWHurst/rpm/project"

func main() {
	project.Generate(flag.Arg(0), flag.Arg(1))
}
