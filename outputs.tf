output "ssm_parameters" {
  value = [
    for parameter in aws_ssm_parameter.this : {
    name = parameter.name
    description = parameter.description
    type = parameter.type
    version = parameter.version
    overwrite = parameter.overwrite
    }
  ]
}