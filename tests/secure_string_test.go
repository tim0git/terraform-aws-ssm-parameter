package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type ssmParameterOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Version     int    `json:"version"`
	Overwrite   bool   `json:"overwrite"`
}

func TestTerraformSingleSecureStringParameter(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../Examples/SingleSecureStringParameter",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputListOfObjects(t, terraformOptions, "ssm_parameters")

	expectedOutput := ssmParameterOutput{
		Name:        "/test/secure_string/ssm_parameter",
		Type:        "SecureString",
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

func TestTerraformMultipleSecureStringParameter(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../Examples/MultipleSecureStringParameter",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputListOfObjects(t, terraformOptions, "ssm_parameters")

	expectedOutputs := []ssmParameterOutput{
		{
			Name:        "/test/secure_string/ssm_parameter_1",
			Type:        "SecureString",
			Overwrite:   false,
			Description: "test description 1",
			Version:     1,
		},
		{
			Name:        "/test/secure_string/ssm_parameter_2",
			Type:        "SecureString",
			Overwrite:   false,
			Description: "test description 2",
			Version:     1,
		},
	}

	for i, expectedOutput := range expectedOutputs {
		assert.Equal(t, expectedOutput.Name, output[i]["name"])
		assert.Equal(t, expectedOutput.Type, output[i]["type"])
		assert.Equal(t, expectedOutput.Overwrite, output[i]["overwrite"])
		assert.Equal(t, expectedOutput.Description, output[i]["description"])
		assert.Equal(t, expectedOutput.Version, output[i]["version"])
	}
}
