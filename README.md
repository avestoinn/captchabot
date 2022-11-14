# Captcha Bot
A Telegram bot that can distinguish robots from humans via simple captcha check made from random blurred text on an
image

## Installation
<hr>
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
<hr>
This part is on development and will be published as soon as possible.

## Example (Demo)
<hr>
You can check a demo version of this bot: <a href="https://captcha_superbot.t.me">@captcha_superbot</a>