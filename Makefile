api_version?=v1

hello-docs:
	docker run -v $(CURDIR)/hello-proto/$(api_version):/out \
	-v $(CURDIR)/:/protos:ro \
	pseudomuto/protoc-gen-doc --doc_opt=markdown,docs.md \
	$(shell ./list_files.sh -v "$(api_version)" -s "hello-proto")
