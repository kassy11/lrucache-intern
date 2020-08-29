# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

all:
	go build -o mylrucache main.go

test:
	go test -v --cover ./lrucache | $(COLORIZE_PASS) | $(COLORIZE_FAIL)
