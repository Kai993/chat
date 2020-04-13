.PHONY: help

go-build: ## ビルド
	@go build -o chat -v

go-test: ## すべてのテストを実行する
	@go test -v ./

go-clean: ## build/testなどを削除する
	@go clean

go-format: ## ソースをフォーマットする
	@goimports -w ./

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
