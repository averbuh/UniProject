#!/bin/bash 

#Connect to EKS


# aws eks update-kubeconfig --region eu-central-1 --name dev-demo
# aws eks update-kubeconfig --region eu-central-1 --name qa-demo 
aws eks update-kubeconfig --region eu-central-1 --name staging-demo

aws eks update-kubeconfig --region eu-central-1 --name test

#change service type to LoadBalancer
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
#get argocd server's url (url=only url)
url=$(kubectl get service --namespace argocd argocd-server | sed '2q;d' | awk '{print $4}')

echo $url
#get argocd generated password (current_pass=only password)
current_pass=$(argocd admin initial-password -n argocd | head -n 1)
#login argocd with service url (past url variable) 

echo $current_pass

#change argocd password to "password"
argocd account update-password --new-password password --current-password "$current_pass"

argocd login "$url" 

#add new clusters
kubectl config get-contexts

argocd cluster add arn:aws:eks:eu-central-1:975050257492:cluster/staging-demo --name staging-demo

# argocd cluster add arn:aws:eks:eu-central-1:975050257492:cluster/qa-demo --name qa-demo

# argocd cluster add arn:aws:eks:eu-central-1:975050257492:cluster/dev-demo --name dev-demo