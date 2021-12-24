.DEFAULT_GOAL:=help

COLOR_ENABLE=$(shell tput colors > /dev/null; echo $$?)
ifeq "$(COLOR_ENABLE)" "0"
CRED:=$(shell tput setaf 1 2>/dev/null)
CGREEN:=$(shell tput setaf 2 2>/dev/null)
CYELLOW:=$(shell tput setaf 3 2>/dev/null)
CBLUE:=$(shell tput setaf 4 2>/dev/null)
CMAGENTA:=$(shell tput setaf 5 2>/dev/null)
CCYAN:=$(shell tput setaf 6 2>/dev/null)
CWHITE:=$(shell tput setaf 7 2>/dev/null)
CEND:=$(shell tput sgr0 2>/dev/null)
endif

.PHONY: fmt
fmt:   ## fmt code
	@echo "$(CGREEN)Run gofmt on all source files ...$(CEND)"
	@echo "gofmt -l -s -w ..."
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -s -w $$d/*.go || ret=$$? ; \
	done ; exit $$ret

.PHONY: deps
deps:		## deps check
	@echo "$(CGREEN)deps check ...$(CEND)"
	@/bin/bash scripts/deps.sh

.PHONY: test
test:  deps fparse   	## go test
	@echo "$(CGREEN)go test ...$(CEND)"
	@go test ./...

.PHONY: fparse
fparse:		## gen fofax parse
	@echo "$(CGREEN)deps check ...$(CEND)"
	@cd internal/fxparser/ && make

.PHONY: build
build:  fmt   	## build current target
	@echo "$(CGREEN)fofax build snapshot no publish ...$(CEND)"
	@goreleaser build --snapshot --rm-dist  --single-target -f scripts/gorelease.yml

.PHONY: buildf
buildf:  fmt fparse test 	## build fparse and build current target
	@echo "$(CGREEN)fofax build fparse and build snapshot no publish ...$(CEND)"
	@goreleaser build --snapshot --rm-dist  --single-target -f scripts/gorelease.yml
.PHONY: buildall
buildall:      	## build all
	@echo "$(CGREEN)beye build snapshot no publish ...$(CEND)"
	@goreleaser build --snapshot --rm-dist -f scripts/gorelease.yml
.PHONY: snapshot

snapshot:   fmt 	## pre snapshot
	@echo "$(CGREEN)fofax release snapshot no publish ...$(CEND)"
	@goreleaser release --skip-publish  --snapshot --rm-dist -f scripts/gorelease.yml
.PHONY: release
release:   fmt	## release no publish
	@echo "$(CGREEN)fofax release no publish ...$(CEND)"
	@goreleaser release --skip-publish  --rm-dist -f scripts/gorelease.yml

.PHONY: clean
clean:      	## clean up
	@echo "$(CGREEN)clean up dist ...$(CEND)"
	@rm -rf ./dist


.PHONY: help
help:			## Show this help.
	@echo "$(CGREEN)fofax project$(CEND)"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make $(CYELLOW)<target>$(CEND) (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(CCYAN)%-12s$(CEND) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
