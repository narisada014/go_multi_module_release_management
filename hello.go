package main

import (
	"fmt"

	"github.com/narisada014/go_multi_module_release_management/print_int/convert"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
	fmt.Println(convert.ConvertStrToInt("12345"))
}
