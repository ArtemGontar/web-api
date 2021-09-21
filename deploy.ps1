kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/config/crd/bases/crdb.cockroachlabs.com_crdbclusters.yaml

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/manifests/operator.yaml

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/examples/example.yaml


kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/examples/client-secure-operator.yaml

kubectl exec -it cockroachdb-client-secure -- ./cockroach sql --certs-dir=/cockroach/cockroach-certs --host=cockroachdb-public