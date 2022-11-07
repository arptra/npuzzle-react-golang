.DEFAULT_GOAL := build

NAME = npuzzle

DOCKER_HUB_LOGIN=dmarsell

all: test build ## Test and build

re: clean build run ## Clean, rebuild and run

test: ## Run all the tests
	go test -covermode=atomic -v -race -timeout=30s ./...
#	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

build: ## Build a exec file
	go build -o $(NAME)

image: ## Build docker image, args: ver=image-version
	docker build --build-arg name=$(NAME) -t $(addprefix $(DOCKER_HUB_LOGIN)/,$(NAME)):$(ver) .

mod: ## Check mod dependencies
	go mod tidy

run: build ## Build and run an exec file
	./$(NAME)

manual: build ## Build and run an exec file
	./$(NAME) manual

clean: ## Remove exec file
	go clean && rm -rf $(NAME)

fclean: clean ## Remove temporary files and exec file
	go clean -modcache

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#setup: ## Install all the build and lint dependencies
#	go get -u github.com/alecthomas/gometalinter
#	go get -u golang.org/x/tools/cmd/cover
#	gometalinter --install --update
#	@$(MAKE) build

#cover: test ## Run all the tests and opens the coverage report
#	go tool cover -html=coverage.txt

#fmt: ## Run goimports on all go files
#	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

#lint: ## Run all the linters
#	gometalinter --vendor --disable-all \
#		--enable=deadcode \
#		--enable=ineffassign \
#		--enable=gosimple \
#		--enable=staticcheck \
#		--enable=gofmt \
#		--enable=goimports \
#		--enable=misspell \
#		--enable=errcheck \
#		--enable=vet \
#		--enable=vetshadow \
#		--deadline=10m \
		./...

#ci: ## Run all the tests and code checks
#	lint test

.PHONY: all re run mod help clean build lint fmt cover test setup ci