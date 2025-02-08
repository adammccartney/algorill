package cli

import (
)

// this basically prints a prompt to stdout telling the user what to do
// so a writer?
type Prompt func(args ...interface{}) []string

// this reads input from the user and transforms input into action
type Transformer func(words []string) interface{}

func InteractivePrompt (pfn Prompt, tfn Transformer) {
    // as long as you don't catch a signal to stop
    // process user input
    for {
        input := pfn()
        tfn(input)
    }
}
