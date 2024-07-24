// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package firebaseappcheck_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirebaseAppCheckDebugToken_firebaseAppCheckDebugTokenBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"token":         "5E728315-E121-467F-BCA1-1FE71130BB98",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckFirebaseAppCheckDebugTokenDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckDebugToken_firebaseAppCheckDebugTokenBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_app_check_debug_token.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_id", "token"},
			},
		},
	})
}

func testAccFirebaseAppCheckDebugToken_firebaseAppCheckDebugTokenBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_web_app" "default" {
  provider = google-beta

  project      = "%{project_id}"
  display_name = "Web App for debug token"
}

# It takes a while for App Check to recognize the new app
# If your app already exists, you don't have to wait 30 seconds.
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_web_app.default]
  create_duration = "30s"
}

resource "google_firebase_app_check_debug_token" "default" {
  provider = google-beta

  project      = "%{project_id}"
  app_id       = google_firebase_web_app.default.app_id
  display_name = "Debug Token%{random_suffix}"
  token        = "%{token}"

  depends_on = [time_sleep.wait_30s]
}
`, context)
}

func testAccCheckFirebaseAppCheckDebugTokenDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_app_check_debug_token" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseAppCheckBasePath}}projects/{{project}}/apps/{{app_id}}/debugTokens/{{debug_token_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("FirebaseAppCheckDebugToken still exists at %s", url)
			}
		}

		return nil
	}
}
