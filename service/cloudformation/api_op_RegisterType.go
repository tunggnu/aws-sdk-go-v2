// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an extension with the CloudFormation service. Registering an
// extension makes it available for use in CloudFormation templates in your Amazon
// Web Services account, and includes:
//   - Validating the extension schema.
//   - Determining which handlers, if any, have been specified for the extension.
//   - Making the extension available for use in your account.
//
// For more information about how to develop extensions and ready them for
// registration, see Creating Resource Providers (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html)
// in the CloudFormation CLI User Guide. You can have a maximum of 50 resource
// extension versions registered at a time. This maximum is per account and per
// Region. Use DeregisterType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeregisterType.html)
// to deregister specific extension versions if necessary. Once you have initiated
// a registration request using RegisterType , you can use DescribeTypeRegistration
// to monitor the progress of the registration request. Once you have registered a
// private extension in your account and Region, use SetTypeConfiguration (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html)
// to specify configuration properties for the extension. For more information, see
// Configuring extensions at the account level (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration)
// in the CloudFormation User Guide.
func (c *Client) RegisterType(ctx context.Context, params *RegisterTypeInput, optFns ...func(*Options)) (*RegisterTypeOutput, error) {
	if params == nil {
		params = &RegisterTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterType", params, optFns, c.addOperationRegisterTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTypeInput struct {

	// A URL to the S3 bucket containing the extension project package that contains
	// the necessary files for the extension you want to register. For information
	// about generating a schema handler package for the extension you want to
	// register, see submit (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html)
	// in the CloudFormation CLI User Guide. The user registering the extension must be
	// able to access the package in the S3 bucket. That's, the user needs to have
	// GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
	// permissions for the schema handler package. For more information, see Actions,
	// Resources, and Condition Keys for Amazon S3 (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html)
	// in the Identity and Access Management User Guide.
	//
	// This member is required.
	SchemaHandlerPackage *string

	// The name of the extension being registered. We suggest that extension names
	// adhere to the following patterns:
	//   - For resource types, company_or_organization::service::type.
	//   - For modules, company_or_organization::service::type::MODULE.
	//   - For hooks, MyCompany::Testing::MyTestHook.
	// The following organization namespaces are reserved and can't be used in your
	// extension names:
	//   - Alexa
	//   - AMZN
	//   - Amazon
	//   - AWS
	//   - Custom
	//   - Dev
	//
	// This member is required.
	TypeName *string

	// A unique identifier that acts as an idempotency key for this registration
	// request. Specifying a client request token prevents CloudFormation from
	// generating more than one version of an extension from the same registration
	// request, even if the request is submitted multiple times.
	ClientRequestToken *string

	// The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume
	// when invoking the extension. For CloudFormation to assume the specified
	// execution role, the role must contain a trust relationship with the
	// CloudFormation service principle ( resources.cloudformation.amazonaws.com ). For
	// more information about adding trust relationships, see Modifying a role trust
	// policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-managingrole-editing-console.html#roles-managingrole_edit-trust-policy)
	// in the Identity and Access Management User Guide. If your extension calls Amazon
	// Web Services APIs in any of its handlers, you must create an IAM execution role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
	// that includes the necessary permissions to call those Amazon Web Services APIs,
	// and provision that execution role in your account. When CloudFormation needs to
	// invoke the resource type handler, CloudFormation assumes this execution role to
	// create a temporary session token, which it then passes to the resource type
	// handler, thereby supplying your resource type with the appropriate credentials.
	ExecutionRoleArn *string

	// Specifies logging configuration information for an extension.
	LoggingConfig *types.LoggingConfig

	// The kind of extension.
	Type types.RegistryType

	noSmithyDocumentSerde
}

type RegisterTypeOutput struct {

	// The identifier for this registration request. Use this registration token when
	// calling DescribeTypeRegistration , which returns information about the status
	// and IDs of the extension registration.
	RegistrationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRegisterType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterType{}, middleware.After)
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
	if err = addRegisterTypeResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "RegisterType",
	}
}

type opRegisterTypeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opRegisterTypeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRegisterTypeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cloudformation"
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
				signingName = "cloudformation"
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
				v4aScheme.SigningName = aws.String("cloudformation")
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

func addRegisterTypeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opRegisterTypeResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
