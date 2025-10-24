import { dag, Container, object, func } from "@dagger.io/dagger";

@object()
export class Example {
  @func()
  containerEcho(stringArg: string): Container {
    return dag.container().from("alpine:latest").withExec(["echo", stringArg]);
  }
}
