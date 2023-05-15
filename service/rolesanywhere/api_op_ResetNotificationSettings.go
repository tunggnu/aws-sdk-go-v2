// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the custom notification setting to IAM Roles Anywhere default setting.
// Required permissions: rolesanywhere:ResetNotificationSettings .
func (c *Client) ResetNotificationSettings(ctx context.Context, params *ResetNotificationSettingsInput, optFns ...func(*Options)) (*ResetNotificationSettingsOutput, error) {
	if params == nil {
		params = &ResetNotificationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetNotificationSettings", params, optFns, c.addOperationResetNotificationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetNotificationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetNotificationSettingsInput struct {

	// A list of notification setting keys to reset. A notification setting key
	// includes the event and the channel.
	//
	// This member is required.
	NotificationSettingKeys []types.NotificationSettingKey

	// The unique identifier of the trust anchor.
	//
	// This member is required.
	TrustAnchorId *string

	noSmithyDocumentSerde
}

type ResetNotificationSettingsOutput struct {

	// The state of the trust anchor after a read or write operation.
	//
	// This member is required.
	TrustAnchor *types.TrustAnchorDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetNotificationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetNotificationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetNotificationSettings{}, middleware.After)
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
	if err = addOpResetNotificationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetNotificationSettings(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opResetNotificationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "ResetNotificationSettings",
	}
}
