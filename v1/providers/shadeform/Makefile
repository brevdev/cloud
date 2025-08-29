generate-shadeform-client:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.8.0 generate \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		-i https://raw.githubusercontent.com/shadeform/docs/refs/heads/main/openapi.yaml \
		-g go \
		--git-user-id brevdev \
		--git-repo-id cloud \
		-o /local/gen/shadeform
	sudo chown -R $(shell id -u):$(shell id -g) gen/shadeform
	gofmt -s -w gen/shadeform
	rm -rf gen/shadeform/go.mod gen/shadeform/go.sum
