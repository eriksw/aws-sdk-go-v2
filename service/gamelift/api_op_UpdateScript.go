// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates Realtime script metadata and content. To update script metadata, specify
// the script ID and provide updated name and/or version values. To update script
// content, provide an updated zip file by pointing to either a local file or an
// Amazon S3 bucket location. You can use either method regardless of how the
// original script was uploaded. Use the Version parameter to track updates to the
// script. If the call is successful, the updated metadata is stored in the script
// record and a revised script is uploaded to the Amazon GameLift service. Once the
// script is updated and acquired by a fleet instance, the new version is used for
// all new game sessions. Learn more Amazon GameLift Realtime Servers
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
// Related operations
//
// * CreateScript
//
// * ListScripts
//
// * DescribeScript
//
// *
// UpdateScript
//
// * DeleteScript
func (c *Client) UpdateScript(ctx context.Context, params *UpdateScriptInput, optFns ...func(*Options)) (*UpdateScriptOutput, error) {
	if params == nil {
		params = &UpdateScriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScript", params, optFns, addOperationUpdateScriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScriptInput struct {

	// A unique identifier for a Realtime script to update. You can use either the
	// script ID or ARN value.
	//
	// This member is required.
	ScriptId *string

	// A descriptive label that is associated with a script. Script names do not need
	// to be unique.
	Name *string

	// The Amazon S3 location of your Realtime scripts. The storage location must
	// specify the S3 bucket name, the zip file name (the "key"), and an IAM role ARN
	// that allows Amazon GameLift to access the S3 storage location. The S3 bucket
	// must be in the same Region as the script you're updating. By default, Amazon
	// GameLift uploads the latest version of the zip file; if you have S3 object
	// versioning turned on, you can use the ObjectVersion parameter to specify an
	// earlier version. To call this operation with a storage location, you must have
	// IAM PassRole permission. For more details on IAM roles and PassRole permissions,
	// see  Set up a role for GameLift access
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html).
	StorageLocation *types.S3Location

	// The version that is associated with a build or script. Version strings do not
	// need to be unique.
	Version *string

	// A data object containing your Realtime scripts and dependencies as a zip file.
	// The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	// When using the AWS CLI tool to create a script, this parameter is set to the zip
	// file name. It must be prepended with the string "fileb://" to indicate that the
	// file data is a binary object. For example: --zip-file
	// fileb://myRealtimeScript.zip.
	ZipFile []byte
}

type UpdateScriptOutput struct {

	// The newly created script record with a unique script ID. The new script's
	// storage location reflects an Amazon S3 location: (1) If the script was uploaded
	// from an S3 bucket under your account, the storage location reflects the
	// information that was provided in the CreateScript request; (2) If the script
	// file was uploaded from a local zip file, the storage location reflects an S3
	// location controls by the Amazon GameLift service.
	Script *types.Script

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateScriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateScript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateScript{}, middleware.After)
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
	if err = addOpUpdateScriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScript(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateScript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateScript",
	}
}
