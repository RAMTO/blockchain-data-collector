# Blockchain data collector

Simple app for extracting data from contracts & MongoDB and making some analytics calculations.

## Settings

App relies on some external configuration that should be present in .env file.

```sh
MONGO_DB=<MONGO_DB_CONNECTION_STRING>
ETH_NODE_URL=<ETHERIUM_NODE_URL OR INFURA_API>
TENANT_DB=<TENAT_DB_NAME>
CLIENT_DB=<CLIENT_DB_NAME>
```

Other part of the configuration should be provided in the main function:

```sh
config := make(map[string]string)
// Tenant data
config["tenant"] = ""
// Address of the Liquidity Pool
config["poolPairAddress"] = ""
// Name of the tenant collection
config["txCollectionName"] = ""
```

## Run the app

```sh
go run main.go
```

## Run the tests

```sh
cd tests
go test
```
