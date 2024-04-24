# aws lambda go hello world
An intro template for aws lambda with api gateway in golang.

## how to use
1. create an aws lambda function go-hello-world on web
2. change build.sh with right target architecture, amd64 or arm64
3. run build.sh
4. create an aws api gateway api and attach the function to it
5. calls the api and see the log in cloud watch

## aws cli command
### check function
```bash
aws lambda get-function --function-name go-hello-world
```

### update function
```bash
aws lambda update-function-code --function-name go-hello-world \
--zip-file fileb://bootstrap.zip
```

see: https://docs.aws.amazon.com/zh_cn/lambda/latest/dg/golang-package.html

## cold start problem
The cold start latency of lambda function in golang is quite small, around 50ms in 128M environment.