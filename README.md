# private-reggie
Private Terraform Provider Registry

### Test
#### With curl
```bash
$ curl http://localhost:8080/terraform/providers/v1/hashicorp/hashicups/versions

http://127.0.0.1:8080/terraform/providers/v1/hashicorp/hashicups/versions

http://127.0.0.1:8080/terraform/providers/v1/hashicorp/hashicups/0.3.2/download/darwin/amd64

```
```javascript
{
  id: "hashicorp/hashicups",
  versions: [{
    version: "0.3.1",
    protocols: [
      "5.0"
    ],
    platforms: [{
      os: "freebsd",
      arch: "386"
    }
  ]}]
}
```
#### With terraform
```terraform
terraform {
  required_providers {
    hashicups = {
      source = "localhost:8081/hashicorp/hashicups"
      version = "0.3.1"
    }
  }
}

provider "hashicups" {
  # Configuration options
}
```

```bash
# we do this because the registry ssl cert is not using a SAN or something
$ export GODEBUG=x509ignoreCN=0
$ terraform init
```