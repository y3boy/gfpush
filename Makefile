all: build check

.PHONY: linter
linter:
	golangci-lint run

.PHONY: build
build:
	@ echo "Building gfpush ..."
	@ go build -o gfpush cli/gfpush/main.go
	@ echo Done ✔️

.PHONY: check
check:
	@ echo
	@ echo "Lets check gfpush version ..."
	@ ./gfpush -v
	@ echo Looks everything works fine 🤩
	@ echo
	@ echo Move gfpush binary file to one of PATH directory.
	@ echo For example /usr/local/bin/.
	@ echo Have fun 🚀
	