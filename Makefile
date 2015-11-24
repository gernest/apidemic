test :
	@go test -cover

deps:
	@go get github.com/mitchellh/gox

dist:
	@gox -output="bin/{{.Dir}}v$(VERSION)_{{.OS}}_{{.Arch}}/{{.Dir}}" ./cmd/apidemic