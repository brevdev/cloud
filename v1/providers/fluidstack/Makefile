generate-fluidstack-client:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.8.0 generate \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		-i https://docs.atlas.fluidstack.io/redocusaurus/infrastructure-api.yaml \
		-g go \
		--git-user-id brevdev \
		--git-repo-id cloud \
		-o /local/gen/fluidstack
	sudo chown -R $(shell id -u):$(shell id -g) gen/fluidstack
	gofmt -s -w gen/fluidstack
	rm -rf gen/fluidstack/go.mod gen/fluidstack/go.sum
