package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_dnspolicy64_func,
		//                Read:          read_dnspolicy64_func,
		//                Update:        update_dnspolicy64_func,
		//                Delete:        delete_dnspolicy64_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
