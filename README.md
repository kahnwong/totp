# totp

Inspired by <https://github.com/yitsushi/totp-cli>. Because I need autocomplete in fish.

## Setup

1. Install

```bash
go install github.com/kahnwong/totp@latest
```

2. create a config file in `~/.config/pgconn/db.yaml`

```yaml
- name: sample-db
  hostname: localhost
  proxy: # this block is optional
    kind: cloud-sql-proxy
    host: $GCP_PROJECT:$GCP_REGION:$INSTANCE_IDENTIFIER
  roles:
    - username: postgres
      password: postgrespassword
      dbname: sample_db

# if using ssh tunnelling
proxy:
  kind: ssh
  host: $SSH_CONFIG_HOST
```

## Available commands

```bash
connect DATABASE ROLE
list
    databases
    roles DATABASE
```

## Examples

`list databases`

```bash
❯ pgconn list databases
Available databases:
    nuc-postgres
    local-postgres
```

`list roles`

```bash
❯ pgconn list roles nuc-map
Database: nuc-postgres
Roles:
    postgres
```

`connect`

```bash
❯ pgconn connect nuc-map postgres
Server: PostgreSQL 15.3
Version: 3.5.0
Home: http://pgcli.com
postgres@192:map>
```
