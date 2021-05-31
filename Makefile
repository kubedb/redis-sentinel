.PHONY: build
build: create

clean:
	kubectl delete sts predis-sts sentinel-sts -n demo
	kubectl delete pvc predis-vol-predis-sts-0 predis-vol-predis-sts-1 predis-vol-predis-sts-2 senti-vol-sentinel-sts-0 senti-vol-sentinel-sts-1 senti-vol-sentinel-sts-2 -n demo
	kubectl delete secret example-com-tls -n demo
	kubectl delete certificate example-com -n demo
create:

	go build .
	./predis create secret
	./predis create cert
	./predis create senti-service
	./predis create redis-service
	./predis create sentistatefulset
	sleep 20
	./predis create sa
	./predis create role
	./predis create binding
	./predis create  statefulset


