

  * `ctrl+shift+p`
  * serach for preferences
  * serach for proto3
  * click "Edith in settings.json"
  * replace "protoc" with : `"protoc": { "options": [ "--proto_path=protos",]}`


## generate API documentation for the gRPC with OpenAPI:
  * `protoc-gen-openapiv2 --help`


### Share documentation as static front-end with [Swagger UI](https://github.com/swagger-api/swagger-ui)

### Embed the static files as binary with [statik](https://github.com/rakyll/statik)
  * `go install github.com/rakyll/statik`
  * `statik -help`