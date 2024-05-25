
<h1 align="center" style="font-weight: bold;">My Balance - telegram bot</h1>

<p align="center">
<a href="#tech">Technologies</a> |
<a href="#started">Getting Started</a> |
<a href="#routes">Commands</a>
</p>


<p align="center">Very simple telegram bot for checking your current balance in Monobank</p>


<p align="center">
<a href="https://github.com/RaxFord1/MyBalance">📱 Visit this Project</a>
</p>

<h2 id="technologies">💻 Technologies</h2>

- viper (for configs)
- tele (as telegram-bot api)
- golang

<h2 id="started">🍕 Prerequisites </h2>

In order to start using this project, you need to have API keys for telegram bot and monobank API.

To get telegram_token use this bot: https://t.me/BotFather

To get monobank api key use https://api.monobank.ua/index.html


<h2 id="started">🚀 Getting started</h2>

0. Clone project
```bash
git clone  https://github.com/RaxFord1/MyBalance
```
1. Download golang
2. Download dependencies
```bash
go mod tidy
```
3. Fill your environment variables as in config/example.env with your values
4. Run application
```bash
go run .\cmd\telegrambot\main.go
```


<h2 id="routes">📍 Commands</h2>

Available commands for bot:

| <kbd>/start</kbd>     -   returns starting message

| <kbd>/balance</kbd>   -   gives current balance info (black card)

| <kbd>/statement</kbd> -   gives statement of your charges for last day (max 500)

| Once a day it automatically sends you statement for the last day


<h2 id="routes"> Notes:</h2>

I host this bot on heroku, so project struct is refined for that deployment environment, but you can change 
cmd/telegrambot/main.go

``config.SetConfigType(ctx, config.TypeOfConfigFromMemory)``
to
``config.SetConfigType(ctx, config.TypeOfConfigFromFile)``

and create `config/local.env` file to get env variables from file. 
