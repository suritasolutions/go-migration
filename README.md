# go-migration
Golang package for database migration

## DB docker command for testing

```bash
docker run --name go-migration -e POSTGRES_USER=user -e POSTGRES_PASSWORD=mysecretpassword -e PGDATA=/var/lib/postgresql/data/pgdata -v ./.db:/var/lib/postgresql/data -p 5432:5432 postgres
```