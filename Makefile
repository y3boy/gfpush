all: build check

.PHONY: linter
linter:
	golangci-lint run

.PHONY: build
build:
	@echo "Building gfpush ..."
	@ go build -o gfpush cli/gfpush/main.go
	@echo done ✔️
.PHONY: check
check:
	@echo
	@echo "Checking gfpush version ..."
	./gfpush -v
	@echo done ✔️
