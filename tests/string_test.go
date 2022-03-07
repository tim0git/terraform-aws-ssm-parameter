package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTerraformSinglePlainTextParameter(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../Examples/SinglePlainTextParameter",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputListOfObjects(t, terraformOptions, "ssm_parameters")

	expectedOutput := ssmParameterOutput{
		Name:        "/test/string/ssm_parameter",
		Type:        "String",
		Overwrite:   false,
		Description: "test description",
		Version:     1,
	}

	assert.Equal(t, expectedOutput.Name, output[0]["name"])
	assert.Equal(t, expectedOutput.Type, output[0]["type"])
	assert.Equal(t, expectedOutput.Overwrite, output[0]["overwrite"])
	assert.Equal(t, expectedOutput.Description, output[0]["description"])
	assert.Equal(t, expectedOutput.Version, output[0]["version"])
}
