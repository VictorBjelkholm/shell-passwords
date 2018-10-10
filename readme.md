# Shell-Passwords (sp)
> Utility for avoiding having password/tokens ending up in your output and shell history

## Purpose

People commonly have various of tokens and passwords locally stored in their
shell history and sometimes in log outputs. This utility gives you a way to
avoid that by combining 1Password's password store with a simply CLI tool for
getting passwords and filtering output.

## Requirements

Install [1Password's CLI client `op`](https://support.1password.com/command-line/) first.

## Usage

### Login to 1Password

`eval $(op signin <subdomain>)`


### Get a password

```console
$ sp g GitHub
ThisIsMyVerySecretPassword
```

### Filter out password from output

```console
$ sp g GitHub | sp f GitHub
**************************

$ export TOKEN=$(sp g GitHub)
$ echo "Hello, this is my token: $TOKEN. It's nice eh?" | sp f GitHub
Hello, this is my token: **************************. It's nice eh?
```

## License

MIT 2018 - Victor Bjelkholm
