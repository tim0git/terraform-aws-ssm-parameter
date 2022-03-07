variable parameters {
  type = list(object({
    name = string
    value = string
    description = string
    type = string #String, StringList and SecureString
    overwrite = bool
  }))
  default = []
}

variable "tags" {
  default = {
    Name = "my-ssm-parameter"
  }
  type        = map(string)
  description = "(Optional) Map of key-value resource tags to associate with the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level."
}