// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Authorizer.
func (c *Client) UpdateAuthorizer(ctx context.Context, params *UpdateAuthorizerInput, optFns ...func(*Options)) (*UpdateAuthorizerOutput, error) {
	if params == nil {
		params = &UpdateAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAuthorizer", params, optFns, addOperationUpdateAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an Authorizer.
type UpdateAuthorizerInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The authorizer identifier.
	//
	// This member is required.
	AuthorizerId *string

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, don't specify this parameter.
	AuthorizerCredentialsArn *string

	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Required for HTTP API Lambda authorizers. Supported values are 1.0 and 2.0. To
	// learn more, see Working with AWS Lambda authorizers for HTTP APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	AuthorizerPayloadFormatVersion *string

	// The time to live (TTL) for cached authorizer results, in seconds. If it equals
	// 0, authorization caching is disabled. If it is greater than 0, API Gateway
	// caches authorizer responses. The maximum value is 3600, or 1 hour. Supported
	// only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds int32

	// The authorizer type. Specify REQUEST for a Lambda function using incoming
	// request parameters. Specify JWT to use JSON Web Tokens (supported only for HTTP
	// APIs).
	AuthorizerType types.AuthorizerType

	// The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers,
	// this must be a well-formed Lambda function URI, for example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form:
	// arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST
	// authorizers.
	AuthorizerUri *string

	// Specifies whether a Lambda authorizer returns a response in a simple format. By
	// default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda
	// authorizer can return a boolean value instead of an IAM policy. Supported only
	// for HTTP APIs. To learn more, see Working with AWS Lambda authorizers for HTTP
	// APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html)
	EnableSimpleResponses bool

	// The identity source for which authorization is requested. For a REQUEST
	// authorizer, this is optional. The value is a set of one or more mapping
	// expressions of the specified request parameters. The identity source can be
	// headers, query string parameters, stage variables, and context parameters. For
	// example, if an Auth header and a Name query string parameter are defined as
	// identity sources, this value is route.request.header.Auth,
	// route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection
	// expressions prefixed with $, for example, $request.header.Auth,
	// $request.querystring.Name. These parameters are used to perform runtime
	// validation for Lambda-based authorizers by verifying all of the identity-related
	// request parameters are present in the request, not null, and non-empty. Only
	// when this is true does the authorizer invoke the authorizer Lambda function.
	// Otherwise, it returns a 401 Unauthorized response without calling the Lambda
	// function. For HTTP APIs, identity sources are also used as the cache key when
	// caching is enabled. To learn more, see Working with AWS Lambda authorizers for
	// HTTP APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	// For JWT, a single entry that specifies where to extract the JSON Web Token (JWT)
	// from inbound requests. Currently only header-based and query parameter-based
	// selections are supported, for example $request.header.Authorization.
	IdentitySource []string

	// This parameter is not used.
	IdentityValidationExpression *string

	// Represents the configuration of a JWT authorizer. Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *types.JWTConfiguration

	// The name of the authorizer.
	Name *string
}

type UpdateAuthorizerOutput struct {

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, don't specify this parameter. Supported only for REQUEST authorizers.
	AuthorizerCredentialsArn *string

	// The authorizer identifier.
	AuthorizerId *string

	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Required for HTTP API Lambda authorizers. Supported values are 1.0 and 2.0. To
	// learn more, see Working with AWS Lambda authorizers for HTTP APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	AuthorizerPayloadFormatVersion *string

	// The time to live (TTL) for cached authorizer results, in seconds. If it equals
	// 0, authorization caching is disabled. If it is greater than 0, API Gateway
	// caches authorizer responses. The maximum value is 3600, or 1 hour. Supported
	// only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds int32

	// The authorizer type. Specify REQUEST for a Lambda function using incoming
	// request parameters. Specify JWT to use JSON Web Tokens (supported only for HTTP
	// APIs).
	AuthorizerType types.AuthorizerType

	// The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers,
	// this must be a well-formed Lambda function URI, for example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form:
	// arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST
	// authorizers.
	AuthorizerUri *string

	// Specifies whether a Lambda authorizer returns a response in a simple format. If
	// enabled, the Lambda authorizer can return a boolean value instead of an IAM
	// policy. Supported only for HTTP APIs. To learn more, see Working with AWS Lambda
	// authorizers for HTTP APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html)
	EnableSimpleResponses bool

	// The identity source for which authorization is requested. For a REQUEST
	// authorizer, this is optional. The value is a set of one or more mapping
	// expressions of the specified request parameters. The identity source can be
	// headers, query string parameters, stage variables, and context parameters. For
	// example, if an Auth header and a Name query string parameter are defined as
	// identity sources, this value is route.request.header.Auth,
	// route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection
	// expressions prefixed with $, for example, $request.header.Auth,
	// $request.querystring.Name. These parameters are used to perform runtime
	// validation for Lambda-based authorizers by verifying all of the identity-related
	// request parameters are present in the request, not null, and non-empty. Only
	// when this is true does the authorizer invoke the authorizer Lambda function.
	// Otherwise, it returns a 401 Unauthorized response without calling the Lambda
	// function. For HTTP APIs, identity sources are also used as the cache key when
	// caching is enabled. To learn more, see Working with AWS Lambda authorizers for
	// HTTP APIs
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	// For JWT, a single entry that specifies where to extract the JSON Web Token (JWT)
	// from inbound requests. Currently only header-based and query parameter-based
	// selections are supported, for example $request.header.Authorization.
	IdentitySource []string

	// The validation expression does not apply to the REQUEST authorizer.
	IdentityValidationExpression *string

	// Represents the configuration of a JWT authorizer. Required for the JWT
	// authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *types.JWTConfiguration

	// The name of the authorizer.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAuthorizer{}, middleware.After)
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
	if err = addOpUpdateAuthorizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuthorizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAuthorizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateAuthorizer",
	}
}