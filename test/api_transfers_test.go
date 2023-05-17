/*
Weavr Multi Product API

Testing TransfersApiService

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

func Test_weavr_TransfersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransfersApiService TransferCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TransfersApi.TransferCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransfersApiService TransferGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id interface{}

		httpRes, err := apiClient.TransfersApi.TransferGet(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransfersApiService TransfersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TransfersApi.TransfersGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}