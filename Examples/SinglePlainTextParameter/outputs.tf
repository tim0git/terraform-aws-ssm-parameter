output "ssm_parameters" {
  value = module.single_plain_text_parameter.ssm_parameters
  description = "ssm parameters"
}