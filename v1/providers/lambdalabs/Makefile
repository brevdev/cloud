generate-lambdalabs-client:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.8.0 generate \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		-i https://cloud.lambdalabs.com/static/api/v1/openapi.yaml \
		-g go \
		--git-user-id brevdev \
		--git-repo-id cloud \
		-o /local/gen/lambdalabs
	sudo chown -R $(shell id -u):$(shell id -g) gen/lambdalabs
	gofmt -s -w gen/lambdalabs
	rm -rf gen/lambdalabs/go.mod gen/lambdalabs/go.sum