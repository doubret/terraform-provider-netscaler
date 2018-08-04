package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAuditsyslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_auditsyslogpolicy_func,
		//                Read:          read_auditsyslogpolicy_func,
		//                Update:        update_auditsyslogpolicy_func,
		//                Delete:        delete_auditsyslogpolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
