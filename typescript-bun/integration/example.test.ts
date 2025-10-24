import { test, expect, beforeAll, describe } from "bun:test";
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

