[![CircleCI](https://circleci.com/gh/PayloadPro/svc.pro.payload.bins/tree/master.svg?style=svg)](https://circleci.com/gh/PayloadPro/svc.pro.payload.bins/tree/master)
[![Maintainability](https://api.codeclimate.com/v1/badges/c6bcadf008abc77d6b91/maintainability)](https://codeclimate.com/github/PayloadPro/svc.pro.payload.bins/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/c6bcadf008abc77d6b91/test_coverage)](https://codeclimate.com/github/PayloadPro/svc.pro.payload.bins/test_coverage)

# svc.pro.payload.bins

Bin micro service 

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
