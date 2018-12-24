include ./hack/help.mk
include ./hack/minikube.mk

.PHONY: environment
environment: ##@setup render service yaml files and apply to current kubernetes namespace
	$(MAKE) -C environment all
	kubectl apply -R -f environment/manifests

lint:
	golangci-lint run -e ./vendor ./...

pkg/profiles.pb.go: pkg/profiles.proto
	protoc --go_out=. pkg/profiles.proto

cli:
	docker build -t test .
	docker run --network=host -ti --rm -w /workspace -v $(PWD):/workspace test sh
