# PyMax Hashes

Сервис метаданных для разных версий MAX, дополняющий [PyMax](https://github.com/MaxApiTeam/PyMax).
Хранит номер сборки, схему подписи и SHA-256-хеши сертификатов, DEX-файлов и нативных библиотек.

Чтение данных доступно без авторизации:

```sh
curl https://hashes.pymax.org/versions.json
```

- [API](./api.md) — запросы, ответы и ошибки.
- [Формат данных](./data.md) — поля метаданных.
- [Самостоятельный запуск](./self-hosting.md) — Docker Compose и настройки.
