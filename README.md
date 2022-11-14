# Captcha Bot
A Telegram bot that can distinguish robots from humans via simple captcha check made from random blurred text on an
image

## Installation
Create a `config.json` in a project's root with the params shown below:

```json5
{
  "bot": {
    "token": "<YOUR_BOT_TOKEN>"
  },
  "database": {
    "dsn": "<DESTINATION_DATABASE_PARAMS>"
    // In case of Sqlite this is the path to a .db file
  }
}
```

## Group admin commands
This part is on development and will be published as soon as possible.

## Example (Demo)
You can check a demo version of this bot: <a href="https://captcha_superbot.t.me">@captcha_superbot</a>