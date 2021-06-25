package main

// PrintFOO is a function that receives a function to fetch a specific environment variable and prints
// the FOO environment variable.
func PrintFOO(getEnv func(string) string) {
	print(getEnv("FOO"))
}

// environment is a helper structure that contains an environment.
type environment struct {
	env map[string]string
}

// GetEnv is a method on the environment structure that fetches a specific environment variable.
func (e *environment) GetEnv(name string) string {
	return e.env[name]
}

// main is the main entry point of this demo application
func main() {
	// Create the environment
	env := &environment{
		// Pass the variables we want to inject
		env: map[string]string{
			"FOO": "bar",
		},
	}
	// We pass the GetEnv method on the env struct pointer to PrintFOO so it can fetch the environment variable.
	PrintFOO(env.GetEnv)
}
