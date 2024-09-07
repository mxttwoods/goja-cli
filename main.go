package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dop251/goja"
)

func main() {
	if len(os.Args) > 1 {
		// Run the specified JS file
		runJSFile(os.Args[1])
	} else {
		// Open REPL
		startREPL()
	}
}

func createConsole(vm *goja.Runtime) {
	console := make(map[string]interface{})
	console["log"] = func(call goja.FunctionCall) goja.Value {
		for _, arg := range call.Arguments {
			fmt.Print(arg.String(), " ")
		}
		fmt.Println()
		return goja.Undefined()
	}
	vm.Set("console", console)
}

func runJSFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	vm := goja.New()
	createConsole(vm)
	_, err = vm.RunString(string(content))
	if err != nil {
		fmt.Printf("Error executing JavaScript: %v\n", err)
	}
}

func startREPL() {
	vm := goja.New()
	createConsole(vm)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Goja REPL (type 'exit' to quit)")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		result, err := vm.RunString(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else if result != goja.Undefined() {
			fmt.Println(result)
		}
	}
}