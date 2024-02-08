PHONY: generate-protos
generate-protos:
	protoc pkg/**/proto/*.proto --go_out=paths=import
