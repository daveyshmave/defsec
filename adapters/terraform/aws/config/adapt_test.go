package config

import (
	"testing"

	"github.com/aquasecurity/defsec/adapters/terraform/testutil"
	"github.com/aquasecurity/defsec/parsers/types"
	"github.com/stretchr/testify/assert"

	"github.com/aquasecurity/defsec/providers/aws/config"
)

func Test_adaptConfigurationAggregrator(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  config.ConfigurationAggregrator
	}{
		{
			name: "configured",
			terraform: `
			resource "aws_config_configuration_aggregator" "example" {
				name = "example"
				  
				account_aggregation_source {
				  account_ids = ["123456789012"]
				  all_regions = true
				}
			}
`,
			expected: config.ConfigurationAggregrator{
				Metadata:         types.NewTestMetadata(),
				SourceAllRegions: types.Bool(true, types.NewTestMetadata()),
				IsDefined:        true,
			},
		},
		{
			name: "defaults",
			terraform: `
			resource "aws_config_configuration_aggregator" "example" {
			}
`,
			expected: config.ConfigurationAggregrator{
				Metadata:         types.NewTestMetadata(),
				SourceAllRegions: types.Bool(false, types.NewTestMetadata()),
				IsDefined:        true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutil.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := adaptConfigurationAggregrator(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

func TestLines(t *testing.T) {
	src := `
	resource "aws_config_configuration_aggregator" "example" {
		name = "example"
		  
		account_aggregation_source {
		  account_ids = ["123456789012"]
		  all_regions = true
		}
	}`

	modules := testutil.CreateModulesFromSource(src, ".tf", t)
	adapted := Adapt(modules)
	aggregator := adapted.ConfigurationAggregrator

	assert.Equal(t, 2, aggregator.GetMetadata().Range().GetStartLine())
	assert.Equal(t, 9, aggregator.GetMetadata().Range().GetEndLine())

	assert.Equal(t, 7, aggregator.SourceAllRegions.GetMetadata().Range().GetStartLine())
	assert.Equal(t, 7, aggregator.SourceAllRegions.GetMetadata().Range().GetEndLine())
}
