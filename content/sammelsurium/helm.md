---
title: "Helm"
date: 2024-01-16T13:07:13+01:00
tags:
- Software
---

<!--more-->

Helm Installationen anzeigen

```
> helm list --all-namespaces
NAME                        	NAMESPACE  	REVISION	UPDATED                             	STATUS  	CHART                             	APP VERSION
aws-load-balancer-controller	kube-system	1       	2024-01-12 12:22:16.430073 +0100 CET	deployed	aws-load-balancer-controller-1.6.2	v2.6.2
```

Helm Chart installieren

```
helm install game-2048 game-2048/ -n game-2048
helm install aws-load-balancer-controller eks/aws-load-balancer-controller -n kube-system
```

Helm Chart upgraden 

    helm upgrade aws-load-balancer-controller eks/aws-load-balancer-controller -n kube-system

Helm Chart entfernen

```
> helm uninstall aws-load-balancer-controller -n kube-system
release "aws-load-balancer-controller" uninstalled
```

Helm Chart verifizieren

```
helm template --dry-run --debug game-2048 game-2048/ -n game-2048
```

Helm Namespaces

sollten nicht mehr in den YAML erstellt werden. 
Seit HELM3 wird auch nicht mehr via CMD automatisch ein Namespace erstellt. 

Stattdessen sollte man, dieser
[Doku](https://helm.sh/docs/faq/changes_since_helm2/#automatically-creating-namespaces) zufolge `--create-namespace` nutzen

```
helm install game-2048 game-2048/ -n game-2048 --create-namespace
```

und Helm selbst die Erstellung überlassen
