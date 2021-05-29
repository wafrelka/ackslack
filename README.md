## slucg

### Overview

`slucg` is a simple command proxy that posts Slack notification when the command exits.

### Build

```
$ make
```

### Usage

```
$ ./build/slucg -- echo "hello"
hello
[*] command `/bin/echo hello` succeeded
```

The message ``command `/bin/echo hello` succeeded`` is also sent to the specified Slack webhook.

### Configuration

#### Sample

```toml
webhook_url = "https://hooks.slack.com/services/AAAAA/BBBBB/0000000000"
```

#### Location

If `slucg` is invoked with the `--config PATH` option, `PATH` will be used.
Otherwise, it searches the following places and uses the first one found.

1. `$USER_CONFIG_DIR/slucg/slucg.toml`
1. `$USER_CONFIG_DIR/slucg/config.toml`
1. `$USER_HOME_DIR/.slucg.toml`
1. `./slucg.toml`

`$USER_CONFIG_DIR` and `$USER_HOME_DIR` are determined by the Golang's
[`os.UserConfigDir`](https://golang.org/pkg/os/#UserConfigDir) and
[`os.UserHomeDir`](https://golang.org/pkg/os/#UserHomeDir) functions.
