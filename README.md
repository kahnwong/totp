# totp

Inspired by <https://github.com/yitsushi/totp-cli>. Because I need autocomplete in fish.

Bonus: it also copies the token to your clipboard.

## Setup

1. Install

```bash
go install github.com/kahnwong/totp@latest
```

2. create a config file in `~/.config/totp/totp.yaml`

```yaml
- org: personal
  accounts:
    - name: foo
      token: XXXXXXXXXXXXXXXXXXXXXXXXX
    - name: bar
      token: XXXXXXXXXXXXXXXXXXXXXXXXX
```

## Available commands

```bash
generate ORG ACCOUNT
list
    orgs
    accounts ORG
```

## Examples

`list orgs`

```bash
❯ totp list orgs
Available organizations:
  - foo
  - bar
```

`list accounts`

```bash
❯ totp list accounts foo
Organization: foo
Accounts:
  - a
  - b
  - c
```

`generate`

```bash
❯ totp generate foo a
XXXXXX
```
