package main

type Command struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}