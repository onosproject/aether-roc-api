<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation

SPDX-License-Identifier: Apache-2.0
-->

## Code generation from Aether Model API specification

Calling on the `makefile` target:

```bash
make oapi-codegen-aether-2.1.0 
```

has the effect of first validating that the OpenAPI 3 specifications contains no errors.

It then proceeds to generate Go code that performs the following functions:

* the types from the OpenAPI spec
* the spec (so that the API specification can be retrieved by clients programmatically)
* the paths and methods from the OpenAPI spec
* conversion code to go from OpenAPI types to YANG model types
* conversion code to go from YANG model types to OpenAPI types
* code to call gNMI methods Get() and Set() from REST methods GET, POST, DELETE
* code to integrate with the [Echo Web Server](https://echo.labstack.com/) built in to this project that serves the API

Breaking the steps down individually:

## Generating Go types
The first step generates [Go type definitions](../pkg/aether_2_1_0/types/aether-2.1.0-types.go) from 
the Schema definitions in the specification:

```bash
oapi-codegen -generate types -package types -o pkg/aether_2_1_0/types/aether-2.1.0-types.go api/aether-2.1.0-openapi3.yaml
```

## Encoding specification in Go
The second run bundles up the API specification as a [blob in base64](../pkg/aether_2_1_0/server/aether-2.1.0-spec.go),
so that it may be extracted programmatically through the REST interface.

```bash
oapi-codegen -generate spec -package server -o pkg/aether_2_1_0/server/aether-2.1.0-spec.go api/aether-2.1.0-openapi3.yaml
```

## Generating implementation of handlers

The third step generates [an implementation of handlers](../pkg/aether_2_1_0/server/aether-2.1.0-impl.go) for accessing
the gNMI client.

```go
    type ServerImpl struct {
        GnmiClient  southbound.GnmiClient
        GnmiTimeout time.Duration
    }
```

and creates a **public** and a **private** method for each of the **paths** in the specification. e.g.

```go
func (i *ServerImpl) GetApplication(ctx echo.Context, enterpriseId externalRef1.EnterpriseId, applicationId string) error
```

> These public functions are an implementation of the `ServerInterface` interface generated in the next step

```go
func (i *ServerImpl) gnmiGetApplication(ctx context.Context,
	openApiPath string, enterpriseId externalRef1.EnterpriseId, args ...string) (*externalRef1.Application, error)
```

The public method is an implementation of a Handler from the 'server' file.

```bash
oapi-codegen \
    -generate types,server \
    -import-mapping imp1:"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types",imp2:"github.com/onosproject/aether-models/models/aether-2.1.x/api" \
    -package server \
    -templates pkg/codegen/templates \
    -o pkg/aether_2_1_0/server/aether-2.1.0-impl.go \
    api/aether-2.1.0-openapi3.yaml
```

> This call does not use the built in template of `oapi-codegen` but custom templates from
> the [codegen/templates](../pkg/codegen/templates) folder.
> 
> Also note the `-import-mapping` here - it is used to force the set of Go imports to include these paths.
> 
> Counterintuitively, the prefixes given here e.g. `imp1` are discarded by oapi-codegen, and imports are ordered
> alphabetically, being given the prefixes `externalRef0` etc inside the code

To correct an omission of an import, a `sed` statement is necessary to fix the file after generation.

> The `sed` statement here is written with Posix compliance, so that it can run on linux and MacOS.

## Generate Handlers

The fourth step generates [the handlers](../pkg/aether_2_1_0/server/aether-2.1.0-server.go) for the Echo Web server.

The `RegisterHandlers` func gives a list of all the paths and their methods and their respective handler functions on
the `serverInterfaceWrapper` struct.

These methods pass only an **echo.Context** argument (as all Echo handlers do) like:

```go
func (w *serverInterfaceWrapper) GetApplication(ctx echo.Context) error
```

so this function extracts parameters like `enterprise-id` and `application-id` from this Context and passes
to another level of function defined in the `ServerInterface` interface:

```go
// ServerInterface represents all server handlers.
type ServerInterface interface {
    // GET /application Container
    // (GET /aether/v2.1.x/{enterprise-id}/application/{application-id})
    GetApplication(ctx echo.Context, enterpriseId externalRef0.EnterpriseId, applicationId string) error
...
```

It is this interface that the methods in the `impl` step above implement.

```bash
oapi-codegen \
    -generate server \
    -import-mapping externalRef0:"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types" \
    -package server \
    -templates pkg/codegen/modified \
    -o pkg/aether_2_1_0/server/aether-2.1.0-server.go \
    api/aether-2.1.0-openapi3.yaml
```

> Note that here again the templates in use are customized ones from [codegen/modified](../pkg/codegen/modified)
> 
> The same kind of `sed` command is used to correct some imports

## Generate OpenAPI schema object structs to gNMI object structs mapping

The fifth step generates code that can transform structs generated from OpenAPI `types` (from the first steps above),
to structs generated by YGOT from the YANG model, during the `aether-models` generation stage.

The YGOT structs can then be serialized to JSON and passed in gNMI calls as values.

It creates a function per OpenAPI type - e.g. for Application

```go
func EncodeToGnmiApplication(
	jsonObj *types.Application, needKey bool, removeIndex bool, enterpriseId types.EnterpriseId, parentPath string, params ...string) (
	[]*gnmi.Update, error) {
```

> In the template `pkg/codegen/convert-oapi-gnmi/typedef.tmpl` there is a variable `$tnbase` which may have white space
> padding at the left hand side - the template syntax `{{-` does not correctly get rid of this leading white space
> properly unless it is on a new line.
> To make it distinguishable we add in the marker ` > !` on the new line, so it can be removed easily later with `sed`.

```bash
oapi-codegen \
    -generate types \
    -import-mapping externalRef0:"github.com/onosproject/aether-models/models/aether-2.1.x/api" \
    -package server \
    -templates pkg/codegen/convert-oapi-gnmi \
    -o pkg/aether_2_1_0/server/aether-2.1.0-convert-oapi-gnmi.go \
    api/aether-2.1.0-openapi3.yaml
```

> Note that here the templates in use are customized ones from [codegen/convert-oapi-gnmi](../pkg/codegen/convert-oapi-gnmi)
>
> The `sed` command used here is to format a particular hack that is used in the templates.
> 
> Specifically it is removing a ` > !` inserted in the code generation

## Generate gNMI object structs to OpenAPI schema object structs mapping

The sixth step generates code that can transform structs created from YGOT (from the YANG model) to OpenAPI structs.

The YGOT structs are unmarshalled from the gNMI payloads (in JSON encoding) coming back in to the system in reply to
gNMI GetRequests.

The template generates two different kinds of struct.

1. with a pointer receiver to a YGOT parent object and
2. a second that's a plain function that breaks apart the specific subsection of the YGOT object

An example of 1. is:

```go
func (d *ModelPluginDevice) toApplication(params ...string) (*types.Application, error)
```

and an example of 2. is:

```go
func toApplication(ygotObjValue *reflect.Value, params ...string) (*types.Application, error)
```

Here the generated code relies on utility functions that can walk the Go struct hierarchy e.g.
* utils.FindModelPluginObject

```bash
oapi-codegen \
    -generate types \
    -import-mapping externalRef0:"github.com/onosproject/aether-models/models/aether-2.1.x/api" \
    -package server \
    -templates pkg/codegen/convert-gnmi-oapi \
    -o pkg/aether_2_1_0/server/aether-2.1.0-convert-gnmi-oapi.go \
    api/aether-2.1.0-openapi3.yaml
```

> Note that here the templates in use are customized ones from [codegen/convert-gnmi-oapi](../pkg/codegen/convert-gnmi-oapi)
> 
> The `sed` command used here is to format a particular hack that is used in the templates.
> 
> Specifically it is removing a ` > !` inserted in the code generation