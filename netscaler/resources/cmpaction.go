package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCmpaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cmpaction_func,
		//                Read:          read_cmpaction_func,
		//                Update:        update_cmpaction_func,
		//                Delete:        delete_cmpaction_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addvaryheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmptype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"deltatype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"varyheadervalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
