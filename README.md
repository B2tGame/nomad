
First, you need to install Go from https://golang.org/dl/ (do not use apt!)

Then:

```bash
#
# This is for Ubuntu 18.04 Bionic Beaver, which is what we have in AWS
# You might need to adjust the packages name for your distro
#
sudo apt install gcc gcc-8-aarch64-linux-gnu

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
```

---

Example how to do a simple build without update the UI:
```bash
export GOPATH=/workspace/go
export PATH=$PATH:/opt/go/bin/:$GOPATH/bin:/opt/node-v14.15.4/bin
make dev-ui
make prerelease
make release
```
