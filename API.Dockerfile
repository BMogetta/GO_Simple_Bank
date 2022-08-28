##########  BUILD   ##########

FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

##########  RUN   ##########

FROM alpine:3.16
#TODO: add user to this container
WORKDIR /app
COPY --from=builder /app/main ./
#TODO: fix env file
COPY ["example.env", "start.sh", "wait-for.sh", "./"]
COPY postgres/migration ./postgres/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]