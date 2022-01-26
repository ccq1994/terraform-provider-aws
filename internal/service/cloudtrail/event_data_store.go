package cloudtrail

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func ResourceEventDataStore() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEventDataStoreCreate,
		ReadContext:   resourceEventDataStoreRead,
		UpdateContext: resourceEventDataStoreUpdate,
		DeleteContext: resourceEventDataStoreDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(eventDataStoreAvailableTimeout),
			Delete: schema.DefaultTimeout(eventDataStoreDeletedTimeout),
		},

		CustomizeDiff: verify.SetTagsDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(3, 128),
			},
			"advanced_event_selector": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_selector": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ends_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"equals": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"field": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.StringInSlice(field_Values(), false),
									},
									"not_ends_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"not_equals": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"not_starts_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"starts_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(0, 1000),
						},
					},
				},
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_region_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"organization_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"retention_period": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: validation.All(
					validation.IntBetween(7, 2555),
				),
			},
			"termination_protection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}

func resourceEventDataStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	name := d.Get("name").(string)
	input := &cloudtrail.CreateEventDataStoreInput{
		Name:                         aws.String(name),
		OrganizationEnabled:          aws.Bool(d.Get("organization_enabled").(bool)),
		MultiRegionEnabled:           aws.Bool(d.Get("multi_region_enabled").(bool)),
		TerminationProtectionEnabled: aws.Bool(d.Get("termination_protection_enabled").(bool)),
		RetentionPeriod:              aws.Int64(int64(d.Get("retention_period").(int))),
	}

	if _, ok := d.GetOk("advanced_event_selector"); ok {
		input.AdvancedEventSelectors = expandAdvancedEventSelector(d.Get("advanced_event_selector").([]interface{}))
	}

	if len(tags) > 0 {
		input.TagsList = Tags(tags.IgnoreAWS())
	}

	if err := input.Validate(); err != nil {
		return diag.Errorf("Error validating CloudTrail Event Data Store (%s): %s", name, err)
	}

	log.Printf("[DEBUG] Creating Event Data Store: %s", input)
	output, err := conn.CreateEventDataStoreWithContext(ctx, input)

	if err != nil {
		return diag.Errorf("error creating CloudTrail Event Data Store (%s): %s", name, err)
	}

	eventDataStoreArn := aws.StringValue(output.EventDataStoreArn)

	if err := waitEventDataStoreAvailable(ctx, conn, eventDataStoreArn, d.Timeout(schema.TimeoutCreate)); err != nil {
		return diag.Errorf("error waiting for CloudTrail Event Data Store (%s) to be created: %s", name, err)
	}

	d.SetId(eventDataStoreArn)

	return resourceEventDataStoreRead(ctx, d, meta)
}

func resourceEventDataStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn

	if d.HasChangesExcept("tags", "tags_all") {
		input := &cloudtrail.UpdateEventDataStoreInput{
			EventDataStore: aws.String(d.Id()),
		}

		if d.HasChange("name") {
			input.Name = aws.String(d.Get("name").(string))
		}

		if d.HasChange("multi_region_enabled") {
			input.MultiRegionEnabled = aws.Bool(d.Get("multi_region_enabled").(bool))
		}

		if d.HasChange("organization_enabled") {
			input.MultiRegionEnabled = aws.Bool(d.Get("organization_enabled").(bool))
		}

		if d.HasChange("termination_protection_enabled") {
			input.MultiRegionEnabled = aws.Bool(d.Get("termination_protection_enabled").(bool))
		}

		if d.HasChange("retention_period") {
			input.RetentionPeriod = aws.Int64(int64(d.Get("retention_period").(int)))
		}

		if !d.IsNewResource() && d.HasChange("advanced_event_selector") {
			input.AdvancedEventSelectors = expandAdvancedEventSelector(d.Get("advanced_event_selector").([]interface{}))
		}

		log.Printf("[DEBUG] Updating CloudTrail Event Data Store (%s)", d.Id())

		if err := input.Validate(); err != nil {
			return diag.Errorf("Error validating CloudTrail Event Data Store (%s): %s", d.Id(), err)
		}

		_, err := conn.UpdateEventDataStoreWithContext(ctx, input)

		if err != nil {
			return diag.Errorf("error updating CloudTrail Event Data Store (%s): %s", d.Id(), err)
		}

		if err := waitEventDataStoreAvailable(ctx, conn, d.Id(), eventDataStoreAvailableTimeout); err != nil {
			return diag.Errorf("error waiting for CloudTrail Event Data Store (%s) to be modified: %s", d.Id(), err)
		}
	}

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")

		log.Printf("[DEBUG] Updating CloudTrail Event Data Store (%s) tags", d.Id())
		if err := UpdateTags(conn, d.Get("arn").(string), o, n); err != nil {
			return diag.Errorf("error updating CloudTrail Event Data Store (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceEventDataStoreRead(ctx, d, meta)
}

func resourceEventDataStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn

	log.Printf("[DEBUG] Deleting CloudTrail Event Data Store: (%s)", d.Id())
	_, err := conn.DeleteEventDataStoreWithContext(ctx, &cloudtrail.DeleteEventDataStoreInput{
		EventDataStore: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, cloudtrail.ErrCodeEventDataStoreNotFoundException) {
		return nil
	}

	if err != nil {
		return diag.Errorf("error deleting CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	if err := waitEventDataStoreDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return diag.Errorf("error waiting for CloudTrail Event Data Store (%s) to be deleted: %s", d.Id(), err)
	}

	return nil
}

func resourceEventDataStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	eventDataStore, err := FindEventDataStoreByArn(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] CloudTrail Event Data Store (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("error reading CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	d.Set("arn", eventDataStore.EventDataStoreArn)

	d.Set("created_timestamp", eventDataStore.CreatedTimestamp)
	d.Set("Updated_timestamp", eventDataStore.UpdatedTimestamp)
	d.Set("multi_region_enabled", eventDataStore.MultiRegionEnabled)
	d.Set("organization_enabled", eventDataStore.OrganizationEnabled)
	d.Set("retention_period", eventDataStore.RetentionPeriod)
	d.Set("termination_protection_enabled", eventDataStore.TerminationProtectionEnabled)
	d.Set("advanced_event_selector", flattenAdvancedEventSelector(eventDataStore.AdvancedEventSelectors))

	tags, err := ListTags(conn, d.Id())

	if err != nil {
		return diag.Errorf("error listing tags for CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	//lintignore:AWSR002
	if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
		return diag.Errorf("error setting tags for CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	if err := d.Set("tags_all", tags.Map()); err != nil {
		return diag.Errorf("error setting tags_all for CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	return nil
}
