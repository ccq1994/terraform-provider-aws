---
subcategory: "CodeArtifact"
layout: "aws"
page_title: "AWS: aws_codeartifact_repository_endpoint"
description: |-
    Provides details about a CodeArtifact Repository Endpoint
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_codeartifact_repository_endpoint

The CodeArtifact Repository Endpoint data source returns the endpoint of a repository for a specific package format.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_codeartifact_repository_endpoint import DataAwsCodeartifactRepositoryEndpoint
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsCodeartifactRepositoryEndpoint(self, "test",
            domain=Token.as_string(aws_codeartifact_domain_test.domain),
            format="npm",
            repository=Token.as_string(aws_codeartifact_repository_test.repository)
        )
```

## Argument Reference

This data source supports the following arguments:

* `domain` - (Required) Name of the domain that contains the repository.
* `repository` - (Required) Name of the repository.
* `format` - (Required) Which endpoint of a repository to return. A repository has one endpoint for each package format: `npm`, `pypi`, `maven`, and `nuget`.
* `domain_owner` - (Optional) Account number of the AWS account that owns the domain.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `repository_endpoint` - URL of the returned endpoint.

<!-- cache-key: cdktf-0.20.0 input-845ec7d27ada856d0ad7aca19277d5091cd5696d048ed754b76ab9909399306c -->