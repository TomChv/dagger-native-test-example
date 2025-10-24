# Dagger native Go Unit Test

This directory contains a simple example of a [Dagger](https://dagger.io) module written in Go tested using
the Go native unit test framework.

## Setup

This section describes how to setup the example project from scratch with native integration tests.  
Use it as a reference to setup tests in your own modules ;)


1. Initialize a new Dagger module (or cd in an existing one)

```shell
dagger init --name=example --sdk=go --source=.
```

2. Generate a client in test directory

```shell
dagger client install go ./integration/client
```

:bulb: This client will also include bindings to call your module's functions.

:warning: The client and the tests files must be in a different directory than the module source code
to avoid loading the module client and panic.

3. Create a test file in the `integration` directory

```go
// integration/example_test.go
package integration

import (
	"context"
	"dagger/example/integration/client/dag"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainerEcho(t *testing.T) {
	ctx := context.Background()
	res, err := dag.Example().ContainerEcho("test").Stdout(ctx)
	require.NoError(t, err)
	require.Equal(t, res, "test\n")
}
```

4. Run the tests

```shell
go test -v ./integration
```

:bulb: The generated client will automatically load your module.

5. Enhance the test with native tracing.

You can also use the [dagger/testctx](https://github.com/dagger/testctx) package to enhance the test with native tracing.

```go
// integration/enhanced_example_test.go
package integration

import (
	"context"
	"dagger/example/integration/client/dag"
	"os"
	"testing"

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
```

```shell
dagger run -v go test ./integration/enhanced_example_test.go
# ...

▼ TestExample 0.7s ✔ 1
╰╴▼ TestContainerEcho 0.7s
  ├╴✔ example: Example! 0.0s
  ├╴▼ .containerEcho(stringArg: "foo"): Container! 0.4s CACHED
  │ ├╴✔ container: Container! 0.0s
  │ ├╴$ .from(address: "alpine:latest"): Container! 0.2s CACHED
  │ ╰╴$ .withExec(args: ["echo", "foo"]): Container! 0.0s CACHED
  ╰╴▼ .stdout: String! 0.0s
    ┃ foo     
```