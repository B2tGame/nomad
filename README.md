
Build instruction for Appland's patched version of Nomad.

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

sudo npm install -g yarn

cd $(go env GOPATH)
git clone git@bitbucket.org:appland/nomad.git
cd nomad

git checkout appland-gpu-patches

go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go get -u github.com/ugorji/go/codec/codecgen

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

