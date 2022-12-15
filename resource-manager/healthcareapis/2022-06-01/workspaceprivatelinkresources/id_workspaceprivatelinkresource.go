package workspaceprivatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = WorkspacePrivateLinkResourceId{}

// WorkspacePrivateLinkResourceId is a struct representing the Resource ID for a Workspace Private Link Resource
type WorkspacePrivateLinkResourceId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	GroupName         string
}

// NewWorkspacePrivateLinkResourceID returns a new WorkspacePrivateLinkResourceId struct
func NewWorkspacePrivateLinkResourceID(subscriptionId string, resourceGroupName string, workspaceName string, groupName string) WorkspacePrivateLinkResourceId {
	return WorkspacePrivateLinkResourceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		GroupName:         groupName,
	}
}

// ParseWorkspacePrivateLinkResourceID parses 'input' into a WorkspacePrivateLinkResourceId
func ParseWorkspacePrivateLinkResourceID(input string) (*WorkspacePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.GroupName, ok = parsed.Parsed["groupName"]; !ok {
		return nil, fmt.Errorf("the segment 'groupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseWorkspacePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a WorkspacePrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParseWorkspacePrivateLinkResourceIDInsensitively(input string) (*WorkspacePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.GroupName, ok = parsed.Parsed["groupName"]; !ok {
		return nil, fmt.Errorf("the segment 'groupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateWorkspacePrivateLinkResourceID checks that 'input' can be parsed as a Workspace Private Link Resource ID
func ValidateWorkspacePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspacePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HealthcareApis/workspaces/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.GroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHealthcareApis", "Microsoft.HealthcareApis", "Microsoft.HealthcareApis"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("groupName", "groupValue"),
	}
}

// String returns a human-readable description of this Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Group Name: %q", id.GroupName),
	}
	return fmt.Sprintf("Workspace Private Link Resource (%s)", strings.Join(components, "\n"))
}