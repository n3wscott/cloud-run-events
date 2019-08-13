# Knative with GCP Metrics

## Prometheus Collection

1. Enable Knatives install of Prometheus to scrape Knative with GCP, run the
   following:

   ```shell
   kubectl get configmap -n  knative-monitoring prometheus-scrape-config -oyaml > tmp.prometheus-scrape-config.yaml
   sed -e 's/^/    /' config/monitoring/metrics/prometheus/prometheus-scrape-kn-gcp.yaml > tmp.prometheus-scrape-kn-gcp.yaml
   sed -e '/    scrape_configs:/r tmp.prometheus-scrape-kn-gcp.yaml' tmp.prometheus-scrape-config.yaml \
     | kubectl apply -f -
   ```

2. To verify, run the following to show the diff between what the resource was
   original and is now (_note_: k8s will update `metadata.annotations`):

   ```shell
   kubectl get configmap -n  knative-monitoring prometheus-scrape-config -oyaml \
     | diff - tmp.prometheus-scrape-config.yaml
   ```

   Or, to just see our changes (without `metadata.annotations`) run:

   ```shell
   CHANGED_LINES=$(echo $(cat  tmp.prometheus-scrape-kn-gcp.yaml | wc -l) + 1 | bc)
   kubectl get configmap -n  knative-monitoring prometheus-scrape-config -oyaml \
     | diff - tmp.prometheus-scrape-config.yaml \
     | head -n $CHANGED_LINES
   ```

3. Restart Prometheus

   To pick up this new config, the pods that run Prometheus need to be
   restarted, run:

   ```shell
   kubectl delete pods -n knative-monitoring prometheus-system-0 prometheus-system-1
   ```

   And they will come back:

   ```shell
   $ kubectl get pods -n knative-monitoring
   NAME                                 READY   STATUS    RESTARTS   AGE
   grafana-d7478555c-8qgf7              1/1     Running   0          22h
   kube-state-metrics-765d876c6-z7dfn   4/4     Running   0          22h
   node-exporter-5m9cz                  2/2     Running   0          22h
   node-exporter-z59gz                  2/2     Running   0          22h
   prometheus-system-0                  1/1     Running   0          32s
   prometheus-system-1                  1/1     Running   0          36s
   ```

4. To remove the temp files, run:

   ```shell
   rm tmp.prometheus-scrape-kn-gcp.yaml
   rm tmp.prometheus-scrape-config.yaml
   ```

## Grafana Dashboard

To enable the Knative with GCP dashboard in Grafana, run the following:

```shell
kubectl patch  configmap grafana-dashboard-definition-knative -n knative-monitoring \
  --patch "$(cat grafana/100-grafana-dash-kn-gcp.yaml)"
```

## Accessing Prometheus and Grafana

kubectl port-forward --namespace knative-monitoring \
 \$(kubectl get pods --namespace knative-monitoring --selector=app=grafana
--output=jsonpath="{.items..metadata.name}") \
 3000

```shell
kubectl port-forward -n knative-monitoring \
  $(kubectl get pods -n knative-monitoring --selector=app=prometheus --output=jsonpath="{.items[0].metadata.name}") \
  9090
```