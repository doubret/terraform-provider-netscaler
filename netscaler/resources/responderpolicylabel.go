package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerResponderpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_responderpolicylabel_func,
		//                Read:          read_responderpolicylabel_func,
		//                Update:        update_responderpolicylabel_func,
		//                Delete:        delete_responderpolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
