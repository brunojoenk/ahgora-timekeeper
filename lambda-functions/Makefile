OS=$(shell uname -s)

build: ## Build staticailly linked binary
	GOOS=linux go build -o main && zip main.zip main && rm main
.PHONY: build