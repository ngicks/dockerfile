export const run = async (cmd: string[]): Promise<number> => {
  if (Deno.env.get("DRY_RUN")) {
    console.log(cmd);
    return 0;
  }

  const p = Deno.run({
    cmd,
  });

  p.stdout?.readable.pipeTo(Deno.stdout.writable);
  p.stderr?.readable.pipeTo(Deno.stderr.writable);

  const { code } = await p.status();
  return code;
};
