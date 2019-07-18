# PacktPub Free Ebook Telegram Bot

This is a simple bot that sends [PacktPub's](https://www.packtpub.com) current [free ebook offer](https://www.packtpub.com/packt/offers/free-learning) to a [Telegram](https://telegram.org) chat.

## Building

On unix based systems, run `./build.sh` on the root of the project.

## Instalation

This bot expects a `config.json` file to be present at the same directory as itself, with the following contents:

```json
{
    "token": "<your-bot-token>",
    "chat_id": <the-chat-id-to-send-the-message-to>,
    "use_eu_url": true/false
}
```

__Note__: The `use_eu_url` flag changes the message's link to the EU free offers website instead.

This bot does not run automatically, but you can configure a service in you operating system of choice to run it periodically.
