# go-cod

go-cod is the _unofficial_ Go SDK of the Call of Duty API used in the callofduty.com site

## Getting Started
```
import "github.com/carlocayos/go-cod"
```

Create a new go-cod client then use the services to access the COD APIs.

A short example:
```
client := go_cod.NewClient(nil)

// enter values to the request params - this may include header, path and query params, or a json payload
leaderBoardParams := service.LeaderBoardParams{
    Page:     1,
    Platform: "battle",
    Title:    "mw",
    Context:  context.Background(),
}

// get leader board list
leaderBoardResponse, err := c.ServiceClient.Operations.LeaderBoard(&leaderBoardParams)
```

The cod client is composed of an authentication client and service client.

### Authentication

Authentication client is for getting a security token that will be used for most of the API requests.

See example below on how to retrieve a token and use it in authenticated requests.

```
	// =======================================================
	// 1) First, is creating a client and sending a register device request
	// =======================================================
	c := gocod.NewClient(nil)

	// send a register device request with a unique device id
	//  ksuid is used here to generate a unique id, but any uid generator would be fine
	deviceId := ksuid.New().String()

	param := authentication.RegisterDeviceParams{
		RegisterDeviceRequest: &models.RegisterDeviceRequest{DeviceID: &deviceId},
		Context:               context.Background(),
	}
	// send register device request
	registerDeviceResponse, _ := c.AuthenticationClient.Operations.RegisterDevice(&param)

	// =======================================================
	// 2) Next is the Login request. Replace it with your email and password
	//      and pass the returned AuthHeader and Device ID
	// =======================================================
	email := "<< CHANGE ME >>"
	password := "<< CHANGE ME >>"

	loginParams := authentication.LoginParams{
		Authorization: "Bearer " + *registerDeviceResponse.Payload.Data.AuthHeader,
		XCodDeviceID:  deviceId,
		LoginRequest: &models.LoginRequest{
			Email:    &email,
			Password: &password,
		},
		Context: context.Background(),
	}
	loginResponse, _ := c.AuthenticationClient.Operations.Login(&loginParams)

	// =======================================================
	// 3) Final step is to use the token to send a Get Gamer Match List
	//    This request requires a token for authentication. Get the ACT_SSO_COOKIE from the Login 
	//    response and append it to the Cookie request as "ACT_SSO_COOKIE=YOUR_TOKEN
	// =======================================================
	// send a gamer match list request
	limit := int32(10)

	matchListParams := service.MatchListParams{
		Start:      0,
		End:        0,
		Cookie:     "ACT_SSO_COOKIE=" + loginResponse.Payload.ACTSSOCOOKIE,
		GameType:   "mp",
		Gamertag:   "MrExcitement#6438", // change this to your gamer tag
		Limit:      &limit,
		LookupType: "gamer",
		Platform:   "battle",
		Title:      "cw",
		Context:    context.Background(),
	}

	// send the Match List request
	matchListResponse, _ := c.ServiceClient.Operations.MatchList(&matchListParams)
	for _, v := range matchListResponse.Payload.Data {
		fmt.Printf("Map = %v\n", v.Map)
		fmt.Printf("MatchID = %v\n", v.MatchID)
		fmt.Printf("Platform = %v\n", v.Platform)
		fmt.Printf("Timestamp = %v\n", v.Timestamp)
		fmt.Printf("Title = %v\n", v.Title)
		fmt.Printf("Type = %v\n", v.Type)
	}
```

### Unmapped fields

Fields that are not mapped to a struct field can be accessed from the `*AdditionalProperties` field of type `map[string]interface{}` 

For example, the weapons values from the loadout request can be found under `Payload.Data.LoadoutResponseDataAdditionalProperties["weapons"]`
Refer to the sample code in [example_unmapped.go](examples/unmapped_fields_loadout/example_unmapped.go) on how to get these field values

## More sample codes

Sample codes are in [examples](examples)

## Generate Models and Client from Swagger Spec

The model and client codes are generated using [go-swagger](https://github.com/go-swagger/go-swagger) and
[OpenAPI Version 2.0](https://swagger.io/specification/v2/)

See [here](https://goswagger.io/generate/client.html) for more information on how to generate an API client

### Generating client and models

For any changes in the COD API, update the [swagger spec](api/specs) and generate a new client and model

Run this command to generate the client and model codes
```shell
make swagger-gen
```

### API Documentation

Run this command to run the [ReDoc](https://github.com/bfirsh/docker-redoc) container
```shell
make swagger-docs
```

Then open http://localhost:9000 to see the list of APIs.

## 🚧 Work in Progress 🚧

- [ ] Add more helper functions
- [ ] Populate response fields for Friend Stats Profile
- [ ] Populate response fields for COD Points
- [ ] Context handling
- [ ] Other missing APIs

## Contribution and Support
Code improvements, suggestions, and updates are most welcome. Please feel free to raise an issue or create a pull
request for your changes.

There is no official COD API and documentation released by Activision. If there is a breaking change on the COD API 
then let me know, so I can update this project. :)

## Credits

Thanks to [Lierrmm](https://github.com/Lierrmm) for his work on the [NodeJS Call of Duty API Wrapper](https://github.com/Lierrmm/Node-CallOfDuty)

## Developer

**Personal Site:** [carlocayos.dev](https://carlocayos.dev)

<a href="https://www.buymeacoffee.com/ccayos" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

**ETH address:** 0x2A17e4031FFeF64C638Dd9B190e05a150b2B8FBc

