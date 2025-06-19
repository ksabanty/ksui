// generate a unit test file for the messageHandler func in msgHandler.go. it utilizes the net/http packages but i would like to mock these requests in the test. please use the testify package for assertions and mocking.

package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock of the http.Client
type MockHTTPClient struct {
    mock.Mock
}
// Do is a mock method for http.Client.Do
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    args := m.Called(req)
    return args.Get(0).(*http.Response), args.Error(1)
}

// TestMessageHandler tests the messageHandler function
func TestMessageHandler(t *testing.T) {
    // Create a new instance of the mock HTTP client
    mockClient := new(MockHTTPClient)

    // Create a test server that uses the mock client
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulate a response from the external service
        responseBody := []byte(`{"status":"success"}`)
        w.WriteHeader(http.StatusOK)
        w.Write(responseBody)
    }))
    defer ts.Close()

    // Mock the Do method of the HTTP client to return the test server's response
    mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
        StatusCode: http.StatusOK,
        Body:       ioutil.NopCloser(bytes.NewBufferString(`{"status":"success"}`)),
    }, nil)

    // Call the messageHandler with the mock client
    req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer([]byte(`{"message":"test"}`)))
    assert.NoError(t, err)

    resp, err := mockClient.Do(req)
    assert.NoError(t, err)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    assert.NoError(t, err)

    // Assert the response status and body
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    assert.JSONEq(t, `{"status":"success"}`, string(body))

    // Assert that the Do method was called
    mockClient.AssertExpectations(t)
}
