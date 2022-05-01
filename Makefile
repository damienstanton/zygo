default: test

test:
	tests/testall.sh && echo "running 'go test'" && cd zygo && go test -v ./...
