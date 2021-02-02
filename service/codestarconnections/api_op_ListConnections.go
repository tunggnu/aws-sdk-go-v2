// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the connections associated with your account.
func (c *Client) ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) {
	if params == nil {
		params = &ListConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnections", params, optFns, addOperationListConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectionsInput struct {

	// Filters the list of connections to those associated with a specified host.
	HostArnFilter *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token that was returned from the previous ListConnections call, which can be
	// used to return the next set of connections in the list.
	NextToken *string

	// Filters the list of connections to those associated with a specified provider,
	// such as Bitbucket.
	ProviderTypeFilter types.ProviderType
}

type ListConnectionsOutput struct {

	// A list of connections and the details for each connection, such as status,
	// owner, and provider type.
	Connections []types.Connection

	// A token that can be used in the next ListConnections call. To view all items in
	// the list, continue to call this operation with each subsequent token until no
	// more nextToken values are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnections(options.Region), middleware.Before); err != nil {
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

// ListConnectionsAPIClient is a client that implements the ListConnections
// operation.
type ListConnectionsAPIClient interface {
	ListConnections(context.Context, *ListConnectionsInput, ...func(*Options)) (*ListConnectionsOutput, error)
}

var _ ListConnectionsAPIClient = (*Client)(nil)

// ListConnectionsPaginatorOptions is the paginator options for ListConnections
type ListConnectionsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectionsPaginator is a paginator for ListConnections
type ListConnectionsPaginator struct {
	options   ListConnectionsPaginatorOptions
	client    ListConnectionsAPIClient
	params    *ListConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListConnectionsPaginator returns a new ListConnectionsPaginator
func NewListConnectionsPaginator(client ListConnectionsAPIClient, params *ListConnectionsInput, optFns ...func(*ListConnectionsPaginatorOptions)) *ListConnectionsPaginator {
	options := ListConnectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListConnectionsInput{}
	}

	return &ListConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListConnections page.
func (p *ListConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListConnections(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "ListConnections",
	}
}