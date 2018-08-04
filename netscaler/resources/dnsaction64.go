package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerDnsaction64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_dnsaction64_func,
		//                Read:          read_dnsaction64_func,
		//                Update:        update_dnsaction64_func,
		//                Delete:        delete_dnsaction64_func,
		Schema: map[string]*schema.Schema{
			"actionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"excluderule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mappedrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
