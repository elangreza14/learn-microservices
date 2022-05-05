#!/bin/bash

1. echo "Creating the volume..."

kubectl apply -f ./kubernetes/persistent-volume.yaml
kubectl apply -f ./kubernetes/persistent-volume-claim.yaml

2. "Creating the database credentials..."

kubectl apply -f ./kubernetes/postgres-secret.yaml

3. "Creating the postgres deployment and service..."

kubectl create -f ./kubernetes/postgres-deployment.yaml

4. "Adding the Golang..."

kubectl apply -f ./kubernetes/go-deployment.yaml

5. "Adding the ingress..."

kubectl apply -f ./kubernetes/ingress.yaml






db secret
// elangreza ZWxhbmdyZXph
// secretauthtest c2VjcmV0YXV0aHRlc3Q=