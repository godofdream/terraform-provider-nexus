/*
Use this resource to create and execute a custom script.

Example Usage

```hcl
resource "nexus_script" "repo_pypi_internal" {
  name    = "create-repo-pypi-internal"
  type    = "groovy"
  content = "repository.createPyPiHosted('pypi-internal')"
}
```
*/
package nexus

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceScriptCreate,
		Read:   resourceScriptRead,
		Update: resourceScriptUpdate,
		Delete: resourceScriptDelete,
		Exists: resourceScriptExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the script.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"content": {
				Description: "The content of this script.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"type": {
				Description: "The type of the script. Default: `groovy`",
				Optional:    true,
				Type:        schema.TypeString,
				Default:     "groovy",
			},
		},
	}
}

func getScriptFromResourceData(d *schema.ResourceData) nexus.Script {
	return nexus.Script{
		Name:    d.Get("name").(string),
		Content: d.Get("content").(string),
		Type:    d.Get("type").(string),
	}
}

func resourceScriptCreate(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)
	script := getScriptFromResourceData(d)

	if err := nexusClient.ScriptCreate(&script); err != nil {
		return err
	}
	// TODO: It should be possible to configure whether to run script or not
	if err := nexusClient.ScriptRun(script.Name); err != nil {
		return err
	}

	d.SetId(script.Name)
	return resourceScriptRead(d, m)
}

func resourceScriptRead(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	script, err := nexusClient.ScriptRead(d.Id())
	if err != nil {
		return err
	}

	if script == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", script.Name)
	d.Set("type", script.Type)
	d.Set("content", script.Content)

	return nil
}

func resourceScriptUpdate(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	if d.HasChange("content") || d.HasChange("type") {
		script := getScriptFromResourceData(d)
		if err := nexusClient.ScriptUpdate(&script); err != nil {
			return err
		}

		if err := nexusClient.ScriptRun(script.Name); err != nil {
			return err
		}
	}

	return resourceScriptRead(d, m)
}

func resourceScriptDelete(d *schema.ResourceData, m interface{}) error {
	service := m.(nexus.NexusService)

	if err := nexusClient.ScriptDelete(d.Id()); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceScriptExists(d *schema.ResourceData, m interface{}) (bool, error) {
	service := m.(nexus.NexusService)

	script, err := nexusClient.ScriptRead(d.Id())
	return script != nil, err
}
