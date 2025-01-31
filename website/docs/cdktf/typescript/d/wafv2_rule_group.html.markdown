---
subcategory: "WAF"
layout: "aws"
page_title: "AWS: aws_wafv2_rule_group"
description: |-
  Retrieves the summary of a WAFv2 Rule Group.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_wafv2_rule_group

Retrieves the summary of a WAFv2 Rule Group.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsWafv2RuleGroup } from "./.gen/providers/aws/data-aws-wafv2-rule-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsWafv2RuleGroup(this, "example", {
      name: "some-rule-group",
      scope: "REGIONAL",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `name` - (Required) Name of the WAFv2 Rule Group.
* `scope` - (Required) Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the entity.
* `description` - Description of the rule group that helps with identification.
* `id` - Unique identifier of the rule group.

<!-- cache-key: cdktf-0.20.0 input-4de733ef434bbfbe5aa4d54ed8db1f5e78db59b5a79986eddb88c9cc5144e9ca -->