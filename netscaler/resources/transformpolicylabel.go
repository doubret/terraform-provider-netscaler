package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerTransformpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_transformpolicylabel_func,
		//                Read:          read_transformpolicylabel_func,
		//                Update:        update_transformpolicylabel_func,
		//                Delete:        delete_transformpolicylabel_func,
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
