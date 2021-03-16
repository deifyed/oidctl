# oidctl

## Requirements
* Go 1.16

## Installation

```shell
make install
```

This will install oidctl in your `~/.local/bin` folder. Ensure this folder is added to your `PATH`.

To install `oidctl` somewhere else, pass `INSTALL_DIR=newlocation` to the make install command.

## Configuration

```shell
export DISCOVERY_URL=https://auth.provider.com/.well-known/openid-configuration
export CLIENT_ID=myclientid
export CLIENT_SECRET=longasssecret
```

## Usage

```shell
oidctl client-credentials AUDIENCE

# Example
oidctl client-credentials https://example.com

# Result
```
