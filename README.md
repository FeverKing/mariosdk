# mario sdk

## Usage

### Setup

#### initialize

```go
client := sdkclient.NewClient()
```

#### Config

```go
client.Config.SetAccessKey(_YOUR_AK_HERE_)
client.Config.SetSecretKey(_YOUR_SK_HERE_)
client.Config.AddEndpoint(_MARIO_ENDPOINT_)
```

#### Auth

```go
err := client.Auth()
if err != nil { 
	// do something
}
```

### Function

#### Get Users' Info From Mario

```go
info, err := client.GetBatchUserInfo([]string{"1811603579241238528"})
if err != nil {
    // do something
}
```