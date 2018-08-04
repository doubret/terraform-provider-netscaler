package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_spilloveraction_func,
		//                Read:          read_spilloveraction_func,
		//                Update:        update_spilloveraction_func,
		//                Delete:        delete_spilloveraction_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
