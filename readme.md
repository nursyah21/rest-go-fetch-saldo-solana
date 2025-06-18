# golang rest fetch saldon solana

## how to run

you need to install air golang

```sh
air
```

## how to build

```sh
go build -ldflags="-s -w"  -o main.exe ./src
```

## testing

```sh
## get solana
curl -X POST localhost:5000/api/get-balances -H "Content-Type: application/json" -d '{"wallets":["2k5AXX4guW9XwRQ1AKCpAuUqgWDpQpwFfpVFh3hnm2Ha","2k5AXX4guW9XwRQ1AKCpAuUqgWDpQpwFfpVFh3hnm2Ha"]}'
```