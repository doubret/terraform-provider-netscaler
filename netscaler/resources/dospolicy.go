package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerDospolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_dospolicy_func,
		//                Read:          read_dospolicy_func,
		//                Update:        update_dospolicy_func,
		//                Delete:        delete_dospolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cltdetectrate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
