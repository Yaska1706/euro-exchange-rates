
## Run Application
Make a .env file and edit the values 

``` cat .env.example > .env ```

The file should look like this:

```
# database configurations
DB_HOST = "localhost"
DB_PORT = "5432"
DB_USER = "postgres"
DB_NAME = "exchange"
DB_PASSWORD = ""

# port for app to listen and serve
LISTEN_ADDRESS = "127.0.0.1"
LISTEN_PORT = "8080"

```

Edit the variables to your preference

After configuring environment variables , run the application with the follow command
#### install dependecies
```make deps```

#### run app

``` make run ```

### Endpoints
`GET /rates/latest`

### Response
```json
{
    "base": "EUR",
    "rates": {
        "AUD": 1.5339,
        "BGN": 1.9558,
        "USD": 1.2023,
        "ZAR": 14.8845
    }
}
```
### Endpoint
`GET /rates/YYYY-MM-DD`

### Response
```json
{
    "base": "EUR",
    "rates": {
        "AUD": 1.5339,
        "BGN": 1.9558,
        "USD": 1.2023,
        "ZAR": 14.8845
    }
}
```

### Endpoint
`GET /rates/analyze`

### Response
```json
{
    "base": "EUR",
    "rates_analyze": {
        "AUD": {
            "min": 1.4994,
            "max": 1.5693,
            "avg": 1.5340524590163933
        },
        "BGN": {
            "min": 1.9558,
            "max": 1.9558,
            "avg": 1.9557999999999973
        },
        "USD": {
            "min": 1.1562,
            "max": 1.2065,
            "avg": 1.1783852459016388
        },
        "ZAR": {
            "min": 14.7325,
            "max": 17.0212,
            "avg": 16.06074426229508
        }
    }
}
```


## A Few SetBacks

The api is not full compact so there are a few errors.
