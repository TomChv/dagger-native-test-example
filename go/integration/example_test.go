package integration

import (
	"context"
	"testing"

	"dagger/example/integration/client/dag"

	"github.com/stretchr/testify/require"
)

func TestContainerEcho(t *testing.T) {
	ctx := context.Background()
	res, err := dag.Example().ContainerEcho("test").Stdout(ctx)
	require.NoError(t, err)
	require.Equal(t, res, "test\n")
}
