package main

import (
	"fmt"
	"mama-lolo/GoCompiler/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Couldnt get User information and therefore terminated.")
		panic(err)
	}
	fmt.Printf("Hell %s! This is a tiny language trial you are taking part in. Good Luck!", user.Username)
	fmt.Println("Feel free to try some commands!")
	repl.Start(os.Stdin, os.Stdout)
}
