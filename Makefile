oapi_codegen_version=v2.4.1
install:
	go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(oapi_codegen_version)
	@mkdir api
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=./swagger/codegen_config.yaml ./swagger/petstore.yaml

clean:
	@echo "Cleaning go mod"
	go mod tidy
	@echo "Cleaning api folder"
	@rm -rf api