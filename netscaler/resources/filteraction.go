package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerFilteraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_filteraction_func,
		//                Read:          read_filteraction_func,
		//                Update:        update_filteraction_func,
		//                Delete:        delete_filteraction_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qual": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"respcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
