PHONY : test
test :
	go test ./...

PHONY : install
install :
	export GO111MODULE=on
	go list ./...