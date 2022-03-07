output "ssm_parameter_name" {
  value = aws_ssm_parameter.this.*.name
}

output "ssm_parameter_description" {
  value = aws_ssm_parameter.this.*.description
}

output "ssm_parameter_type" {
  value = aws_ssm_parameter.this.*.type
}