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

package compute_test

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

func TestAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetIpPortExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetIpPortExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint.region-internet-ip-port-endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "region", "region_network_endpoint_group"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetIpPortExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_endpoint" "region-internet-ip-port-endpoint" {
  region_network_endpoint_group = google_compute_region_network_endpoint_group.group.name
  region                = "us-central1"

  ip_address  = "8.8.8.8"
  port        = 443
}


resource "google_compute_region_network_endpoint_group" "group" {
  name         = "tf-test-ip-port-neg%{random_suffix}"
  network      = google_compute_network.default.id

  region         = "us-central1"
  network_endpoint_type = "INTERNET_IP_PORT"
}

resource "google_compute_network" "default" {
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetFqdnPortExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetFqdnPortExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint.region-internet-fqdn-port-endpoint",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "region", "region_network_endpoint_group"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpoint_regionNetworkEndpointInternetFqdnPortExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_endpoint" "region-internet-fqdn-port-endpoint" {
  region_network_endpoint_group = google_compute_region_network_endpoint_group.group.name
  region                = "us-central1"

  fqdn  = "backend.example.com"
  port        = 443
}


resource "google_compute_region_network_endpoint_group" "group" {
  name         = "tf-test-fqdn-port-neg%{random_suffix}"
  network      = google_compute_network.default.id

  region         = "us-central1"
  network_endpoint_type = "INTERNET_FQDN_PORT"
}

resource "google_compute_network" "default" {
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeRegionNetworkEndpointDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_network_endpoint" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{region_network_endpoint_group}}/listNetworkEndpoints")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "POST",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeRegionNetworkEndpoint still exists at %s", url)
			}
		}

		return nil
	}
}
