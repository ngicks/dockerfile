import { defaultRun } from "./internal/default_param.ts";
import { run } from "./internal/run.ts";

const trimProtocolScheme = (url?: string): string | undefined => {
  if (typeof url === "undefined") {
    return void 0;
  }
  try {
    const u = new URL(url);
    // must be local path.
    // only pathname should suffice.
    return u.pathname;
  } catch {
    // new URL throws TypeError when input is invalid url.
    return url;
  }
};

Deno.exit(
  await run([
    ...defaultRun,
    // prefer v option for that
    // rancher-desktop is not currently able to handle `--mount` option for single file mount.
    "-v",
    `${
      trimProtocolScheme(Deno.env.get("DOCKER_HOST")) ?? "/var/run/docker.sock"
    }:/var/run/docker.sock`,
    ...Deno.args,
  ]),
);
