# k8s-external-secrets-manager
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://img.shields.io/badge/License-GPL%20v2-blue.svg) 
[![Build Status](https://cloud.drone.io/api/badges/jgavinray/k8s-external-secret-manager/status.svg)](https://cloud.drone.io/jgavinray/k8s-external-secret-manager) 
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-1.2-0baaaa.svg)](code_of_conduct.md)

## What is this?
This is my attempt at solving the problem of a kubernetes operator/administrator. 
As an operator I need to provide a mechanism where users can self-serve their deployments. 
One of the challenges that is faced is where do secrets get stored? 
There are great projects already out there that solve this problem, such as [kubernetes-external-secrets](https://github.com/external-secrets/kubernetes-external-secrets). 
Where I see this falling short is not the secrets management, but rather the ability for a user of kubernetes to manage their own secrets. 
A user/developer will know what variables and secrets are required for their application, they need to ability to create said secrets in a self service manner. 

By providing a custom resource for a user, they can generate/manage secrets without having to know what backend the operations team has chosen for their secrets manager.    

The goal is simple, create a customer resource definition that operates like the aforementioned kubernetes-external-secrets, but also add in the ability to generate new secrets on the fly.

## Based scaffolding provided by operator-sdk v1.13.1