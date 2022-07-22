## ackslack

### Overview

`ackslack` is a simple command wrapper that enables Slack notification when a command completes.

### Build

prerequisite: Go 1.13+

```sh-session
$ make
```

### Usage

```sh-session
$ ./build/ackslack -- echo "hello"
hello
[*] succeeded: `/bin/echo hello`
```

The message `` succeeded: `/bin/echo hello` `` is also sent to the specified Slack webhook url.

### Configuration

#### Sample

```toml
webhook_url = "https://hooks.slack.com/services/AAAAA/BBBBB/0000000000"
```

#### Location

You can pass the config file path via the `--config PATH` option.
Otherwise, `ackslack` searches the following paths and uses the first one it finds.

1. `$USER_CONFIG_DIR/ackslack/ackslack.toml`
1. `$USER_CONFIG_DIR/ackslack/config.toml`
1. `$USER_HOME_DIR/.ackslack.toml`
1. `./ackslack.toml`

`$USER_CONFIG_DIR` and `$USER_HOME_DIR` are determined by the Golang's
[`os.UserConfigDir`](https://golang.org/pkg/os/#UserConfigDir) and
[`os.UserHomeDir`](https://golang.org/pkg/os/#UserHomeDir) functions.
