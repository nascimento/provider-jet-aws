# Rede Annotations

```bash

export BUILD_REGISTRY=nascimento
export DOCKER_REGISTRY=nascimento
export REGISTRIES=nascimento
export VERSION=v0.6.0-rede
export BUILDER_HOME=/tmp/upbound
export OS=linux
export ARCH=amd64
export TARGETARCH=amd64
export HOSTOS=linux


crds=( dynamodb glue kinesis kms lambda rds s3 secretsmanager sns sqs ec2 elasticache )

## Backup and Disable CRDs
find package/crds -type f -name '*.yaml' -exec rename 's/\.yaml$/.backup/' {} +
for i in "${crds[@]}"
do
	find package/crds -type f -name "*$i.aws*.backup" -exec rename 's/\.backup$/.yaml/' {} +
done
find package/crds -type f -name 'aws.jet*.backup' -exec rename 's/\.backup$/.yaml/' {} +


## Backup and Disable Controllers
sed -e '/github.com\/crossplane-contrib\/provider-jet-aws\/internal\/controller/ s/^#*/\/\//' -i.backup ./internal/controller/zz_setup.go
for i in "${crds[@]}"
do
	sed -e "/controller\/$i/ s/\/\///" -i.backup ./internal/controller/zz_setup.go
done
sed -e "/controller\/providerconfig/ s/\/\///" -i.backup ./internal/controller/zz_setup.go

## Setup disable
sed -e '/.Setup,/ s/^#*/\/\//' -i.backup ./internal/controller/zz_setup.go
for i in "${crds[@]}"
do
	tempcrds=(`grep -e "/controller\/$i/" ./internal/controller/zz_setup.go | awk '{print $1}'`)
  for j in "${tempcrds[@]}"
  do
    sed -e "/	$j\.Setup/ s/\/\///" -i.backup ./internal/controller/zz_setup.go
  done
done
sed -e "/	providerconfig\.Setup/ s/\/\///" -i.backup ./internal/controller/zz_setup.go


> Verificar se existem imports perdidos no `./internal/controller/zz_setup.go`

PLATFORM=linux_amd64 build
PLATFORMS=linux_amd64 make publish

docker manifest create "nascimento/provider-jet-aws-controller:${VERSION}" --amend "nascimento/provider-jet-aws-controller-amd64:${VERSION}"
docker manifest annotate --arch amd64 --os linux "nascimento/provider-jet-aws-controller:${VERSION}" "nascimento/provider-jet-aws-controller-amd64:${VERSION}"
docker manifest inspect "nascimento/provider-jet-aws-controller:${VERSION}"
docker manifest push "nascimento/provider-jet-aws-controller:${VERSION}"

docker manifest create "nascimento/provider-jet-aws:${VERSION}" --amend "nascimento/provider-jet-aws-amd64:${VERSION}" 
docker manifest annotate --arch amd64 --os linux "nascimento/provider-jet-aws:${VERSION}" "nascimento/provider-jet-aws-amd64:${VERSION}"
docker manifest inspect "nascimento/provider-jet-aws:${VERSION}"
docker manifest push "nascimento/provider-jet-aws:${VERSION}"

```

# Terrajet AWS Provider

`provider-jet-aws` is a [Crossplane](https://crossplane.io/) provider that is
built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for
[Amazon AWS](https://aws.amazon.com/).

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/crossplane-contrib/provider-jet-aws/releases):

```
kubectl crossplane install provider crossplane/provider-jet-aws:v0.2.1
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-aws).

## Contributing

Please see the [Adding New Resources](/docs/adding-resources.md) guide.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane/provider-jet-aws/issues).

## Contact

Please use the following to reach members of the community:

- Slack: Join our [slack channel](https://slack.crossplane.io)
- Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
- Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
- Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-jet-aws is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-aws adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-aws is under the Apache 2.0 license.
