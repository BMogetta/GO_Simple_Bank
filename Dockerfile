##########  BUILD   ##########

FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

##########  RUN   ##########

FROM alpine:3.16
WORKDIR /app
RUN adduser \
    --disabled-password \
    --no-create-home \
    "go"
COPY --from=builder /app/main ./
COPY ["app.env", "start.sh", "wait-for.sh", "./"]
COPY postgres/migration ./postgres/migration
RUN chown -R go:go /app

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]