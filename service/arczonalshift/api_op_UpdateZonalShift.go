// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update an active zonal shift in Amazon Route 53 Application Recovery Controller
// in your Amazon Web Services account. You can update a zonal shift to set a new
// expiration, or edit or replace the comment for the zonal shift.
func (c *Client) UpdateZonalShift(ctx context.Context, params *UpdateZonalShiftInput, optFns ...func(*Options)) (*UpdateZonalShiftOutput, error) {
	if params == nil {
		params = &UpdateZonalShiftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateZonalShift", params, optFns, c.addOperationUpdateZonalShiftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateZonalShiftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateZonalShiftInput struct {

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. A new comment overwrites any
	// existing comment string.
	Comment *string

	// The length of time that you want a zonal shift to be active, which Route 53 ARC
	// converts to an expiry time (expiration time). Zonal shifts are temporary. You
	// can set a zonal shift to be active initially for up to three days (72 hours). If
	// you want to still keep traffic away from an Availability Zone, you can update
	// the zonal shift and set a new expiration. You can also cancel a zonal shift,
	// before it expires, for example, if you're ready to restore traffic to the
	// Availability Zone. To set a length of time for a zonal shift to be active,
	// specify a whole number, and then one of the following, with no space:
	//   - A lowercase letter m: To specify that the value is in minutes.
	//   - A lowercase letter h: To specify that the value is in hours.
	// For example: 20h means the zonal shift expires in 20 hours. 120m means the
	// zonal shift expires in 120 minutes (2 hours).
	ExpiresIn *string

	noSmithyDocumentSerde
}

type UpdateZonalShiftOutput struct {

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the Amazon Web
	// Services Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. A new comment overwrites any
	// existing comment string.
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for a customer-started zonal shift. A zonal
	// shift is temporary and must be set to expire when you start the zonal shift. You
	// can initially set a zonal shift to expire in a maximum of three days (72 hours).
	// However, you can update a zonal shift to set a new expiration at any time. When
	// you start a zonal shift, you specify how long you want it to be active, which
	// Route 53 ARC converts to an expiry time (expiration time). You can cancel a
	// zonal shift when you're ready to restore traffic to the Availability Zone, or
	// just wait for it to expire. Or you can update the zonal shift to specify another
	// length of time to expire in.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource to shift away traffic for. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, supported
	// resources are Network Load Balancers and Application Load Balancers with
	// cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift starts.
	//
	// This member is required.
	StartTime *time.Time

	// A status for a zonal shift. The Status for a zonal shift can have one of the
	// following values:
	//   - ACTIVE: The zonal shift has been started and active.
	//   - EXPIRED: The zonal shift has expired (the expiry time was exceeded).
	//   - CANCELED: The zonal shift was canceled.
	//
	// This member is required.
	Status types.ZonalShiftStatus

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateZonalShiftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateZonalShift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateZonalShift{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateZonalShift"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateZonalShiftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateZonalShift(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateZonalShift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateZonalShift",
	}
}
