# **AWS + KUBERNETES + K9S**

## **Install [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html#cliv2-linux-install) interface**
```
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

unzip awscliv2.zip

sudo ./aws/install

aws --version
```
The AWS CLI stores this information in a profile (a collection of settings) named default in the credentials file.

Insert your default credentials with: 

`aws configure`

---
## **Install [kubectl](https://kubernetes.io/docs/tasks/tools/)**
```
snap install kubectl --classic
```

Configure files:

`aws eks update-kubeconfig --name goBank --region sa-east-1`

retrieve information: 

`kubectl cluster-info`


  switch between aws profiles with: 
  
  `export AWS_PROFILE=<your-profile-name>`

Example of manually applying changes to the cluster
  * `kubectl apply -f eks/aws-auth.yaml`
  * `kubectl apply -f eks/deployment.yaml`

---

## **Install [k9s](https://k9scli.io/) to make easier kubectl interactions**
   ```
   curl -Lo k9s.tgz https://github.com/derailed/k9s/releases/download/v0.26.3/k9s_Linux_x86_64.tar.gz

   tar -xf k9s.tgz

   sudo install k9s /usr/local/bin/
   ```

  run with `k9s`

---
## **Deploy an [ingress-nginx](https://kubernetes.github.io/ingress-nginx/deploy/#aws) instance on AWS**

On most Kubernetes clusters, the ingress controller will work without requiring any extra configuration. 

`kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.0/deploy/static/provider/aws/deploy.yaml`

---
## **Manage certificates in Kubertenes with [cert-manager](https://cert-manager.io/)**

cert-manager adds certificates and certificate issuers as resource types in Kubernetes clusters, and simplifies the process of obtaining, renewing and using those certificates.

The default static configuration can be installed as follows:

`kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.9.1/cert-manager.yaml`