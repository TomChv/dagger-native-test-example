package integration

import (
	"context"
	"os"
	"testing"

	"dagger/example/integration/client/dag"

	"github.com/dagger/testctx"
	"github.com/dagger/testctx/oteltest"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(oteltest.Main(m)) // auto-wires OTel exporters
}

func TestExample(t *testing.T) {
	testctx.New(t,
		testctx.WithParallel(),             // run tests in parallel
		oteltest.WithTracing[*testing.T](), // trace each test and subtest
		oteltest.WithLogging[*testing.T](), // direct t.Log etc. to span logs
	).RunTests(&ExampleSuite{})
}

type ExampleSuite struct{}

func (ExampleSuite) TestContainerEcho(ctx context.Context, t *testctx.T) {
	res, err := dag.Example().ContainerEcho("foo").Stdout(ctx)
	require.NoError(t, err)
	require.Equal(t, "foo\n", res)
}
