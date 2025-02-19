// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies the specified DB snapshot. The source DB snapshot must be in the
// available state. You can copy a snapshot from one AWS Region to another. In that
// case, the AWS Region where you call the CopyDBSnapshot action is the destination
// AWS Region for the DB snapshot copy. For more information about copying
// snapshots, see Copying a DB Snapshot
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopyDBSnapshot)
// in the Amazon RDS User Guide.
func (c *Client) CopyDBSnapshot(ctx context.Context, params *CopyDBSnapshotInput, optFns ...func(*Options)) (*CopyDBSnapshotOutput, error) {
	if params == nil {
		params = &CopyDBSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyDBSnapshot", params, optFns, addOperationCopyDBSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CopyDBSnapshotInput struct {

	// The identifier for the source DB snapshot. If the source snapshot is in the same
	// AWS Region as the copy, specify a valid DB snapshot identifier. For example, you
	// might specify rds:mysql-instance1-snapshot-20130805. If the source snapshot is
	// in a different AWS Region than the copy, specify a valid DB snapshot ARN. For
	// example, you might specify
	// arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805.
	// If you are copying from a shared manual DB snapshot, this parameter must be the
	// Amazon Resource Name (ARN) of the shared DB snapshot. If you are copying an
	// encrypted snapshot this parameter must be in the ARN format for the source AWS
	// Region, and must match the SourceDBSnapshotIdentifier in the PreSignedUrl
	// parameter. Constraints:
	//
	// * Must specify a valid system snapshot in the
	// "available" state.
	//
	// Example: rds:mydb-2012-04-02-00-01 Example:
	// arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20130805
	//
	// This member is required.
	SourceDBSnapshotIdentifier *string

	// The identifier for the copy of the snapshot. Constraints:
	//
	// * Can't be null,
	// empty, or blank
	//
	// * Must contain from 1 to 255 letters, numbers, or hyphens
	//
	// *
	// First character must be a letter
	//
	// * Can't end with a hyphen or contain two
	// consecutive hyphens
	//
	// Example: my-db-snapshot
	//
	// This member is required.
	TargetDBSnapshotIdentifier *string

	// A value that indicates whether to copy all tags from the source DB snapshot to
	// the target DB snapshot. By default, tags are not copied.
	CopyTags *bool

	// The AWS KMS key identifier for an encrypted DB snapshot. The AWS KMS key
	// identifier is the key ARN, key ID, alias ARN, or alias name for the AWS KMS
	// customer master key (CMK). If you copy an encrypted DB snapshot from your AWS
	// account, you can specify a value for this parameter to encrypt the copy with a
	// new AWS KMS CMK. If you don't specify a value for this parameter, then the copy
	// of the DB snapshot is encrypted with the same AWS KMS key as the source DB
	// snapshot. If you copy an encrypted DB snapshot that is shared from another AWS
	// account, then you must specify a value for this parameter. If you specify this
	// parameter when you copy an unencrypted snapshot, the copy is encrypted. If you
	// copy an encrypted snapshot to a different AWS Region, then you must specify a
	// AWS KMS key identifier for the destination AWS Region. AWS KMS CMKs are specific
	// to the AWS Region that they are created in, and you can't use CMKs from one AWS
	// Region in another AWS Region.
	KmsKeyId *string

	// The name of an option group to associate with the copy of the snapshot. Specify
	// this option if you are copying a snapshot from one AWS Region to another, and
	// your DB instance uses a nondefault option group. If your source DB instance uses
	// Transparent Data Encryption for Oracle or Microsoft SQL Server, you must specify
	// this option when copying across AWS Regions. For more information, see Option
	// group considerations
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CopySnapshot.html#USER_CopySnapshot.Options)
	// in the Amazon RDS User Guide.
	OptionGroupName *string

	// The URL that contains a Signature Version 4 signed request for the
	// CopyDBSnapshot API action in the source AWS Region that contains the source DB
	// snapshot to copy. You must specify this parameter when you copy an encrypted DB
	// snapshot from another AWS Region by using the Amazon RDS API. Don't specify
	// PreSignedUrl when you are copying an encrypted DB snapshot in the same AWS
	// Region. The presigned URL must be a valid request for the CopyDBSnapshot API
	// action that can be executed in the source AWS Region that contains the encrypted
	// DB snapshot to be copied. The presigned URL request must contain the following
	// parameter values:
	//
	// * DestinationRegion - The AWS Region that the encrypted DB
	// snapshot is copied to. This AWS Region is the same one where the CopyDBSnapshot
	// action is called that contains this presigned URL. For example, if you copy an
	// encrypted DB snapshot from the us-west-2 AWS Region to the us-east-1 AWS Region,
	// then you call the CopyDBSnapshot action in the us-east-1 AWS Region and provide
	// a presigned URL that contains a call to the CopyDBSnapshot action in the
	// us-west-2 AWS Region. For this example, the DestinationRegion in the presigned
	// URL must be set to the us-east-1 AWS Region.
	//
	// * KmsKeyId - The AWS KMS key
	// identifier for the customer master key (CMK) to use to encrypt the copy of the
	// DB snapshot in the destination AWS Region. This is the same identifier for both
	// the CopyDBSnapshot action that is called in the destination AWS Region, and the
	// action contained in the presigned URL.
	//
	// * SourceDBSnapshotIdentifier - The DB
	// snapshot identifier for the encrypted snapshot to be copied. This identifier
	// must be in the Amazon Resource Name (ARN) format for the source AWS Region. For
	// example, if you are copying an encrypted DB snapshot from the us-west-2 AWS
	// Region, then your SourceDBSnapshotIdentifier looks like the following example:
	// arn:aws:rds:us-west-2:123456789012:snapshot:mysql-instance1-snapshot-20161115.
	//
	// To
	// learn how to generate a Signature Version 4 signed request, see Authenticating
	// Requests: Using Query Parameters (AWS Signature Version 4)
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process
	// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). If you
	// are using an AWS SDK tool or the AWS CLI, you can specify SourceRegion (or
	// --source-region for the AWS CLI) instead of specifying PreSignedUrl manually.
	// Specifying SourceRegion autogenerates a pre-signed URL that is a valid request
	// for the operation that can be executed in the source AWS Region.
	PreSignedUrl *string

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag

	// The external custom Availability Zone (CAZ) identifier for the target CAZ.
	// Example: rds-caz-aiqhTgQv.
	TargetCustomAvailabilityZone *string

	// Used by the SDK's PresignURL autofill customization to specify the region the of
	// the client's request.
	destinationRegion *string
}

type CopyDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB snapshot. This data type is used as a
	// response element in the DescribeDBSnapshots action.
	DBSnapshot *types.DBSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyDBSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBSnapshot{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addCopyDBSnapshotPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopyDBSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBSnapshot(options.Region), middleware.Before); err != nil {
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

func copyCopyDBSnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopyDBSnapshotPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCopyDBSnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopyDBSnapshotPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCopyDBSnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCopyDBSnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCopyDBSnapshotPreSignedUrl,

			GetSourceRegion: getCopyDBSnapshotSourceRegion,

			CopyInput: copyCopyDBSnapshotInputForPresign,

			SetDestinationRegion: setCopyDBSnapshotdestinationRegion,

			SetPresignedURL: setCopyDBSnapshotPreSignedUrl,
		},
		Presigner: &presignAutoFillCopyDBSnapshotClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCopyDBSnapshotClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCopyDBSnapshotClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CopyDBSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBSnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCopyDBSnapshot(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCopyDBSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBSnapshot",
	}
}

// PresignCopyDBSnapshot is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCopyDBSnapshot(ctx context.Context, params *CopyDBSnapshotInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CopyDBSnapshotInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CopyDBSnapshot", params, clientOptFns,
		addOperationCopyDBSnapshotMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
