---
subcategory: "Detective"
layout: "aws"
page_title: "AWS: aws_detective_organization_admin_account"
description: |-
  Manages a Detective Organization Admin Account
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_detective_organization_admin_account

Manages a Detective Organization Admin Account. The AWS account utilizing this resource must be an Organizations primary account. More information about Organizations support in Detective can be found in the [Detective User Guide](https://docs.aws.amazon.com/detective/latest/adminguide/accounts-orgs-transition.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DetectiveOrganizationAdminAccount } from "./.gen/providers/aws/detective-organization-admin-account";
import { OrganizationsOrganization } from "./.gen/providers/aws/organizations-organization";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new OrganizationsOrganization(this, "example", {
      awsServiceAccessPrincipals: ["detective.amazonaws.com"],
      featureSet: "ALL",
    });
    const awsDetectiveOrganizationAdminAccountExample =
      new DetectiveOrganizationAdminAccount(this, "example_1", {
        accountId: "123456789012",
        dependsOn: [example],
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDetectiveOrganizationAdminAccountExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are supported:

* `accountId` - (Required) AWS account identifier to designate as a delegated administrator for Detective.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - AWS account identifier.

## Import

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_detective_organization_admin_account` using `accountId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

Using `terraform import`, import `aws_detective_organization_admin_account` using `accountId`. For example:

```console
% terraform import aws_detective_organization_admin_account.example 123456789012
```

<!-- cache-key: cdktf-0.20.0 input-780ffda473384073f83f5f613539908079196f4002299af4677d3d4bf560d48a -->