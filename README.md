# svc.pro.payload.bins

## Environment

When running the container you can use the following environment variables:

| ENV             | Description                           | Default
| --------------- | ------------------------------------- | -------
| `DB_PROTO`      | Postgres compatible DSN               | `postgresql`
| `DB_HOST`       | Postgres compatible DSN               | `localhost`
| `DB_PORT`       | Postgres compatible DSN               | `5432`
| `DB_USER`       | Postgres compatible DSN               |
| `DB_PASS`       | Postgres compatible DSN               |
| `DB_DATABASE`   | Postgres compatible DSN               | `payloadpro`

## Testing

The source code can be tested with:

```
go test ./...
``` 
