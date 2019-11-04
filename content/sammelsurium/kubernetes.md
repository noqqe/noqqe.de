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

    kubectl get deployments;

Löschen eines Deployments

    kubectl delete deployment hello-node

Neues Deployment anlegen

    kubectl create deployment hello-node --image=gcr.io/hello-minikube-zero-install/hello-node

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

### Nodes

Alle Nodes mit allen Details anzeigen

    kubectl get nodes -o wide

### Kubernetes Cluster

Cluster Log

    kubectl get events;

Cluster config

    kubectl config view;

Show kubeadmin init config

    kubectl -n kube-system get cm kubeadm-config -oyaml

Show all avialable resources in the cluster

    kubectl api-resource

Cluster Neuinstallation. Auf jedem Node:

    kubeadm reset
    rm -rf /var/lib/etcd

und dann die Prozedur aus dem Setup nochmal machen. Siehe dazu
[hobby-kube](https://github.com/hobby-kube/guide#choosing-a-cloud-provider)
