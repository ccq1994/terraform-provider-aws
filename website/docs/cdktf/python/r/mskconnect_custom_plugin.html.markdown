---
subcategory: "Managed Streaming for Kafka Connect"
layout: "aws"
page_title: "AWS: aws_mskconnect_custom_plugin"
description: |-
  Provides an Amazon MSK Connect custom plugin resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_mskconnect_custom_plugin

Provides an Amazon MSK Connect Custom Plugin Resource.

## Example Usage

### Basic configuration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.mskconnect_custom_plugin import MskconnectCustomPlugin
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_object import S3Object
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = S3Bucket(self, "example",
            bucket="example"
        )
        aws_s3_object_example = S3Object(self, "example_1",
            bucket=example.id,
            key="debezium.zip",
            source="debezium.zip"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_object_example.override_logical_id("example")
        aws_mskconnect_custom_plugin_example = MskconnectCustomPlugin(self, "example_2",
            content_type="ZIP",
            location=MskconnectCustomPluginLocation(
                s3=MskconnectCustomPluginLocationS3(
                    bucket_arn=example.arn,
                    file_key=Token.as_string(aws_s3_object_example.key)
                )
            ),
            name="debezium-example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_mskconnect_custom_plugin_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `name` - (Required) The name of the custom plugin..
* `content_type` - (Required) The type of the plugin file. Allowed values are `ZIP` and `JAR`.
* `location` - (Required) Information about the location of a custom plugin. See below.

The following arguments are optional:

* `description` - (Optional) A summary description of the custom plugin.

### location Argument Reference

* `s3` - (Required) Information of the plugin file stored in Amazon S3. See below.

#### location s3 Argument Reference

* `bucket_arn` - (Required) The Amazon Resource Name (ARN) of an S3 bucket.
* `file_key` - (Required) The file key for an object in an S3 bucket.
* `object_version` - (Optional) The version of an object in an S3 bucket.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - the Amazon Resource Name (ARN) of the custom plugin.
* `latest_revision` - an ID of the latest successfully created revision of the custom plugin.
* `state` - the state of the custom plugin.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import MSK Connect Custom Plugin using the plugin's `arn`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import MSK Connect Custom Plugin using the plugin's `arn`. For example:

```console
% terraform import aws_mskconnect_custom_plugin.example 'arn:aws:kafkaconnect:eu-central-1:123456789012:custom-plugin/debezium-example/abcdefgh-1234-5678-9abc-defghijklmno-4'
```

<!-- cache-key: cdktf-0.20.0 input-fa6ff218273ad9d6a29b1b6d47850b35e3eb2ccea8af6818ee7abaf4e10e3981 -->