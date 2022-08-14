# dockerfile

my dockerfiles

## Use deno

I'm tired of constantly googling bashscript syntax.

For installation, follow instruction:

https://deno.land/#installation

```shell
# Mac/Linux
curl -fsSL https://deno.land/install.sh | sh
# Windows
iwr https://deno.land/install.ps1 -useb | iex
```

## links

- [Docker: Guides](https://docs.docker.com/get-started/overview/)
- [Docker: Reference](https://docs.docker.com/reference/)
- [Docker Buildkit(syntax)](https://github.com/docker/buildx)

## Set up docker (things you always do when you are to set up your new gear.)

### Installation

- [Docker: Install Docker Engine on Debian](https://docs.docker.com/engine/install/debian/)

```shell
sudo apt-get remove docker docker-engine docker.io containerd runc

sudo apt-get update
sudo apt-get install ca-certificates curl gnupg lsb-release

sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

### Enable rootless mode

- [Docker: Run the Docker daemon as a non-root user (Rootless mode)](https://docs.docker.com/engine/security/rootless/)

```shell
sudo apt-get update && sudo apt-get install -y uidmap dbus-user-session

sudo systemctl disable --now docker.service docker.socket

dockerd-rootless-setuptool.sh install

# This env var may change in future or by each environment.
# Just use final line of output of dockerd-rootless-setuptool.sh.
echo "export DOCKER_HOST=unix:///run/user/$(id -u)/docker.sock" >> ~/.bashrc
```

### Set default Buildkit usage to true

- [Docker: Build images with BuildKit](https://docs.docker.com/develop/develop-images/build_enhancements/)

edit `/etc/docker/daemon.json` or in rootless mode, `~/.config/docker/daemon.json`.

```json
{ "features": { "buildkit": true } }
```

## Opinions

### Base image should be ubuntu.

Base should be ubuntu. And it should be determiticly-fixed version.

Famous, and popular. Attack surface is relatively narrow, and fixes should be delivered fast.

### Deno as a script engine.

Easy for scripting usage, for both developers and users.
Dependencies are small, single binary.

### Go as a template engine / code generation

Go has really strong template package.
Go itself is simple and compiler is fast. Very good for development helper, also production.