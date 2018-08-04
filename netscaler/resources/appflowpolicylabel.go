package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAppflowpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_appflowpolicylabel_func,
		//                Read:          read_appflowpolicylabel_func,
		//                Update:        update_appflowpolicylabel_func,
		//                Delete:        delete_appflowpolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
