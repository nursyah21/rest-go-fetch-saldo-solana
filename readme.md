# golang rest fetch saldo solana

## prepare
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

```sh
## optional, if you need local mongo
docker compose up -d
```

## how to run

```sh
## for windows you can use git bash to run this
./watch.sh
```

## testing

```sh
./rest.sh
```