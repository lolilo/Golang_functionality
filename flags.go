// https://gobyexample.com/command-line-flags

package main

import "flag"
import "fmt"


// Basic flag declarations are available for string, integer, and boolean options. 
// Here we declare a string flag "word" 
// with a default value "foo" 
// and a short description, "a string".

// This flag.String function returns a string pointer 
// (not a string value); we’ll see how to use this pointer below.

func main() {
	wordPtr := flag.String("word", "foo", "a string")



	// decalre numb and fork flags
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")



	// It's also possible to declare an option that uses 
	// and existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag declaration function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")


	// Once all flags are declared, call flag.Prase() to execute
	// the command-line parsing.	
	flag.Parse()

	// Here, we just dump out parsed options and 
	// any trailing positional arguments. 
	// Note that we need to dereference the points with
	// e.g. *wordPtr to get the actual option values.

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}




