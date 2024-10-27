package main

type VMState struct {
	Running   bool
	Suspended bool
	Stopped   bool
}

type VMConfig struct {
	Name   string
	Memory int
	CPU    int
}

type VM struct {
	State  VMState
	Config VMConfig
}

func main() {
	// Start the VM Manager
}
