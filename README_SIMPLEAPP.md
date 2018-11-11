# Simple App

I hate having to manage infra and am all for automating it away. Luckily we can easily do that with Kubernetes, Docker and a Microservice architecture.

## First Step

Get a cluster up and running in AWS using KOPS (which utilizes terraform)

### Installing kops -

So lets install kops and some other tools first:

- kubernetes-cli --- Includes the kubectl command and other utilities for working with kubernetes from the command line.
- awscli --- Command line interface to aws api.

```bash
brew update && brew install kops kubernetes-cli awscli
```

## Sign Up For AWS Account

Go to the aws homepage and signup for a new account and validate your account.

Sign into your freshly minted AWS account and via the IAM service Create a new user (I usually refer to these users as mule users because they will do the actual work for us).

Your newly created user will provide access to create the kops user (give the user you create full admin privilages)

You should now have 2 users ---

- The root account user
  - This usually corresponds to your password and email address you used to create the aws account, this account will not usually show up in the IAM console

- The mule user

## awscli requires that you have some things accessible via the environment

Add the following to your ~/.bashrc :

```bash
export AWS_REGION="THE_AWS_REGION_YOU_PLAN_ON_USING"
alias access-personal-aws-0='export AWS_ACCESS_KEY_ID="YOUR_KEY_ID_HERE"'
alias access-personal-aws-1='export AWS_SECRET_ACCESS_KEY="YOUR_KEY_SECRET_HERE"'
```

Then reload your .bashrc file in your terminal and execute your newly defined commands:

```bash
access-personal-aws-0
access-personal-aws-1
```

### [From here on out the process is basically the same as found here](https://github.com/kubernetes/kops/blob/master/docs/aws.md)

Create the kops IAM user from the command line using the following

```bash
aws iam create-group --group-name kops

aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonRoute53FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/IAMFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess --group-name kops

aws iam create-user --user-name kops

aws iam add-user-to-group --user-name kops --group-name kops

aws iam create-access-key --user-name kops
```

You should record the SecretAccessKey and AccessKeyID in the returned JSON output, and then use them below:

### Because "aws configure" doesn't export these vars for kops to use, we export them now

### Setup something along these lines here (in your ~/.bashrc)

```bash
alias access-personal-aws-kops-0='export AWS_ACCESS_KEY_ID="YOUR NEWLY CREATED KEY ID HERE"'
alias access-personal-aws-kops-1='export AWS_SECRET_ACCESS_KEY="YOUR NEWLY CREATED KEY SECRET HERE"'
```

### Then reload your .bashrc file in your terminal and execute your newly defined commands

```bash
access-personal-aws-kops-0
access-personal-aws-kops-1
```

### configure the aws client to use your new IAM user

```bash
aws configure
aws iam list-users
```

### Now terraform records the state in which it last left the environment

### So we need to make sure to create a place to store that

```bash
aws s3api create-bucket \
    --bucket SOME_NAME_HERE_0 \
    --region us-east-1
```

### Also we want to enable versioning in our newly created environment

```bash
aws s3api put-bucket-versioning --bucket SOME_NAME_HERE_0 --versioning-configuration Status=Enabled
```

### Makes it easier to abstract some additional things to the environment

### We have the `.k8s.local` on the end of it to indicate that our cluster is a gossip based cluster instead of a DNS based cluster

```bash
export KUBE_CLUSTER_NAME=SOME_NAME_HERE_1.k8s.local
```

### your giving it the name of your s3 bucket you just created

```bash
export KOPS_STATE_STORE=s3://SOME_NAME_HERE_0
```

### The below command will generate a cluster configuration, but not start building it

### Make sure that you have generated SSH key pair before creating the cluster

[When first starting I found this page useful.](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/)

```bash
kops create cluster \
    --zones YOUR_SELECTED_AWS_AVAILABILITY_ZONE \
    ${KUBE_CLUSTER_NAME}
```

### Now we have a cluster configuration, we can look at every aspect that defines our cluster by editing the description

```bash
kops edit cluster ${KUBE_CLUSTER_NAME}
```

### This opens your editor (as defined by $EDITOR) and allows you to edit the configuration. The configuration is loaded from the S3 bucket we created earlier, and automatically updated when we save and exit the editor

### All instances created by kops will be built within ASG (Auto Scaling Groups), which means each instance will be automatically monitored and rebuilt by AWS if it suffers any failure

### Now we take the final step of actually building the cluster

### This'll take a while

### Once it finishes you'll have to wait longer while the booted instances finish downloading Kubernetes components and reach a "ready" state

```bash
kops update cluster ${KUBE_CLUSTER_NAME} --yes
```

### The configuration for your cluster was automatically generated and written to ~/.kube/config for you

### kops ships with a handy validation tool that can be ran to ensure your cluster is working as expected

```bash
kops validate cluster
```

### Running a Kubernetes cluster within AWS obviously costs money, and so you may want to delete your cluster if you are finished

### You can preview all of the AWS resources that will be destroyed when the cluster is deleted by issuing the following command

```bash
kops delete cluster --name ${KUBE_CLUSTER_NAME}
```

### When you are sure you want to delete your cluster, issue the delete command with the --yes flag

### Note that this command is very destructive, and will delete your cluster and everything contained within it

```bash
kops delete cluster --name ${KUBE_CLUSTER_NAME} --yes
```

### Now at this point you probably want a UI to checkout what you have deployed. So lets go ahead and deploy the simple kubernetes dashboard

```bash
kubectl create -f https://raw.githubusercontent.com/kubernetes/kops/master/addons/kubernetes-dashboard/v1.8.3.yaml
```

### Then set up correct permissions by saving the following to dash-rbac.yaml

### *********START_dash-rbac.yaml*********

```yaml
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
  labels:
    k8s-app: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
```

### *********END_dash-rbac.yaml*********

### Then create it in your kubernetes cluster

```bash
kubectl create -f PATH-TO/dash-rbac.yaml
```

### Then connect to your newly created Kubernetes Dashboard via

```bash
kubectl proxy
```

### and to view your dashboard open your web browser and navigate to

```bash
http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/
```

### If it asks for login credentials then The login credentials are

```bash
Username: admin
Password: # get by running: kops get secrets kube --type secret -oplaintext or kubectl config view --minify
```

### However the dashboard is not a very good monitoring solution so instead we will use prometheus and grafana setup [found here](https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus)

```bash
mkdir -p ~/tmp/kube/
cd ~/tmp/kube/

git clone git@github.com:coreos/prometheus-operator.git
```

### To create the stack

```bash
kubectl create -f ./prometheus-operator/contrib/kube-prometheus/manifests/ || true
```

### It can take a few seconds for the above 'create manifests' command to fully create the following resources, so verify the resources are ready before proceeding

```bash
until kubectl get customresourcedefinitions servicemonitors.monitoring.coreos.com ; do date; sleep 1; echo ""; done
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done

# This command sometimes may need to be done twice (to workaround a race condition).
kubectl create -f ./prometheus-operator/contrib/kube-prometheus/manifests/ 2>/dev/null || true

# When your ready (if you want to...) just delete the stack:
kubectl delete -f ./prometheus-operator/contrib/kube-prometheus/manifests/ || true
```

### Access the dashboards --- Prometheus, Grafana, and Alertmanager dashboards can be accessed quickly using kubectl port-forward after running the quickstart via the commands below

### Kubernetes 1.10 or later is required

### Note: There are instructions on how to route to these pods behind an ingress controller in the Exposing Prometheus/Alermanager/Grafana via Ingress section

## Prometheus

---

```bash
kubectl --namespace monitoring port-forward svc/prometheus-k8s 9090
# Then access via http://localhost:9090

# Grafana

kubectl --namespace monitoring port-forward svc/grafana 3000
# Then access via http://localhost:3000 and use the default grafana user:password of admin:admin.

# Alert Manager

kubectl --namespace monitoring port-forward svc/alertmanager-main 9093
# Then access via http://localhost:9093
```

### Ok so now that we have some visibility into our kubernetes cluster we can install our package manager

### Please note I will be using a dev installation method for helm... this installation will not be production grade

### To find instruction for production grade installations of Helm please [see here](https://docs.helm.sh/using_helm/#securing-your-helm-installation)

```bash
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash
```

### Also it is important to realize that helm follows a server-client interaction model

### `helm` is the client component and `tiller` is the server component

### additionally we have an RBAC enabled cluster so we need to do some additional steps

```bash
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller

helm init \
--override 'spec.template.spec.containers[0].command'='{/tiller,--storage=secret}' \
--service-account=tiller

## To finish of the set up:
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

### Ok now that we have helm deployed and tiller installed into our kubernetes cluster we just need to deploy an ingress controller and the cert-manager addon

### Ingress controller (req'ed for all env's)

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
```

### Additional Provider Specific instructions (level 4) specifically for AWS

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/aws/service-l4.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/aws/patch-configmap-l4.yaml
```

### Go into aws console and click into the load balancer that has been created for you nginx controller and scroll down in the description until you arrive at

### the access logging portion where you will click on configure access logs and then in the popup that appears you will

- click to enable the logs
- set log delivery to the deisred interval
- define an s3 bucket name and select to have the location create for you

### then you are done

### Ok so now that he application is built...we just have to worry about running it...and we do that via docker/helm/kubernetes

### if you want to build the dockerfile locally you will need to be in the root directory and run

### Targeting Local

```bash
docker build -t noah-heil/ibm-simple-app:latest .
```

### Targeting gitlab at some point before all this you will need to do a docker login

```bash
docker build -t registry.gitlab.com/noah-heil/ibm-simple-app:latest .
docker push registry.gitlab.com/noah-heil/ibm-simple-app:latest
```

### Targeting dockerhub at some point before all this you will need to do a docker login

```bash
docker build -t noahheil/ibm-simple-app:latest .
docker push noahheil/ibm-simple-app
```

### To run only the docker image by it's self try

```bash
docker run -p 127.0.0.1:80:8080/tcp noahheil/ibm-simple-app:latest
```

### Now we need to create a quick docker-compose.yml file to then use for generating the helm chart for us

```bash
kompose convert -c
```

### Make sure everything is the way it should be in your helm chart and then we can run helm lint to double check that everything is functioning as expected

```bash
helm lint
# or you can run
# helm --debug lint

# If everything is good, you can package the chart as a release by running (from your repository root) :

helm package ibm-simple-app --debug
# I like to add the --debug flag to see the output of the packaged chart.

# From that we can see that the chart is placed in our current directory as well as in our local helm repository.
# To deploy this release, we can point helm directly to the chart file as follows :
helm install ibm-simple-app-0.0.1.tgz
```

### There is only one more thing left to do to get everything up and running the way we expect it to be

### Now that we have our Ingress controller and ingress set up we need to make sure to update our react components can make requests to the api and that it is routed correctly via the ingress

### First we need to get the correct ingress address

```bash
kubectl describe ingress ibm-simple-app | grep [A]ddress
```

### Then copy the value of the field which is on your screen

### No we are going to take that field and swap some values in our app.jsx file

### So open up app.jsx and replace the 4 occurances of `localhost` with the value we retrieved above

### after that all we have left to do is rebuild the docker container and push it up to a docker registry where our helm chart can pick up the change we made

### (to expedite the process we might consider scaling our deployment down to 0 and then back up to 1)

### That should be everything...If you have any questions...Please feel free to reach out to me
