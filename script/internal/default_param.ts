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
  ...(Deno.env.get("TMP_AS_TMPFS")
    ? [
      "--mount",
      "type=tmpfs,dst=/tmp",
    ]
    : []),
];
