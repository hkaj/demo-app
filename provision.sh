#! /bin/bash
kubectl create -f user_ns.yaml
kubectl create -f monitoring/agent-ds.yaml
kubectl create -f monitoring/kube-state-metrics.yaml
kubectl create -f postgres/kubernetes.yaml
kubectl create -f app/kubernetes.yaml
kubectl create -f nginx/kubernetes.yaml
