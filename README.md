Build instruction for Appland's patched version of Nomad.
Nomad
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)
[![Discuss](https://img.shields.io/badge/discuss-nomad-00BC7F?style=flat)](https://discuss.hashicorp.com/c/nomad)
===

<p align="center" style="text-align:center;">
  <a href="https://nomadproject.io">
    <img alt="HashiCorp Nomad logo" src="website/public/img/logo-hashicorp.svg" width="500" />
  </a>
</p>

Nomad is a simple and flexible workload orchestrator to deploy and manage containers ([docker](https://www.nomadproject.io/docs/drivers/docker.html), [podman](https://www.nomadproject.io/docs/drivers/podman)), non-containerized applications ([executable](https://www.nomadproject.io/docs/drivers/exec.html), [Java](https://www.nomadproject.io/docs/drivers/java)), and virtual machines ([qemu](https://www.nomadproject.io/docs/drivers/qemu.html)) across on-prem and clouds at scale.

While the instruction are expressed as a bash script, it's probably better that
you read it and execute the commands manually.

```bash
cd


#
# First, we need to install Go from https://golang.org/dl/
# (The apt version is super old and won't work)
#
wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz


#
# This is for Ubuntu 18.04 Bionic Beaver, which is what we have in AWS
# You might need to adjust the packages name for your distro
#
sudo apt install gcc gcc-8-aarch64-linux-gnu


#
# Ensure that your PATH can reach both the compiler binaries, by
# default in `/usr/local/bin/go`, and the Go packages binaries
# which instead default to `$(go env GOPATH)/bin`
#
# You might need to add this command to your ~/.profile
#
export PATH=$PATH:/usr/local/go/bin/:$(/usr/local/go/bin/go env GOPATH)/bin

* **Simple & Reliable**:  Nomad runs as a single binary and is entirely self contained - combining resource management and scheduling into a single system.  Nomad does not require any external services for storage or coordination.  Nomad automatically handles application, node, and driver failures.  Nomad is distributed and resilient, using leader election and state replication to provide high availability in the event of failures.

* **Device Plugins & GPU Support**: Nomad offers built-in support for GPU workloads such as machine learning (ML) and artificial intelligence (AI).  Nomad uses device plugins to automatically detect and utilize resources from hardware devices such as GPU, FPGAs, and TPUs.

* **Federation for Multi-Region, Multi-Cloud**: Nomad was designed to support infrastructure at a global scale.  Nomad supports federation out-of-the-box and can deploy applications across multiple regions and clouds.

* **Proven Scalability**: Nomad is optimistically concurrent, which increases throughput and reduces latency for workloads.  Nomad has been proven to scale to clusters of 10K+ nodes in real-world production environments.

make bootstrap
make dev-ui
make prerelease
make release


#
# These haven't been tested the slightest
#
# mv /usr/local/bin/nomad /usr/local/bin/nomad.old
# sudo cp pkg/linux_amd64/nomad /usr/local/bin/nomad
#
```
