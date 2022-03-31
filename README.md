# rakuten-interview
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

## A Few Hiccups

In some places i have used third party libraries which does not concur with the task description.
The api is not full compact so there are a few errors.