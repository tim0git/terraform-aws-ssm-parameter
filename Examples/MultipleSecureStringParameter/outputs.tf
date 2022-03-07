output "ssm_parameters" {
  value = module.multiple_secure_string_ssm_parameter.ssm_parameters
  description = "ssm parameters"
}