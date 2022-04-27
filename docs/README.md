<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation

SPDX-License-Identifier: Apache-2.0
-->
# Project structure and Tooling

The Aether ROC API consists of

* one or more Aether model driven REST APIs 
* a Top Level API,
* an Application Gateway API

all defined as OpenAPI 3 specifications and implemented as Go code.

## Aether model API

The Aether Model API specifications are generated from the original Aether YANG models in
[../../aether-models] using the Model Plugin Compiler. For example 
for [Aether 2.1.x](https://github.com/onosproject/aether-models/blob/master/models/aether-2.1.x/openapi.yaml)

> For each of the leafs, containers and lists in YANG there is a corresponding PATH in the REST model, each with a
> GET, POST and DELETE. Additionally, for lists there is another Path to retrieve all instances, which has only
> a GET method. 

Once generated, they can be simply copied across here in to the [../api](../api) folder and named appropriately.

Once in-place, the corresponding code generation can be run to create Go code that implements:
* the types from the OpenAPI spec
* the spec (so that the API specification can be retrieved by clients programmatically)
* the paths and methods from the OpenAPI spec
* conversion code to go from OpenAPI types to YANG model types
* conversion code to go from YANG model types to OpenAPI types
* code to call gNMI methods Get() and Set() from REST methods GET, POST, DELETE
* code to integrate with the [Echo Web Server](https://echo.labstack.com/) built in to this project that serves the API

The code generation is performed through [Deepmap OAPI Codegen](https://github.com/deepmap/oapi-codegen) tool.
This takes the OpenAPI 3 specification and generates Go code according to templates, defined
in the format of Go's [text/template](https://pkg.go.dev/text/template).

The complete detail of how these templates are structured is in the [code-generation.md](./code-generation.md) page.

## Aether Model changes

If for example you have an updated version of the [aether-2.1.0-openapi3.yaml](../api/aether-2.1.0-openapi3.yaml),
it is possible to re-generate all the corresponding Go code by running:

```bash
make oapi-codegen-aether-2.1.0
```

## Top level API

The [Top Level API](../api/aether-top-level-openapi3.yaml) contains some convenience functions that give further
convenience methods on top of the Aether Model APIs.

* A method `GET /transactions` to get the list of transactions from the underlying onos-config
* A method `GET /targets` to get the list of `enterprises` stored in `onos-topo`
* A method `POST /sdcore/synchronize/` to call on the `sdcore-adapter` to synchronize its contents down to the SD-CORE
* A method `GET /spec` to give access to the OpenAPI specification programmatically
* A method `PATCH /aether-roc-api` to allow multiple configuration elements of the Aether model API to be 
    written as one transaction
  * To offer this method, the Aether Model APIs are referenced in the top level spec using a syntax like 
    `$ref: './aether-2.1.0-openapi3.yaml#/components/schemas/Site_List'`

The Top Level API implementation is only *partially* code-generated. Running the command:
```bash
make aether-top-level
```

will generate code for the:
* types defined in the Top Level API schema
* the spec (so that the API specification can be retrieved by clients programmatically)

The remaining code in [toplevel](../pkg/toplevel) is written by hand, and so must be manually updated if
any change is made to this part of the API.

## Application Gateway API

The [Application Gateway API](../api/aether-app-gtwy-openapi3.yaml) allows retrieval of Device status attributes.

Devices can be retrieved:

* individually by Enterprise, Site and Device Id or
* by Enterprise and Site alone, returning all device in scope

The Application Gateway API is fully generated from code. Running the command

```bash
make 
```