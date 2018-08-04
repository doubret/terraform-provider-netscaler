package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAuditnslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_auditnslogpolicy_func,
		//                Read:          read_auditnslogpolicy_func,
		//                Update:        update_auditnslogpolicy_func,
		//                Delete:        delete_auditnslogpolicy_func,
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
