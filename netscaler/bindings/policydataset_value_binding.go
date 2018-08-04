package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicydatasetValueBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policydataset_value_binding_func,
		//                Read:          read_policydataset_value_binding_func,
		//                Update:        update_policydataset_value_binding_func,
		//                Delete:        delete_policydataset_value_binding_func,
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
