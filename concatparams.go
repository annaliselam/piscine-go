package piscine

// Write a function that takes the arguments received in parameters and returns them as a string. 
// The string is the result of all the arguments concatenated with a newline (\n) between.

func ConcatParams(args []string) string {
	// answer := make([]int, len(args))
	answerString := ""
	
	for i := 0; i < len(args); i++ {
		for _, x := range args[i] {
			answerString += string(x)
		}
	}

	return answerString

}
