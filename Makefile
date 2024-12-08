DIR = "./src/$(YEAR)/days/$(DAY)"

FILES = $(DIR)/code.go \
				$(DIR)/code_test.go \
				$(DIR)/code_benchmark_test.go

all:
	@echo "Available commands:"
	@echo "  make YEAR=<year> DAY=<day> gen    - Generate template files"
	@echo "  make YEAR=<year> DAY=<day> run    - Run the code"
	@echo "  make YEAR=<year> DAY=<day> test   - Run tests"
	@echo "  make YEAR=<year> DAY=<day> bench  - Run benchmarks"

gen: $(DIR) $(FILES)

$(DIR):
	mkdir -p $(DIR)

$(DIR)/%.go:
	@if [ ! -f $@ ]; then \
		echo "Creating $@"; \
		cp "./src/harness/templates/$*.go.tmpl" $@; \
	else \
		echo "File $@ already exists, skipping..."; \
	fi

run:
	 @cd $(DIR) && AOC_HARNESS=1 AOC_PART2=true go run ./code.go

watch:
	 @cd $(DIR) && AOC_PART2=true go run ./code.go

test:
	@cd $(DIR) && go test

bench:
	@cd $(DIR) && go test -bench=. -run=^#

.PHONY: gen run watch test bench

