// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TransactWriteItemsInput struct {
	_ struct{} `type:"structure"`

	// Providing a ClientRequestToken makes the call to TransactWriteItems idempotent,
	// meaning that multiple identical calls have the same effect as one single
	// call.
	//
	// Although multiple identical calls using the same client request token produce
	// the same result on the server (no side effects), the responses to the calls
	// might not be the same. If the ReturnConsumedCapacity> parameter is set, then
	// the initial TransactWriteItems call returns the amount of write capacity
	// units consumed in making the changes. Subsequent TransactWriteItems calls
	// with the same client token return the number of read capacity units consumed
	// in reading the item.
	//
	// A client request token is valid for 10 minutes after the first request that
	// uses it is completed. After 10 minutes, any request with the same client
	// token is treated as a new request. Do not resubmit the same request with
	// the same client token for more than 10 minutes, or the result might not be
	// idempotent.
	//
	// If you submit a request with the same client token but a change in other
	// parameters within the 10-minute idempotency window, DynamoDB returns an IdempotentParameterMismatch
	// exception.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// Determines the level of detail about provisioned throughput consumption that
	// is returned in the response:
	//
	//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
	//    operation, together with ConsumedCapacity for each table and secondary
	//    index that was accessed. Note that some operations, such as GetItem and
	//    BatchGetItem, do not access any indexes at all. In these cases, specifying
	//    INDEXES will only return ConsumedCapacity information for table(s).
	//
	//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
	//    the operation.
	//
	//    * NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity ReturnConsumedCapacity `type:"string" enum:"true"`

	// Determines whether item collection metrics are returned. If set to SIZE,
	// the response includes statistics about item collections (if any), that were
	// modified during the operation and are returned in the response. If set to
	// NONE (the default), no statistics are returned.
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics `type:"string" enum:"true"`

	// An ordered array of up to 25 TransactWriteItem objects, each of which contains
	// a ConditionCheck, Put, Update, or Delete object. These can operate on items
	// in different tables, but the tables must reside in the same AWS account and
	// Region, and no two of them can operate on the same item.
	//
	// TransactItems is a required field
	TransactItems []TransactWriteItem `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TransactWriteItemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TransactWriteItemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TransactWriteItemsInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.TransactItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransactItems"))
	}
	if s.TransactItems != nil && len(s.TransactItems) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransactItems", 1))
	}
	if s.TransactItems != nil {
		for i, v := range s.TransactItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TransactItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TransactWriteItemsOutput struct {
	_ struct{} `type:"structure"`

	// The capacity units consumed by the entire TransactWriteItems operation. The
	// values of the list are ordered according to the ordering of the TransactItems
	// request parameter.
	ConsumedCapacity []ConsumedCapacity `type:"list"`

	// A list of tables that were processed by TransactWriteItems and, for each
	// table, information about any item collections that were affected by individual
	// UpdateItem, PutItem, or DeleteItem operations.
	ItemCollectionMetrics map[string][]ItemCollectionMetrics `type:"map"`
}

// String returns the string representation
func (s TransactWriteItemsOutput) String() string {
	return awsutil.Prettify(s)
}

const opTransactWriteItems = "TransactWriteItems"

// TransactWriteItemsRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// TransactWriteItems is a synchronous write operation that groups up to 25
// action requests. These actions can target items in different tables, but
// not in different AWS accounts or Regions, and no two actions can target the
// same item. For example, you cannot both ConditionCheck and Update the same
// item. The aggregate size of the items in the transaction cannot exceed 4
// MB.
//
// The actions are completed atomically so that either all of them succeed,
// or all of them fail. They are defined by the following objects:
//
//    * Put — Initiates a PutItem operation to write a new item. This structure
//    specifies the primary key of the item to be written, the name of the table
//    to write it in, an optional condition expression that must be satisfied
//    for the write to succeed, a list of the item's attributes, and a field
//    indicating whether to retrieve the item's attributes if the condition
//    is not met.
//
//    * Update — Initiates an UpdateItem operation to update an existing item.
//    This structure specifies the primary key of the item to be updated, the
//    name of the table where it resides, an optional condition expression that
//    must be satisfied for the update to succeed, an expression that defines
//    one or more attributes to be updated, and a field indicating whether to
//    retrieve the item's attributes if the condition is not met.
//
//    * Delete — Initiates a DeleteItem operation to delete an existing item.
//    This structure specifies the primary key of the item to be deleted, the
//    name of the table where it resides, an optional condition expression that
//    must be satisfied for the deletion to succeed, and a field indicating
//    whether to retrieve the item's attributes if the condition is not met.
//
//    * ConditionCheck — Applies a condition to an item that is not being
//    modified by the transaction. This structure specifies the primary key
//    of the item to be checked, the name of the table where it resides, a condition
//    expression that must be satisfied for the transaction to succeed, and
//    a field indicating whether to retrieve the item's attributes if the condition
//    is not met.
//
// DynamoDB rejects the entire TransactWriteItems request if any of the following
// is true:
//
//    * A condition in one of the condition expressions is not met.
//
//    * An ongoing operation is in the process of updating the same item.
//
//    * There is insufficient provisioned capacity for the transaction to be
//    completed.
//
//    * An item size becomes too large (bigger than 400 KB), a local secondary
//    index (LSI) becomes too large, or a similar validation error occurs because
//    of changes made by the transaction.
//
//    * The aggregate size of the items in the transaction exceeds 4 MB.
//
//    * There is a user error, such as an invalid data format.
//
//    // Example sending a request using TransactWriteItemsRequest.
//    req := client.TransactWriteItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/TransactWriteItems
func (c *Client) TransactWriteItemsRequest(input *TransactWriteItemsInput) TransactWriteItemsRequest {
	op := &aws.Operation{
		Name:       opTransactWriteItems,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TransactWriteItemsInput{}
	}

	req := c.newRequest(op, input, &TransactWriteItemsOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

	return TransactWriteItemsRequest{Request: req, Input: input, Copy: c.TransactWriteItemsRequest}
}

// TransactWriteItemsRequest is the request type for the
// TransactWriteItems API operation.
type TransactWriteItemsRequest struct {
	*aws.Request
	Input *TransactWriteItemsInput
	Copy  func(*TransactWriteItemsInput) TransactWriteItemsRequest
}

// Send marshals and sends the TransactWriteItems API request.
func (r TransactWriteItemsRequest) Send(ctx context.Context) (*TransactWriteItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TransactWriteItemsResponse{
		TransactWriteItemsOutput: r.Request.Data.(*TransactWriteItemsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TransactWriteItemsResponse is the response type for the
// TransactWriteItems API operation.
type TransactWriteItemsResponse struct {
	*TransactWriteItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TransactWriteItems request.
func (r *TransactWriteItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
