// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This section contains the Amazon Managed Workflows for Apache Airflow (MWAA)
// API reference documentation to create an environment. For more information, see
// Get started with Amazon Managed Workflows for Apache Airflow (https://docs.aws.amazon.com/mwaa/latest/userguide/get-started.html)
// .
type CreateEnvironmentInput struct {

	// The relative path to the DAGs folder on your Amazon S3 bucket. For example, dags
	// . For more information, see Adding or updating DAGs (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html)
	// .
	//
	// This member is required.
	DagS3Path *string

	// The Amazon Resource Name (ARN) of the execution role for your environment. An
	// execution role is an Amazon Web Services Identity and Access Management (IAM)
	// role that grants MWAA permission to access Amazon Web Services services and
	// resources used by your environment. For example,
	// arn:aws:iam::123456789:role/my-execution-role . For more information, see
	// Amazon MWAA Execution role (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html)
	// .
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment .
	//
	// This member is required.
	Name *string

	// The VPC networking components used to secure and enable network traffic between
	// the Amazon Web Services resources for your environment. For more information,
	// see About networking on Amazon MWAA (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html)
	// .
	//
	// This member is required.
	NetworkConfiguration *types.NetworkConfiguration

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and
	// supporting files are stored. For example,
	// arn:aws:s3:::my-airflow-bucket-unique-name . For more information, see Create
	// an Amazon S3 bucket for Amazon MWAA (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html)
	// .
	//
	// This member is required.
	SourceBucketArn *string

	// A list of key-value pairs containing the Apache Airflow configuration options
	// you want to attach to your environment. For more information, see Apache
	// Airflow configuration options (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html)
	// .
	AirflowConfigurationOptions map[string]string

	// The Apache Airflow version for your environment. If no value is specified, it
	// defaults to the latest version. Valid values: 1.10.12 , 2.0.2 , 2.2.2 , 2.4.3 ,
	// and 2.5.1 . For more information, see Apache Airflow versions on Amazon Managed
	// Workflows for Apache Airflow (MWAA) (https://docs.aws.amazon.com/mwaa/latest/userguide/airflow-versions.html)
	// .
	AirflowVersion *string

	// The environment class type. Valid values: mw1.small , mw1.medium , mw1.large .
	// For more information, see Amazon MWAA environment class (https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html)
	// .
	EnvironmentClass *string

	// The Amazon Web Services Key Management Service (KMS) key to encrypt the data in
	// your environment. You can use an Amazon Web Services owned CMK, or a Customer
	// managed CMK (advanced). For more information, see Create an Amazon MWAA
	// environment (https://docs.aws.amazon.com/mwaa/latest/userguide/create-environment.html)
	// .
	KmsKey *string

	// Defines the Apache Airflow logs to send to CloudWatch Logs.
	LoggingConfiguration *types.LoggingConfigurationInput

	// The maximum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. For example, 20 . When there are no more tasks running, and no
	// more in the queue, MWAA disposes of the extra workers leaving the one worker
	// that is included with your environment, or the number you specify in MinWorkers .
	MaxWorkers *int32

	// The minimum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. When there are no more tasks running, and no more in the
	// queue, MWAA disposes of the extra workers leaving the worker count you specify
	// in the MinWorkers field. For example, 2 .
	MinWorkers *int32

	// The version of the plugins.zip file on your Amazon S3 bucket. You must specify
	// a version each time a plugins.zip file is updated. For more information, see
	// How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// .
	PluginsS3ObjectVersion *string

	// The relative path to the plugins.zip file on your Amazon S3 bucket. For
	// example, plugins.zip . If specified, then the plugins.zip version is required.
	// For more information, see Installing custom plugins (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html)
	// .
	PluginsS3Path *string

	// The version of the requirements.txt file on your Amazon S3 bucket. You must
	// specify a version each time a requirements.txt file is updated. For more
	// information, see How S3 Versioning works (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// .
	RequirementsS3ObjectVersion *string

	// The relative path to the requirements.txt file on your Amazon S3 bucket. For
	// example, requirements.txt . If specified, then a version is required. For more
	// information, see Installing Python dependencies (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html)
	// .
	RequirementsS3Path *string

	// The number of Apache Airflow schedulers to run in your environment. Valid
	// values:
	//   - v2 - Accepts between 2 to 5. Defaults to 2.
	//   - v1 - Accepts 1.
	Schedulers *int32

	// The version of the startup shell script in your Amazon S3 bucket. You must
	// specify the version ID (https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html)
	// that Amazon S3 assigns to the file every time you update the script. Version IDs
	// are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than
	// 1,024 bytes long. The following is an example:
	// 3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo For more
	// information, see Using a startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html)
	// .
	StartupScriptS3ObjectVersion *string

	// The relative path to the startup shell script in your Amazon S3 bucket. For
	// example, s3://mwaa-environment/startup.sh . Amazon MWAA runs the script as your
	// environment starts, and before running the Apache Airflow process. You can use
	// this script to install dependencies, modify Apache Airflow configuration
	// options, and set environment variables. For more information, see Using a
	// startup script (https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html)
	// .
	StartupScriptS3Path *string

	// The key-value tag pairs you want to associate to your environment. For example,
	// "Environment": "Staging" . For more information, see Tagging Amazon Web
	// Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags map[string]string

	// The Apache Airflow Web server access mode. For more information, see Apache
	// Airflow access modes (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html)
	// .
	WebserverAccessMode types.WebserverAccessMode

	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour
	// standard time to start weekly maintenance updates of your environment in the
	// following format: DAY:HH:MM . For example: TUE:03:30 . You can specify a start
	// time in 30 minute increments only.
	WeeklyMaintenanceWindowStart *string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) returned in the response for the environment.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateEnvironmentMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateEnvironmentResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateEnvironmentMiddleware struct {
}

func (*endpointPrefix_opCreateEnvironmentMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateEnvironmentMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateEnvironmentMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateEnvironmentMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "CreateEnvironment",
	}
}

type opCreateEnvironmentResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateEnvironmentResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateEnvironmentResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "airflow"
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
				signingName = "airflow"
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
				v4aScheme.SigningName = aws.String("airflow")
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

func addCreateEnvironmentResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateEnvironmentResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
