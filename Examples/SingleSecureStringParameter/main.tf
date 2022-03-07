module "single_secure_string_ssm_parameter" {
  source = "../../"

  parameters = [{
    name = "/test/secure_string/ssm_parameter"
    description = "test description"
    value = "test value"
    type = "SecureString"
    overwrite = false
  }]

  tags = {
    Name = "test-secure-string-ssm-parameter"
  }
}