check:
	go mod tidy
	golangci-lint run
	staticcheck -f stylish ./...

format:
	gofumpt -w -l .
	gofmt -w -l -s .
	gci write --skip-generated -s standard,default .

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(HOME)/go/bin latest
	go install github.com/daixiang0/gci@latest
	go install mvdan.cc/gofumpt@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

pre-commit: format check

test:
	go test ./...