package project

import "fmt"
import "github.com/RobertWHurst/rpm/config"

func Generate(name string, path string) error {
	config := config.New()
	fmt.Println(config.Get("", "test"))
	return nil
}
