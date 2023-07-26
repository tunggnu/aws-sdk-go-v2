// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified user pool with the specified attributes. You can get a
// list of the current user pool settings using DescribeUserPool (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html)
// . If you don't provide a value for an attribute, it will be set to the default
// value. This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with Amazon Pinpoint (https://console.aws.amazon.com/pinpoint/home/)
// . Amazon Cognito uses the registered number automatically. Otherwise, Amazon
// Cognito users who must receive SMS messages might not be able to sign up,
// activate their accounts, or sign in. If you have never used SMS text messages
// with Amazon Cognito or any other Amazon Web Service, Amazon Simple Notification
// Service might place your account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-sms-userpool-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) UpdateUserPool(ctx context.Context, params *UpdateUserPoolInput, optFns ...func(*Options)) (*UpdateUserPoolOutput, error) {
	if params == nil {
		params = &UpdateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserPool", params, optFns, c.addOperationUpdateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user pool.
type UpdateUserPoolInput struct {

	// The user pool ID for the user pool you want to update.
	//
	// This member is required.
	UserPoolId *string

	// The available verified method a user can use to recover their password when
	// they call ForgotPassword . You can use this setting to define a preferred method
	// when a user has more than one method available. With this setting, SMS doesn't
	// qualify for a valid password recovery mechanism if the user also has SMS
	// multi-factor authentication (MFA) activated. In the absence of this setting,
	// Amazon Cognito uses the legacy behavior to determine the recovery method where
	// SMS is preferred through email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// The attributes that are automatically verified when Amazon Cognito requests to
	// update user pools.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// When active, DeletionProtection prevents accidental deletion of your user pool.
	// Before you can delete a user pool that you have protected against deletion, you
	// must deactivate this feature. When you try to delete a protected user pool in a
	// DeleteUserPool API request, Amazon Cognito returns an InvalidParameterException
	// error. To delete a protected user pool, send a new DeleteUserPool request after
	// you deactivate deletion protection in an UpdateUserPool API request.
	DeletionProtection types.DeletionProtectionType

	// The device-remembering configuration for a user pool. A null value indicates
	// that you have deactivated device remembering in your user pool. When you provide
	// a value for any DeviceConfiguration field, you activate the Amazon Cognito
	// device-remembering feature.
	DeviceConfiguration *types.DeviceConfigurationType

	// The email configuration of your user pool. The email configuration type sets
	// your preferred sending method, Amazon Web Services Region, and sender for email
	// invitation and verification messages from your user pool.
	EmailConfiguration *types.EmailConfigurationType

	// This parameter is no longer used. See VerificationMessageTemplateType (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html)
	// .
	EmailVerificationMessage *string

	// This parameter is no longer used. See VerificationMessageTemplateType (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html)
	// .
	EmailVerificationSubject *string

	// The Lambda configuration information from the request to update the user pool.
	LambdaConfig *types.LambdaConfigType

	// Possible values include:
	//   - OFF - MFA tokens aren't required and can't be specified during user
	//   registration.
	//   - ON - MFA tokens are required for all user registrations. You can only
	//   specify ON when you're initially creating a user pool. You can use the
	//   SetUserPoolMfaConfig (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserPoolMfaConfig.html)
	//   API operation to turn MFA "ON" for existing user pools.
	//   - OPTIONAL - Users have the option when registering to create an MFA token.
	MfaConfiguration types.UserPoolMfaType

	// A container with the policies you want to update in a user pool.
	Policies *types.UserPoolPolicyType

	// The contents of the SMS authentication message.
	SmsAuthenticationMessage *string

	// The SMS configuration with the settings that your Amazon Cognito user pool must
	// use to send an SMS message from your Amazon Web Services account through Amazon
	// Simple Notification Service. To send SMS messages with Amazon SNS in the Amazon
	// Web Services Region that you want, the Amazon Cognito user pool uses an Identity
	// and Access Management (IAM) role in your Amazon Web Services account.
	SmsConfiguration *types.SmsConfigurationType

	// This parameter is no longer used. See VerificationMessageTemplateType (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html)
	// .
	SmsVerificationMessage *string

	// The settings for updates to user attributes. These settings include the
	// property AttributesRequireVerificationBeforeUpdate , a user-pool setting that
	// tells Amazon Cognito how to handle changes to the value of your users' email
	// address and phone number attributes. For more information, see Verifying
	// updates to email addresses and phone numbers (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates)
	// .
	UserAttributeUpdateSettings *types.UserAttributeUpdateSettingsType

	// Enables advanced security risk detection. Set the key AdvancedSecurityMode to
	// the value "AUDIT".
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string

	// The template for verification messages.
	VerificationMessageTemplate *types.VerificationMessageTemplateType

	noSmithyDocumentSerde
}

// Represents the response from the server when you make a request to update the
// user pool.
type UpdateUserPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPool{}, middleware.After)
	if err != nil {
		return err
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addUpdateUserPoolResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPool(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateUserPool",
	}
}

type opUpdateUserPoolResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateUserPoolResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateUserPoolResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "cognito-idp"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "cognito-idp"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("cognito-idp")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateUserPoolResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateUserPoolResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
