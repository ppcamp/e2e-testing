# SEE https://www.google.com/search?q=.phony+makefile&oq=.phony+makefile&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIGCAEQRRg70gEIMjc3MmowajSoAgCwAgA&sourceid=chrome&ie=UTF-8


default: check_bash_version help

# Variables
SHELL := /bin/bash

# Methods aliases
get_first_digits = sed -En 's/^([0-9]+).*$$/\1/p'

install: install-go ## Install everything
install-go: install-go-pkgs install-go-extra-deps ## [GO] install all deps and tools


install-go-pkgs: ## [GO] install all dependencies packages
	@echo "Installing GO deps"
	@go get -u -t && go mod tidy


install-go-extra-deps: ## [GO] install all dependencies packages (*tools*)
	@echo "Installing Playwright (browser deps)"
	@go install github.com/playwright-community/playwright-go/cmd/playwright@latest
	@playwright install --with-deps $(MAUI_TEST_BROWSER)


list: ## List all available tags under the feature files
	@printf "\e[33mCommon description\n p1:Blocker\n p2:Functional Issue\n p3:Non-blocking functional issue\n p4:Content Only\033[0m\n\n"
    # 1. list features
    # 2. print the current feature file
    # 3. read the file
    # 4. searh for the lines that has a flag
    # 5. get only the flag name
    # 6. sort and remove duplicates
    # 7. print a double line to split between tag files
	@for file in `ls $(FEATURES_FOLDER)`; do \
		printf "\e[32m Flags for $${file%.feature}:\033[0m\n\n"; \
		cat $(FEATURES_FOLDER)/$$file | \
		grep -E "^\s*@.*$$" | \
		sed -rn "s/^.*@(.*)[\s]*$$/\1/p" | \
		sort -u; \
		printf "\n\n"; \
	done;


example: ## Show some example cases
	@printf "\n\e[33mExample:\033[0m\n"
	@echo "export MAUI_TEST_ENV=http://localhost:5000/"
	@echo "export MAUI_TEST_BROWSER=msedge"
	@echo "export MAUI_TEST_TAGS=playwright"
	@echo "make run"
	@printf "\e[2m OR:\033[0m\n"
	@echo "MAUI_TEST_ENV=http://localhost:5000/ MAUI_TEST_BROWSER=msedge MAUI_TEST_TAGS=playwright make run"


help:
	@printf "\x1b[2m Available methods:\x1b[0m\n\n"
        # - See https://dev.to/ifenna__/adding-colors-to-bash-scripts-48g4
        # - read makefile
        # - get lines that can have a method description
        # - color method names and description
        # - colour backticks (``)
        # - colour brackets ([])
        # - make it bold
        # - make it italic
	@cat $(MAKEFILE_LIST) | \
	 	grep -E '^[a-zA-Z_-]+:.* ## .*$$' | \
		sed -En 's/(.*):.* ## (.*)/\x1b[32m\1:\x1b[0m\2/p' | \
		sed -En 's/`([a-zA-Z0-9=\_\ \-]+)`/\x1b[33m\1\x1b[0m/g;p' | \
		sed -En 's/\[([a-zA-Z0-9=\_\ \-]+)\]/\x1b[36m\1\x1b[0m/g;p' | \
		sed -En 's/\*{2}([a-zA-Z0-9=\_\ \-]+)\*{2}/\x1b[1m\1\x1b[0m/g;p' | \
		sed -En 's/\*{1}([a-zA-Z0-9=\_\ \-]+)\*{1}/\x1b[3m\1\x1b[0m/g;p' | \
		column -t -s ":"


check_bash_version:
	@current_bash_version=`echo $$BASH_VERSION|${get_first_digits}`; \
	if [ $$current_bash_version -lt 4 ]; then \
		echo "Bash version must be equal or higher than 4"; \
		exit 1; \
	fi
