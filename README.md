# Go API client for douyin

douyin open api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./douyin"
```

## Documentation for API Endpoints

All URIs are relative to *https://open.douyin.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OauthApi* | [**Connect**](docs/OauthApi.md#connect) | **Get** /platform/oauth/connect | 获取授权码(code)


## Documentation For Models

 - [ApiResponse](docs/ApiResponse.md)
 - [ApiResponseData](docs/ApiResponseData.md)
 - [ApiResponseExtra](docs/ApiResponseExtra.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



