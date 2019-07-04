.PHONY: clean build deploy

S3_BUCKET := sam-test-fukui-190628
REGION := ap-northeast-1
TEMPLATE_FILE := template.yml
PACKAGE_FILE := package.yml
STACK_NAME := stepfunctions-sample-fukui-190703
PROFILE := sam-test

deps:
	dep ensure

build:
	make -C src/sample_setting build
	make -C src/sample_addition build
	make -C src/sample_multiply build

package:
	sam package \
		--template-file $(TEMPLATE_FILE) \
		--s3-bucket $(S3_BUCKET) \
		--s3-prefix stepfunctions-sample \
		--output-template-file $(PACKAGE_FILE) \
		--profile $(PROFILE)

deploy: build package
	sam deploy \
		--template-file $(PACKAGE_FILE) \
		--stack-name $(STACK_NAME) \
		--region $(REGION) \
		--profile  $(PROFILE) \
		--capabilities CAPABILITY_IAM
