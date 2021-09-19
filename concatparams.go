package piscine

// Write a function that takes the arguments received in parameters and returns them as a string.
// The string is the result of all the arguments concatenated with a newline (\n) between.

func ConcatParams(args []string) string {
	emptyString := ""
	count := 0

	for range args {
		count++
	}

	for i := range args {
		if i == count-1 {
			emptyString += args[i]
		} else {
			emptyString += args[i] + "\n"
		}
	}
	return emptyString
}
