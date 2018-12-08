PHONY : test
test : set-environment
	go test ./...

PHONY : install
install : set-environment
	go list ./...

PHONY : set-environment
set-environment :
	export GO111MODULE=on
