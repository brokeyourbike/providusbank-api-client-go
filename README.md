# providusbank-api-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/brokeyourbike/providusbank-api-client-go.svg)](https://pkg.go.dev/github.com/brokeyourbike/providusbank-api-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/providusbank-api-client-go)](https://goreportcard.com/report/github.com/brokeyourbike/providusbank-api-client-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/7764dfd1735596f6e9c1/maintainability)](https://codeclimate.com/github/brokeyourbike/providusbank-api-client-go/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7764dfd1735596f6e9c1/test_coverage)](https://codeclimate.com/github/brokeyourbike/providusbank-api-client-go/test_coverage)

Providus Bank API Client for Go

## Installation

```bash
go get github.com/brokeyourbike/providusbank-api-client-go
```

## Usage

### Account

```go
accountClient := providusbank.NewAccountClient("providusbank.com", "client_id", "client_secret")
accountClient.CreateDynamicAccount(context.TODO(), "John Doe")
```

### Transfer

```go
transferClient := providusbank.NewTransferClient("providusbank.com", "username", "password")
transferClient.GetNIPBanks(context.TODO())
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/brokeyourbike/providusbank-api-client-go/blob/main/LICENSE)
