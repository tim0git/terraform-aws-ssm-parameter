TO DO.
#tier - (Optional) The tier of the parameter. If not specified, will default to Standard. Valid tiers are Standard, Advanced, and Intelligent-Tiering. For more information on parameter tiers, see the AWS SSM Parameter tier comparison and guide.

#key_id - (Optional) The KMS key id or arn for encrypting a SecureString.

#allowed_pattern - (Optional) A regular expression used to validate the parameter value.

#data_type - (Optional) The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the Native parameter support for Amazon Machine Image IDs