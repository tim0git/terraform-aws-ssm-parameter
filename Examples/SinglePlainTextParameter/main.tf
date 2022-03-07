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