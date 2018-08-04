package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCmppolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cmppolicylabel_func,
		//                Read:          read_cmppolicylabel_func,
		//                Update:        update_cmppolicylabel_func,
		//                Delete:        delete_cmppolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
