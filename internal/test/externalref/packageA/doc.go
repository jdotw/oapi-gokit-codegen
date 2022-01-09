package packageA

//go:generate go run github.com/jdotw/oapi-gokit-codegen/cmd/oapi-gokit-codegen -generate types,skip-prune,spec --package=packageA -o externalref.gen.go --import-mapping=../packageB/spec.yaml:github.com/jdotw/oapi-gokit-codegen/internal/test/externalref/packageB spec.yaml
