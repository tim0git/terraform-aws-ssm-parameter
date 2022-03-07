module "multiple_secure_string_ssm_parameter" {
  source = "../../"

  parameters = [{
    name = "/test/secure_string/ssm_parameter_1"
    description = "test description 1"
    value = "test value one"
    type = "SecureString"
    overwrite = false
  },{
    name = "/test/secure_string/ssm_parameter_2"
    description = "test description 2"
    value = "test value two"
    type = "SecureString"
    overwrite = false
  }]

  tags = {
    Name = "test-secure-string-ssm-parameter"
  }
}