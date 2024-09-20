# cutshort-url

## Getting Started

#### Prerequisites
- Docker

After the Docker is set up clone the repository, create an .env file and paste the following contents in it.

```
APP_PORT={YOUR_PORT}
DB_ADDR="cache:{YOUR_PORT}"
OBJECT_DB_STORE_ADDR="db:{YOUR_PORT}"
OBJECT_DB_NAME="url_store"
```

Once done with .env run the following command and the compose service should be up.

```
docker compose up -d
```

Fire up the below curl command to test if the service is up and running fine.

```
curl \
--request GET 'localhost:{YOUR_PORT}/api/shorten' \
--header 'Content-Type: application/json' \
--data '{"url": "https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/insert/"
}'
```

It should return the below response

```
{"url":"hBxM5A6","custom_short":"","expiry":0,"x_rate_remaining":0,"x_rate_limit_reset":0}
```

Note: Replace `YOUR_PORT` placeholder with whatever port you have assigned to your api service (`APP_PORT` variable in `.env`)
