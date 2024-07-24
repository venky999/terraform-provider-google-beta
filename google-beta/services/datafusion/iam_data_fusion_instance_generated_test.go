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

package datafusion_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataFusionInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"prober_test_run": `options = { prober_test_run = "true" }`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_data_fusion_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataFusionInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_data_fusion_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataFusionInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"prober_test_run": `options = { prober_test_run = "true" }`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataFusionInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_data_fusion_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataFusionInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/viewer",
		"prober_test_run": `options = { prober_test_run = "true" }`,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstanceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_data_fusion_instance_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_data_fusion_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataFusionInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_data_fusion_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataFusionInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}

resource "google_data_fusion_instance_iam_member" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataFusionInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_data_fusion_instance_iam_policy" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_data_fusion_instance_iam_policy" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  depends_on = [
    google_data_fusion_instance_iam_policy.foo
  ]
}
`, context)
}

func testAccDataFusionInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}

data "google_iam_policy" "foo" {
}

resource "google_data_fusion_instance_iam_policy" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataFusionInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}

resource "google_data_fusion_instance_iam_binding" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataFusionInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_fusion_instance" "basic_instance" {
  name   = "tf-test-my-instance%{random_suffix}"
  region = "us-central1"
  type   = "BASIC"
  %{prober_test_run}
}

resource "google_data_fusion_instance_iam_binding" "foo" {
  project = google_data_fusion_instance.basic_instance.project
  region = google_data_fusion_instance.basic_instance.region
  name = google_data_fusion_instance.basic_instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
