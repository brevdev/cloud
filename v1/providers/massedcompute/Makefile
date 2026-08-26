SPEC_VERSION    ?= v1.0.0
SPEC_FILE       ?= openapi-${SPEC_VERSION}.yaml
SPEC_FILE_FINAL ?= openapi-${SPEC_VERSION}.final.yaml
SPEC_PATCH      ?= openapi-${SPEC_VERSION}.patch
OUTPUT_DIR      ?= massedcompute

.PHONY: finalize-massedcompute-openapi generate-massedcompute-client
finalize-massedcompute-openapi:
	cp ${SPEC_FILE} ${SPEC_FILE_FINAL}
	patch -s -F 0 ${SPEC_FILE_FINAL} ${SPEC_PATCH}

generate-massedcompute-client: finalize-massedcompute-openapi
	rm -rf gen/${OUTPUT_DIR}
	mkdir -p gen/${OUTPUT_DIR}
	docker run --rm -v "${CURDIR}:/local" openapitools/openapi-generator-cli:v7.8.0 generate \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		-i /local/${SPEC_FILE_FINAL} \
		-g go \
		--git-user-id brevdev \
		--git-repo-id cloud \
		-o /local/gen/${OUTPUT_DIR}
	find gen/${OUTPUT_DIR} -name "*.go" -type f -exec sed -i.bak 's|openapiclient "github.com/brevdev/cloud"|openapiclient "github.com/brevdev/cloud/v1/providers/massedcompute/gen/massedcompute"|g' {} \; && find gen/${OUTPUT_DIR} -name "*.go.bak" -delete
	find gen/${OUTPUT_DIR} -name "*.md" -type f -exec sed -i.bak 's/[[:space:]]*$$//' {} \; && find gen/${OUTPUT_DIR} -name "*.md.bak" -delete
	gofmt -s -w gen/${OUTPUT_DIR}
	rm -f gen/${OUTPUT_DIR}/go.mod gen/${OUTPUT_DIR}/go.sum
