models:
	oapi-codegen -package=api -generate "types" docs/openapi.yaml > models.gen.go

gen:
	go generate ./...

builds:
	go build -o build/ ./...

test:
	go test -v ./...