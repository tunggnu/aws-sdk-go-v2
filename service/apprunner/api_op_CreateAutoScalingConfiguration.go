// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an App Runner automatic scaling configuration resource. App Runner
// requires this resource when you create or update App Runner services and you
// require non-default auto scaling settings. You can share an auto scaling
// configuration across multiple services. Create multiple revisions of a
// configuration by calling this action multiple times using the same
// AutoScalingConfigurationName . The call returns incremental
// AutoScalingConfigurationRevision values. When you create a service and configure
// an auto scaling configuration resource, the service uses the latest active
// revision of the auto scaling configuration by default. You can optionally
// configure the service to use a specific revision. Configure a higher MinSize to
// increase the spread of your App Runner service over more Availability Zones in
// the Amazon Web Services Region. The tradeoff is a higher minimal cost. Configure
// a lower MaxSize to control your cost. The tradeoff is lower responsiveness
// during peak demand.
func (c *Client) CreateAutoScalingConfiguration(ctx context.Context, params *CreateAutoScalingConfigurationInput, optFns ...func(*Options)) (*CreateAutoScalingConfigurationOutput, error) {
	if params == nil {
		params = &CreateAutoScalingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoScalingConfiguration", params, optFns, c.addOperationCreateAutoScalingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoScalingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoScalingConfigurationInput struct {

	// A name for the auto scaling configuration. When you use it for the first time
	// in an Amazon Web Services Region, App Runner creates revision number 1 of this
	// name. When you use the same name in subsequent calls, App Runner creates
	// incremental revisions of the configuration. The name DefaultConfiguration is
	// reserved (it's the configuration that App Runner uses if you don't provide a
	// custome one). You can't use it to create a new auto scaling configuration, and
	// you can't create a revision of it. When you want to use your own auto scaling
	// configuration for your App Runner service, create a configuration with a
	// different name, and then provide it when you create or update your service.
	//
	// This member is required.
	AutoScalingConfigurationName *string

	// The maximum number of concurrent requests that you want an instance to process.
	// If the number of concurrent requests exceeds this limit, App Runner scales up
	// your service. Default: 100
	MaxConcurrency *int32

	// The maximum number of instances that your service scales up to. At most MaxSize
	// instances actively serve traffic for your service. Default: 25
	MaxSize *int32

	// The minimum number of instances that App Runner provisions for your service.
	// The service always has at least MinSize provisioned instances. Some of them
	// actively serve traffic. The rest of them (provisioned and inactive instances)
	// are a cost-effective compute capacity reserve and are ready to be quickly
	// activated. You pay for memory usage of all the provisioned instances. You pay
	// for CPU usage of only the active subset. App Runner temporarily doubles the
	// number of provisioned instances during deployments, to maintain the same
	// capacity for both old and new code. Default: 1
	MinSize *int32

	// A list of metadata items that you can associate with your auto scaling
	// configuration resource. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAutoScalingConfigurationOutput struct {

	// A description of the App Runner auto scaling configuration that's created by
	// this request.
	//
	// This member is required.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoScalingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateAutoScalingConfiguration{}, middleware.After)
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
	if err = addCreateAutoScalingConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAutoScalingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoScalingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAutoScalingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "CreateAutoScalingConfiguration",
	}
}

type opCreateAutoScalingConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateAutoScalingConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateAutoScalingConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "apprunner"
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
				signingName = "apprunner"
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
				v4aScheme.SigningName = aws.String("apprunner")
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

func addCreateAutoScalingConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateAutoScalingConfigurationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
