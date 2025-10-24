# Dagger native TypeScript Bun Unit Test

This directory contains a simple example of a [Dagger](https://dagger.io) module written in TypeScript with the Bun
runtime and tested using the Bun native unit test framework.

## Setup

This section describes how to setup the example project from scratch with native integration tests.  
Use it as a reference to setup tests in your own modules ;)


1. Initialize a new Bun project (or cd in an existing one)

```shell
bun init -y
# If you tsconfig.json has comment, dagger will fail to init so we either remove it or clean it.
rm index.ts tsconfig.json
```

2. Initialize a new Dagger module (or cd in an existing one)

```shell
dagger init --name=example --sdk=typescript --source=.
```

3. Create a new integration directory with a empty dagger module that has a dependency on the module to test, then generate a client.

```shell
mkdir integration && cd integration
dagger init
dagger install ../
dagger client install typescript ./client
```

:bulb: This is required because the module's `tsconfig.json` is configured to work on the module so to avoid
conficts, we create a light submodule that just proxies the client. 

:bulb: This client will also include bindings to call your module's functions.

4. Create a test file in the `integration` directory

```typescript
// integration/example.test.ts
import { test, expect, describe } from "bun:test";
import { connection, dag } from "@dagger.io/client";

describe("Example", () => {
  test(
    "containerEcho",
    async () => {
      let res!: string;

      await connection(async () => {
        res = await dag.example().containerEcho("test").stdout();
      });

      expect(res).toBe("test\n");
    },
    { timeout: 30 * 1000 }
  );
});

```

5. Run the tests

```shell
bun test
#  integration/example.test.ts:
#  ✓ Example > containerEcho [7193.76ms]
# ...
```

:bulb: The generated client will automatically load your module.

:bulb: You may need to run `dagger develop` and `bun install` before running the tests.