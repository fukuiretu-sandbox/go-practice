.PHONY: clean build deploy

S3_BUCKET := sam-test-fukui-190628
REGION := ap-northeast-1
TEMPLATE_FILE := template.yml
PACKAGE_FILE := package.yml
STACK_NAME := lambda-go-sample-fukui-190628
PROFILE := sam-test

deps:
	dep ensure

clean:
	rm -rf dest/main

build: clean
	env GOARCH="amd64" GOOS="linux" go build -o dest/main

package:
	sam package \
		--template-file $(TEMPLATE_FILE) \
		--s3-bucket $(S3_BUCKET) \
		--s3-prefix lambda-go-sample \
		--output-template-file $(PACKAGE_FILE) \
		--profile $(PROFILE)

deploy: build package
	sam deploy \
		--template-file $(PACKAGE_FILE) \
		--stack-name  $(STACK_NAME) \
		--region $(REGION) \
		--profile  $(PROFILE) \
		--capabilities CAPABILITY_IAM

tail:
	saw watch /aws/lambda/main \
		--profile $(PROFILE) \
		--region $(REGION)
