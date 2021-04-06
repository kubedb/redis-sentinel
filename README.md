
###### First step
* `go build .` it'll create a binary called predis
* `kc apply -f file/configmap.yaml`
* `./predis create service`
* `./predis create svc1`
* `./predis create statefulset`
* `go run test/demo.go` it'll add another label field 'role' with first pod 'predis-sts-0' 
   & will get the svc1 created before



###### Second step to run master :
* `kc exec -it predis-sts-0 bash`
* `cd /data/db/tst`
* `redis-server /data/predis-data/master.conf --port 6379`

###### Third step to run replica :
* `kc exec -it predis-sts-1 bash`
* `cd /data/db/tst`
* `redis-server /data/predis-data/replica.conf --port 6379`

###### Forth step to run master :
* `kc exec -it predis-sts-2 bash`
* `cd /data/db/tst`
* `redis-server /data/predis-data/replica.conf --port 6379`


###### Output:
* Total 5 replica with 1 master pod 
