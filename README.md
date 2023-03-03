# GoCloudPlaylist
### Тестовое задание GoCloudCamp.
Сервис для управления и изменения музыкального плейлиста.
#### Сервис обладает следующими возможностями:

* Play - начинает воспроизведение
* Pause - приостанавливает воспроизведение
* AddSong - добавляет в конец плейлиста песню
* Next - воспроизвести след песню
* Prev - воспроизвести предыдущую песню
* Status - статус воспроизведения

#### Стек технологий
`GO`, `Docker`, `PostgreSQL`, `gRPC`, `HTTP`, `Swagger`, `Make`

___
## Сервис

## Запуск

Для запуска сервиса в Docker-контейнере выполните `make build`, для перезапуска `make restart`

Для запуска локально поднимите `PostgreSQL` на вашей машине на порту 5432(Или измените `dsn` в `config.yaml`
,с базой данных `PlaylistService` и выполните команду `go run cmd/main.go`.

Для для генерации .pb файлов - `make gen`. Для повторной генерации `make re-gen`.

Для генерации `swagger`-документации - `make swagger`.

Для тестирования - `make test`

## Доступ через HTTP

Для HTTP обработчиков доступна `swagger-документация` и сервис удобно использовать через нее(http://localhost:8080/swagger/) 

### Play
GET-запрос который ничего не принимает на вход, а в ответе возвращает `json` ответ или код и сообщение ошибки

Пример:

Ответ:

```json
{
  "song_name": "Bohemian Rhapsody - Queen",
  "song_duration": "00:00:30",
  "playback_status": "Bohemian Rhapsody - Queen plays at 00:00:00 of 00:00:30"
}
```

### Pause
GET-запрос который ничего не принимает на вход, а в ответе возвращает `json` ответ или код и сообщение ошибки

Или код и сообщение ошибки

Пример:

`curl -X 'GET' \
'http://localhost:8080/pause' \
-H 'accept: application/json'`

Ответ:

```json
{
  "song_name": "Waka Waka - Shakira",
  "song_duration": "00:00:25",
  "playback_status": "Waka Waka - Shakira paused at 00:00:06 of 00:00:25"
}
```

### Status
GET-запрос который ничего не принимает на вход, а в ответе возвращает `json` ответ или код и сообщение ошибки

Или код и сообщение ошибки

Пример:

`curl -X 'GET' \
'http://localhost:8080/status' \
-H 'accept: application/json'`

Ответ:

```json
{
  "song_name": "Waka Waka - Shakira",
  "song_duration": "00:00:25",
  "playback_status": "Playback status: Waka Waka - Shakira paused on 00:00:06 of 00:00:25"
}
```

### Next song
GET-запрос который ничего не принимает на вход, а в ответе возвращает `json` ответ или код и сообщение ошибки или код и сообщение ошибки

Или код и сообщение ошибки

Пример:

`curl -X 'GET' \
'http://localhost:8080/next_song' \
-H 'accept: application/json'`

Ответ:

```json
{
  "song_name": "Nothing Else Matters - Metallica",
  "song_duration": "00:00:32",
  "playback_status": "Switched to next song: Nothing Else Matters - Metallica"
}
```
### Previous song
GET-запрос который ничего не принимает на вход, а в ответе возвращает `json` ответ или код и сообщение ошибки или код и сообщение ошибки

Пример:

`curl -X 'GET' \
'http://localhost:8080/prev_song' \
-H 'accept: application/json'`

Ответ:

```json
{
  "song_name": "Waka Waka - Shakira",
  "song_duration": "00:00:25",
  "playback_status": "Switched to previous song: Waka Waka - Shakira"
}
```

### Add song
POST-запрос, принимающий на вход json формата  
```json
{
"song_duration": "00:00:10",
"song_name": "string1"
}
```
Длительность обязательно дожна быть в формате `00:01:30`.

A в ответе возвращает `json` со списком всех композиция или код и сообщение ошибки или код и сообщение ошибки

Пример:

`curl -X 'POST' \
'http://localhost:8080/add_song' \
-H 'accept: application/json' \
-H 'Content-Type: application/json' \
-d '{
"song_duration": "00:00:30",
"song_name": "Bohemian Rhapsody - Queen"
}'`

Ответ:

```json
[
  {
   "song_name": "Bohemian Rhapsody - Queen",
    "Duration": 30
  }

]
```

### Delete song
GET-запрос который принимает на вход название песни в параметре `name`.

A в ответе возвращает `json` со списком всех композиция или код и сообщение ошибки или код и сообщение ошибки



Пример:

`curl -X 'GET' \
'http://localhost:8080/delete_song?name=Bohemian Rhapsody - Queen' \
-H 'accept: application/json'`

Ответ:

```json
[
  {
    "song_name": "Bohemian Rhapsody - Queen",
    "Duration": 30
  }
]
```

`Swagger`-докуметация:

![Image alt](pictures/Screenshot%202023-03-02%20at%2015.03.09.png)

## Доступ через gRPC

<details>
<summary>protobuf</summary>

```proto
syntax = "proto3";

option go_package = "./";

package api;

service GoCloudPlaylist {
  rpc AddSong(AddRequest) returns (PlaylistResponse) {}
  rpc DeleteSong(SongNameForDelete) returns (PlaylistResponse) {}
  rpc PlaySong(Empty) returns (SongProc) {}
  rpc PauseSong(Empty) returns (SongProc) {}
  rpc Next(Empty) returns (SongProc) {}
  rpc Prev(Empty) returns (SongProc) {}
  rpc Status(Empty) returns (SongProc) {}
}

message SongProc {
  string name = 1;
  string time = 2;
  string status = 3;
}

message Song {
  string name = 1;
  string duration = 2;
}
message PlaylistResponse {
  repeated Song Playlist = 1;
}
message AddRequest {
  string name = 1;
  string time = 2;
}

message SongNameForDelete {
  string name = 1;
}

message Empty {
}

```
</details>

gRPC эндпоинты реализованы аналогично. Имеются следующие методы:
* AddSong
* DeleteSong
* PlaySong
* PauseSong
* Next
* Prev
* Status

gRPC эндпоинты протестированы при помощи утилиты `Evans`:

![Image alt](pictures/Screenshot%202023-03-02%20at%2015.02.11.png)

### Тесты
В сервисе покрыты тестами все `HTTP` и `gRPC` эндпоинты, а также переиспользуемый код

### Пример использования
Пример использования сервиса находится в директори `client`

Для запуска - `go run client/client.go` (При запущенном сервисе)

В коде поочередно отправляются `gRPC` и `HTTP` запросы показывая работу сервиса.

В файле `client/example.http` есть примеры `HTTP`-запросов
