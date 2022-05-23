## Docker

### How to build, push and pull Docker Images
 
Build
```sh
docker build -t image:version .
```

Push
```sh
docker push image:version
```

Pull
```sh
docker pull image:version
```

## Kubernetes

### How to Setup a Local Kubernetes Cluster

Abaixo seguem os passos para instalar o kind que é uma ferramenta para executar clusters locais do Kubernetes usando “nós” de contêiner do Docker. O kind foi projetado principalmente para testar o próprio Kubernetes, mas pode ser usado para desenvolvimento local ou CI.

```sh
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-linux-amd64

chmod +x ./kind

mv ./kind /some-dir-in-your-PATH/kind
```

Após realizar a instalação do kind crie o seu cluster utilizando o comando abaixo

```sh
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF
```

Para que o recurso Ingress funcione, o cluster deve ter um ingress controller em execução. Como exemplo utilizaremos o NGIX Ingress Controller, utilize o comando abaixo para instalar

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

Agora o Ingress está todo configurado. Aguarde até que esteja pronto para processar as solicitações em execução:

```sh
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```






### How to install a kubernetes components

Namespace
```sh
kubectl apply -f ./k8s/namespace.yml
```

ConfigMap
```sh
kubectl apply -f ./k8s/config-map.yml
```

Dockerhub Secret
```sh
kubectl apply -f ./k8s/dockerhub-secret.yml
```

Secret
```sh
kubectl apply -f ./k8s/secret.yml
```

Deployment
```sh
kubectl apply -f ./k8s/deployment.yml
```

Service
```sh
kubectl apply -f ./k8s/service.yml
```

Ingress Controller
```sh
kubectl apply -f ./k8s/ingress.yml
```
