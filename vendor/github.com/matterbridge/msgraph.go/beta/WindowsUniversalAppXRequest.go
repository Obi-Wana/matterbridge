// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// WindowsUniversalAppXRequestBuilder is request builder for WindowsUniversalAppX
type WindowsUniversalAppXRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsUniversalAppXRequest
func (b *WindowsUniversalAppXRequestBuilder) Request() *WindowsUniversalAppXRequest {
	return &WindowsUniversalAppXRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsUniversalAppXRequest is request for WindowsUniversalAppX
type WindowsUniversalAppXRequest struct{ BaseRequest }

// Get performs GET request for WindowsUniversalAppX
func (r *WindowsUniversalAppXRequest) Get(ctx context.Context) (resObj *WindowsUniversalAppX, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsUniversalAppX
func (r *WindowsUniversalAppXRequest) Update(ctx context.Context, reqObj *WindowsUniversalAppX) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsUniversalAppX
func (r *WindowsUniversalAppXRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CommittedContainedApps returns request builder for MobileContainedApp collection
func (b *WindowsUniversalAppXRequestBuilder) CommittedContainedApps() *WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder {
	bb := &WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/committedContainedApps"
	return bb
}

// WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder is request builder for MobileContainedApp collection
type WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileContainedApp collection
func (b *WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder) Request() *WindowsUniversalAppXCommittedContainedAppsCollectionRequest {
	return &WindowsUniversalAppXCommittedContainedAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileContainedApp item
func (b *WindowsUniversalAppXCommittedContainedAppsCollectionRequestBuilder) ID(id string) *MobileContainedAppRequestBuilder {
	bb := &MobileContainedAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsUniversalAppXCommittedContainedAppsCollectionRequest is request for MobileContainedApp collection
type WindowsUniversalAppXCommittedContainedAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileContainedApp collection
func (r *WindowsUniversalAppXCommittedContainedAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileContainedApp, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileContainedApp
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MobileContainedApp
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for MobileContainedApp collection
func (r *WindowsUniversalAppXCommittedContainedAppsCollectionRequest) Get(ctx context.Context) ([]MobileContainedApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileContainedApp collection
func (r *WindowsUniversalAppXCommittedContainedAppsCollectionRequest) Add(ctx context.Context, reqObj *MobileContainedApp) (resObj *MobileContainedApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}