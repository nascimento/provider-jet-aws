## CUSTOM BUILD

> ! PERSONAL USAGE FOR TESTS

My notes:

```bash
export BUILD_REGISTRY=nascimento
export DOCKER_REGISTRY=nascimento
export REGISTRIES=nascimento
export VERSION=v0.4.0-rede
export BUILDER_HOME=/tmp/upbound

crds=( glue kafka kinesis kms lambda rds s3 secretsmanager sns sqs )

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



make build -j4 # make build.all -j4 # AMD and ARM
make publish -j4

docker manifest create nascimento/provider-jet-aws-controller:${VERSION} --amend nascimento/provider-jet-aws-controller-amd64:${VERSION} --amend nascimento/provider-jet-aws-controller-arm64:${VERSION}
docker manifest push nascimento/provider-jet-aws-controller:${VERSION}

docker manifest create nascimento/provider-jet-aws:${VERSION} --amend nascimento/provider-jet-aws-amd64:${VERSION} --amend nascimento/provider-jet-aws-arm64:${VERSION}
docker manifest push nascimento/provider-jet-aws:${VERSION}
```

Run:

```bash
export TERRAFORM_VERSION=1.0.5
export TERRAFORM_PROVIDER_SOURCE="hashicorp/aws"
export TERRAFORM_PROVIDER_VERSION=3.56.0


--terraform-version=3.56.0 --terraform-provider-source="hashicorp/aws" --terraform-provider-version=1.0.5
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
