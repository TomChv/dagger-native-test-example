package main

import (
	"dagger/example/internal/dagger"
)

type Example struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Example) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}
