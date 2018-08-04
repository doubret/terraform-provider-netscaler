package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicydataset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policydataset_func,
		//                Read:          read_policydataset_func,
		//                Update:        update_policydataset_func,
		//                Delete:        delete_policydataset_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
