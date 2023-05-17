/*
Weavr Multi Product API

Testing AdditionalFactorsApiService

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

func Test_weavr_AdditionalFactorsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdditionalFactorsApiService AuthFactorsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdditionalFactorsApi.AuthFactorsGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdditionalFactorsApiService EnrolDeviceUsingOtpStepOne", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAOtpChannel

		httpRes, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingOtpStepOne(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdditionalFactorsApiService EnrolDeviceUsingOtpStepTwo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAOtpChannel

		httpRes, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingOtpStepTwo(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdditionalFactorsApiService EnrolDeviceUsingPush", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAPushChannel

		httpRes, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingPush(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdditionalFactorsApiService UnlinkDeviceUsingPush", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel SCAPushChannel

		httpRes, err := apiClient.AdditionalFactorsApi.UnlinkDeviceUsingPush(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}