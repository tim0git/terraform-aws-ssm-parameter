resource "aws_ssm_parameter" "this" {
  count = length(var.parameters)
  name  = var.parameters[count.index].name
  type  = var.parameters[count.index].type
  value = var.parameters[count.index].value
  description = var.parameters[count.index].description
  overwrite = var.parameters[count.index].overwrite
}