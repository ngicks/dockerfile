import "./dep/dot_env_load.ts";

import { defaultRun } from "./internal/default_param.ts";
import { run } from "./internal/run.ts";

Deno.exit(
  await run([
    ...defaultRun(),
    ...Deno.args,
  ]),
);
