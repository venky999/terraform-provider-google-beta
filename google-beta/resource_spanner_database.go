// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSpannerDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpannerDatabaseCreate,
		Read:   resourceSpannerDatabaseRead,
		Delete: resourceSpannerDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSpannerDatabaseImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRegexp(`^(?:[a-z](?:[-_a-z0-9]{0,28}[a-z0-9])?)$`),
			},
			"ddl": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSpannerDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandSpannerDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	extraStatementsProp, err := expandSpannerDatabaseDdl(d.Get("ddl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ddl"); !isEmptyValue(reflect.ValueOf(extraStatementsProp)) && (ok || !reflect.DeepEqual(v, extraStatementsProp)) {
		obj["extraStatements"] = extraStatementsProp
	}
	instanceProp, err := expandSpannerDatabaseInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	obj, err = resourceSpannerDatabaseEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://spanner.googleapis.com/v1/projects/{{project}}/instances/{{instance}}/databases")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Database: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Database: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{instance}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Database %q: %#v", d.Id(), res)

	return resourceSpannerDatabaseRead(d, meta)
}

func resourceSpannerDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://spanner.googleapis.com/v1/projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("SpannerDatabase %q", d.Id()))
	}

	res, err = resourceSpannerDatabaseDecoder(d, meta, res)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	if err := d.Set("name", flattenSpannerDatabaseName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("state", flattenSpannerDatabaseState(res["state"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("instance", flattenSpannerDatabaseInstance(res["instance"], d)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	return nil
}

func resourceSpannerDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://spanner.googleapis.com/v1/projects/{{project}}/instances/{{instance}}/databases/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Database %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Database")
	}

	log.Printf("[DEBUG] Finished deleting Database %q: %#v", d.Id(), res)
	return nil
}

func resourceSpannerDatabaseImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)", "instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<name>[^/]+)", "(?P<instance>[^/]+)/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{instance}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSpannerDatabaseName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSpannerDatabaseState(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenSpannerDatabaseInstance(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandSpannerDatabaseName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseDdl(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSpannerDatabaseInstance(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("instances", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for instance: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceSpannerDatabaseEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["createStatement"] = fmt.Sprintf("CREATE DATABASE `%s`", obj["name"])
	delete(obj, "name")
	delete(obj, "instance")
	return obj, nil
}

func resourceSpannerDatabaseDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	d.SetId(res["name"].(string))
	log.Printf("[DEBUG] name = %s", res["name"])
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/databases/(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}
	res["project"] = d.Get("project").(string)
	res["instance"] = d.Get("instance").(string)
	res["name"] = d.Get("name").(string)
	log.Printf("[DEBUG] result %#v", res)
	id, err := replaceVars(d, config, "{{instance}}/{{name}}")
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	return res, nil
}
