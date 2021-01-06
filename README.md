# go-cod

go-cod is the _unofficial_ Go SDK of the Call of Duty API used in the callofduty.com site

## Getting Started
```
import "github.com/carlocayos/go-cod"
```

Create a new go-cod client then use the services to access the COD APIs.

A short example:
```
// create the client
c := go_cod.NewClient(nil)

// get leader board list
leaderBoardResp, _ := c.LeaderBoard(context.Background(), go_cod.ModernWarfare, go_cod.Battlenet, 3)
fmt.Println(leaderBoardResp.Status)
fmt.Println(leaderBoardResp.Data.Title)
```

The cod client is composed of an authentication client and service client.

### Authentication

Authentication client is for getting a security token that will be used for most of the API requests.

See example below on how to retrieve a token and use it in authenticated requests.

```
// =======================================================
// 1) First, is creating a client and sending a register device request
// =======================================================
c := go_cod.NewClient(nil)

// send a register device request with a unique device id
//  ksuid is used here to generate a unique id, but any uid generator would be fine
deviceId := ksuid.New().String()
registerDeviceRes, _ := c.RegisterDevice(context.Background(), deviceId)

// =======================================================
// 2) Next is the Login request. Replace it with your email and password
//      and pass the returned AuthHeader and Device ID
// =======================================================
email := "<< CHANGE ME >>"
password := "<< CHANGE ME >>"
loginRes, _ := c.Login(context.Background(), deviceId, email, password, *registerDeviceRes.Data.AuthHeader)
fmt.Println(loginRes.ACTSSOCOOKIE)

// =======================================================
// 3) Final step is to send a Get Gamer Match List.
//    You need to call Login() before using this authenticated request.
//	  The token is stored in the client and will be implicitly sent along each Authenticated request.
// =======================================================
// create Gamer struct - MUST CHANGE this to your own account
gamer := &go_cod.Gamer{
    Platform:   go_cod.Battlenet,
    LookupType: go_cod.BattlenetLookup,
    GamerTag:   "MrExcitement#6438",
}

gamerMatchListResp, _ := c.GamerMatchList(context.Background(), go_cod.ModernWarfare, gamer,
    go_cod.Multiplayer, 0, 0, 3)

fmt.Println(gamerMatchListResp.Status)
for _, v := range gamerMatchListResp.Data {
    fmt.Printf("\tMap = %v\n", v.Map)
    fmt.Printf("\tMatchID = %v\n", v.MatchID)
    fmt.Printf("\tPlatform = %v\n", v.Platform)
    fmt.Printf("\tTimestamp = %v\n", v.Timestamp)
    fmt.Printf("\tTitle = %v\n", v.Title)
    fmt.Printf("\tType = %v\n\n", v.Type)
}
```

### Unmapped fields

Fields that are not mapped to a struct field can be accessed from the `*AdditionalProperties` field of type `map[string]interface{}` 

For example, the weapons values from the loadout request can be found under `Data.LoadoutResponseDataAdditionalProperties["weapons"]`
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

- [X] Populate response fields for Friend Stats Profile
- [X] Populate response fields for COD Points
- [X] Added facade to simplify API call process
- [ ] Add more example codes and helper functions (e.g. Get Uno ID...)
- [ ] Some fields can be mapped better. See [issue#3](https://github.com/carlocayos/go-cod/issues/3)
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

