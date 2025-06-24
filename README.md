# Todo-list
Проект реализует todo-list с авторизацией для создания задач, настройки их авто переназначения. Реализованы возможности изменения статуса задачи(выполнение, удаление). Реализованы все задачи со звёздочками.

## Содержание
- [Сборка через Docker](#сборка-через-docker)
- [Тестирование](#тестирование)



## Сборка через Docker
Настройки используемые в проекте для .env файла (
    TODO_PORT=7540
TODO_DBFILE=./scheduler.db
TODO_PASSWORD=123
)

Для сборки через Docker:
docker build -t todo-list .

docker run -d   --env-file .env -v $(pwd)/scheduler.db:/app/scheduler.db  -p 7540:7540   --name todo-app   todo-list


переходим на сайт [localhost](http://localhost:7540/)


## Тестирование
В файле настроек тестов (settings.go) указаны следующие параметры:

var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoYXNoIjpbMTY2LDEwMSwxNjQsODksMzIsNjYsNDcsMTU3LDY1LDEyNiw3MiwxMDMsMjM5LDIyMCw3OSwxODQsMTYwLDc0LDMxLDYzLDI1NSwzMSwxNjAsMTI2LDE1MywxNDIsMTM0LDI0NywyNDcsMTYyLDEyMiwyMjddfQ.pTkNTcofPoARNxIxJOui58YLCbx1vh0GguOO7ZXL7qA`

Для запуска тестов используем go test ./tests
