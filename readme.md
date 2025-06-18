# golang rest fetch saldon solana

## clone and install
```sh
## clone from github
git clone --depth 1 https://github.com/nursyah21/rest-go-fetch-saldo-solana
cd rest-go-fetch-saldo-solana
```

```sh
## installing dependecies
go mod tidy
```

```sh
## after create .env you need to set environment in .env
cp .env.example .env
```

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