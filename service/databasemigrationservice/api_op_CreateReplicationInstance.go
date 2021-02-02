// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the replication instance using the specified parameters. AWS DMS
// requires that your account have certain roles with appropriate permissions
// before you can create a replication instance. For information on the required
// roles, see Creating the IAM Roles to Use With the AWS CLI and AWS DMS API
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.APIRole).
// For information on the required permissions, see IAM Permissions Needed to Use
// AWS DMS
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#CHAP_Security.IAMPermissions).
func (c *Client) CreateReplicationInstance(ctx context.Context, params *CreateReplicationInstanceInput, optFns ...func(*Options)) (*CreateReplicationInstanceOutput, error) {
	if params == nil {
		params = &CreateReplicationInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationInstance", params, optFns, addOperationCreateReplicationInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateReplicationInstanceInput struct {

	// The compute and memory capacity of the replication instance as defined for the
	// specified replication instance class. For example to specify the instance class
	// dms.c4.large, set this parameter to "dms.c4.large". For more information on the
	// settings and capacities for the available replication instance classes, see
	// Selecting the right AWS DMS replication instance for your migration
	// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth).
	//
	// This member is required.
	ReplicationInstanceClass *string

	// The replication instance identifier. This parameter is stored as a lowercase
	// string. Constraints:
	//
	// * Must contain 1-63 alphanumeric characters or hyphens.
	//
	// *
	// First character must be a letter.
	//
	// * Can't end with a hyphen or contain two
	// consecutive hyphens.
	//
	// Example: myrepinstance
	//
	// This member is required.
	ReplicationInstanceIdentifier *string

	// The amount of storage (in gigabytes) to be initially allocated for the
	// replication instance.
	AllocatedStorage *int32

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the replication instance during the maintenance window. This parameter
	// defaults to true. Default: true
	AutoMinorVersionUpgrade *bool

	// The Availability Zone where the replication instance will be created. The
	// default value is a random, system-chosen Availability Zone in the endpoint's AWS
	// Region, for example: us-east-1d
	AvailabilityZone *string

	// A list of custom DNS name servers supported for the replication instance to
	// access your on-premise source or target database. This list overrides the
	// default name servers supported by the replication instance. You can specify a
	// comma-separated list of internet addresses for up to four on-premise DNS name
	// servers. For example: "1.1.1.1,2.2.2.2,3.3.3.3,4.4.4.4"
	DnsNameServers *string

	// The engine version number of the replication instance. If an engine version
	// number is not specified when a replication instance is created, the default is
	// the latest engine version available.
	EngineVersion *string

	// An AWS KMS key identifier that is used to encrypt the data on the replication
	// instance. If you don't specify a value for the KmsKeyId parameter, then AWS DMS
	// uses your default encryption key. AWS KMS creates the default encryption key for
	// your AWS account. Your AWS account has a different default encryption key for
	// each AWS Region.
	KmsKeyId *string

	// Specifies whether the replication instance is a Multi-AZ deployment. You can't
	// set the AvailabilityZone parameter if the Multi-AZ parameter is set to true.
	MultiAZ *bool

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi Default: A 30-minute
	// window selected at random from an 8-hour block of time per AWS Region, occurring
	// on a random day of the week. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// Specifies the accessibility options for the replication instance. A value of
	// true represents an instance with a public IP address. A value of false
	// represents an instance with a private IP address. The default value is true.
	PubliclyAccessible *bool

	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier *string

	// A friendly name for the resource identifier at the end of the EndpointArn
	// response parameter that is returned in the created Endpoint object. The value
	// for this parameter can have up to 31 characters. It can contain only ASCII
	// letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain
	// two consecutive hyphens, and can only begin with a letter, such as
	// Example-App-ARN1. For example, this value might result in the EndpointArn value
	// arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1. If you don't specify a
	// ResourceIdentifier value, AWS DMS generates a default identifier value for the
	// end of EndpointArn.
	ResourceIdentifier *string

	// One or more tags to be assigned to the replication instance.
	Tags []types.Tag

	// Specifies the VPC security group to be used with the replication instance. The
	// VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds []string
}

//
type CreateReplicationInstanceOutput struct {

	// The replication instance that was created.
	ReplicationInstance *types.ReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateReplicationInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationInstance{}, middleware.After)
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
	if err = addOpCreateReplicationInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "CreateReplicationInstance",
	}
}