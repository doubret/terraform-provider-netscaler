package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_caaction_func,
		//                Read:          read_caaction_func,
		//                Update:        update_caaction_func,
		//                Delete:        delete_caaction_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"accumressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
