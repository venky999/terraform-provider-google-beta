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

package networkconnectivity_test

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

func TestAccNetworkConnectivitySpoke_networkConnectivitySpokeLinkedVpcNetworkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivitySpokeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivitySpoke_networkConnectivitySpokeLinkedVpcNetworkBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_spoke.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"hub", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivitySpoke_networkConnectivitySpokeLinkedVpcNetworkBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_connectivity_hub" "basic_hub" {
  name        = "hub1%{random_suffix}"
  description = "A sample hub"
  labels = {
    label-two = "value-one"
  }
}

resource "google_network_connectivity_spoke" "primary"  {
  name = "spoke1%{random_suffix}"
  location = "global"
  description = "A sample spoke with a linked router appliance instance"
  labels = {
    label-one = "value-one"
  }
  hub = google_network_connectivity_hub.basic_hub.id
  linked_vpc_network {
    exclude_export_ranges = [
      "198.51.100.0/24",
      "10.10.0.0/16"
    ]
    uri = google_compute_network.network.self_link
  }
}
`, context)
}

func TestAccNetworkConnectivitySpoke_networkConnectivitySpokeRouterApplianceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivitySpokeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivitySpoke_networkConnectivitySpokeRouterApplianceBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_spoke.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"hub", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivitySpoke_networkConnectivitySpokeRouterApplianceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  name                    = "tf-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "tf-test-subnet%{random_suffix}"
  ip_cidr_range = "10.0.0.0/28"
  region        = "us-central1"
  network       = google_compute_network.network.self_link
}

resource "google_compute_instance" "instance" {
  name         = "tf-test-instance%{random_suffix}"
  machine_type = "e2-medium"
  can_ip_forward = true
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "projects/debian-cloud/global/images/debian-10-buster-v20210817"
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.subnetwork.name
    network_ip = "10.0.0.2"
    access_config {
        network_tier = "PREMIUM"
    }
  }
}

resource "google_network_connectivity_hub" "basic_hub" {
  name        = "tf-test-hub%{random_suffix}"
  description = "A sample hub"
  labels = {
    label-two = "value-one"
  }
}

resource "google_network_connectivity_spoke" "primary" {
  name = "tf-test-name%{random_suffix}"
  location = "us-central1"
  description = "A sample spoke with a linked routher appliance instance"
  labels = {
    label-one = "value-one"
  }
  hub =  google_network_connectivity_hub.basic_hub.id
  linked_router_appliance_instances {
    instances {
        virtual_machine = google_compute_instance.instance.self_link
        ip_address = "10.0.0.2"
    }
    site_to_site_data_transfer = true
  }
}
`, context)
}

func testAccCheckNetworkConnectivitySpokeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_spoke" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/spokes/{{name}}")
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
				return fmt.Errorf("NetworkConnectivitySpoke still exists at %s", url)
			}
		}

		return nil
	}
}
