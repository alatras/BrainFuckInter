package BrainFuckInter

/*
/ validateArguments validates length and type of input.
*/
func validateArguments(args []string) bool {
	return len(args) == 3 && (args[1] == "--file" || args[1] == "--script")
}
