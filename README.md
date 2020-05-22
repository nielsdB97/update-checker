# Update checker

## Building Docker image

`docker build -t update-checker .`

## Running Docker image

Make sure to create a `.env` file containing the keys present in the `.env.example` file.

You can run the image directly:

```docker run -it --rm --name update-checker update-checker```

or you can run it by using docker compose:

```docker-compose up```
