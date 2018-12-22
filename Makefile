go: FORTH.go
	gofmt -w $<
	go run $<
