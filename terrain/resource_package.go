package terrain

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os/exec"
)

func resourcePackage() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackageCreate,
		Read:   resourcePackageRead,
		Update: resourcePackageUpdate,
		Delete: resourcePackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourcePackageCreate(d *schema.ResourceData, m interface{}) error {
	// install the package
	c := m.(providerConfiguration)
	cmd := exec.Command(c.installPkgCommand)
}

func resourcePackageRead(*schema.ResourceData, interface{}) {
	// read in whether it's installed or not?
}

func resourcePackageUpdate(*schema.ResourceData, interface{}) {
	// change version?
}

func resourcePackageDelete(*schema.ResourceData, interface{}) {
	// uninstall
}
