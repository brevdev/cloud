SPEC_VERSION    ?= v2.40.9
SPEC_FILE       ?= swagger-${SPEC_VERSION}.yaml
SPEC_FILE_FINAL ?= swagger-${SPEC_VERSION}.final.yaml
OUTPUT_DIR      ?= launchpad
# To understand disallowAdditionalPropertiesIfNotPresent=false
# Check out https://openapi-generator.tech/docs/generators/go/#config-options
generate-launchpad-client:
	go run ./scripts/finalize-swagger ${SPEC_FILE} ${SPEC_FILE_FINAL}
	rm -rf gen/${OUTPUT_DIR}
	mkdir -p gen/${OUTPUT_DIR}
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.8.0 generate \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		-i /local/${SPEC_FILE_FINAL} \
		-g go \
		--git-user-id brevdev \
		--git-repo-id cloud \
		--skip-validate-spec \
		-o /local/gen/${OUTPUT_DIR}
	sudo chown -R $(shell id -u):$(shell id -g) gen/${OUTPUT_DIR}//
	find gen/${OUTPUT_DIR} -name "*.go" -type f -exec sed -i.bak 's|openapiclient "github.com/brevdev/cloud"|openapiclient "github.com/brevdev/cloud/v1/providers/launchpad/gen/launchpad"|g' {} \; && find gen/${OUTPUT_DIR} -name "*.go.bak" -delete
	gofmt -s -w gen/${OUTPUT_DIR}
	rm -rf gen/${OUTPUT_DIR}/go.mod gen/${OUTPUT_DIR}/go.sum
