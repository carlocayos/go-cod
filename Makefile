.PHONY: swagger-gen
swagger-gen:
	docker run --rm -v $(PWD):/go/src/github.com/carlocayos/go-cod \
		-w /go/src/github.com/carlocayos/go-cod/v2/api \
		--entrypoint scripts/generate-swagger-client.sh \
		quay.io/goswagger/swagger:v0.25.0

.PHONY: swagger-docs
swagger-docs: ## preview the API documentation
	@echo "API docs preview will be running at http://localhost:9000"
	@docker run --rm \
		-v $(PWD)/api/specs/v1.0.0/cod-openapi-v1.0.0.yaml:/usr/share/nginx/html/swagger.yaml \
		-v $(PWD)/api/specs/v1.0.0/json-schema:/usr/share/nginx/html/json-schema \
		-e 'REDOC_OPTIONS=hide-hostname="true" lazy-rendering' \
		-p 9000:80 \
		bfirsh/redoc:1.14.0