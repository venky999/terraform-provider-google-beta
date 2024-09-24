// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// TestAccSdkProvider_credentials is a series of acc tests asserting how the SDK provider handles credentials arguments
// It is SDK specific because the HCL used provisions SDK-implemented resources
// It is a counterpart to TestAccFwProvider_credentials
func TestAccSdkProvider_credentials(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"credentials can be configured as a path to a credentials JSON file":                                       testAccSdkProvider_credentials_validJsonFilePath,
		"configuring credentials as a path to a non-existent file results in an error":                             testAccSdkProvider_credentials_badJsonFilepathCausesError,
		"config takes precedence over environment variables":                                                       testAccSdkProvider_credentials_configPrecedenceOverEnvironmentVariables,
		"when credentials is unset in the config, environment variables are used in a given order":                 testAccSdkProvider_credentials_precedenceOrderEnvironmentVariables, // GOOGLE_CREDENTIALS, GOOGLE_CLOUD_KEYFILE_JSON, GCLOUD_KEYFILE_JSON, GOOGLE_APPLICATION_CREDENTIALS
		"when credentials is set to an empty string in the config the value isn't ignored and results in an error": testAccSdkProvider_credentials_emptyStringValidation,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccSdkProvider_credentials_validJsonFilePath(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	// unset all credentials env vars
	for _, v := range envvar.CredsEnvVars {
		t.Setenv(v, "")
	}

	credentials := transport_tpg.TestFakeCredentialsPath

	context := map[string]interface{}{
		"credentials": credentials,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Credentials set as what we expect
				Config: testAccSdkProvider_credentialsInProviderBlock(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_sdk.default", "credentials", credentials),
				),
			},
		},
	})
}

func testAccSdkProvider_credentials_badJsonFilepathCausesError(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	// unset all credentials env vars
	for _, v := range envvar.CredsEnvVars {
		t.Setenv(v, "")
	}

	pathToMissingFile := "./this/path/does/not/exist.json" // Doesn't exist

	context := map[string]interface{}{
		"credentials": pathToMissingFile,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Apply-time error due to the file not existing
				Config:      testAccSdkProvider_credentialsInProviderBlock(context),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("JSON credentials are not valid"),
			},
		},
	})
}

func testAccSdkProvider_credentials_configPrecedenceOverEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	credentials := envvar.GetTestCredsFromEnv()

	// ensure all possible credentials env vars set; show they aren't used
	for _, v := range envvar.CredsEnvVars {
		t.Setenv(v, credentials)
	}

	pathToMissingFile := "./this/path/does/not/exist.json" // Doesn't exist

	context := map[string]interface{}{
		"credentials": pathToMissingFile,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Apply-time error; bad value in config is used over of good values in ENVs
				Config:      testAccSdkProvider_credentialsInProviderBlock(context),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("JSON credentials are not valid"),
			},
		},
	})
}

func testAccSdkProvider_credentials_precedenceOrderEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API
	/*
		These are all the ENVs for credentials, and they are in order of precedence.
		GOOGLE_CREDENTIALS
		GOOGLE_CLOUD_KEYFILE_JSON
		GCLOUD_KEYFILE_JSON
		GOOGLE_APPLICATION_CREDENTIALS
		GOOGLE_USE_DEFAULT_CREDENTIALS
	*/

	GOOGLE_CREDENTIALS := acctest.GenerateFakeCredentialsJson("GOOGLE_CREDENTIALS")
	GOOGLE_CLOUD_KEYFILE_JSON := acctest.GenerateFakeCredentialsJson("GOOGLE_CLOUD_KEYFILE_JSON")
	GCLOUD_KEYFILE_JSON := acctest.GenerateFakeCredentialsJson("GCLOUD_KEYFILE_JSON")
	GOOGLE_APPLICATION_CREDENTIALS := "./fake/file/path/nonexistent/a/credentials.json" // GOOGLE_APPLICATION_CREDENTIALS needs to be a path, not JSON

	context := map[string]interface{}{}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// GOOGLE_CREDENTIALS is used 1st if set
				PreConfig: func() {
					t.Setenv("GOOGLE_CREDENTIALS", GOOGLE_CREDENTIALS) //used
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", GOOGLE_CLOUD_KEYFILE_JSON)
					t.Setenv("GCLOUD_KEYFILE_JSON", GCLOUD_KEYFILE_JSON)
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GOOGLE_APPLICATION_CREDENTIALS)
				},
				Config: testAccSdkProvider_credentialsInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_sdk.default", "credentials", GOOGLE_CREDENTIALS),
				),
			},
			{
				// GOOGLE_CLOUD_KEYFILE_JSON is used 2nd
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					// set
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", GOOGLE_CLOUD_KEYFILE_JSON) //used
					t.Setenv("GCLOUD_KEYFILE_JSON", GCLOUD_KEYFILE_JSON)
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GOOGLE_APPLICATION_CREDENTIALS)

				},
				Config: testAccSdkProvider_credentialsInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_sdk.default", "credentials", GOOGLE_CLOUD_KEYFILE_JSON),
				),
			},
			{
				// GOOGLE_CLOUD_KEYFILE_JSON is used 3rd
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", "")
					// set
					t.Setenv("GCLOUD_KEYFILE_JSON", GCLOUD_KEYFILE_JSON) //used
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GOOGLE_APPLICATION_CREDENTIALS)
				},
				Config: testAccSdkProvider_credentialsInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_sdk.default", "credentials", GCLOUD_KEYFILE_JSON),
				),
			},
			{
				// GOOGLE_APPLICATION_CREDENTIALS is used 4th
				PreConfig: func() {
					// unset
					t.Setenv("GOOGLE_CREDENTIALS", "")
					t.Setenv("GOOGLE_CLOUD_KEYFILE_JSON", "")
					t.Setenv("GCLOUD_KEYFILE_JSON", "")
					// set
					t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", GOOGLE_APPLICATION_CREDENTIALS) //used
				},
				Config:      testAccSdkProvider_credentialsInEnvsOnly(context),
				ExpectError: regexp.MustCompile("no such file or directory"),
			},
			// Need a step to help post-test destroy run without error from GOOGLE_APPLICATION_CREDENTIALS
			{
				PreConfig: func() {
					t.Setenv("GOOGLE_CREDENTIALS", GOOGLE_CREDENTIALS)
				},
				Config: "// Need a step to help post-test destroy run without error",
			},
		},
	})
}

func testAccSdkProvider_credentials_emptyStringValidation(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	credentials := envvar.GetTestCredsFromEnv()

	// ensure all credentials env vars set
	for _, v := range envvar.CredsEnvVars {
		t.Setenv(v, credentials)
	}

	context := map[string]interface{}{
		"credentials": "", // empty string used
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccSdkProvider_credentialsInProviderBlock(context),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("expected a non-empty string"),
			},
		},
	})
}

// testAccSdkProvider_credentialsInProviderBlock allows setting the credentials argument in a provider block.
// This function uses data.google_provider_config_sdk because it is implemented with the SDKv2
func testAccSdkProvider_credentialsInProviderBlock(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	credentials = "%{credentials}"
}

data "google_provider_config_sdk" "default" {}
`, context)
}

// testAccSdkProvider_credentialsInEnvsOnly allows testing when the credentials argument
// is only supplied via ENVs
func testAccSdkProvider_credentialsInEnvsOnly(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_provider_config_sdk" "default" {}
`, context)
}
