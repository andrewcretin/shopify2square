build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/shopify2square shopify2square/main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build schema
	sls deploy --verbose --stage=staging --alias=staging

.PHONY: deployProd
deploy: clean build schema
	sls deploy --verbose --stage=prod --alias=prod

.PHONY: plugins
plugins:
	sls plugin install -n serverless-prune-plugin
	sls plugin install -n serverless-plugin-log-retention
	sls plugin install -n serverless-aws-alias
