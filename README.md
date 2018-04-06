# Go API client for tsuru

Open source, extensible and Docker-based Platform as a Service (PaaS)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.6
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./tsuru"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppApi* | [**AppList**](docs/AppApi.md#applist) | **Get** /1.0/apps | 
*AppApi* | [**Create**](docs/AppApi.md#create) | **Post** /1.0/apps | 
*AppApi* | [**EnvSet**](docs/AppApi.md#envset) | **Post** /1.0/apps/{app}/env | 
*VolumeApi* | [**VolumeBind**](docs/VolumeApi.md#volumebind) | **Post** /1.4/volumes/{volume}/bind | 
*VolumeApi* | [**VolumeCreate**](docs/VolumeApi.md#volumecreate) | **Post** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeDelete**](docs/VolumeApi.md#volumedelete) | **Delete** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeGet**](docs/VolumeApi.md#volumeget) | **Get** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeList**](docs/VolumeApi.md#volumelist) | **Get** /1.4/volumes | 
*VolumeApi* | [**VolumePlansList**](docs/VolumeApi.md#volumeplanslist) | **Get** /1.4/volumeplans | 
*VolumeApi* | [**VolumeUnbind**](docs/VolumeApi.md#volumeunbind) | **Delete** /1.4/volumes/{volume}/bind | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppCreateResponse](docs/AppCreateResponse.md)
 - [Env](docs/Env.md)
 - [EnvSetData](docs/EnvSetData.md)
 - [EnvSetResponse](docs/EnvSetResponse.md)
 - [EnvSetResponseInner](docs/EnvSetResponseInner.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [Lock](docs/Lock.md)
 - [MiniApp](docs/MiniApp.md)
 - [Plan](docs/Plan.md)
 - [Router](docs/Router.md)
 - [Unit](docs/Unit.md)
 - [Url](docs/Url.md)
 - [Volume](docs/Volume.md)
 - [VolumeBind](docs/VolumeBind.md)
 - [VolumeBindData](docs/VolumeBindData.md)
 - [VolumeBindId](docs/VolumeBindId.md)
 - [VolumeListResponse](docs/VolumeListResponse.md)
 - [VolumePlan](docs/VolumePlan.md)
 - [VolumePlansListResponse](docs/VolumePlansListResponse.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```

## Author



