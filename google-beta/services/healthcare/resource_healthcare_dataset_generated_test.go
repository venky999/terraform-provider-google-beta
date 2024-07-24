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

package healthcare_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccHealthcareDataset_healthcareDatasetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDataset_healthcareDatasetBasicExample(context),
			},
			{
				ResourceName:            "google_healthcare_dataset.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "self_link"},
			},
		},
	})
}

func testAccHealthcareDataset_healthcareDatasetBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_dataset" "default" {
  name      = "tf-test-example-dataset%{random_suffix}"
  location  = "us-central1"
  time_zone = "UTC"
}
`, context)
}

func TestAccHealthcareDataset_healthcareDatasetCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDataset_healthcareDatasetCmekExample(context),
			},
			{
				ResourceName:            "google_healthcare_dataset.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "self_link"},
			},
		},
	})
}

func testAccHealthcareDataset_healthcareDatasetCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {}

resource "google_healthcare_dataset" "default" {
  name      = "tf-test-example-dataset%{random_suffix}"
  location  = "us-central1"
  time_zone = "UTC"

  encryption_spec {
    kms_key_name = google_kms_crypto_key.crypto_key.id
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.healthcare_cmek_keyuser
  ]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-example-key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-example-keyring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key_iam_binding" "healthcare_cmek_keyuser" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-healthcare.iam.gserviceaccount.com",
  ]
}
`, context)
}

func testAccCheckHealthcareDatasetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_dataset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.HealthcareDatasetNotInitialized},
			})
			if err == nil {
				return fmt.Errorf("HealthcareDataset still exists at %s", url)
			}
		}

		return nil
	}
}
