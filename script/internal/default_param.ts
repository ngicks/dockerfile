export const defaultRun = () => [
  "docker",
  "run",
  "-itd",
  "--mount",
  `type=volume,src=git,dst=${Deno.env.get("GIT_MOUNT_DST") ?? "/mnt/git"}`,
  "--mount",
  `type=volume,src=config,dst=${
    Deno.env.get("CONFIG_MOUNT_DST") ?? "/mnt/config"
  }`,
  "--mount",
  "type=volume,src=certs,dst=/usr/local/share/ca-certificates/external",
  ...(envIsSetAndNonZero("TMP_AS_TMPFS")
    ? [
      "--mount",
      "type=tmpfs,dst=/tmp",
    ]
    : []),
  ...(envIsSetAndNonZero("MOUNT_VSCODE_SERVER")
    ? [
      "--mount",
      "type=volume,dst=/root/.vscode-server",
    ]
    : []),
];

const envIsSetAndNonZero = (envStr: string): boolean => {
  const env = Deno.env.get(envStr);
  return typeof env === "string" && env !== "" && env !== "0";
};
