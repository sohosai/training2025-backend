# Go Web Backend Template

This uses:

- Go
- Gin
  - Web(HTTP) server library for Go
- sqlx
  - SQL interface library for Go (Also supports Rust)
- goose
  - Migration tool written in Go
  - To create migration file:
  ```
  docker exec -it container-name sh
  goose create migration-file-name sql
  ```
  - Then, file will be created in ./migrations/, edit it.
  - You can also use Go-based migration file.
- PostgreSQL
