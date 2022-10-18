# livereload on target dev
FROM golang:1.19 AS dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /home/app

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /home/wait
RUN chmod +x /home/wait

EXPOSE 3000

CMD ["/bin/bash", "-c", "/home/wait && air"]

# build for prod
FROM golang:alpine3.16 AS build

WORKDIR /home/app

COPY . .

RUN go build -o /home/go.app

# target prod
FROM alpine:3.16 AS prod

WORKDIR /home

COPY --from=build /home/go.app /home/go.app

# !important for github.com/joho/godotenv file .env
COPY .env /home/.env

EXPOSE 3000

CMD ["/home/go.app"]

## DOCS
FROM node:16.16-alpine as docs

# set working directory
WORKDIR /app

COPY ./docs ./

RUN npx redoc-cli bundle -o index.html swagger.json

# start app
CMD ["node", "index.js"]