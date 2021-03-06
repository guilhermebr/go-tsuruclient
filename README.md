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
*AppApi* | [**AppCreate**](docs/AppApi.md#appcreate) | **Post** /1.0/apps | 
*AppApi* | [**AppDelete**](docs/AppApi.md#appdelete) | **Delete** /1.0/apps/{app} | 
*AppApi* | [**AppGet**](docs/AppApi.md#appget) | **Get** /1.0/apps/{app} | 
*AppApi* | [**AppList**](docs/AppApi.md#applist) | **Get** /1.0/apps | 
*AppApi* | [**EnvGet**](docs/AppApi.md#envget) | **Get** /1.0/apps/{app}/env | 
*AppApi* | [**EnvSet**](docs/AppApi.md#envset) | **Post** /1.0/apps/{app}/env | 
*AppApi* | [**EnvUnset**](docs/AppApi.md#envunset) | **Delete** /1.0/apps/{app}/env | 
*AuthApi* | [**AssignRoleToToken**](docs/AuthApi.md#assignroletotoken) | **Post** /1.6/roles/{role_name}/token | 
*AuthApi* | [**DissociateRoleFromToken**](docs/AuthApi.md#dissociaterolefromtoken) | **Delete** /1.6/roles/{role_name}/token/{token_id} | 
*AuthApi* | [**TeamTokenCreate**](docs/AuthApi.md#teamtokencreate) | **Post** /1.6/tokens | 
*AuthApi* | [**TeamTokenDelete**](docs/AuthApi.md#teamtokendelete) | **Delete** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokenUpdate**](docs/AuthApi.md#teamtokenupdate) | **Put** /1.6/tokens/{token_id} | 
*AuthApi* | [**TeamTokensList**](docs/AuthApi.md#teamtokenslist) | **Get** /1.6/tokens | 
*EventApi* | [**EventCancel**](docs/EventApi.md#eventcancel) | **Post** /1.1/events/{eventid}/cancel | 
*EventApi* | [**WebhookCreate**](docs/EventApi.md#webhookcreate) | **Post** /1.6/events/webhooks | 
*EventApi* | [**WebhookDelete**](docs/EventApi.md#webhookdelete) | **Delete** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookGet**](docs/EventApi.md#webhookget) | **Get** /1.6/events/webhooks/{name} | 
*EventApi* | [**WebhookList**](docs/EventApi.md#webhooklist) | **Get** /1.6/events/webhooks | 
*EventApi* | [**WebhookUpdate**](docs/EventApi.md#webhookupdate) | **Put** /1.6/events/webhooks/{name} | 
*NodeApi* | [**NodeAdd**](docs/NodeApi.md#nodeadd) | **Post** /1.2/node | 
*NodeApi* | [**NodeDelete**](docs/NodeApi.md#nodedelete) | **Delete** /1.2/node/{address} | 
*NodeApi* | [**NodeGet**](docs/NodeApi.md#nodeget) | **Get** /1.2/node/{address} | 
*NodeApi* | [**NodeList**](docs/NodeApi.md#nodelist) | **Get** /1.2/node | 
*NodeApi* | [**NodeUpdate**](docs/NodeApi.md#nodeupdate) | **Put** /1.2/node | 
*PlatformApi* | [**PlatformAdd**](docs/PlatformApi.md#platformadd) | **Post** /1.0/platforms | 
*PlatformApi* | [**PlatformDelete**](docs/PlatformApi.md#platformdelete) | **Delete** /1.0/platforms/{platform} | 
*PlatformApi* | [**PlatformList**](docs/PlatformApi.md#platformlist) | **Get** /1.0/platforms | 
*PlatformApi* | [**PlatformUpdate**](docs/PlatformApi.md#platformupdate) | **Put** /1.0/platforms/{platform} | 
*PoolApi* | [**PoolCreate**](docs/PoolApi.md#poolcreate) | **Post** /1.0/pools | 
*PoolApi* | [**PoolDelete**](docs/PoolApi.md#pooldelete) | **Delete** /pools/{pool} | 
*PoolApi* | [**PoolList**](docs/PoolApi.md#poollist) | **Get** /1.0/pools | 
*PoolApi* | [**PoolUpdate**](docs/PoolApi.md#poolupdate) | **Put** /pools/{pool} | 
*ServiceApi* | [**InstanceDelete**](docs/ServiceApi.md#instancedelete) | **Delete** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstanceGet**](docs/ServiceApi.md#instanceget) | **Get** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstanceUpdate**](docs/ServiceApi.md#instanceupdate) | **Put** /1.0/services/{service}/instances/{instance} | 
*ServiceApi* | [**InstancesList**](docs/ServiceApi.md#instanceslist) | **Get** /1.0/services/instances | 
*ServiceApi* | [**ServiceBrokerCreate**](docs/ServiceApi.md#servicebrokercreate) | **Post** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerDelete**](docs/ServiceApi.md#servicebrokerdelete) | **Delete** /1.7/brokers/{name} | 
*ServiceApi* | [**ServiceBrokerList**](docs/ServiceApi.md#servicebrokerlist) | **Get** /1.7/brokers | 
*ServiceApi* | [**ServiceBrokerUpdate**](docs/ServiceApi.md#servicebrokerupdate) | **Put** /1.7/brokers/{name} | 
*ServiceApi* | [**ServicesList**](docs/ServiceApi.md#serviceslist) | **Get** /1.0/services | 
*TeamApi* | [**TeamCreate**](docs/TeamApi.md#teamcreate) | **Post** /1.0/teams | 
*TeamApi* | [**TeamDelete**](docs/TeamApi.md#teamdelete) | **Delete** /1.0/teams/{team} | 
*TeamApi* | [**TeamGet**](docs/TeamApi.md#teamget) | **Get** /1.4/teams/{team} | 
*TeamApi* | [**TeamUpdate**](docs/TeamApi.md#teamupdate) | **Put** /1.6/teams/{team} | 
*TeamApi* | [**TeamsList**](docs/TeamApi.md#teamslist) | **Get** /1.0/teams | 
*UserApi* | [**APITokenGet**](docs/UserApi.md#apitokenget) | **Get** /1.0/users/api-key | 
*UserApi* | [**APITokenRegenerate**](docs/UserApi.md#apitokenregenerate) | **Post** /1.0/users/api-key | 
*UserApi* | [**ChangePassword**](docs/UserApi.md#changepassword) | **Put** /1.0/users/password | 
*UserApi* | [**ResetPassword**](docs/UserApi.md#resetpassword) | **Post** /1.0/users/{email}/password | 
*UserApi* | [**SSHKeyAdd**](docs/UserApi.md#sshkeyadd) | **Post** /1.0/users/keys | 
*UserApi* | [**SSHKeyList**](docs/UserApi.md#sshkeylist) | **Get** /1.0/users/keys | 
*UserApi* | [**SSHKeyRemove**](docs/UserApi.md#sshkeyremove) | **Delete** /1.0/users/keys/{key} | 
*UserApi* | [**UserCreate**](docs/UserApi.md#usercreate) | **Post** /1.0/users | 
*UserApi* | [**UserDelete**](docs/UserApi.md#userdelete) | **Delete** /1.0/users | 
*UserApi* | [**UserGet**](docs/UserApi.md#userget) | **Get** /1.0/users/info | 
*UserApi* | [**UserQuotaChange**](docs/UserApi.md#userquotachange) | **Put** /1.0/users/{email}/quota | 
*UserApi* | [**UserQuotaGet**](docs/UserApi.md#userquotaget) | **Get** /1.0/users/{email}/quota | 
*UserApi* | [**UserTokenDelete**](docs/UserApi.md#usertokendelete) | **Delete** /1.0/users/tokens | 
*UserApi* | [**UsersList**](docs/UserApi.md#userslist) | **Get** /1.0/users | 
*VolumeApi* | [**VolumeBind**](docs/VolumeApi.md#volumebind) | **Post** /1.4/volumes/{volume}/bind | 
*VolumeApi* | [**VolumeCreate**](docs/VolumeApi.md#volumecreate) | **Post** /1.4/volumes | 
*VolumeApi* | [**VolumeDelete**](docs/VolumeApi.md#volumedelete) | **Delete** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeGet**](docs/VolumeApi.md#volumeget) | **Get** /1.4/volumes/{volume} | 
*VolumeApi* | [**VolumeList**](docs/VolumeApi.md#volumelist) | **Get** /1.4/volumes | 
*VolumeApi* | [**VolumePlansList**](docs/VolumeApi.md#volumeplanslist) | **Get** /1.4/volumeplans | 
*VolumeApi* | [**VolumeUnbind**](docs/VolumeApi.md#volumeunbind) | **Delete** /1.4/volumes/{volume}/bind | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppCreateResponse](docs/AppCreateResponse.md)
 - [AssignTokenArgs](docs/AssignTokenArgs.md)
 - [ChangePasswordData](docs/ChangePasswordData.md)
 - [Env](docs/Env.md)
 - [EnvSetData](docs/EnvSetData.md)
 - [EnvSetResponse](docs/EnvSetResponse.md)
 - [EnvSetResponseInner](docs/EnvSetResponseInner.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [EventCancelArgs](docs/EventCancelArgs.md)
 - [Lock](docs/Lock.md)
 - [Machine](docs/Machine.md)
 - [MiniApp](docs/MiniApp.md)
 - [Node](docs/Node.md)
 - [NodeAddData](docs/NodeAddData.md)
 - [NodeCheck](docs/NodeCheck.md)
 - [NodeCheckResult](docs/NodeCheckResult.md)
 - [NodeGetResponse](docs/NodeGetResponse.md)
 - [NodeListResponse](docs/NodeListResponse.md)
 - [NodeStatus](docs/NodeStatus.md)
 - [NodeUpdateData](docs/NodeUpdateData.md)
 - [PermissionUser](docs/PermissionUser.md)
 - [Plan](docs/Plan.md)
 - [Platform](docs/Platform.md)
 - [Pool](docs/Pool.md)
 - [PoolCreateData](docs/PoolCreateData.md)
 - [PoolUpdateData](docs/PoolUpdateData.md)
 - [RoleInstance](docs/RoleInstance.md)
 - [RoleUser](docs/RoleUser.md)
 - [Router](docs/Router.md)
 - [Service](docs/Service.md)
 - [ServiceBroker](docs/ServiceBroker.md)
 - [ServiceBrokerConfig](docs/ServiceBrokerConfig.md)
 - [ServiceBrokerConfigAuthConfig](docs/ServiceBrokerConfigAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBasicAuthConfig](docs/ServiceBrokerConfigAuthConfigBasicAuthConfig.md)
 - [ServiceBrokerConfigAuthConfigBearerConfig](docs/ServiceBrokerConfigAuthConfigBearerConfig.md)
 - [ServiceBrokerList](docs/ServiceBrokerList.md)
 - [ServiceInstance](docs/ServiceInstance.md)
 - [ServiceInstanceBoundUnit](docs/ServiceInstanceBoundUnit.md)
 - [ServiceInstanceInfo](docs/ServiceInstanceInfo.md)
 - [ServiceInstanceUpdateData](docs/ServiceInstanceUpdateData.md)
 - [SshKeyAddData](docs/SshKeyAddData.md)
 - [SshKeyListResponse](docs/SshKeyListResponse.md)
 - [Team](docs/Team.md)
 - [TeamData](docs/TeamData.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamToken](docs/TeamToken.md)
 - [TeamTokenCreateArgs](docs/TeamTokenCreateArgs.md)
 - [TeamTokenUpdateArgs](docs/TeamTokenUpdateArgs.md)
 - [Unit](docs/Unit.md)
 - [UpdateData](docs/UpdateData.md)
 - [Url](docs/Url.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)
 - [UserQuotaViewResponse](docs/UserQuotaViewResponse.md)
 - [Volume](docs/Volume.md)
 - [VolumeBind](docs/VolumeBind.md)
 - [VolumeBindData](docs/VolumeBindData.md)
 - [VolumeBindId](docs/VolumeBindId.md)
 - [VolumePlan](docs/VolumePlan.md)
 - [VolumePlansListResponse](docs/VolumePlansListResponse.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookEventFilter](docs/WebhookEventFilter.md)


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



