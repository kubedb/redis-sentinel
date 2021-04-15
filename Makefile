.PHONY: build
build: apply create

clean:
	kubectl delete sts predis-sts sentinel-sts
	kubectl delete pvc predis-vol-predis-sts-0 predis-vol-predis-sts-1 predis-vol-predis-sts-2 senti-vol-sentinel-sts-0 senti-vol-sentinel-sts-1 senti-vol-sentinel-sts-2
create:
	go build .
	./predis create sentistatefulset
	sleep 10
	./predis create  statefulset
apply:
	kubectl apply -f file/configmap-scripts.yaml
	kubectl apply -f file/sentiConfigmap.yaml
	kubectl apply -f file/configmap.yaml
	kubectl apply -f file/sentinel-scrpts.yaml