package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPqpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_pqpolicy_func,
		//                Read:          read_pqpolicy_func,
		//                Update:        update_pqpolicy_func,
		//                Delete:        delete_pqpolicy_func,
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"polqdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
