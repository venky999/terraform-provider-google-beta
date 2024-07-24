// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebaseappcheck_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseAppCheckServiceConfigDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(context, "UNENFORCED"),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(context, "ENFORCED"),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
			{
				Config: testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(context, ""),
			},
			{
				ResourceName:            "google_firebase_app_check_service_config.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id"},
			},
		},
	})
}

func testAccFirebaseAppCheckServiceConfig_firebaseAppCheckServiceConfigUpdate(context map[string]interface{}, enforcementMode string) string {
	context["enforcement_mode"] = enforcementMode
	return acctest.Nprintf(`
resource "google_project" "default" {
  provider   = google-beta
  project_id = "tf-test-appcheck%{random_suffix}"
  name       = "tf-test-appcheck%{random_suffix}"
  org_id     = "%{org_id}"
  labels     = {
    "firebase" = "enabled"
  }
}

resource "google_project_service" "firebase" {
  provider = google-beta
  project  = google_project.default.project_id
  service  = "firebase.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "database" {
  provider = google-beta
  project  = google_project.default.project_id
  service  = "firebasedatabase.googleapis.com"
  disable_on_destroy = false
  depends_on = [
    google_project_service.firebase,
  ]
}

resource "google_project_service" "appcheck" {
  provider = google-beta
  project  = google_project.default.project_id
  service  = "firebaseappcheck.googleapis.com"
  disable_on_destroy = false
  depends_on = [
    google_project_service.database,
  ]
}

resource "google_firebase_project" "default" {
  provider = google-beta
  project  = google_project.default.project_id

  depends_on = [
    google_project_service.appcheck,
  ]
}

# It takes a while for the new project to be ready for a database
resource "time_sleep" "wait_30s" {
  depends_on      = [google_firebase_project.default]
  create_duration = "30s"
}

resource "google_firebase_database_instance" "default" {
  provider = google-beta
  project  = google_firebase_project.default.project
  region   = "us-central1"
  instance_id = "tf-test-appcheck%{random_suffix}-default-rtdb"
  type     = "DEFAULT_DATABASE"

  depends_on = [time_sleep.wait_30s]
}

resource "google_firebase_app_check_service_config" "default" {
  provider = google-beta
  project = google_firebase_project.default.project
  service_id = "firebasedatabase.googleapis.com"
  enforcement_mode = "%{enforcement_mode}"

  depends_on = [google_firebase_database_instance.default]
}
`, context)
}
