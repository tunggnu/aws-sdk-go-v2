// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the instance limits and current utilization for an Amazon Web
// Services Region or location. Instance limits control the number of instances,
// per instance type, per location, that your Amazon Web Services account can use.
// Learn more at Amazon EC2 Instance Types (http://aws.amazon.com/ec2/instance-types/)
// . The information returned includes the maximum number of instances allowed and
// your account's current usage across all fleets. This information can affect your
// ability to scale your Amazon GameLift fleets. You can request a limit increase
// for your account by using the Service limits page in the Amazon GameLift
// console. Instance limits differ based on whether the instances are deployed in a
// fleet's home Region or in a remote location. For remote locations, limits also
// differ based on the combination of home Region and remote location. All requests
// must specify an Amazon Web Services Region (either explicitly or as your default
// settings). To get the limit for a remote location, you must also specify the
// location. For example, the following requests all return different results:
//   - Request specifies the Region ap-northeast-1 with no location. The result is
//     limits and usage data on all instance types that are deployed in us-east-2 ,
//     by all of the fleets that reside in ap-northeast-1 .
//   - Request specifies the Region us-east-1 with location ca-central-1 . The
//     result is limits and usage data on all instance types that are deployed in
//     ca-central-1 , by all of the fleets that reside in us-east-2 . These limits do
//     not affect fleets in any other Regions that deploy instances to ca-central-1 .
//   - Request specifies the Region eu-west-1 with location ca-central-1 . The
//     result is limits and usage data on all instance types that are deployed in
//     ca-central-1 , by all of the fleets that reside in eu-west-1 .
//
// This operation can be used in the following ways:
//   - To get limit and usage data for all instance types that are deployed in an
//     Amazon Web Services Region by fleets that reside in the same Region: Specify the
//     Region only. Optionally, specify a single instance type to retrieve information
//     for.
//   - To get limit and usage data for all instance types that are deployed to a
//     remote location by fleets that reside in different Amazon Web Services Region:
//     Provide both the Amazon Web Services Region and the remote location. Optionally,
//     specify a single instance type to retrieve information for.
//
// If successful, an EC2InstanceLimits object is returned with limits and usage
// data for each requested instance type. Learn more Setting up Amazon GameLift
// fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
func (c *Client) DescribeEC2InstanceLimits(ctx context.Context, params *DescribeEC2InstanceLimitsInput, optFns ...func(*Options)) (*DescribeEC2InstanceLimitsOutput, error) {
	if params == nil {
		params = &DescribeEC2InstanceLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEC2InstanceLimits", params, optFns, c.addOperationDescribeEC2InstanceLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEC2InstanceLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEC2InstanceLimitsInput struct {

	// Name of an Amazon EC2 instance type that is supported in Amazon GameLift. A
	// fleet instance type determines the computing resources of each instance in the
	// fleet, including CPU, memory, storage, and networking capacity. Do not specify a
	// value for this parameter to retrieve limits for all instance types.
	EC2InstanceType types.EC2InstanceType

	// The name of a remote location to request instance limits for, in the form of an
	// Amazon Web Services Region code such as us-west-2 .
	Location *string

	noSmithyDocumentSerde
}

type DescribeEC2InstanceLimitsOutput struct {

	// The maximum number of instances for the specified instance type.
	EC2InstanceLimits []types.EC2InstanceLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEC2InstanceLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEC2InstanceLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEC2InstanceLimits{}, middleware.After)
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
	if err = addDescribeEC2InstanceLimitsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEC2InstanceLimits(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEC2InstanceLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeEC2InstanceLimits",
	}
}

type opDescribeEC2InstanceLimitsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeEC2InstanceLimitsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeEC2InstanceLimitsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "gamelift"
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
				signingName = "gamelift"
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
				v4aScheme.SigningName = aws.String("gamelift")
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

func addDescribeEC2InstanceLimitsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeEC2InstanceLimitsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
