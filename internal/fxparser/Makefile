

.PHONY: build
build:   	## run snapshot
	@echo "$(CGREEN) build language golang ...$(CEND)"
	java -jar script/antlr-4.9.2-complete.jar -Dlanguage=Go -no-visitor -listener FOFA.g4 -o parser

.PHONY: help
help:			## Show this help.
	@echo "$(CGREEN)goblin project$(CEND)"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make $(CYELLOW)<target>$(CEND) (default: help)\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  $(CCYAN)%-12s$(CEND) %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
