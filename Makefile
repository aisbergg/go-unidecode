VERSION 			:= ""

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

FUZZ_TIME ?= 1m

.PHONY: dev-setup fmt-lint test test-coverage fuzz display-coverage bench bench-compare release-tag release-push

## Setup dev dependencies
dev-setup:
	@echo Setup Dev Dependencies
	@bash tools/install.sh
	@pre-commit install

## Format and lint using pre-commit hooks
fmt-lint:
	@echo Formatting, linting and checking code
	@pre-commit run -a

## Run the tests
test:
	@echo Running tests
	@go test -race -v -gcflags '-N -l' ./pkg/...

## Run the tests with coverage
test-coverage:
	@echo Running tests with coverage
	@# we use "-gcflags '-N -l'", because that stops optimization from eating
	@# creating stacks with identical frames (PCs) and thus ruin our test results
	@go test -short -coverprofile cover.out -covermode=atomic -gcflags '-N -l' $$(go list ./pkg/... | grep -v testutils)

## Run fuzz testing
fuzz:
	@echo Running fuzz tests
	@bash tools/fuzz_all.sh $(FUZZ_TIME)

## Display test coverage
display-coverage:
	@echo Displaying test coverage
	@go tool cover -html=cover.out

## Run benchmarks
bench:
	@echo Running benchmark tests
	@cd benchmarks && go test -benchmem -bench=. && cd ..

## Run and compare benchmarks to previous run. Use `make release-tag PREVIOUS=bench-compare` to compare against a specifc run
bench-compare:
	@echo Running and comparing benchmarks against previous run
	@bash benchmarks/benchstat.sh "${PREVIOUS}"

## Stage a release (usage: make release-tag VERSION=v0.0.0)
release-tag: fmt-lint test
	@test -n "${VERSION}" || (echo -e "\nERR: You have to specify the next release version (e.g. 'VERSION=v0.0.0')!\n" && exit 1)
	@git diff-index --quiet HEAD -- || (echo -e "\nERR: You have uncommited changes. Please stash them before creating a release!\n" && exit 1)
	@echo Generating changelog
	@git-chglog -o CHANGELOG.md --next-tag "${VERSION}"
	@git add CHANGELOG.md
	@git commit -m "chore: update changelog for version ${VERSION}"
	@echo "Tagging release with version ${VERSION}"
	@git tag -s -a "${VERSION}" -m "${VERSION}"

## Push a release (warning: make sure the release was staged properly before doing this)
release-push:
	@echo Publishing release
	@git push --follow-tags
