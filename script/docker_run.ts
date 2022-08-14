import { defaultRun } from "./internal/default_param.ts";
import { run } from "./internal/run.ts";

Deno.exit(
  await run([
    ...defaultRun,
    ...Deno.args,
  ]),
);
