# Ratelet

---

## Overview

A service that provides exchange rate calculations and displays currency rates using real-time data from the Open Exchange Rates API.

## Development Dependencies
- Golang `1.24`

- Mock library
```
go install github.com/vektra/mockery/v3@v3.2.5
```


## Generate mocks

```
make mocks
```

## Unit Tests
```
make ut
```

## Run in development mode
```
make run
```

## Run the service in a Docker container
```
make image
make run-image
```

## How to consume the API
> \> Install [HTTPie CLI](https://httpie.io/cli)

- Getting rates
```
http localhost:8080/rates currencies==USD,PLN
```

- Exchanging cryptocurrencies
```
http localhost:8080/exchange from==USDT to==BEER amount==1
```