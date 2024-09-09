package mariosdk

import (
	"mariosdk/sdk/sdkclient"
	"testing"
)

func TestDefaultClient_Auth(t *testing.T) {

	client := sdkclient.NewClient()
	client.Config.SetAccessKey("Z9r1DyumUwpUYaNg")
	client.Config.SetSecretKey("qoTYWJM88pBb8H-2qDTI2ayqswFkKYQj")
	client.Config.AddEndpoint("https://mario-syclover.geesec.com/api")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	res, err := client.GetBatchUserInfo([]string{"1811603579241238528"})
	if err != nil {
		t.Errorf("GetBatchUserInfo() failed: %v", err)
	}
	t.Logf("GetBatchUserInfo() result: %v", res)
}
