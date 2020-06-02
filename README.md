## watcher

Watcher следит за указанными ему tcp/udp серверами, и если какие-то из серверов прекратили работу, watcher предупредит Вас об этом в ВКонтакте

## Установка

```shell script
go get github.com/GreenWix/watcher
go build -o main
```

## Запуск

```shell script
./main config=/абсолютный/путь/до/конфига.json
```

## TODO список
+ Поддержка UDP серверов
+ Указание байтов, после отправки которых UDP серверам, должен быть получен ответ

## Конфиг

Поле | Описание
------------ | -------------
vk_token | токен группы ВКонтакте
vk_chat_id | id чата ВКонтакте, в который будут присылаться уведомления
time | Раз в это время (в секундах) будет осуществляться проверка серверов (при условии, что последняя проверка watcher'a была успешной)
servers | сервера

Описание полей сервера

Поле | Описание
------------ | -------------
name | имя сервера (нужно для уведомлений, чтобы Вы поняли какой именно сервер упал)
addr | адрес сервера 
protocol | tcp/udp. Протокол, используемый сервером
mentions_text | строка, содержащая список упоминаний пользователей, которые ответственны за данный сервер

Если последняя проверка watcher'а была неуспешной, то watcher будет проверять каждую секунду, пока результат проверки не станет успешным

Пример конфига:
```json
{
  "vk_token": "token",
  "vk_chat_id": 1,
  "time": 10,
  "servers": [
    {
      "name": "my http server",
      "addr": "127.0.0.1:8080",
      "protocol": "tcp",
      "mentions_text": "@all"
    },
    {
      "name": "my udp server",
      "addr": "127.0.0.1:1234",
      "protocol": "udp",
      "mentions_text": "@online"
    }
  ]
}
```