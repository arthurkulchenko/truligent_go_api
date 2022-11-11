FROM golang:1.19.0-alpine3.15

WORKDIR /home/app

COPY . .

EXPOSE 8080

RUN chmod +x ./docker/go/docker-entrypoint.sh

ENTRYPOINT "./docker/go/docker-entrypoint.sh"
