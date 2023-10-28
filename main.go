package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chzyer/readline"
)

func mustReadPassword(rl *readline.Instance, prompt string) string {
	line, err := rl.ReadPassword(prompt)
	if err != nil {
		if err == readline.ErrInterrupt || err == io.EOF {
			rl.Close()
			os.Exit(0)
		}
		panic(err)
	}
	return string(line)
}

func sameFirstChars(s1, s2 string, count int) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) < count || len(r2) < count {
		return false
	}

	for i := 0; i < count; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	password := mustReadPassword(rl, "Enter the password to learn: ")
	password2 := mustReadPassword(rl, "Confirm it: ")

	if password != password2 {
		fmt.Println("Passwords do not match")
		return
	}

	// Loop until there is enough successful input
	target := 10
	success := 0
	failed := 0
	for i := 3; i < len(password); i++ {

		result := mustReadPassword(rl, fmt.Sprintf("Type %d first chars: ", i))

		if sameFirstChars(result, password, i) {
			success++
			fmt.Printf("Nice! (success: %d)\n", success)
			continue
		}

		i--
		failed++
		fmt.Printf("Sorry, try again (failures: %d)\n", failed)
	}

	success = 0
	for {
		if success >= target {
			break
		}

		if mustReadPassword(rl, "Type full password: ") == password {
			success++
			fmt.Printf("Nice! (success: %d)\n", success)
			continue
		}

		failed++
		fmt.Printf("Sorry, try again (failures: %d)\n", failed)
	}

	fmt.Printf("Congrats, you're done!\nSuccess: %d\nFailures: %d\n", success, failed)

}
