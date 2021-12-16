## Clone provider from github
```bash
git clone https://github.com/sigurdkb/terraform-provider-canvaslms.git
```

## Build provider
```bash
cd terraform-provider-canvaslms 
go install
```
## Install provider in local cache
The example shows folder structure for macos, for linux replace __darwin__ with __linux__
```bash
mkdir -p ~/.terraform.d/plugins/github.com/sigurdkb/canvaslms/0.1.0/darwin_amd64/
cp ~/go/bin/terraform-provider-canvaslms ~/.terraform.d/plugins/github.com/sigurdkb/canvaslms/0.1.0/darwin_amd64/terraform-provider-canvaslms_v0.1.0
```

