/*
Weavr Multi Product API

Testing AccessApiService

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

func Test_weavr_AccessApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessApiService LoginWithPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AccessApi.LoginWithPassword(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService Logout", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AccessApi.Logout(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService RequestAccessToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AccessApi.RequestAccessToken(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService StepupSCAChallenge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAOtpChannel

		httpRes, err := apiClient.AccessApi.StepupSCAChallenge(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService StepupSCAChallengePush", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAPushChannel

		httpRes, err := apiClient.AccessApi.StepupSCAChallengePush(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService StepupSCAVerify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAOtpChannel

		httpRes, err := apiClient.AccessApi.StepupSCAVerify(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessApiService UserIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AccessApi.UserIdentities(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
