package server

//go:generate go run github.com/jdotw/oapi-gokit-codegen/cmd/oapi-gokit-codegen --generate=types,chi-server --package=server -o server.gen.go ../test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
