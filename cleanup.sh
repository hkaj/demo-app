#! /bin/bash
kubectl delete -f nginx/kubernetes.yaml
kubectl delete -f app/kubernetes.yaml
kubectl delete -f postgres/kubernetes.yaml
kubectl delete -f monitoring/agent-ds.yaml
kubectl delete -f monitoring/kube-state-metrics.yaml
kubectl delete -f user_ns.yaml
