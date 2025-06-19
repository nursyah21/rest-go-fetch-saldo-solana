# golang rest fetch saldo solana

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

```sh
## linux
sh watch.sh
```

```sh
## windows
watch.bat
```

## how to build

```sh
go build -ldflags="-s -w"  -o main.exe ./src
```

## testing

you can use httpie, and import `httpie-collection-api.json`