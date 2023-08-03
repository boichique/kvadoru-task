# Тестовое задание Kvado.ru

## Задание
Спроектировать базу данных, в которой содержится авторы книг и сами книги. Необходимо написать сервис, который будет по автору искать книги, а по книге искать её авторов.

## Требования к сервису
- Сервис должен принимать запрос по GRPC
- Должна быть использована база данных MySQL
- Код сервиса должен быть хорошо откомментирован
- Код должен быть покрыт unit тестами
- В сервисе должен лежать Dockerfile, для запуска базы данных с тестовыми данными
- Должна быть написана документация, как запустить сервис
- Плюсом будет, если в документации будут указания на команды, для запуска сервиса и его окружения, через Makefile

## Необходимые инструменты для запуска сервиса
На компьютере должны быть установлены:
- Docker (с возможностью использования docker compose)
- go
- [grpcui](https://github.com/fullstorydev/grpcui "grpcui") (устанавливается консольной командой `go install github.com/fullstorydev/grpcui/cmd/grpcui@latest`)

## Команды Makefile
Запуск сервиса и сборка клиента:
- `make run-service`

Остановка сервиса и удаление клиента:
- `make stop-service`

Запуск GUI-консоли для тестирования gRPC:
- `make grpc-gui`

Форматирование, проверка линтерами и прогон тестов:
- `make before-push`

Компиляция protobuf:
- `make protoc-books`

## Работа с сервисом
Сервис стартует с небольшим набором тестовых данных, что дает возможность сразу запускать запросы. Для запуска запросов можно использовать клиент. Ниже приведены примеры команд для запуска из директории проекта:

`./client/cli -author "J.K. Rowling"`

`./client/cli -book "Harry Potter and the Chamber of Secrets"`

#### Также можно использовать Postman:
Поиск авторов по книге
![GetAuthorsByBook](https://github.com/boichique/kvadoru_task/assets/87061629/ab064ff4-e2e7-4aaa-b0cc-e8c447c63214)

Поиск книг по автору
![GetBooksByAuthor](https://github.com/boichique/kvadoru_task/assets/87061629/10cac055-d051-45e3-8a0c-9b07d0f15489)

#### И можно использовать GUI-консоль grpcui:
Поиск авторов по книге
![grpcuiGetAuthorsByBook](https://github.com/boichique/kvadoru_task/assets/87061629/2281ac08-6ebe-49e1-8316-38bcc5f7b069)

Поиск книг по автору
![grpcuiGetBooksByAuthors](https://github.com/boichique/kvadoru_task/assets/87061629/5e67dde2-9e0a-46c7-ada1-78091252efda)
