// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get lens version differences.
func (c *Client) GetLensVersionDifference(ctx context.Context, params *GetLensVersionDifferenceInput, optFns ...func(*Options)) (*GetLensVersionDifferenceOutput, error) {
	if params == nil {
		params = &GetLensVersionDifferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLensVersionDifference", params, optFns, addOperationGetLensVersionDifferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLensVersionDifferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLensVersionDifferenceInput struct {

	// The base version of the lens.
	//
	// This member is required.
	BaseLensVersion *string

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string
}

type GetLensVersionDifferenceOutput struct {

	// The base version of the lens.
	BaseLensVersion *string

	// The latest version of the lens.
	LatestLensVersion *string

	// The alias of the lens, for example, serverless. Each lens is identified by its
	// LensSummary$LensAlias.
	LensAlias *string

	// The differences between the base and latest versions of the lens.
	VersionDifferences *types.VersionDifferences

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLensVersionDifferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLensVersionDifference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLensVersionDifference{}, middleware.After)
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
	if err = addOpGetLensVersionDifferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLensVersionDifference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLensVersionDifference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "GetLensVersionDifference",
	}
}
