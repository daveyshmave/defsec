package s3

import (
	"github.com/aquasecurity/defsec/providers"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckVersioningIsEnabled = rules.Register(
	rules.Rule{
		AVDID:      "AVD-AWS-0090",
		Provider:   providers.AWSProvider,
		Service:    "s3",
		ShortCode:  "enable-versioning",
		Summary:    "S3 Data should be versioned",
		Impact:     "Deleted or modified data would not be recoverable",
		Resolution: "Enable versioning to protect against accidental/malicious removal or modification",
		Explanation: `
Versioning in Amazon S3 is a means of keeping multiple variants of an object in the same bucket. 
You can use the S3 Versioning feature to preserve, retrieve, and restore every version of every object stored in your buckets. 
With versioning you can recover more easily from both unintended user actions and application failures.
`,
		Links: []string{
			"https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html",
		},
		Terraform: &rules.EngineMetadata{
			GoodExamples:        terraformEnableVersioningGoodExamples,
			BadExamples:         terraformEnableVersioningBadExamples,
			Links:               terraformEnableVersioningLinks,
			RemediationMarkdown: terraformEnableVersioningRemediationMarkdown,
		},
		CloudFormation: &rules.EngineMetadata{
			GoodExamples:        cloudFormationEnableVersioningGoodExamples,
			BadExamples:         cloudFormationEnableVersioningBadExamples,
			Links:               cloudFormationEnableVersioningLinks,
			RemediationMarkdown: cloudFormationEnableVersioningRemediationMarkdown,
		},
		Severity: severity.Medium,
	},
	func(s *state.State) (results rules.Results) {
		for _, bucket := range s.AWS.S3.Buckets {
			if !bucket.Versioning.Enabled.IsTrue() {
				results.Add(
					"Bucket does not have versioning enabled",
					bucket.Versioning.Enabled,
				)
			} else {
				results.AddPassed(&bucket)
			}
		}
		return results
	},
)
