package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerTransformprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_transformprofile_func,
		//                Read:          read_transformprofile_func,
		//                Update:        update_transformprofile_func,
		//                Delete:        delete_transformprofile_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"onlytransformabsurlinbody": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
