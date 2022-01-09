package packageB

//go:generate go run github.com/jdotw/oapi-gokit-codegen/cmd/oapi-gokit-codegen -generate types,skip-prune,spec --package=packageB -o externalref.gen.go spec.yaml
