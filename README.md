


![Arch_Diagram](https://github.com/rachakondadharmendra/Ops-Knowledge-Base/blob/main/Arch-Daigrams/apisix-istio-arch.gif)
```markdown




# Minikube Setup
```bash
minikube start --cpus=4 --memory=8192
```

# Istio Installation
```bash
curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$HOME/.istioctl/bin:$PATH 

istioctl x precheck
istioctl install --set profile=minimal  -y
kubectl -n istio-system get pods
```

# Apisix Setup
```bash
kubectl apply -f apisix_yamls/01_ns.yaml

helm repo add apisix https://charts.apiseven.com
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

ADMIN_API_VERSION=v3
helm install apisix apisix/apisix \
   --set gateway.type=NodePort \
   --set ingress-controller.enabled=true \
   --set apisix.ssl.enabled=true \
   --set ingress-controller.config.apisix.serviceNamespace=apisix  \
   --set ingress-controller.config.apisix.serviceName=apisix-admin \
   --set ingress-controller.config.apisix.adminAPIVersion=$ADMIN_API_VERSION \
   --namespace apisix

export NODE_PORT=$(kubectl get --namespace apisix -o jsonpath="{.spec.ports[0].nodePort}" services apisix-gateway)
export NODE_IP=$(kubectl get nodes --namespace apisix -o jsonpath="{.items[0].status.addresses[0].address}")
echo http://$NODE_IP:$NODE_PORT

kubectl get po -n apisix
```

# Metallb Setup (commented out)
```bash
# helm repo add metallb https://metallb.github.io/metallb
# helm uninstall metallb metallb/metallb \
#    --namespace metallb-system
```

# SSL Certificates
```bash
export DOMAIN_NAME=rachakondadharmendra.info
export SUBDOMAIN_NAME=apisix

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=$DOMAIN_NAME Inc./CN=$DOMAIN_NAME' -keyout $DOMAIN_NAME.key -out $DOMAIN_NAME.crt

openssl req -out $SUBDOMAIN_NAME.$DOMAIN_NAME.csr -newkey rsa:2048 -nodes -keyout $SUBDOMAIN_NAME.$DOMAIN_NAME.key -subj "/CN=$SUBDOMAIN_NAME.$DOMAIN_NAME/O=$SUBDOMAIN_NAME is subdomain of $DOMAIN_NAME"

openssl x509 -req -days 365 -CA $DOMAIN_NAME.crt -CAkey $DOMAIN_NAME.key -set_serial 0 -in $SUBDOMAIN_NAME.$DOMAIN_NAME.csr -out $SUBDOMAIN_NAME.$DOMAIN_NAME.crt
```

# Base64 Conversion
```bash
base64 -w 0 $SUBDOMAIN_NAME.$DOMAIN_NAME.crt
base64 -w 0 $SUBDOMAIN_NAME.$DOMAIN_NAME.key

base64 -w 0 $DOMAIN_NAME.crt >> test 
```

# Kubernetes Secrets Creation
```bash
kubectl create -n istio-system secret tls tetratelabs-credential --key=hello.tetratelabs.dev.key --cert=hello.tetratelabs.dev.crt
export CERT_PATH=/home/ubuntu/production/rachakondadharmendra_certs
export NODE_PORT=31234
```

# Testing
```bash
curl -H "Host:$SUBDOMAIN_NAME.$DOMAIN_NAME" --resolve "$SUBDOMAIN_NAME.$DOMAIN_NAME:$NODE_PORT:$NODE_IP" --cacert $CERT_PATH/$DOMAIN_NAME.crt "https://$SUBDOMAIN_NAME.$DOMAIN_NAME:$NODE_PORT"

curl -H "Host:hello.tetratelabs.dev" --resolve "hello.tetratelabs.dev:443:$INGRESS_IP" --cacert tetratelabs.dev.crt "https://hello.tetratelabs.dev:443"

curl -I -H "HOST: apisix.rachakondadharmendra.info" http://$NODE_IP:$NODE_PORT/
curl -I -H "HOST: apisix.rachakondadharmendra.info" http://192.168.49.2:31962

curl -H "Host: apisix.rachakondadharmendra.info" --resolve "apisix.rachakondadharmendra.info:8443:$NODE_IP" --cacert rachakondadharmendra.info.crt "https://apisix.rachakondadharmendra.info:8443/"

curl https://apisix.rachakondadharmendra.info:8443/srv1 --resolve 'apisix.rachakondadharmendra.info:8443:192.168.49.2' -sk

curl -v -H "Host: apisix.rachakondadharmendra.info" --resolve "apisix.rachakondadharmendra.info:8443:192.168.49.2" --cacert ~/production/rachakondadharmendra_certs/rachakondadharmendra.info.crt "https://apisix.rachakondadharmendra.info:8443/"
```

# Miscellaneous
```bash
kubectl create -n production secret tls apisix-rachakondadharmendra-info-cert-new --key=apisix.rachakondadharmendra.info.key --cert=apisix.rachakondadharmendra.info.crt --dry-run=client 

kubectl create -n production secret tls apisix-rachakondadharmendra-info-cert --key=apisix.rachakondadharmendra.info.key --cert=apisix.rachakondadharmendra.info.crt --cacert=rachakondadharmendra.info.crt

kubectl port-forward --address 0.0.0.0 service/apisix-gateway -n apisix 8443:443 

curl -H "Host:hello.tetratelabs.dev" --resolve "hello.tetratelabs.dev:443:$NODE_IP" --cacert tetratelabs.dev.crt "https://hello.tetratelabs.dev:443"

openssl req -new -key apisix.rachakondadharmendra.info.key -out apisix.rachakondadharmendra.info.csr -subj "/C=Your Country/ST=Your State/L=Your City/O=Your Organization/CN=apisix.rachakondadharmendra.info/emailAddress=Your Email"

##############
5U1wmJ5!cV\+
##############

curl http://192.168.49.2:30789/ -H 
kubectl logs apisix-ingress-controller-777cc48747-b9r8m -c istio-proxy -n apisix
kubectl exec -it client -n client-nammespace -- sh

while true; do curl http://service1.production.svc.cluster.local:8080 && echo "" && sleep 1; done

curl http://192.168.49.2:32004  -H 'HOST: apisix.rachakondadharmendra.info' 
```
