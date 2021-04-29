.PHONY: build
build: create

clean:
	kubectl delete sts predis-sts sentinel-sts
	kubectl delete pvc predis-vol-predis-sts-0 predis-vol-predis-sts-1 predis-vol-predis-sts-2 senti-vol-sentinel-sts-0 senti-vol-sentinel-sts-1 senti-vol-sentinel-sts-2
create:

	go build .
	./predis create senti-service
	./predis create redis-service
	./predis create sentistatefulset
	sleep 10
	./predis create sa
	./predis create role
	./predis create binding
	./predis create  statefulset


