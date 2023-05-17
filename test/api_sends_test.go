/*
Weavr Multi Product API

Testing SendsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package weavr

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_weavr_SendsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SendsApiService SendCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SendsApi.SendCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SendsApiService SendGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id interface{}

		httpRes, err := apiClient.SendsApi.SendGet(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SendsApiService SendSCAChallenge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id interface{}
		var channel SCAOtpChannel

		httpRes, err := apiClient.SendsApi.SendSCAChallenge(context.Background(), id, channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SendsApiService SendSCAChallengePush", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id interface{}
		var channel SCAPushChannel

		httpRes, err := apiClient.SendsApi.SendSCAChallengePush(context.Background(), id, channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SendsApiService SendSCAVerify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id interface{}
		var channel SCAOtpChannel

		httpRes, err := apiClient.SendsApi.SendSCAVerify(context.Background(), id, channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SendsApiService SendsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SendsApi.SendsGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}