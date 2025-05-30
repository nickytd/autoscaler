// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVcnsRequest wrapper for the ListVcns operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListVcns.go.html to see an example of how to use ListVcnsRequest.
type ListVcnsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListVcnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListVcnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle
	// state. The state value is case-insensitive.
	LifecycleState VcnLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVcnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVcnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVcnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVcnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVcnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVcnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVcnsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVcnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVcnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVcnLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVcnLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVcnsResponse wrapper for the ListVcns operation
type ListVcnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Vcn instances
	Items []Vcn `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVcnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVcnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVcnsSortByEnum Enum with underlying type: string
type ListVcnsSortByEnum string

// Set of constants representing the allowable values for ListVcnsSortByEnum
const (
	ListVcnsSortByTimecreated ListVcnsSortByEnum = "TIMECREATED"
	ListVcnsSortByDisplayname ListVcnsSortByEnum = "DISPLAYNAME"
)

var mappingListVcnsSortByEnum = map[string]ListVcnsSortByEnum{
	"TIMECREATED": ListVcnsSortByTimecreated,
	"DISPLAYNAME": ListVcnsSortByDisplayname,
}

var mappingListVcnsSortByEnumLowerCase = map[string]ListVcnsSortByEnum{
	"timecreated": ListVcnsSortByTimecreated,
	"displayname": ListVcnsSortByDisplayname,
}

// GetListVcnsSortByEnumValues Enumerates the set of values for ListVcnsSortByEnum
func GetListVcnsSortByEnumValues() []ListVcnsSortByEnum {
	values := make([]ListVcnsSortByEnum, 0)
	for _, v := range mappingListVcnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVcnsSortByEnumStringValues Enumerates the set of values in String for ListVcnsSortByEnum
func GetListVcnsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListVcnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVcnsSortByEnum(val string) (ListVcnsSortByEnum, bool) {
	enum, ok := mappingListVcnsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVcnsSortOrderEnum Enum with underlying type: string
type ListVcnsSortOrderEnum string

// Set of constants representing the allowable values for ListVcnsSortOrderEnum
const (
	ListVcnsSortOrderAsc  ListVcnsSortOrderEnum = "ASC"
	ListVcnsSortOrderDesc ListVcnsSortOrderEnum = "DESC"
)

var mappingListVcnsSortOrderEnum = map[string]ListVcnsSortOrderEnum{
	"ASC":  ListVcnsSortOrderAsc,
	"DESC": ListVcnsSortOrderDesc,
}

var mappingListVcnsSortOrderEnumLowerCase = map[string]ListVcnsSortOrderEnum{
	"asc":  ListVcnsSortOrderAsc,
	"desc": ListVcnsSortOrderDesc,
}

// GetListVcnsSortOrderEnumValues Enumerates the set of values for ListVcnsSortOrderEnum
func GetListVcnsSortOrderEnumValues() []ListVcnsSortOrderEnum {
	values := make([]ListVcnsSortOrderEnum, 0)
	for _, v := range mappingListVcnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVcnsSortOrderEnumStringValues Enumerates the set of values in String for ListVcnsSortOrderEnum
func GetListVcnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVcnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVcnsSortOrderEnum(val string) (ListVcnsSortOrderEnum, bool) {
	enum, ok := mappingListVcnsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
