# Update checker

## Prerequisites

Make sure to create a `.env` file containing the keys present in the `.env.example` file.
An API token can be retrieved from Telegram itself [following their steps](https://core.telegram.org/api/obtaining_api_id).
A chat ID can be retrieved by starting `outputChatID.go` in the `setup` folder and sending a message to your bot username:

```go run setup/outputChatID.go```

## Building Docker image

`docker build -t update-checker .`

## Running Docker image

You can run the image directly:

```docker run -it --rm --name update-checker update-checker```

or you can run it by using docker compose:

```docker-compose up```
