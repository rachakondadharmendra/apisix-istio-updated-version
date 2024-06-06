
# Instance Creation with Configuration

- **OS**: Ubuntu 22.04
- **CPU**: 4
- **RAM**: 16GB
- **Storage**: 25GB
- **Open Ports**: 80, 8080, 22 (default for SSH)
- **IAM Instance Profile**: Admin_access_role (To access AWS resources securely)

[Installation Configuration Details](Installation_Config.md)

---

# ISTIOCTL Setup

```bash
cd ~/
curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$HOME/.istioctl/bin:$PATH 

minikube start --cpus=4 --memory=8192

istioctl x precheck
istioctl install --set profile=minimal -y
kubectl -n istio-system get pods
```

## APIsix Installation

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

## Applying ISTIO and APIsix Configuration

```bash
cd ~/apisix-istio-updated-version/
kubectl apply -f 01_strict_mtls.yaml
kubectl apply -f 02_client.yaml
```

## Verification of Pods and Services

```bash
kubectl get po,svc -n apisix
kubectl get po,svc -n istio-system
kubectl get ns
```

---

# Application Deployment

Now, let's deploy the applications. We have 2 services where service1, at path "/", redirects traffic to service2 at path "/srv2/".

- Deploy the 2 services into the production namespace.
- Ensure Istio envoy sidecar containers are added for mutual TLS.

```bash
kubectl apply -f deployment/
```

## Setting Up APIsixRoute

- Use APIsix gateway as the frontline for managing client-side requests.
- Create an ApisixRoute to route traffic from the APIsix gateway to internal backends.

```bash
kubectl apply -f apisix_yamls/02_apisix_route.yaml
```

---

## Testing Configuration

```bash
minikube service apisix-gateway --url -n apisix
```

You can access it now. Make sure to check the port and IP used as per the above command results.

```bash
curl http://192.168.49.2:30825/ -H 'Host: local.rachakonda.me'
```

You should now be able to make requests to the backend services via the APIsix gateway.

---

## Securing Connection

To secure the connection between the APIsix gateway and the client side, we'll use OpenSSL.

```bash
mkdir certs/ && cd certs/
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=local.rachakonda.me/O=local.rachakonda.me"

kubectl create secret tls local-rachakonda-me-cert --cert=tls.crt --key=tls.key -n production
```

Create a new file with the name `tls-local-rachakonda-me.yaml` with the following content:

```yaml
apiVersion: apisix.apache.org/v2
kind: ApisixTls
metadata:
  name: tls-local-rachakonda-me
  namespace: production
spec:
  hosts:
    - local.rachakonda.me
  secret:
    name: local-rachakonda-me-cert
    namespace: production
```

Configuration of ApisixTLS YAML with the secret should be done by now.

```bash
minikube service apisix-gateway --url -n apisix
```

You can access it now. Make sure to check the port and IP used as per the above command results.

```bash
curl https://local.rachakonda.me:30336 --resolve 'local.rachakonda.me:30336:192.168.49.2' -k -v
```

# Troubleshooting and Log Checking

## APIsix Troubleshooting

- **Check APIsix Pods Status:**
```bash
kubectl get pods -n apisix
```

- **Check APIsix Services:**
```bash
kubectl get services -n apisix
```

- **Check APIsix Ingress Controller Logs:**
```bash
kubectl logs -n apisix <apisix-ingress-controller-pod-name>
```

## Istio Troubleshooting

- **Check Istio Pods Status:**
```bash
kubectl get pods -n istio-system
```

- **Check Istio Services:**
```bash
kubectl get services -n istio-system
```

- **Check Istio Ingress Gateway Logs:**
```bash
kubectl logs -n istio-system <istio-ingress-gateway-pod-name>
```

## Istio Sidecar Troubleshooting

- **Check Sidecar Injection Status:**
```bash
kubectl get pods -n <namespace> -o=jsonpath='{range .items[*]}{"\n"}{.metadata.name}{": "}{.spec.containers[*].name}{": "}{.status.phase}{": "}{.status.containerStatuses[*].restartCount}{end}'
```

- **Check Sidecar Logs:**
```bash
kubectl logs -n <namespace> <pod-name> -c istio-proxy
```

- **Check Sidecar Config:**
```bash
kubectl exec -it <pod-name> -n <namespace> -c istio-proxy -- pilot-agent request GET config_dump | grep listener
```
