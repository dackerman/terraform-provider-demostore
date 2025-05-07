# Stainless Store Terraform Provider

The [Stainless Store Terraform provider](https://registry.terraform.io/providers/dackerman/demostore/latest/docs) provides convenient access to
[the Stainless Store REST API](https://docs.dackerman-store.com) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    demostore = {
      source  = "dackerman/demostore"
      version = "~> 0.4.0"
    }
  }
}

# Initialize the provider
provider "demostore" {
  # The token to use for authentication
  auth_token = "123e4567-e89b-12d3-a456-426614174000" # or set DEMOSTORE_API_KEY env variable
}

# Configure a resource
resource "demostore_product" "example_product" {
  description = "description"
  image_url = "image_url"
  name = "name"
  price = 0
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/dackerman/demostore/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property   | Environment variable | Required | Default value |
| ---------- | -------------------- | -------- | ------------- |
| auth_token | `DEMOSTORE_API_KEY`  | true     | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/dackerman/terraform-provider-demostore/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
