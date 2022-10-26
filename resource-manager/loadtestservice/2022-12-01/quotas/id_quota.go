package quotas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = QuotaId{}

// QuotaId is a struct representing the Resource ID for a Quota
type QuotaId struct {
	SubscriptionId  string
	Location        string
	QuotaBucketName string
}

// NewQuotaID returns a new QuotaId struct
func NewQuotaID(subscriptionId string, location string, quotaBucketName string) QuotaId {
	return QuotaId{
		SubscriptionId:  subscriptionId,
		Location:        location,
		QuotaBucketName: quotaBucketName,
	}
}

// ParseQuotaID parses 'input' into a QuotaId
func ParseQuotaID(input string) (*QuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.QuotaBucketName, ok = parsed.Parsed["quotaBucketName"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaBucketName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseQuotaIDInsensitively parses 'input' case-insensitively into a QuotaId
// note: this method should only be used for API response data and not user input
func ParseQuotaIDInsensitively(input string) (*QuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(QuotaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := QuotaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.QuotaBucketName, ok = parsed.Parsed["quotaBucketName"]; !ok {
		return nil, fmt.Errorf("the segment 'quotaBucketName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateQuotaID checks that 'input' can be parsed as a Quota ID
func ValidateQuotaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQuotaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Quota ID
func (id QuotaId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.LoadTestService/locations/%s/quotas/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.QuotaBucketName)
}

// Segments returns a slice of Resource ID Segments which comprise this Quota ID
func (id QuotaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLoadTestService", "Microsoft.LoadTestService", "Microsoft.LoadTestService"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticQuotas", "quotas", "quotas"),
		resourceids.UserSpecifiedSegment("quotaBucketName", "quotaBucketValue"),
	}
}

// String returns a human-readable description of this Quota ID
func (id QuotaId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Quota Bucket Name: %q", id.QuotaBucketName),
	}
	return fmt.Sprintf("Quota (%s)", strings.Join(components, "\n"))
}