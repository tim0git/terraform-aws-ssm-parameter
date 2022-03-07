output "ssm_parameters" {
  value = [{
    name: join("", aws_ssm_parameter.this[*].name),
    description: join("", aws_ssm_parameter.this[*].description),
    type: join("", aws_ssm_parameter.this[*].type),
    version: join("", aws_ssm_parameter.this[*].version)
    overwrite: join("", aws_ssm_parameter.this[*].overwrite)
  }]
}