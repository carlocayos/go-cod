# go-cod

go-cod is the _unofficial_ Go SDK of the Call of Duty API used in the callofduty.com site

- [Getting Started](#getting-started)
    * [Authentication](#authentication)
    * [Unmapped fields](#unmapped-fields)
- [Sample codes and Payload](#sample-codes-and-payload)
- [Generate Models and Client from Swagger Spec](#generate-models-and-client-from-swagger-spec)
    * [Generating client and models](#generating-client-and-models)
    * [API Documentation](#api-documentation)
- [Roadmap](#roadmap)
- [Request for missing APIs and fields](#request-for-missing-apis-and-fields)
- [Contribution](#contribution)
- [Credits](#credits)
- [Developer](#developer)

## Getting Started
```
import "github.com/carlocayos/go-cod/v2"
```

Create a new go-cod client then use the services to access the COD APIs.

A short example:
```go
// create the client
c := go_cod.NewClient(nil)

// get leader board list
leaderBoardResp, _ := c.LeaderBoard(context.Background(), 
    go_cod.ModernWarfare, go_cod.Battlenet, 3)
fmt.Println(leaderBoardResp.Status)
fmt.Println(leaderBoardResp.Data.Title)
```

The cod client is composed of an authentication client and service client.

### Authentication

Authentication client is for getting a security token that will be used for most of the API requests.

See example on how to login and send authenticated requests.

```go
// =======================================================
// 1) First, is creating a client and sending a register device request
// =======================================================
c := go_cod.NewClient(nil)

// send a register device request with a unique device id
// ksuid is used here to generate a unique id, but any uid generator would be fine
deviceId := ksuid.New().String()
registerDeviceRes, _ := c.RegisterDevice(context.Background(), deviceId)

// =======================================================
// 2) Next is the Login request. Replace it with your email and password
//    and pass the returned AuthHeader and Device ID
// =======================================================
email := "<< CHANGE ME >>"
password := "<< CHANGE ME >>"
loginRes, _ := c.Login(context.Background(), deviceId, email, 
    password, *registerDeviceRes.Data.AuthHeader)
fmt.Println(loginRes.ACTSSOCOOKIE)

// =======================================================
// 3) Final step is to send a Get Gamer Match List.
//    You need to call Login() before using this authenticated request.
//    The token is stored in the client and will be implicitly sent along 
//    each Authenticated request.
// =======================================================
// create Gamer struct - MUST CHANGE this to your own account
gamer := &go_cod.Gamer{
    Platform:   go_cod.Battlenet,
    LookupType: go_cod.BattlenetLookup,
    GamerTag:   "MrExcitement#6438",
}

gamerMatchListResp, _ := c.GamerMatchList(context.Background(), 
    go_cod.ModernWarfare, gamer, go_cod.Multiplayer, 0, 0, 3)

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

As there is no official API documentation, the swagger spec and JSON schema mapping is inferred through the 
actual JSON response payload.

Fields that are not mapped to a struct field can be accessed from the `*AdditionalProperties` field of type `map[string]interface{}`

```go
missingField := response.Data.SampleResponseDataAdditionalProperties["missing_field_name"].(map[string]interface{})
anotherMissingField := missingField["another_missing_field"]
fmt.Printf("anotherMissingField")
```

See [How to request missing APIs and fields](#request-for-missing-apis-and-fields)

## Sample codes and Payload

Sample codes are in [examples](examples)

Actual JSON response payload `*_sample.json` can be found in [json-schema](api/specs/v1.0.0/json-schema)

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

## Roadmap

- [X] ~~Added facade to simplify API call process~~
- [X] ~~Improve field mapping for Friend Stats API Response~~
- [X] ~~Improve field mapping for Gamer Stats API Response~~
- [X] ~~Improve field mapping for Battlepass Loot API Response~~
- [X] ~~Improve field mapping for Match Details API Response~~
- [X] ~~Improve field mapping for Loadout API Response~~
- [X] ~~Improve field mapping for COD Points API Response~~
- [X] ~~Improve field mapping for Purchasable API Response~~
- [X] ~~Imoproved field mapping.~~ See [issue#3](https://github.com/carlocayos/go-cod/issues/3)
- [ ] Add more example codes and helper functions (e.g. Get Uno ID...)
- [ ] Context handling

## Request for missing APIs and fields

Report new or missing COD APIs and fields [here](https://github.com/carlocayos/go-cod/issues/new?assignees=&labels=enhancement&template=add-api-or-field.md&title=Reques+to+add+a+new+field+or+API)

## Contribution
Code improvements, suggestions, and updates are most welcome. Please feel free to raise an issue or create a pull
request for your changes. 🙂

There is no official COD API and documentation released by Activision. If there is a breaking change on the COD API 
then let me know, so I can update this project. 

## Credits

Thanks to [Lierrmm](https://github.com/Lierrmm) for his work on the [NodeJS Call of Duty API Wrapper](https://github.com/Lierrmm/Node-CallOfDuty)

## Developer

**Personal Site:** [carlocayos.com](https://carlocayos.com)

<a href="https://www.buymeacoffee.com/ccayos" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

**BTC address:** 32zunH725N7PjBYj2TfbVoC3jVCyhqyn5h

**ETH address:** 0x2A17e4031FFeF64C638Dd9B190e05a150b2B8FBc
