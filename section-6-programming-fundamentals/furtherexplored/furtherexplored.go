package furtherexplored

var planetSpeed = 67000

// Fascinating is in "package block" scope
// Capitalized functions are exported automatically
func Fascinating() {
	// planetSped is in "block scope"
	// and is accessible here
	println(planetSpeed)
}
