## Postgres DB setup
```bash
docker pull postgres
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 8002:5432 -d --rm postgres
```
### Migration
```bash
migrate create -ext sql -dir schema -seq init
migrate -path schema -database 'postgres://postgres:qwerty@0.0.0.0:8002/postgres?sslmode=disable' up
migrate -path schema -database 'postgres://postgres:qwerty@0.0.0.0:8002/postgres?sslmode=disable' down
```
### Troubleshooting
```bash
migrate -path schema -database 'postgres://postgres:qwerty@0.0.0.0:8002/postgres?sslmode=disable' force 1
```
or
```bash
update schema_migrations set version='000001', dirty=false;
```
