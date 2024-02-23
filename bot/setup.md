# Setup

- There are two ways to setup the bot.

## Method 1: From source

### Prerequisites

- Ruby 3.2.2
- Bundler
- Git

1. Clone the repository

```bash
git clone https://github.com/ruby-network/corlink.git 
```

2. Get to where the bot is located

```bash
cd corlink/bot/
```

3. Install the dependencies

```bash
bundle install
```

4. Create a `.env` file in the root of the bot directory and add the following:

```env
DISCORD_TOKEN=changeme
GUILD_ID=123456789012345678
LICENSING_SERVER_URL=http://changeme.com
LICENSING_SERVER_KEY=changeme
OWNER_ID=changeme
```

5. Run the bot

```bash
bundle exec ruby main.rb
```

## Method 2: Docker & Docker Compose (Recommended)

### Prerequisites
- Docker
- If you want to use Docker Compose, you will need to install it as well.
- If you want to build the image yourself, you will need to install Git, Ruby, and Bundler.

#### Docker (prebuilt image)

This one is simple just run the following command:

1. Create a `.env` file in the root of the bot directory and add the following:

```env 
DISCORD_TOKEN=changeme
GUILD_ID=123456789012345678
LICENSING_SERVER_URL=http://changeme.com
LICENSING_SERVER_KEY=changeme
OWNER_ID=changeme
```

2. Run the bot

```bash
docker run -d --name corlink-bot --env-file .env ghcr.io/ruby-network/corlink-bot:latest
```

#### Docker (build image yourself)

1. Clone the repository

```bash
git clone https://github.com/ruby-network/corlink.git 
```

2. Get to where the bot is located

```bash
cd corlink/bot/
```

3. Build the image

```bash
docker build -t corlink-bot .
```

4. Edit the `.env` file in the root of the bot directory and add the following:

5. Run the bot

```bash
docker run -d --name corlink-bot --env-file .env corlink-bot
```

#### Docker Compose (prebuilt image)

1. Create a `docker-compose.yml` file in the root of a directory and copy the file [here](./docker-compose.yml).
2. Edit the ENV variables in the `docker-compose.yml` file.
3. Run the bot

```bash
docker-compose up -d
```

#### Docker Compose (build image yourself)

1. Clone the repository

```bash
git clone https://github.com/ruby-network/corlink.git 
```

2. Get to where the bot is located

```bash
cd corlink/bot/
```

3. Edit the `.env` env file 

4. Build the image

```bash
docker compose -f ./docker-compose.build.yml build
```

5. Run the bot

```bash
docker compose -f ./docker-compose.build.yml up -d
```
