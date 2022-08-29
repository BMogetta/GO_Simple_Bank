## Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
  * `snap install kubectl --classic`

## Configure files:
  * `aws eks update-kubeconfig --name goBank --region sa-east-1`
  * `kubectl cluster-info`
  * switch between aws profiles with: `export AWS_PROFILE=<your-profile-name>`
  * `kubectl apply -f eks/aws-auth.yaml`
  * `kubectl apply -f eks/deployment.yaml`

## Install [k9s](https://k9scli.io/) to make easier kubectl interactions
  * `curl -Lo k9s.tgz https://github.com/derailed/k9s/releases/download/v0.26.3/k9s_Linux_x86_64.tar.gz`
  * `tar -xf k9s.tgz`
  * `sudo install k9s /usr/local/bin/`
  * run with `k9s`
