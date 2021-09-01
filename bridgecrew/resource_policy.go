package bridgecrew

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePolicyCreate,
		ReadContext:   resourcePolicyRead,
		UpdateContext: resourePolicyUpdate,
		DeleteContext: resourcePolicyDelete,
		Schema: map[string]*schema.Schema{
			"policies": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Computed: false,
							Required: true,
						},
						//"id": {
						//  Type:     schema.TypeString,
						//  Computed: true,
						//},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"descriptivetitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"constructivetitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Required: true,
						},
						"category": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resourcetypes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type:    schema.TypeString,
								Default: "",
							},
						},
						"accountsdata": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository": {
										Required: true,
										Type:     schema.TypeString,
									},
									"amounts": {
										Required: true,
										Type:     schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"lastupdatedate": {
										Computed: true,
										Type:     schema.TypeString,
									},
								},
							},
						},
						"guideline": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iscustom": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"conditionquery": {
							Type:     schema.TypeString,
							Required: true,
						},
						"benchmarks": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"benchmark": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"createdby": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"code": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourePolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePolicyRead(ctx, d, m)
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
