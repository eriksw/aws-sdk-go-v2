// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new Model resource to an existing RestApi resource.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add a new Model to an existing RestApi resource.
type CreateModelInput struct {

	// [Required] The content-type for the model.
	//
	// This member is required.
	ContentType *string

	// [Required] The name of the model. Must be alphanumeric.
	//
	// This member is required.
	Name *string

	// [Required] The RestApi identifier under which the Model will be created.
	//
	// This member is required.
	RestApiId *string

	// The description of the model.
	Description *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model.
	Schema *string
}

// Represents the data structure of a method's request or response payload. A
// request model defines the data structure of the client-supplied request payload.
// A response model defines the data structure of the response payload returned by
// the back end. Although not required, models are useful for mapping payloads
// between the front end and back end. A model is used for generating an API's SDK,
// validating the input request body, and creating a skeletal mapping template.
// Method, MethodResponse, Models and Mappings
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type CreateModelOutput struct {

	// The content-type for the model.
	ContentType *string

	// The description of the model.
	Description *string

	// The identifier for the model resource.
	Id *string

	// The name of the model. Must be an alphanumeric string.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Do
	// not include "\*/" characters in the description of any properties because such
	// "\*/" characters may be interpreted as the closing marker for comments in some
	// languages, such as Java or JavaScript, causing the installation of your API's
	// SDK generated by API Gateway to fail.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateModel{}, middleware.After)
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
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateModel",
	}
}
