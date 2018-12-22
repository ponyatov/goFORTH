go: FORTH.go
	gofmt -w $<
	go run $<

local:
	python C:\Google\Cloud\SDK\google-cloud-sdk\bin\dev_appserver.py app.yaml 
deploy:
	gcloud app deploy 
browse:
	gcloud app browse
