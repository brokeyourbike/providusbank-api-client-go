# providusbank-api-client-go
Providus Bank API Client for Go

## Installation

```bash
go get github.com/brokeyourbike/providusbank-api-client-go
```

## Usage

```go
client := providusbank.NewClient("token", signer)

err := client.Test(context.TODO(), "hello")
require.NoError(t, err)
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/brokeyourbike/providusbank-api-client-go/blob/main/LICENSE)
