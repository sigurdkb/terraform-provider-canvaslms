package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sigurdkb/canvaslms-client-go"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("CANVAS_AUTH_TOKEN", ""),
				Required:    true,
				Sensitive:   true,
			},

			"course_code": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	baseUrl := d.Get("base_url").(string)
	token := d.Get("token").(string)
	courseCode := d.Get("course_code").(int)

	client, err := canvaslms.NewClient(&baseUrl, &token)
	if err != nil {
		return append(diags, diag.Errorf("Error creating rest client: %s", err)...)
	}

	course, err := client.GetCourse(courseCode)
	if err != nil {
		return append(diags, diag.Errorf("Error retreiving course: %s", err)...)
	}

	result, err := json.Marshal(course)
	if err != nil {
		return append(diags, diag.Errorf("error marshaling json: %s", err)...)
	}

	d.Set("body", string(result))

	// set ID as something more stable than time
	d.SetId(baseUrl)

	return diags
}
