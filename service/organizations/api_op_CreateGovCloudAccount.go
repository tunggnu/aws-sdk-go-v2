// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is available if all of the following are true:
//   - You're authorized to create accounts in the Amazon Web Services GovCloud
//     (US) Region. For more information on the Amazon Web Services GovCloud (US)
//     Region, see the Amazon Web Services GovCloud User Guide. (https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/welcome.html)
//   - You already have an account in the Amazon Web Services GovCloud (US) Region
//     that is paired with a management account of an organization in the commercial
//     Region.
//   - You call this action from the management account of your organization in
//     the commercial Region.
//   - You have the organizations:CreateGovCloudAccount permission.
//
// Organizations automatically creates the required service-linked role named
// AWSServiceRoleForOrganizations . For more information, see Organizations and
// Service-Linked Roles (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs)
// in the Organizations User Guide. Amazon Web Services automatically enables
// CloudTrail for Amazon Web Services GovCloud (US) accounts, but you should also
// do the following:
//   - Verify that CloudTrail is enabled to store logs.
//   - Create an Amazon S3 bucket for CloudTrail log storage. For more
//     information, see Verifying CloudTrail Is Enabled (https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/verifying-cloudtrail.html)
//     in the Amazon Web Services GovCloud User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission. The tags are attached to the commercial
// account associated with the GovCloud account, rather than the GovCloud account
// itself. To add tags to the GovCloud account, call the TagResource operation in
// the GovCloud Region after the new GovCloud account exists. You call this action
// from the management account of your organization in the commercial Region to
// create a standalone Amazon Web Services account in the Amazon Web Services
// GovCloud (US) Region. After the account is created, the management account of an
// organization in the Amazon Web Services GovCloud (US) Region can invite it to
// that organization. For more information on inviting standalone accounts in the
// Amazon Web Services GovCloud (US) to join an organization, see Organizations (https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-organizations.html)
// in the Amazon Web Services GovCloud User Guide. Calling CreateGovCloudAccount
// is an asynchronous request that Amazon Web Services performs in the background.
// Because CreateGovCloudAccount operates asynchronously, it can return a
// successful completion message even though account initialization might still be
// in progress. You might need to wait a few minutes before you can successfully
// access the account. To check the status of the request, do one of the following:
//
//   - Use the OperationId response element from this operation to provide as a
//     parameter to the DescribeCreateAccountStatus operation.
//   - Check the CloudTrail log for the CreateAccountResult event. For information
//     on using CloudTrail with Organizations, see Monitoring the Activity in Your
//     Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_monitoring.html)
//     in the Organizations User Guide.
//
// When you call the CreateGovCloudAccount action, you create two accounts: a
// standalone account in the Amazon Web Services GovCloud (US) Region and an
// associated account in the commercial Region for billing and support purposes.
// The account in the commercial Region is automatically a member of the
// organization whose credentials made the request. Both accounts are associated
// with the same email address. A role is created in the new account in the
// commercial Region that allows the management account in the organization in the
// commercial Region to assume it. An Amazon Web Services GovCloud (US) account is
// then created and associated with the commercial account that you just created. A
// role is also created in the new Amazon Web Services GovCloud (US) account that
// can be assumed by the Amazon Web Services GovCloud (US) account that is
// associated with the management account of the commercial organization. For more
// information and to view a diagram that explains how account access works, see
// Organizations (https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-organizations.html)
// in the Amazon Web Services GovCloud User Guide. For more information about
// creating accounts, see Creating an Amazon Web Services account in Your
// Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html)
// in the Organizations User Guide.
//   - When you create an account in an organization using the Organizations
//     console, API, or CLI commands, the information required for the account to
//     operate as a standalone account is not automatically collected. This includes a
//     payment method and signing the end user license agreement (EULA). If you must
//     remove an account from your organization later, you can do so only after you
//     provide the missing information. Follow the steps at To leave an organization
//     as a member account (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
//     in the Organizations User Guide.
//   - If you get an exception that indicates that you exceeded your account
//     limits for the organization, contact Amazon Web Services Support (https://console.aws.amazon.com/support/home#/)
//     .
//   - If you get an exception that indicates that the operation failed because
//     your organization is still initializing, wait one hour and then try again. If
//     the error persists, contact Amazon Web Services Support (https://console.aws.amazon.com/support/home#/)
//     .
//   - Using CreateGovCloudAccount to create multiple temporary accounts isn't
//     recommended. You can only close an account from the Amazon Web Services Billing
//     and Cost Management console, and you must be signed in as the root user. For
//     information on the requirements and process for closing an account, see
//     Closing an Amazon Web Services account (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html)
//     in the Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether to
// create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable it,
// only the account root user can access billing information. For information about
// how to disable this switch for an account, see Granting Access to Your Billing
// Information and Tools (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html)
// .
func (c *Client) CreateGovCloudAccount(ctx context.Context, params *CreateGovCloudAccountInput, optFns ...func(*Options)) (*CreateGovCloudAccountOutput, error) {
	if params == nil {
		params = &CreateGovCloudAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGovCloudAccount", params, optFns, c.addOperationCreateGovCloudAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGovCloudAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGovCloudAccountInput struct {

	// The friendly name of the member account. The account name can consist of only
	// the characters [a-z],[A-Z],[0-9], hyphen (-), or dot (.) You can't separate
	// characters with a dash (–).
	//
	// This member is required.
	AccountName *string

	// Specifies the email address of the owner to assign to the new member account in
	// the commercial Region. This email address must not already be associated with
	// another Amazon Web Services account. You must use a valid email address to
	// complete account creation. The rules for a valid email address:
	//   - The address must be a minimum of 6 and a maximum of 64 characters long.
	//   - All characters must be 7-bit ASCII characters.
	//   - There must be one and only one @ symbol, which separates the local name
	//   from the domain name.
	//   - The local name can't contain any of the following characters: whitespace, "
	//   ' ( ) < > [ ] : ; , \ | % &
	//   - The local name can't begin with a dot (.)
	//   - The domain name can consist of only the characters [a-z],[A-Z],[0-9],
	//   hyphen (-), or dot (.)
	//   - The domain name can't begin or end with a hyphen (-) or dot (.)
	//   - The domain name must contain at least one dot
	// You can't access the root user of the account or remove an account that was
	// created with an invalid email address. Like all request parameters for
	// CreateGovCloudAccount , the request for the email address for the Amazon Web
	// Services GovCloud (US) account originates from the commercial Region, not from
	// the Amazon Web Services GovCloud (US) Region.
	//
	// This member is required.
	Email *string

	// If set to ALLOW , the new linked account in the commercial Region enables IAM
	// users to access account billing information if they have the required
	// permissions. If set to DENY , only the root user of the new account can access
	// account billing information. For more information, see Activating Access to the
	// Billing and Cost Management Console (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
	// in the Amazon Web Services Billing and Cost Management User Guide. If you don't
	// specify this parameter, the value defaults to ALLOW , and IAM users and roles
	// with the required permissions can access billing information for the new
	// account.
	IamUserAccessToBilling types.IAMUserAccessToBilling

	// (Optional) The name of an IAM role that Organizations automatically
	// preconfigures in the new member accounts in both the Amazon Web Services
	// GovCloud (US) Region and in the commercial Region. This role trusts the
	// management account, allowing users in the management account to assume the role,
	// as permitted by the management account administrator. The role has administrator
	// permissions in the new member account. If you don't specify this parameter, the
	// role name defaults to OrganizationAccountAccessRole . For more information about
	// how to use this role to access the member account, see Accessing and
	// Administering the Member Accounts in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role)
	// in the Organizations User Guide and steps 2 and 3 in Tutorial: Delegate Access
	// Across Amazon Web Services accounts Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)
	// in the IAM User Guide. The regex pattern (http://wikipedia.org/wiki/regex) that
	// is used to validate this parameter. The pattern can include uppercase letters,
	// lowercase letters, digits with no spaces, and any of the following characters:
	// =,.@-
	RoleName *string

	// A list of tags that you want to attach to the newly created account. These tags
	// are attached to the commercial account associated with the GovCloud account, and
	// not to the GovCloud account itself. To add tags to the actual GovCloud account,
	// call the TagResource operation in the GovCloud region after the new GovCloud
	// account exists. For each tag in the list, you must specify both a tag key and a
	// value. You can set the value to an empty string, but you can't set it to null .
	// For more information about tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)
	// in the Organizations User Guide. If any one of the tags is not valid or if you
	// exceed the maximum allowed number of tags for an account, then the entire
	// request fails and the account is not created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateGovCloudAccountOutput struct {

	// Contains the status about a CreateAccount or CreateGovCloudAccount request to
	// create an Amazon Web Services account or an Amazon Web Services GovCloud (US)
	// account in an organization.
	CreateAccountStatus *types.CreateAccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGovCloudAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGovCloudAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGovCloudAccount{}, middleware.After)
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
	if err = addCreateGovCloudAccountResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateGovCloudAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGovCloudAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGovCloudAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "CreateGovCloudAccount",
	}
}

type opCreateGovCloudAccountResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateGovCloudAccountResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateGovCloudAccountResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "organizations"
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
				signingName = "organizations"
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
				v4aScheme.SigningName = aws.String("organizations")
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

func addCreateGovCloudAccountResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateGovCloudAccountResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
