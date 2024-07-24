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

package bigtable_test

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

func TestAccBigtableAppProfile_bigtableAppProfileAnyclusterExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileAnyclusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "ignore_warnings", "ignore_warnings", "instance"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileAnyclusterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "us-central1-a"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-2"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-3"
    zone         = "us-central1-c"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  // Requests will be routed to any of the 3 clusters.
  multi_cluster_routing_use_any = true

  ignore_warnings               = true
}
`, context)
}

func TestAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "ignore_warnings", "ignore_warnings", "instance"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  // Requests will be routed to the following cluster.
  single_cluster_routing {
    cluster_id                 = "cluster-1"
    allow_transactional_writes = true
  }

  ignore_warnings = true
}
`, context)
}

func TestAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "ignore_warnings", "ignore_warnings", "instance"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "us-central1-a"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-2"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-3"
    zone         = "us-central1-c"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  // Requests will be routed to the following 2 clusters.
  multi_cluster_routing_use_any = true
  multi_cluster_routing_cluster_ids = ["cluster-1", "cluster-2"]

  ignore_warnings               = true
}
`, context)
}

func TestAccBigtableAppProfile_bigtableAppProfilePriorityExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigtableAppProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfilePriorityExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"app_profile_id", "ignore_warnings", "ignore_warnings", "instance"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfilePriorityExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-bt-instance%{random_suffix}"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }

  deletion_protection  = "%{deletion_protection}"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-bt-profile%{random_suffix}"

  // Requests will be routed to the following cluster.
  single_cluster_routing {
    cluster_id                 = "cluster-1"
    allow_transactional_writes = true
  }

  standard_isolation {
    priority = "PRIORITY_LOW"
  }

  ignore_warnings = true
}
`, context)
}

func testAccCheckBigtableAppProfileDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigtable_app_profile" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
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
				return fmt.Errorf("BigtableAppProfile still exists at %s", url)
			}
		}

		return nil
	}
}
