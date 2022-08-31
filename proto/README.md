# **What are protocol buffers?**

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

---
## **[proto compiler](https://grpc.io/docs/protoc-installation/)**
```
sudo apt install -y protobuf-compiler

protoc --version
```

---
## **[Go plugins](https://grpc.io/docs/languages/go/quickstart/)**

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

  1. `ctrl+shift+p`
  2. serach for preferences
  3. serach for proto3
  4. click "Edith in settings.json"
  5. replace "protoc" with : 
  ```
  "protoc": { 
    "options": [ 
      "--proto_path=protos",
    ]
  }
  ```

---
## **generate API documentation for the gRPC with OpenAPI:**
`protoc-gen-openapiv2 --help`

---
## **Share documentation as static front-end with [Swagger UI](https://github.com/swagger-api/swagger-ui)**

Swagger UI allows anyone — be it your development team or your end consumers — to visualize and interact with the API’s resources without having any of the implementation logic in place. It’s automatically generated from your OpenAPI (formerly known as Swagger) Specification, with the visual documentation making it easy for back-end implementation and client-side consumption.

---
## **Embed the static files as binary with [statik](https://github.com/rakyll/statik)**

statik allows you to embed a directory of static files into your Go binary to be later served from an http.FileSystem.

`go install github.com/rakyll/statik`

`statik -help`