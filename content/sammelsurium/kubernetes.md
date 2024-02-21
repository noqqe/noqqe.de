---
title: Kubernetes
date: 2016-03-22T16:20:21
tags:
- kubernetes
- Software
---

### Deployments

Beschreibt welche Container gespawned werden, welche Ports davon benutzt
werden usw.

Alle anzeigen

    kubectl get deployments

Neues Deployment anlegen

    kubectl create deployment hello-node --image=gcr.io/hello-minikube-zero-install/hello-node

Löschen eines Deployments

    kubectl delete deployment hello-node

### Services

Ein Service mapped ContainerIP:ContainerPort nach "extern". Macht also einen
Dienst nach außen verfügbar.

Alle aktuell laufenden Services anzeigen

    kubectl get services

### Pods

Pods sind ein oder mehrere Container die auf einem Host platziert werden.
Also die ausführbaren Runtimes der Docker Container.

Alle Pods anzeigen

    kubectl get pods;

Auch system pods anzeigen

    kubectl get pods --all-namespaces

Mit noch mehr details

    kubectl get pods -o wide  --all-namespaces

Details für einen Pod anzeigen

    kubectl describe pod hello-node-78cd77d68f-d7lhz

Pod zerstören

    kubectl delete pod hello-node-78cd77d68f-7q6tn

Log von Container eines Pods abholen

    kubectl logs <pod> <container> -n=<name-space>
    kubectl logs weave-net-fcprg weave  -n=kube-system

Ports und IPs von Pod anzeigen

    kubectl get pod vault-0 -o jsonpath='Host:{.status.hostIP} PodIP:{.status.podIP} PodPorts:{.spec.containers[*].ports[*].containerPort}'

### Ingress 

Services via Ingress Controller nach aussen sichtbar gemacht. 
Via nginx-ingress controller oder aws-alb-ingress bei EKS

Ingress Endpoints / URLs anzeigen

    kubectl get ingress --all-namespaces

Speziellen Ingress anzeigen

     kubectl get ingress/<name> -n <namespace>

Ingress löschen
    
    kubectl delete ingress <name> -n <namespace>

Stuck Ingress löschen (Hintergrund: aws-alb-ingress broken)

    kubectl patch ingress <name> -n <namespace> -p '{"metadata":{"finalizers":[]}}' --type=merge

### Nodes

Alle Nodes mit allen Details anzeigen

    kubectl get nodes -o wide

### Secrets

Secret anlegen (string)

    kubectl create secret generic db-password --from-literal=password='secret123'

Secret anlegen (file)

    kubectl create secret generic license --from-file=nexus-repo-license.lic=nexus.lic --dry-run=client --output=yaml

### Kubernetes Cluster

Cluster Log

    kubectl get events

Cluster config

    kubectl config view

Show kubeadmin init config

    kubectl -n kube-system get cm kubeadm-config -oyaml

Show all avialable resources in the cluster

    kubectl api-resource

### EKS 

kubeconfig aus aws-cli generieren lassen

    aws eks update-kubeconfig --name dev 
    kubectl get nodes
