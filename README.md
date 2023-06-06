# ozon_test

## Репозиторий содержит микросервис на языке программирования GO.

## Установка и запуск

Для установки микросервиса необходимо склонировать репозиторий на свой компьютер,
скопировать файл .env.example (с названием .env)
и запустить, используя команду:


docker compose up --build


## API

Микросервис предоставляет следующие API-методы:

POST - http://localhost:5006/api/
{
"link": "https://example.fun/docs"
}

GET - http://localhost:5006/api/
query : short_link 2COdFJmAmy
