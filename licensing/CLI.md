# CLI Usage

### Table of Contents:

- [Introduction](#introduction)
- [Ways to install](#ways-to-install)
- [Usage](#usage)

### Introduction:

The corlink server comes with a CLI to make it easier to run the server. The CLI is written in Go and is available as a binary, npm package, and docker image.

### Ways to install:

- **npm**:
  ```bash
  npm install -g corlink-server
  ```

- **go**:
  ```bash
  go install github.com/ruby-network/corlink/license@latest
  ```

- **docker**:
    ```bash
    docker run -d -p 8080:8080 motortruck1221/corlink
    ```
    Or you can use the `docker-compose.yml` file: [./docker-compose.yml](./docker-compose.yml)

- **release**:
  You can download the latest release from the [releases page](https://github.com/ruby-network/corlink/releases).

### Usage:

First, create a `.env` file in the directory you want to run the server in. The `.env` file should look like this:
```env
ADMIN_KEY=1234
DB_HOST=localhost
DB_USER=changeme
DB_PASS=changeme
DB_NAME=changeme
#default postgres port
DB_PORT=5432
```

> [!WARNING]
> Make sure to have POSTGRES installed and running. If you don't have POSTGRES installed, you can use SQLITE instead.

> [!IMPORTANT]
> If you want, you can use SQLITE instead of POSTGRES. If you choose to do this (not recommended), you don't need to add the `DB_HOST`, `DB_USER`, `DB_PASS`, `DB_NAME`, or `DB_PORT` to the `.env` file. The server will also, automatically create the db file.

Edit the values to match your database and set the `ADMIN_KEY` to a random string.


To run with default settings:
```bash
corlink-server start
```

| Flag | Short Flag | Description | Default | Usage | Short Usage |
| ---- | ---------- | ----------- | ------- | ----- | ------------ |
| `--port` | `-p` | The port to run the server on | `8080` | `corlink-server start --port 8080` | `corlink-server start -p 8080` |
| `--host` | `-H` | The host to run the server on | `0.0.0.0` | `corlink-server start --host localhost` | `corlink-server start -H localhost` |
| `--sqlite` | `-s` | Use sqlite instead of postgres | `false` | `corlink-server start --sqlite` | `corlink-server start -s` |
| `--directory` | `-d` | The directory to run the server under | `/` | `corlink-server start --directory /directory` | `corlink-server start -d /directory` |
