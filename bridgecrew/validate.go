package bridgecrew

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

// ValidateThreshold ensure the enum is enforced
func ValidateThreshold(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"CRITICAL", "HIGH", "MEDIUM", "LOW":
		return
	}
	errs = append(errs, fmt.Errorf("%q Must be one of CRITICAL, HIGH, MEDIUM, LOW", val))
	return
}

// ValidateOperator looks at valid logic operator inputs
func ValidateOperator(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"contains",
		"ending_with",
		"equals",
		"exists",
		"greater_than",
		"greater_than_or_equal",
		"jsonpath_equals",
		"jsonpath_exists",
		"less_than",
		"less_than_or_equal",
		"not_contains",
		"not_equals",
		"not_ending_with",
		"not_exists",
		"not_reqex_match",
		"not_starting_with",
		"not_subset",
		"regex_match",
		"starting_with",
		"subset",
		"within":
		return
	}
	errs = append(errs, fmt.Errorf("%q Must be one of contains, ending_with, equals, exists,greater_than,"+
		" greater_than_or_equal, jsonpath_equals, jsonpath_exists, less_than, less_than_or_equal, not_contains,"+
		" not_equals, not_ending_with, not_exists, not_reqex_match, not_starting_with, not_subset, regex_match,"+
		" starting_with, subset, within", val))
	return
}

// ValidateCloudProvider checks that only supported cloud providers are added
func ValidateCloudProvider(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"aws",
		"gcp",
		"azure",
		"oci",
		"alibabacloud",
		"kubernetes",
		"openstack":
		return
	}
	errs = append(errs, fmt.Errorf("%q Must be one of aws, gcp, linode, azure, oci, alicloud or digitalocean", val))
	return
}

// ValidateSeverity checks that only supported severities can be added.
func ValidateSeverity(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"critical",
		"high",
		"low",
		"medium":
		return
	}
	errs = append(errs, fmt.Errorf("%q Must be one of critical, high, medium or low", val))
	return
}

// ValidateCategory permits only supported Categories
func ValidateCategory(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"compute",
		"elasticsearch",
		"general",
		"storage",
		"encryption",
		"networking",
		"monitoring",
		"kubernetes",
		"serverless",
		"backup_and_recovery",
		"iam",
		"secrets",
		"public":
		return
	}
	errs = append(errs,
		fmt.Errorf("%q Must be one of elasticsearch, general, storage, encryption,"+
			" networking, monitoring, kubernetes, serverless, backup_and_recovery, backup_and_recovery, public,"+
			" or iam", val))
	return
}

// ValidateIsYAMLFile is this YAML?
func ValidateIsYAMLFile(val interface{}, key string) (warns []string, errors []error) {

	code, err := loadFileContent(val.(string))
	if err != nil {
		errors = append(errors, fmt.Errorf("unable to load %q: %w", val.(string), err))
		return
	}

	if _, err := CheckYAMLString(string(code)); err != nil {
		errors = append(errors, fmt.Errorf("%q contains an invalid YAML: %s", key, err))
	}
	return
}

// ValidateGuidelines is a length check - 50 characters or more please.
func ValidateGuidelines(val interface{}, key string) (warns []string, errs []error) {
	if len(val.(string)) < 50 {
		errs = append(errs, fmt.Errorf("%q Guideline should attempt be helpful (gt 50 chars)", val))
	}
	return
}

// ValidatePolicyTitle is a length check - 20 characters or more please.
func ValidatePolicyTitle(val interface{}, key string) (warns []string, errs []error) {
	if len(val.(string)) < 20 {
		errs = append(errs, fmt.Errorf("%q Title should attempt be meaningful (gt 20 chars)", val))
	}
	return
}

// ValidPolicyJSON checks that a string contains JSON
func ValidPolicyJSON(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) < 1 {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON policy", k))
		return
	}
	if value[:1] != "{" {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON policy", k))
		return
	}
	if _, err := structure.NormalizeJsonString(v); err != nil {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON: %s", k, err))
	}
	return
}

// keyExists looks to see if an item exists in the map
func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}

// ValidateRepository checks that only supported repositories are added
func ValidateRepository(val interface{}, key string) (warns []string, errs []error) {
	switch val.(string) {
	case
		"Github", "Bitbucket",
		"Gitlab", "AzureRepos",
		"cli", "AWS",
		"Azure", "GCP",
		"githubEnterprise", "gitlabEnterprise",
		"bitbucketEnterprise, terraformCloud",
		"tfcRunTasks, githubActions",
		"circleci", "codebuild",
		"jenkins", "kubernetesWorkloads",
		"Kubernetes", "admissionController":
		return
	}
	errs = append(errs, fmt.Errorf("%q Must be one of aws, gcp, linode, azure, oci, alicloud or digitalocean", val))
	return
}
