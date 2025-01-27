# Challengefile

**Challengefile** is a lightweight tool to deploy CTF challenges on Kubernetes. 

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/chxmxii/challengefile.git
   cd challengefile
   ```
2. Install the tool:
   ```bash
   make install
   ```

## Usage

1. Create a `challengefile` for your challenge like the example below:
   ```yaml
   library:

     Metadata:
       Author: chxmxii
       Namespace: lib-pwn
       category: pwn
     
     Deployment:
       name: library
       image: chxmxii/library
       replicas: 2 # ignored when hpa is true
       hpa: true
       healthCheck: true

     Service:
       name: library-svc
       port: 30330
       protocol: TCP
   ```
2. Run commands to manage challenges:

   - **Deploy** a challenge:
     ```bash
     challengefile deploy -f <challengefile> -c <challenge_name> -k <path/to/kubeconfig> 
     ```
   - **Destroy** a challenge:
     ```bash
     challengefile destroy -f <challengefile> -c <challenge_name> -k <path/to/kubeconfig>
     ```
   - **Validate** the `challengefile`:
     ```bash
     challengefile validate -f <file>
     ```
   - View the full command list:
     ```bash
     challengefile help
     ```
- **Note**: The `challengefile` and `kubeconfig` are optional arguments. If not provided, the tool will look for the `./challengefile` in the current directory and use the default kubeconfig.
## Features

- **Deploy challenges** to Kubernetes with ease.
- **Validate configurations** before deployment.
- **Support for HPA** and health checks.
- **Service management** with flexible options.

## Roadmap

- [ ] Support for more challenge types.
- [ ] Integration with CI/CD pipelines.
- [ ] Support for challenge updates.


## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.