
⚠️ **A work in progress**

## Golytics

A poorly named minimal self-hosted analytics service powered by Go and Sqlite. Very pre `v0.0.0`.  

## Setup

### Enviornment

A root `.env` is required

```
DATABASE_URL="sqlite:db/database.sqlite3"
ADMIN_PASSWORD=FOO_BARADMIN_PASSWORD
SESSION_KEY=FOO_BAR_SESSION_KEY
DATABASE_LOCATION="./ddbdatabase.sqlite3"
```
## Build

**Seed**

```
go build -o seed ./pkg/cmd/seed/main.go
```

**Server**

```
go build -o seed ./pkg/cmd/server/main.go
```

## Development

**Seed**

```
go run ./pkg/cmd/seed/main.go
```

**Server**

```
go run ./pkg/cmd/server/main.go
```

## Feature / todo list

- [x] Admin login via `POST`
- [x] In-memory session persistance
- [ ] Basic JS lib to load on client. ES5 compliant.
- [ ] Update `Dockerfile` for server build / deployment
- [ ] Figure out seed check, i.e. should seed be manually run, run controlled by env var, or either.
