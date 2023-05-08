default: check_bash_version help

# Variables
SHELL := /bin/bash
ESC = \x1b
FEATURES_FOLDER := features

REGEX_FEATURE_TAG = ^.*@(.*)[\s]*$$
REGEX_COLUMN_SEP = :
REGEX_MAKEFILE_DOC = ^([a-zA-Z_-]+):.* \#\# (.*)$$
REGEX_MD_ITALIC = \*{1}([a-zA-Z0-9=\_\ \-]+)\*{1}
REGEX_MD_BOLD = \*{2}([a-zA-Z0-9=\_\ \-]+)\*{2}
REGEX_MD_MONO = `([a-zA-Z0-9=\_\ \-]+)`
REGEX_MD_LINK = \[([a-zA-Z0-9=\_\ \-]+)\]

# fonts
LD = ${ESC}[0m#   default
LB = ${ESC}[1m#   bold
FF = ${ESC}[2m#   faint
FI = ${ESC}[3m#   italic
FU = ${ESC}[4m#   underline
# foreground colors
F0 = ${ESC}[30m#  black
F1 = ${ESC}[31m#  red
F2 = ${ESC}[32m#  green
F3 = ${ESC}[33m#  yellow
F4 = ${ESC}[34m#  blue
F5 = ${ESC}[35m#  magenta
F6 = ${ESC}[36m#  cyan
F7 = ${ESC}[37m#  light gray
F8 = ${ESC}[90m#  gray
F9 = ${ESC}[91m#  light red
F10 = ${ESC}[92m# light green
F11 = ${ESC}[93m# light yellow
F12 = ${ESC}[93m# light blue
F13 = ${ESC}[94m# light blue
F14 = ${ESC}[95m# light magenta
F15 = ${ESC}[96m# light cyan
F16 = ${ESC}[97m# white
# background colors
B0 = ${ESC}[40m#   black
B1 = ${ESC}[41m#   red
B2 = ${ESC}[42m#   green
B3 = ${ESC}[43m#   yellow
B4 = ${ESC}[44m#   blue
B5 = ${ESC}[45m#   magenta
B6 = ${ESC}[46m#   cyan
B7 = ${ESC}[47m#   light gray
B8 = ${ESC}[100m#  gray
B9 = ${ESC}[101m#  light red
B10 = ${ESC}[102m# light green
B11 = ${ESC}[103m# light yellow
B12 = ${ESC}[103m# light blue
B13 = ${ESC}[104m# light blue
B14 = ${ESC}[105m# light magenta
B15 = ${ESC}[106m# light cyan
B16 = ${ESC}[107m# white


# Methods aliases
get_first_digits = sed -En 's/^([0-9]+).*$$/\1/p'


check_bash_version:
	@current_bash_version=`echo $$BASH_VERSION|${get_first_digits}`; \
	if [ $$current_bash_version -lt 4 ]; then \
		printf "$(B3) Bash version must be equal or higher than 4 $(LD) "; \
		exit 1; \
	fi


install: install-go ## Install [everything]
install-go: install-go-pkgs install-go-extra-deps ## [GO] install all deps and tools


install-go-pkgs: ## [GO] install all dependencies packages
	@echo "Installing GO deps"
	@go get -u -t && go mod tidy


install-go-extra-deps: ## [GO] install all dependencies packages (*tools*)
	@echo "Installing Playwright (browser deps)"
	@go install github.com/playwright-community/playwright-go/cmd/playwright@latest
	@playwright install --with-deps $(MAUI_TEST_BROWSER)


list: ## List all available tags under the feature files
    # 1. list features
    # 2. print the current feature file
    # 3. read the file
    # 4. searh for the lines that has a flag
    # 5. get only the flag name
    # 6. sort and remove duplicates
    # 7. print a double line to split between tag files
	@for file in `ls $(FEATURES_FOLDER)`; do \
		printf "$(F2)Flags for $${file%.feature}: $(LD) \n"; \
		cat $(FEATURES_FOLDER)/$$file | \
		grep -E "$(REGEX_FEATURE_TAG)" | \
		sed -En "s/$(REGEX_FEATURE_TAG)/  \1/p" | \
		sort -u; \
		printf "\n"; \
	done;


example: ## Show some example cases
	@printf "$(F1) Example: $(LD) \n"
	@echo "export MAUI_TEST_ENV=http://localhost:5000/"
	@echo "export MAUI_TEST_BROWSER=msedge"
	@echo "export MAUI_TEST_TAGS=playwright"
	@echo "make run"
	@printf "$(FF) OR: $(LD) \n"
	@echo "MAUI_TEST_ENV=http://localhost:5000/ MAUI_TEST_BROWSER=msedge MAUI_TEST_TAGS=playwright make run"


help:
	@printf "$(FF) Available methods: $(LD)\n\n"
        # 1. read makefile
        # 2. get lines that can have a method description
        # 3. color method names and add a COLUMN_SEPARATOR
        # 4. colour backticks (``)
        # 5. colour brackets ([])
        # 6. make it bold
        # 7. make it italic
        # 8. show as table
	@cat $(MAKEFILE_LIST) | \
	 	grep -E "$(REGEX_MAKEFILE_DOC)" | \
		sed -En 's/$(REGEX_MAKEFILE_DOC)/$(F2)\1$(REGEX_COLUMN_SEP)$(LD)\2/p' | \
		sed -En 's/$(REGEX_MD_MONO)/$(F3)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_LINK)/$(F6)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_BOLD)/$(LB)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_ITALIC)/$(FI)\1$(LD)/g;p' | \
		column -t -s "$(REGEX_COLUMN_SEP)"

