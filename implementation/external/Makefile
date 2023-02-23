IMG := arnobkumarsaha/external-data-provider:dev

all:
	go build -o bin/provider main.go
	docker build -t ${IMG} .
	kind load docker-image ${IMG}
	kubectl apply -f yamls/install.yaml

del:
	kubectl delete -f yamls/install.yaml
	sleep 2

pod:
	kubectl run nginx --image=error_nginx --dry-run=server -ojson