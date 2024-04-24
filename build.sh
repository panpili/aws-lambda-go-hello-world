GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
zip bootstrap.zip bootstrap
aws lambda update-function-code --function-name go-hello-world \
--zip-file fileb://bootstrap.zip