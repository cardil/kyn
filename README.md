# kyn
Kubernetes YAML Namespace changer

## Usage

Rename namespace for directory:

```bash
kyn --namespace acme ./yamls/ \
  | kubectl apply -f -
```

Rename for directory, for specific namespace:

```bash
kyn --namespace default=acme ./yamls/ \
  | kubectl apply -f -
```

Rename namespace in standard input:

```bash
cat kube.yaml | \
  kyn --namespace acme - | \
  kubectl apply -f -
```

## Installation

```bash
go install github.com/cardil/kyn@latest
```

Or use directly (Go 1.22+):

```bash
go run github.com/cardil/kyn@latest \
  --namespace acme ./yamls/ | \
  kubectl apply -f -
```
