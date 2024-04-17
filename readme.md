# I am All Ears

A simple stub to log HTTP requests and respond with OK.

## Develop

```bash
go get
go run .
```

## Run

```bash
go build
PORT=3004 ./allears
# or
PORT=3004 nohup ./allears >>log.txt 2>&1 &
```
