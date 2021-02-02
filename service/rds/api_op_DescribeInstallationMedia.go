// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the available installation media for a DB engine that requires an
// on-premises customer provided license, such as Microsoft SQL Server.
func (c *Client) DescribeInstallationMedia(ctx context.Context, params *DescribeInstallationMediaInput, optFns ...func(*Options)) (*DescribeInstallationMediaOutput, error) {
	if params == nil {
		params = &DescribeInstallationMediaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstallationMedia", params, optFns, addOperationDescribeInstallationMediaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstallationMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstallationMediaInput struct {

	// A filter that specifies one or more installation media to describe. Supported
	// filters include the following:
	//
	// * custom-availability-zone-id - Accepts custom
	// Availability Zone (AZ) identifiers. The results list includes information about
	// only the custom AZs identified by these identifiers.
	//
	// * engine - Accepts
	// database engines. The results list includes information about only the database
	// engines identified by these identifiers. For more information about the valid
	// engines for installation media, see ImportInstallationMedia.
	Filters []types.Filter

	// The installation medium ID.
	InstallationMediaId *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// An optional pagination token provided by a previous DescribeInstallationMedia
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	MaxRecords *int32
}

type DescribeInstallationMediaOutput struct {

	// The list of InstallationMedia objects for the AWS account.
	InstallationMedia []types.InstallationMedia

	// An optional pagination token provided by a previous DescribeInstallationMedia
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstallationMediaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstallationMedia{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstallationMedia{}, middleware.After)
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
	if err = addOpDescribeInstallationMediaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstallationMedia(options.Region), middleware.Before); err != nil {
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

// DescribeInstallationMediaAPIClient is a client that implements the
// DescribeInstallationMedia operation.
type DescribeInstallationMediaAPIClient interface {
	DescribeInstallationMedia(context.Context, *DescribeInstallationMediaInput, ...func(*Options)) (*DescribeInstallationMediaOutput, error)
}

var _ DescribeInstallationMediaAPIClient = (*Client)(nil)

// DescribeInstallationMediaPaginatorOptions is the paginator options for
// DescribeInstallationMedia
type DescribeInstallationMediaPaginatorOptions struct {
	// An optional pagination token provided by a previous DescribeInstallationMedia
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstallationMediaPaginator is a paginator for DescribeInstallationMedia
type DescribeInstallationMediaPaginator struct {
	options   DescribeInstallationMediaPaginatorOptions
	client    DescribeInstallationMediaAPIClient
	params    *DescribeInstallationMediaInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstallationMediaPaginator returns a new
// DescribeInstallationMediaPaginator
func NewDescribeInstallationMediaPaginator(client DescribeInstallationMediaAPIClient, params *DescribeInstallationMediaInput, optFns ...func(*DescribeInstallationMediaPaginatorOptions)) *DescribeInstallationMediaPaginator {
	options := DescribeInstallationMediaPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeInstallationMediaInput{}
	}

	return &DescribeInstallationMediaPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstallationMediaPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeInstallationMedia page.
func (p *DescribeInstallationMediaPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstallationMediaOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeInstallationMedia(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstallationMedia(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeInstallationMedia",
	}
}