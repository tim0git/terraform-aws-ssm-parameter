# terraform-aws-ssm-parameter
Terraform module which creates AWS ssm parameters

The following resources will be created

- SSM parameter

Inputs are provided as Terraform variables

parameters:
 - name: The name of the parameter
 - value: The value of the parameter
 - type: The type of the parameter `"String", "StringList" or "SecureString"`
 - description: The description of the parameter
 - overwrite: Whether to overwrite the parameter if it already exists `true or false`
 
tags: A map of tags to assign to the parameter

### Example 1 Basic 
```hcl
module "single_plain_text_parameter" {
  source = "../../"

  parameters = [{
    name = "/test/string/ssm_parameter"
    description = "test description"
    value = "test value one"
    type = "String"
    overwrite = false
  }]

  tags = {
    Name = "test-string-ssm-parameter"
  }
}
```

Example 2 Basic with overwrite

```hcl
module "single_plain_text_parameter" {
  source = "../../"

  parameters = [{
    name = "/test/string/ssm_parameter"
    description = "test description"
    value = "test value one"
    type = "String"
    overwrite = true
  }]

  tags = {
    Name = "test-string-ssm-parameter"
  }
}
```

Example 3 Secure String

```hcl
module "single_secure_string_parameter" {
  source = "../../"

  parameters = [{
    name = "/test/secure_string/ssm_parameter"
    description = "test description"
    value = "test value one"
    type = "SecureString"
    overwrite = false
  }]

  tags = {
    Name = "test-secure-string-ssm-parameter"
  }
}
```

Example 4 Multiple Secure String Parameters

```hcl
module "multiple_secure_string_parameters" {
  source = "../../"

  parameters = [{
    name = "/test/secure_string/ssm_parameter_one"
    description = "test description"
    value = "test value one"
    type = "SecureString"
    overwrite = false
  },
  {
    name = "/test/secure_string/ssm_parameter_two"
    description = "test description"
    value = "test value two"
    type = "SecureString"
    overwrite = false
  }]

  tags = {
    Name = "test-secure-string-ssm-parameters"
  }
}
```

Outputs are available as a list of ssm parameters

Example outputs.tf
```hcl
output "ssm_parameters" {
  value = module.single_plain_text_parameter.ssm_parameters
  description = "ssm parameters"
}
```


