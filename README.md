# challengefile
A simple tool to help you deploy your challenge on kubernetes.

# Installation

1. Clone the repository.
2. Run `make install` in the root of the repository.

## How to use?
1. Create a `challengefile` in the root of your challenge repository.
2. Add the following content to the file:

```yaml
Metadata:
  Author: chxmxii
  name: library
  Namespace: lib-pwn 
  category: pwn

Deployment:
  name: library
  image: chxmxii/library
  replicas: 2 #ignore when hpa is true.
  hpa: true
  healthCheck: true 

Service:
  name: library-svc
  port: 30330
  endpoint: library.ctf.xyz
  protocol: TCP
```

3. Run `challengefile -f <file> -c <challenge_name> -k <path/to/kubeconfing> (optional)` in the root of your challenge repository.