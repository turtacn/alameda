#kubectl apply -f example/deployment/kubernetes/0alameda-operator/
#kubectl apply -f example/deployment/kubernetes/0alameda-datahub/
#kubectl apply -f example/deployment/kubernetes/0alameda-ai/
#kubectl apply -f example/deployment/kubernetes/0alameda-evictioner/
#kubectl apply -f example/deployment/kubernetes/0admission-controller/
kubectl apply -f example/deployment/kubernetes/0alameda-influxdb/ -n alameda
kubectl apply  -f example/deployment/kubernetes/0alameda-grafana/ -n alameda
kubectl create -f example/deployment/kubernetes/0alameda-grafana/dashboards-json-configmap.yaml  -n alameda
