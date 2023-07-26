// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified endpoint. For a MySQL source or target endpoint, don't
// explicitly specify the database using the DatabaseName request parameter on the
// ModifyEndpoint API call. Specifying DatabaseName when you modify a MySQL
// endpoint replicates all the task tables to this single database. For MySQL
// endpoints, you specify the database only when you specify the schema in the
// table-mapping rules of the DMS task.
func (c *Client) ModifyEndpoint(ctx context.Context, params *ModifyEndpointInput, optFns ...func(*Options)) (*ModifyEndpointOutput, error) {
	if params == nil {
		params = &ModifyEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyEndpoint", params, optFns, c.addOperationModifyEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyEndpointInput struct {

	// The Amazon Resource Name (ARN) string that uniquely identifies the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// The Amazon Resource Name (ARN) of the certificate used for SSL connection.
	CertificateArn *string

	// The name of the endpoint database. For a MySQL source or target endpoint, do
	// not specify DatabaseName.
	DatabaseName *string

	// The settings in JSON format for the DMS transfer type of source endpoint.
	// Attributes include the following:
	//   - serviceAccessRoleArn - The Amazon Resource Name (ARN) used by the service
	//   access IAM role. The role must allow the iam:PassRole action.
	//   - BucketName - The name of the S3 bucket to use.
	// Shorthand syntax for these settings is as follows: ServiceAccessRoleArn=string
	// ,BucketName=string JSON syntax for these settings is as follows: {
	// "ServiceAccessRoleArn": "string", "BucketName": "string"}
	DmsTransferSettings *types.DmsTransferSettings

	// Settings in JSON format for the source DocumentDB endpoint. For more
	// information about the available settings, see the configuration properties
	// section in Using DocumentDB as a Target for Database Migration Service  (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html)
	// in the Database Migration Service User Guide.
	DocDbSettings *types.DocDbSettings

	// Settings in JSON format for the target Amazon DynamoDB endpoint. For
	// information about other available settings, see Using Object Mapping to Migrate
	// Data to DynamoDB (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping)
	// in the Database Migration Service User Guide.
	DynamoDbSettings *types.DynamoDbSettings

	// Settings in JSON format for the target OpenSearch endpoint. For more
	// information about the available settings, see Extra Connection Attributes When
	// Using OpenSearch as a Target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration)
	// in the Database Migration Service User Guide.
	ElasticsearchSettings *types.ElasticsearchSettings

	// The database endpoint identifier. Identifiers must begin with a letter and must
	// contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or
	// contain two consecutive hyphens.
	EndpointIdentifier *string

	// The type of endpoint. Valid values are source and target .
	EndpointType types.ReplicationEndpointTypeValue

	// The database engine name. Valid values, depending on the EndpointType, include
	// "mysql" , "oracle" , "postgres" , "mariadb" , "aurora" , "aurora-postgresql" ,
	// "redshift" , "s3" , "db2" , "db2-zos" , "azuredb" , "sybase" , "dynamodb" ,
	// "mongodb" , "kinesis" , "kafka" , "elasticsearch" , "documentdb" , "sqlserver" ,
	// "neptune" , and "babelfish" .
	EngineName *string

	// If this attribute is Y, the current call to ModifyEndpoint replaces all
	// existing endpoint settings with the exact settings that you specify in this
	// call. If this attribute is N, the current call to ModifyEndpoint does two
	// things:
	//   - It replaces any endpoint settings that already exist with new values, for
	//   settings with the same names.
	//   - It creates new endpoint settings that you specify in the call, for settings
	//   with different names.
	// For example, if you call create-endpoint ... --endpoint-settings '{"a":1}' ... ,
	// the endpoint has the following endpoint settings: '{"a":1}' . If you then call
	// modify-endpoint ... --endpoint-settings '{"b":2}' ... for the same endpoint, the
	// endpoint has the following settings: '{"a":1,"b":2}' . However, suppose that you
	// follow this with a call to modify-endpoint ... --endpoint-settings '{"b":2}'
	// --exact-settings ... for that same endpoint again. Then the endpoint has the
	// following settings: '{"b":2}' . All existing settings are replaced with the
	// exact settings that you specify.
	ExactSettings *bool

	// The external table definition.
	ExternalTableDefinition *string

	// Additional attributes associated with the connection. To reset this parameter,
	// pass the empty string ("") as an argument.
	ExtraConnectionAttributes *string

	// Settings in JSON format for the source GCP MySQL endpoint.
	GcpMySQLSettings *types.GcpMySQLSettings

	// Settings in JSON format for the source IBM Db2 LUW endpoint. For information
	// about other available settings, see Extra connection attributes when using Db2
	// LUW as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	IBMDb2Settings *types.IBMDb2Settings

	// Settings in JSON format for the target Apache Kafka endpoint. For more
	// information about the available settings, see Using object mapping to migrate
	// data to a Kafka topic (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping)
	// in the Database Migration Service User Guide.
	KafkaSettings *types.KafkaSettings

	// Settings in JSON format for the target endpoint for Amazon Kinesis Data
	// Streams. For more information about the available settings, see Using object
	// mapping to migrate data to a Kinesis data stream (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping)
	// in the Database Migration Service User Guide.
	KinesisSettings *types.KinesisSettings

	// Settings in JSON format for the source and target Microsoft SQL Server
	// endpoint. For information about other available settings, see Extra connection
	// attributes when using SQL Server as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib)
	// and Extra connection attributes when using SQL Server as a target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	MicrosoftSQLServerSettings *types.MicrosoftSQLServerSettings

	// Settings in JSON format for the source MongoDB endpoint. For more information
	// about the available settings, see the configuration properties section in
	// Endpoint configuration settings when using MongoDB as a source for Database
	// Migration Service (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration)
	// in the Database Migration Service User Guide.
	MongoDbSettings *types.MongoDbSettings

	// Settings in JSON format for the source and target MySQL endpoint. For
	// information about other available settings, see Extra connection attributes
	// when using MySQL as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib)
	// and Extra connection attributes when using a MySQL-compatible database as a
	// target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	MySQLSettings *types.MySQLSettings

	// Settings in JSON format for the target Amazon Neptune endpoint. For more
	// information about the available settings, see Specifying graph-mapping rules
	// using Gremlin and R2RML for Amazon Neptune as a target (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings)
	// in the Database Migration Service User Guide.
	NeptuneSettings *types.NeptuneSettings

	// Settings in JSON format for the source and target Oracle endpoint. For
	// information about other available settings, see Extra connection attributes
	// when using Oracle as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib)
	// and Extra connection attributes when using Oracle as a target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	OracleSettings *types.OracleSettings

	// The password to be used to login to the endpoint database.
	Password *string

	// The port used by the endpoint database.
	Port *int32

	// Settings in JSON format for the source and target PostgreSQL endpoint. For
	// information about other available settings, see Extra connection attributes
	// when using PostgreSQL as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib)
	// and Extra connection attributes when using PostgreSQL as a target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	PostgreSQLSettings *types.PostgreSQLSettings

	// Settings in JSON format for the Redis target endpoint.
	RedisSettings *types.RedisSettings

	// Provides information that defines an Amazon Redshift endpoint.
	RedshiftSettings *types.RedshiftSettings

	// Settings in JSON format for the target Amazon S3 endpoint. For more information
	// about the available settings, see Extra Connection Attributes When Using Amazon
	// S3 as a Target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring)
	// in the Database Migration Service User Guide.
	S3Settings *types.S3Settings

	// The name of the server where the endpoint database resides.
	ServerName *string

	// The Amazon Resource Name (ARN) for the IAM role you want to use to modify the
	// endpoint. The role must allow the iam:PassRole action.
	ServiceAccessRoleArn *string

	// The SSL mode used to connect to the endpoint. The default value is none .
	SslMode types.DmsSslModeValue

	// Settings in JSON format for the source and target SAP ASE endpoint. For
	// information about other available settings, see Extra connection attributes
	// when using SAP ASE as a source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib)
	// and Extra connection attributes when using SAP ASE as a target for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib)
	// in the Database Migration Service User Guide.
	SybaseSettings *types.SybaseSettings

	// Settings in JSON format for the target Amazon Timestream endpoint.
	TimestreamSettings *types.TimestreamSettings

	// The user name to be used to login to the endpoint database.
	Username *string

	noSmithyDocumentSerde
}

type ModifyEndpointOutput struct {

	// The modified endpoint.
	Endpoint *types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyEndpoint{}, middleware.After)
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
	if err = addModifyEndpointResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyEndpoint",
	}
}

type opModifyEndpointResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opModifyEndpointResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opModifyEndpointResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "dms"
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
				signingName = "dms"
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
				v4aScheme.SigningName = aws.String("dms")
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

func addModifyEndpointResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opModifyEndpointResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
