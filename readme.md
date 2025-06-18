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

you can use httpie, and import `httpie-collection-api.json`