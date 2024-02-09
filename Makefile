PHONY: generate-protos
generate-protos:
	protoc pkg/**/proto/*.proto --go_out=./ --go_opt=paths=import
