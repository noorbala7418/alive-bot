# Alive Bot
Simple Telegram Bot for Getting Status of your server with Golang.

### Environment Variables
- `TELEGRAM_BOT_DEBUG_MODE`: Enable or Disable Debug mode. Possible values: `true` or `false`
- `TELEGRAM_BOT_ADMIN_ID`: Telegram ID of admin user
- `TELEGRAM_BOT_TOKEN`: API Token


## Build docker image
```bash
docker build -t alive-bot:latest
```

## Run docker container

```bash
docker run -d --name alive-bot -e TELEGRAM_BOT_DEBUG_MODE=false -e TELEGRAM_BOT_ADMIN_ID=YOUR_TELEGRAM_USER_ID -e TELEGRAM_BOT_TOKEN="YOUR_BOT_TOKEN" --network host  alive-bot:latest
```