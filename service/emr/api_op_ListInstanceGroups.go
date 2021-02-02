// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides all available details about the instance groups in a cluster.
func (c *Client) ListInstanceGroups(ctx context.Context, params *ListInstanceGroupsInput, optFns ...func(*Options)) (*ListInstanceGroupsOutput, error) {
	if params == nil {
		params = &ListInstanceGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceGroups", params, optFns, addOperationListInstanceGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which instance groups to retrieve.
type ListInstanceGroupsInput struct {

	// The identifier of the cluster for which to list the instance groups.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
}

// This input determines which instance groups to retrieve.
type ListInstanceGroupsOutput struct {

	// The list of instance groups for the cluster and given filters.
	InstanceGroups []types.InstanceGroup

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstanceGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstanceGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstanceGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListInstanceGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceGroups(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListInstanceGroupsAPIClient is a client that implements the ListInstanceGroups
// operation.
type ListInstanceGroupsAPIClient interface {
	ListInstanceGroups(context.Context, *ListInstanceGroupsInput, ...func(*Options)) (*ListInstanceGroupsOutput, error)
}

var _ ListInstanceGroupsAPIClient = (*Client)(nil)

// ListInstanceGroupsPaginatorOptions is the paginator options for
// ListInstanceGroups
type ListInstanceGroupsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceGroupsPaginator is a paginator for ListInstanceGroups
type ListInstanceGroupsPaginator struct {
	options   ListInstanceGroupsPaginatorOptions
	client    ListInstanceGroupsAPIClient
	params    *ListInstanceGroupsInput
	nextToken *string
	firstPage bool
}

// NewListInstanceGroupsPaginator returns a new ListInstanceGroupsPaginator
func NewListInstanceGroupsPaginator(client ListInstanceGroupsAPIClient, params *ListInstanceGroupsInput, optFns ...func(*ListInstanceGroupsPaginatorOptions)) *ListInstanceGroupsPaginator {
	options := ListInstanceGroupsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListInstanceGroupsInput{}
	}

	return &ListInstanceGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListInstanceGroups page.
func (p *ListInstanceGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListInstanceGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInstanceGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstanceGroups",
	}
}