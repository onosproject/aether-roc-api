// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	configapi "github.com/onosproject/onos-api/go/onos/config/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/gnmi/proto/gnmi_ext"
	"github.com/openconfig/ygot/ygot"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	splitCapsAndNums = regexp.MustCompile(`([A-Z]|[0-9]*)[^0-9A-Z]*`)
	splitNumsThenCap = regexp.MustCompile(`([A-Z]|([0-9]+[A-Z])|([0-9]+[a-z]*))[^0-9A-Z]*`)
	splitYgotStruct  = regexp.MustCompile(`(_[0-9]+[A-Z])[^A-Z_]*`) // underscore followed by one or more digits, followed by a single capital letter
	log              = logging.GetLogger("gnmi_utils")
)

// NewGnmiGetRequest creates a GetRequest from a REST call
func NewGnmiGetRequest(openapiPath string, target string, pathParams ...string) (*gnmi.GetRequest, error) {
	gnmiGet := new(gnmi.GetRequest)
	gnmiGet.Path = make([]*gnmi.Path, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}

	gnmiGet.Path[0] = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	return gnmiGet, nil
}

// GetResponseUpdate -- extract the single Update from the GetResponse
func GetResponseUpdate(gr *gnmi.GetResponse, err error) (*gnmi.TypedValue, error) {
	if err != nil {
		return nil, err
	}
	if len(gr.Notification) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notifications %d", len(gr.Notification))
	}
	n0 := gr.Notification[0]
	if len(n0.Update) != 1 {
		return nil, fmt.Errorf("unexpected number of GetResponse notification updates %d", len(n0.Update))
	}
	u0 := n0.Update[0]
	if u0.Val == nil {
		return nil, nil
	}
	return &gnmi.TypedValue{
		Value: u0.Val.Value,
	}, nil
}

// NewGnmiSetDeleteRequest a single delete in a Set request
func NewGnmiSetDeleteRequest(openapiPath string, target string, pathParams ...string) (*gnmi.SetRequest, error) {
	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Delete = make([]*gnmi.Path, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}

	gnmiSet.Delete[0] = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	return gnmiSet, nil
}

// NewGnmiSetUpdateRequestUpdates a single update in a Set request
func NewGnmiSetUpdateRequestUpdates(openapiPath string, target string,
	update []*gnmi.Update, pathParams ...string) (*gnmi.SetRequest, error) {

	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Extension = buildExtensions(openapiPath)
	gnmiSet.Update = make([]*gnmi.Update, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}
	gnmiSet.Prefix = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	gnmiSet.Update = update

	return gnmiSet, nil
}

// NewGnmiSetUpdateRequestUpdatesDeletes Set request with update and delete
func NewGnmiSetUpdateRequestUpdatesDeletes(openapiPath string, target string,
	update []*gnmi.Update, delete []*gnmi.Path, pathParams ...string) (*gnmi.SetRequest, error) {

	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Extension = buildExtensions(openapiPath)
	gnmiSet.Update = make([]*gnmi.Update, 1)
	gnmiSet.Delete = make([]*gnmi.Path, 1)
	elems, err := BuildElems(openapiPath, 4, pathParams...)
	if err != nil {
		return nil, fmt.Errorf("error creating new update set request %v", err)
	}
	gnmiSet.Prefix = &gnmi.Path{
		Elem:   elems,
		Target: target,
	}
	gnmiSet.Update = update
	gnmiSet.Delete = delete

	return gnmiSet, nil
}

// NewGnmiSetRequest -- new set request including updates and deletes
func NewGnmiSetRequest(updates []*gnmi.Update, deletes []*gnmi.Path,
	ext100Name *string, ext101Version *string, ext102Type *string,
	ext111Strategy *int) (*gnmi.SetRequest, error) {
	gnmiSet := new(gnmi.SetRequest)
	gnmiSet.Update = updates
	gnmiSet.Delete = deletes

	gnmiSet.Extension = make([]*gnmi_ext.Extension, 0)
	if ext100Name != nil {
		gnmiSet.Extension = append(gnmiSet.Extension, &gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  100,
					Msg: []byte(*ext100Name),
				},
			},
		})
	}
	if ext101Version != nil {
		gnmiSet.Extension = append(gnmiSet.Extension, &gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  101,
					Msg: []byte(*ext101Version),
				},
			},
		})
	}
	if ext102Type != nil {
		gnmiSet.Extension = append(gnmiSet.Extension, &gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  102,
					Msg: []byte(*ext102Type),
				},
			},
		})
	}

	if ext111Strategy != nil {
		ext := configapi.TransactionStrategy{
			Synchronicity: configapi.TransactionStrategy_Synchronicity(uint32(*ext111Strategy)),
		}
		b, err := ext.Marshal()
		if err != nil {
			return nil, err
		}
		gnmiSet.Extension = append(gnmiSet.Extension, &gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  111,
					Msg: b,
				},
			},
		})
	}

	return gnmiSet, nil
}

// ExtractResponseID - the name of the change will be returned as extension 100
func ExtractResponseID(gnmiResponse *gnmi.SetResponse) (*string, error) {
	for _, ext := range gnmiResponse.Extension {
		switch extTyped := ext.Ext.(type) {
		case *gnmi_ext.Extension_RegisteredExt:
			// NOTE this is used in ONOS config
			if extTyped.RegisteredExt.Id == 100 {
				changeName := string(extTyped.RegisteredExt.Msg)
				return &changeName, nil
			}
			if extTyped.RegisteredExt.Id == configapi.TransactionInfoExtensionID {
				bytes := extTyped.RegisteredExt.Msg
				transactionInfo := &configapi.TransactionInfo{}
				err := proto.Unmarshal(bytes, transactionInfo)
				if err != nil {
					log.Errorw("cannot unmarshal transactionInfo", "err", err)
					return nil, err
				}
				changeName := string(transactionInfo.ID)
				return &changeName, nil
			}
		}
	}

	return nil, fmt.Errorf("cannot find transaction ID in response")
}

// BuildElems - create a set of gnmi PathElems
// For start at this is the element in the path at the ith position (remembering that 0 is empty)
func BuildElems(openapiPath string, startAt int, pathParams ...string) ([]*gnmi.PathElem, error) {
	if !strings.HasPrefix(openapiPath, "/") {
		return nil, fmt.Errorf("openapipath must begin with '/'. Got %s", openapiPath)
	}
	oapiParts := strings.Split(openapiPath, "/")
	if len(oapiParts) < startAt+1 {
		return nil, fmt.Errorf("expected path to have >= %d parts e.g. api,ver,device,path Got %v", startAt, oapiParts)
	}
	elemCount := 0
	paramCount := 0
	elems := make([]*gnmi.PathElem, 0)

	for i := startAt; i < len(oapiParts); i++ {
		if strings.Contains(oapiParts[i], "{") { // Is a key
			keyName := oapiParts[i]
			keyName = keyName[1 : len(keyName)-1]
			if elems[elemCount-1].Key == nil {
				elems[elemCount-1].Key = make(map[string]string)
			}
			elems[elemCount-1].Key[keyName] = pathParams[paramCount]
			paramCount++
		} else {
			pathElem := gnmi.PathElem{
				Name: oapiParts[i],
			}
			elems = append(elems, &pathElem)
			elemCount++
		}
	}

	return elems, nil
}

func buildExtensions(openapiPath string) []*gnmi_ext.Extension {
	oapiParts := strings.Split(openapiPath, "/")
	if len(oapiParts) < 3 {
		return nil
	}
	// First 2 fields should give us the modelType and modelVersion
	caser := cases.Title(language.English)
	modelType := caser.String(oapiParts[1]) // Change to title case
	modelVersion := oapiParts[2][1:]        // Remove the "v" at the start of "v1.0.0"

	extensions := make([]*gnmi_ext.Extension, 0)

	// always make direct rest call synchronous transactions
	ext := configapi.TransactionStrategy{
		Synchronicity: configapi.TransactionStrategy_SYNCHRONOUS,
	}
	b, err := ext.Marshal()
	if err != nil {
		log.Error(err)
	} else {
		ext111 := gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  111,
					Msg: b,
				},
			},
		}
		extensions = append(extensions, &ext111)
	}

	if modelVersion != "" {
		ext101 := gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  101,
					Msg: []byte(modelVersion),
				},
			},
		}
		extensions = append(extensions, &ext101)
	}
	if modelType != "" {
		ext102 := gnmi_ext.Extension{
			Ext: &gnmi_ext.Extension_RegisteredExt{
				RegisteredExt: &gnmi_ext.RegisteredExtension{
					Id:  102,
					Msg: []byte(modelType),
				},
			},
		}
		extensions = append(extensions, &ext102)
	}
	return extensions
}

// DeleteForElement -- create a gnmi.Delete for a Json element
func DeleteForElement(path string, pathParams ...string) (*gnmi.Path, error) {
	gnmiDelete := new(gnmi.Path)
	var err error
	if gnmiDelete.Elem, err = BuildElems(path, 1, pathParams...); err != nil {
		return nil, err
	}
	return gnmiDelete, nil
}

// UpdateForElement -- create a gnmi.Update for a Json element
func UpdateForElement(value interface{}, path string, pathParams ...string) (*gnmi.Update, error) {
	reflectValue := reflect.ValueOf(value)
	update := new(gnmi.Update)
	update.Path = new(gnmi.Path)
	var err error
	if update.Path.Elem, err = BuildElems(path, 1, pathParams...); err != nil {
		return nil, err
	}
	update.Val = new(gnmi.TypedValue)
	switch reflectValue.Type().String() {
	case "*string":
		update.Val.Value = &gnmi.TypedValue_StringVal{StringVal: reflect.Indirect(reflectValue).String()}
	case "*uint8", "*uint16", "*uint32", "*uint64":
		update.Val.Value = &gnmi.TypedValue_UintVal{UintVal: reflect.Indirect(reflectValue).Uint()}
	case "*int8", "*int16", "*int32", "*int64":
		update.Val.Value = &gnmi.TypedValue_IntVal{IntVal: reflect.Indirect(reflectValue).Int()}
	case "*bool":
		update.Val.Value = &gnmi.TypedValue_BoolVal{BoolVal: reflect.Indirect(reflectValue).Bool()}
	default:
		switch reflectValue.Kind() {
		case reflect.Int64:
			update.Val.Value = &gnmi.TypedValue_IntVal{IntVal: reflect.Indirect(reflectValue).Int()}
		case reflect.Ptr:
			update.Val, err = extractEnum(reflectValue.Elem())
			if err != nil {
				return nil, err
			}
		case reflect.Slice:
			leafListVal := &gnmi.TypedValue_LeaflistVal{
				LeaflistVal: new(gnmi.ScalarArray),
			}
			for i := 0; i < reflectValue.Len(); i++ {
				val, err := extractEnum(reflectValue.Index(i))
				if err != nil && err.Error() == "no enum method" {
					// could be string, integer, bool
					switch itemKind := reflectValue.Index(i).Kind(); itemKind {
					case reflect.String:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_StringVal{
								StringVal: reflectValue.Index(i).Interface().(string),
							},
						}
					case reflect.Uint8:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{
								UintVal: uint64(reflectValue.Index(i).Interface().(uint8)),
							},
						}
					case reflect.Uint16:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{
								UintVal: uint64(reflectValue.Index(i).Interface().(uint16)),
							},
						}
					case reflect.Uint32:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{
								UintVal: uint64(reflectValue.Index(i).Interface().(uint32)),
							},
						}
					case reflect.Uint64:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_UintVal{
								UintVal: reflectValue.Index(i).Interface().(uint64),
							},
						}
					case reflect.Int8:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_IntVal{
								IntVal: int64(reflectValue.Index(i).Interface().(int8)),
							},
						}
					case reflect.Int16:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_IntVal{
								IntVal: int64(reflectValue.Index(i).Interface().(int16)),
							},
						}
					case reflect.Int32:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_IntVal{
								IntVal: int64(reflectValue.Index(i).Interface().(int32)),
							},
						}
					case reflect.Int64:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_IntVal{
								IntVal: reflectValue.Index(i).Interface().(int64),
							},
						}
					case reflect.Bool:
						val = &gnmi.TypedValue{
							Value: &gnmi.TypedValue_BoolVal{
								BoolVal: reflectValue.Index(i).Interface().(bool),
							},
						}
					default:
						return nil, fmt.Errorf("unhandled leaf-list type %v", itemKind)
					}
				} else if err != nil {
					return nil, err
				}
				leafListVal.LeaflistVal.Element = append(leafListVal.LeaflistVal.Element, val)
			}
			update.Val = &gnmi.TypedValue{
				Value: leafListVal,
			}
		default:
			n := reflectValue.Type().String()
			k := reflectValue.Kind()
			return nil, fmt.Errorf("unhandled type %s %v", n, k)
		}
	}

	return update, nil
}

func extractEnum(reflectValue reflect.Value) (*gnmi.TypedValue, error) {
	// It might be an enum
	enumListMethod := reflectValue.MethodByName("ΛMap")
	if enumListMethod.IsValid() && !enumListMethod.IsZero() {
		returnMap := enumListMethod.Call(nil)
		if len(returnMap) != 1 {
			return nil, fmt.Errorf("error reading enum values")
		}
		enumMap := returnMap[0].Interface().(map[string]map[int64]ygot.EnumDefinition)
		enumValues, ok := enumMap[reflectValue.Type().Name()]
		if !ok {
			return nil, fmt.Errorf("could not find enum %s", reflectValue.Type().Elem().Name())
		}
		enumValue, ok := enumValues[reflectValue.Int()]
		if !ok {
			return nil, fmt.Errorf("unexpected enum value %d", reflect.Indirect(reflectValue).Int())
		}
		return &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: enumValue.Name}}, nil
	}
	return nil, fmt.Errorf("no enum method")
}

// ExtractGnmiListKeyMap - get the keys of a map
func ExtractGnmiListKeyMap(gnmiElement interface{}) (map[string]interface{}, error) {
	value := reflect.ValueOf(gnmiElement)
	keysMethod := value.MethodByName("ΛListKeyMap")
	if !keysMethod.IsValid() {
		return nil, fmt.Errorf("could not find method 'ΛListKeyMap' on %v", gnmiElement)
	}
	methodReturn := keysMethod.Call(make([]reflect.Value, 0))
	if len(methodReturn) != 2 {
		return nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
	}
	if !methodReturn[1].IsNil() {
		return nil, fmt.Errorf("error calling ΛListKeyMap %s %v", methodReturn[1].Interface().(error).Error(), value.Elem().Interface())
	}
	return methodReturn[0].Interface().(map[string]interface{}), nil
}

// ExtractGnmiEnumMap - extract an enum value from YGOT
func ExtractGnmiEnumMap(gnmiValuePtr *reflect.Value, path string, oaiValue interface{}) (string, *ygot.EnumDefinition, error) {
	submatchall := splitCapsAndNums.FindAllString(path, -1)
	enumPath := fmt.Sprintf("/%s", strings.ToLower(strings.Join(submatchall, "/")))
	gnmiPtrValue := reflect.New(gnmiValuePtr.Type()) // Create a pointer to struct, so we can access pointer receiver methods

	keysMethod := gnmiPtrValue.MethodByName("ΛEnumTypeMap")
	if !keysMethod.IsZero() && keysMethod.IsValid() && !keysMethod.IsNil() {
		methodReturn := keysMethod.Call(make([]reflect.Value, 0))
		if len(methodReturn) != 1 {
			return "", nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
		}
		yangEnumTypeMapIf := methodReturn[0].Interface()
		yangEnumTypeMap, ok := yangEnumTypeMapIf.(map[string][]reflect.Type)
		if !ok {
			return "", nil, fmt.Errorf("unable to cast to a map")
		}
		var t1 []reflect.Type
		for key, val := range yangEnumTypeMap {
			s1 := strings.ReplaceAll(key, "-", "/")
			if s1 == enumPath {
				t1 = val
			}
		}
		if t1 == nil {
			return "", nil, fmt.Errorf("could not find Enum %s in device enum map", enumPath)
		}
		for _, e := range t1 {
			eVal := reflect.Zero(e)
			lambdaMap := eVal.MethodByName("ΛMap")
			if !lambdaMap.IsZero() {
				lambdaMapReturn := lambdaMap.Call(make([]reflect.Value, 0))
				if len(lambdaMapReturn) != 1 {
					return e.Name(), nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
				}
				lambdaMapIf := lambdaMapReturn[0].Interface()
				yangEnumTypeMap, ok := lambdaMapIf.(map[string]map[int64]ygot.EnumDefinition)
				if !ok {
					return e.Name(), nil, fmt.Errorf("unable to cast to a map")
				}
				mapDefs, ok := yangEnumTypeMap[e.Name()]
				if !ok {
					return e.Name(), nil, fmt.Errorf("enum %s not present", e.Name())
				}
				oaiVal := reflect.ValueOf(oaiValue).Int()
				def, ok := mapDefs[oaiVal]
				if !ok {
					return e.Name(), nil, nil
				}
				return e.Name(), &def, nil
			}
		}
		return "", nil, fmt.Errorf("expected to find enum values")
	}

	return "", nil, fmt.Errorf("expected to find enum values")
}

// FindModelPluginObject - iterate through model plugin model structure to build object
func FindModelPluginObject(modelPluginPtr interface{}, path string, params ...string) (*reflect.Value, error) {
	submatchall := splitPath(path)
	return recurseFindMp(modelPluginPtr, submatchall, params)
}

func splitPath(path string) []string {
	return splitCapsAndNums.FindAllString(path, -1)
}

// splitPathYgotStruct -- undo the horros of goyang CamelCase which converts core-4g in to Core_4G, but converts cont1a in to Cont1A
func splitPathYgotStruct(path string) []string {
	matches := splitYgotStruct.FindAllString(path, -1)
	for _, m := range matches {
		path = strings.Replace(path, m, strings.ToLower(strings.TrimPrefix(m, "_")), 1)
	}
	return splitNumsThenCap.FindAllString(strings.ReplaceAll(path, "_", ""), -1)
}

func recurseFindMp(element interface{}, pathParts []string, params []string) (*reflect.Value, error) {
	skipPathParts := 0
	skipParams := 0
	value := reflect.ValueOf(element)
	var field reflect.Value
	switch value.Kind() {
	case reflect.String, reflect.Bool, reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64:
		return &value, nil
	case reflect.Struct:
		if len(pathParts) == 0 {
			return &value, nil
		}
		structField, skipped, err := findChildByParamNames(reflect.TypeOf(element), pathParts)
		if err != nil {
			return nil, err
		}
		field = value.FieldByName(structField.Name)
		skipPathParts = skipped
		if !field.IsValid() {
			return nil, fmt.Errorf("error getting fieldname %v on %v", pathParts, element)
		}
	case reflect.Ptr:
		field = value.Elem()
		// might be nil
		if !field.IsValid() {
			return nil, nil
		}
	case reflect.Map:
		if len(params) == 0 {
			return &value, nil
		}
		var matchingKey reflect.Value
	iterateKeys:
		for _, mapKey := range value.MapKeys() {
			matchingKeys := 0
			switch mapKey.Kind() {
			case reflect.String:
				if mapKey.Interface().(string) == params[0] {
					matchingKey = mapKey
					break iterateKeys
				}
			case reflect.Uint8:
				if fmt.Sprintf("%d", mapKey.Interface()) == params[0] || (mapKey.Interface().(uint8) == math.MaxUint8) {
					matchingKey = mapKey
					break iterateKeys
				}
			case reflect.Uint16:
				if fmt.Sprintf("%d", mapKey.Interface()) == params[0] || (mapKey.Interface().(uint16) == math.MaxUint16) {
					matchingKey = mapKey
					break iterateKeys
				}
			case reflect.Struct: // A compound key - e.g. double keyed list
				for i := 0; i < mapKey.NumField(); i++ {
					keyFieldVal := fmt.Sprintf("%v", mapKey.Field(i).Interface())
					// In the case of multi part key, there might only be 1 unknown_id param
					if (len(params) > i && params[i] == UnknownID) || (params[0] == UnknownID && len(params) == 1) {
						matchingKeys++
					} else if keyFieldVal == params[i] {
						matchingKeys++
					} else {
						matchingKeys = 0
					}
				}
				if matchingKeys == mapKey.NumField() {
					matchingKey = mapKey
				}
				break iterateKeys
			default:
				return nil, fmt.Errorf("unsupported Kind %s %v", mapKey.Kind(), mapKey.Type().Name())
			}
		}
		field = value.MapIndex(matchingKey)
		if !field.IsValid() {
			return nil, fmt.Errorf("error getting map index %s on %v", params[0], element)
		}
		skipParams++
	case reflect.Slice:
		values := make([]string, value.Len())
		for i := 0; i < value.Len(); i++ {
			res, err := recurseFindMp(value.Index(i).Interface(), pathParts[skipPathParts:], params[skipParams:])
			if err != nil {
				return nil, err
			}
			switch res.Kind() {
			case reflect.String, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				values[i] = fmt.Sprintf("%v", res.Interface())
			default:
				return nil, fmt.Errorf("unhandled %v", res.Kind())
			}
		}
		return &value, nil
	default:
		return nil, fmt.Errorf("unhandled %v", value.Kind())
	}

	return recurseFindMp(field.Interface(), pathParts[skipPathParts:], params[skipParams:])
}

// CreateModelPluginObject - iterate through model plugin model structure to build object
func CreateModelPluginObject(modelPluginPtr ygot.GoStruct, path string, params ...string) (interface{}, error) {
	submatchall := splitPath(path)
	return recurseCreateMp(modelPluginPtr, submatchall, params)
}

func recurseCreateMp(mpObjectPtr interface{}, pathParts []string, params []string) (interface{}, error) {
	skipParam := 0
	mpType := reflect.TypeOf(mpObjectPtr)
	if len(pathParts) == 0 {
		if len(params) == 0 {
			return nil, fmt.Errorf("expected a remaining param")
		}
		if err := setReflectValue(mpType.Elem(), reflect.ValueOf(mpObjectPtr).Elem(), params[0]); err != nil {
			return nil, err
		}
		return mpObjectPtr, nil
	}
	structField, _, err := findChildByParamNames(mpType, pathParts)
	ep := structField.Name
	epParts := len(splitPathYgotStruct(structField.Name))
	if err != nil {
		return nil, fmt.Errorf("cannot find child %s", pathParts[0])
	}
	switch structField.Type.Kind() {
	case reflect.Ptr:
		fieldValue := reflect.ValueOf(mpObjectPtr).Elem().FieldByName(ep)
		if fieldValue.IsNil() {
			fieldValue := reflect.New(structField.Type.Elem())
			reflect.ValueOf(mpObjectPtr).Elem().FieldByName(ep).Set(fieldValue)
		}
		return recurseCreateMp(fieldValue.Interface(), pathParts[epParts:], params[skipParam:])
	case reflect.Map:
		theMap := reflect.ValueOf(mpObjectPtr).Elem().FieldByName(ep)
		secondField := false
		if (!theMap.IsValid() && len(pathParts) > 1) || (theMap.IsValid() && len(pathParts) > 1 && !checkValue(pathParts[1:], theMap)) {
			theMap = reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1]))
			if !theMap.IsValid() {
				return nil, fmt.Errorf("unexpected field name %s", pathParts[0])
			}
			secondField = true
		}
		if theMap.IsNil() {
			theMap = reflect.MakeMap(structField.Type)
			if !secondField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(ep).Set(theMap)
			} else {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1])).Set(theMap)
			}
		}
		valueType := reflect.TypeOf(theMap.Interface()).Elem().Elem()
		skipParam++
		existingValue := reflect.Zero(valueType)
		keyType := reflect.TypeOf(theMap.Interface()).Key()
		for _, existingkey := range theMap.MapKeys() {
			existingKeyStr := fmt.Sprintf("%v", existingkey.Interface())
			switch keyType.Kind() {
			case reflect.Struct:
				testValues := make([]string, keyType.NumField())
				for i := 0; i < keyType.NumField(); i++ {
					if params[0] == UnknownID {
						testValues[i] = params[0]
						switch kft := keyType.Field(i).Type.Kind(); kft {
						case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
							reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							testValues[i] = fmt.Sprintf("%d", maxForIntKind(kft))
						}
					} else {
						testValues[i] = params[i]
					}
				}
				testValuesStr := fmt.Sprintf("{%s}", strings.Join(testValues, " "))
				if testValuesStr == existingKeyStr {
					existingValue = theMap.MapIndex(existingkey)
				}
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if params[0] == UnknownID {
					existingValue = theMap.MapIndex(existingkey)
				}
			default:
				if strings.Contains(existingKeyStr, params[0]) {
					existingValue = theMap.MapIndex(existingkey)
				}
			}
		}
		if existingValue.IsZero() {
			key := reflect.New(keyType)
			if keyType.Kind() == reflect.Struct {
				if len(params) < keyType.NumField() {
					return nil, fmt.Errorf("not enough params to create Key. Expected %d space separated parts, got %d. Value %v",
						keyType.NumField(), len(params), params)
				}
				for i := 0; i < keyType.NumField(); i++ {
					param := params[i]
					if i >= 1 && params[0] == UnknownID {
						param = UnknownID
					} else if i > 0 {
						skipParam++
					}
					if err = setReflectValue(keyType.Field(i).Type, key.Elem().Field(i), param); err != nil {
						return nil, err
					}
				}
			} else {
				if err := setReflectValue(keyType, key.Elem(), params[0]); err != nil {
					return nil, err
				}
			}
			existingValue = reflect.New(valueType)
			theMap.SetMapIndex(key.Elem(), existingValue)
		} else {
			if keyType.Kind() == reflect.Struct && len(params) > 0 && params[0] != UnknownID {
				skipParam += keyType.NumField() - 1 // Handle multi keyed lists
			}
		}
		return recurseCreateMp(existingValue.Interface(), pathParts[epParts:], params[skipParam:])
	case reflect.Slice:
		newSlice := reflect.MakeSlice(structField.Type, 0, 0)
		reflect.ValueOf(mpObjectPtr).Elem().FieldByName(ep).Set(newSlice)
		valueType := reflect.TypeOf(newSlice.Interface()).Elem()
		for range params {
			sliceEntry := reflect.New(valueType)
			sliceValue, err := recurseCreateMp(sliceEntry.Interface(), pathParts[epParts:], params[skipParam:])
			if err != nil {
				return nil, err
			}
			skipParam++
			newSlice = reflect.Append(newSlice, reflect.ValueOf(sliceValue).Elem())
		}
		return newSlice.Interface(), nil
	case reflect.Int64: // For enums
		newValue := reflect.New(structField.Type)
		return recurseCreateMp(newValue.Interface(), pathParts[epParts:], params[skipParam:])
	default:
		return nil, fmt.Errorf("recurseCreateMp unhandled %v %v", structField.Type.Kind(), mpObjectPtr)
	}
}

// findChildByParamNames find a child of the YANG object, using the param names got from the OpenAPI objects
// This has to be able to differentiate between 2 child fields that start with the same name
// e.g. Device and DeviceGroup - this recurses all the way down in to the hierarchy to guarantee that
// it returns the correct type
func findChildByParamNames(mpType reflect.Type, pathParts []string) (reflect.StructField, int, error) {
	if len(pathParts) == 0 {
		return reflect.StructField{}, 0, nil
	}
	pathPartsJoined := strings.ToLower(strings.Join(pathParts, ""))
	switch mpType.Kind() {
	case reflect.Ptr:
		return findChildByParamNames(mpType.Elem(), pathParts)
	case reflect.Map:
		return findChildByParamNames(mpType.Elem(), pathParts)
	case reflect.Struct:
		for i := 0; i < mpType.NumField(); i++ {
			childField := mpType.Field(i)
			path := strings.ReplaceAll(childField.Tag.Get("path"), "-", "")
			if strings.HasPrefix(pathPartsJoined, path) {
				skipped := 0
				remainder := path
				for _, p := range pathParts {
					pNoDash := strings.ToLower(strings.ReplaceAll(p, "-", ""))
					if strings.HasPrefix(remainder, pNoDash) {
						skipped++
						remainder = strings.TrimPrefix(remainder, pNoDash)
					} else {
						break
					}
				}
				if _, _, err := findChildByParamNames(childField.Type, pathParts[skipped:]); err != nil {
					continue
				}
				return childField, skipped, nil
			}
		}
	}
	return reflect.StructField{}, 0, fmt.Errorf("no field matching %s", pathPartsJoined)
}

func checkValue(pathParts []string, value reflect.Value) bool {
	if value.Kind() == reflect.Ptr {
		var ppStr = ""
		for _, pp := range pathParts {
			ppStr = ppStr + pp
			elem := value.Type().Elem()
			if elem == nil || elem.Kind() != reflect.Struct {
				return false
			}
			_, ok := elem.FieldByName(ppStr)
			if ok {
				return true
			}
		}
		return false
	}
	return true
}

func setReflectValue(theType reflect.Type, theStruct reflect.Value, theValue string) error {
	switch kt := theType.Kind(); kt {
	case reflect.String:
		theStruct.SetString(theValue)
		return nil
	case reflect.Bool:
		boolVal := false
		if theValue == "true" {
			boolVal = true
		}
		theStruct.SetBool(boolVal)
		return nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var uintVal uint64
		if theValue == UnknownID {
			uintVal = maxForIntKind(kt)
		} else {
			intVal, err := strconv.Atoi(theValue)
			if err != nil {
				return err
			}
			uintVal = uint64(intVal)
		}
		theStruct.SetUint(uintVal)
		return nil
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.Atoi(theValue)
		if err != nil {
			enumListMethod := theStruct.MethodByName("ΛMap")
			if !enumListMethod.IsZero() {
				returnMap := enumListMethod.Call(nil)
				if len(returnMap) != 1 {
					return fmt.Errorf("error reading enum values")
				}
				enumMap := returnMap[0].Interface().(map[string]map[int64]ygot.EnumDefinition)
				enumValues, ok := enumMap[theType.Name()]
				if !ok {
					return fmt.Errorf("could not find enum %s", theType.Elem().Name())
				}
				for k, v := range enumValues {
					if strings.EqualFold(v.Name, theValue) {
						theStruct.SetInt(k)
						return nil
					}
				}
			} else {
				return err
			}
		}
		theStruct.SetInt(int64(intVal))
		return nil
	default:
		return fmt.Errorf("unhandled type %s", kt.String())
	}
}

func maxForIntKind(kt reflect.Kind) uint64 {
	switch kt {
	case reflect.Int8:
		return math.MaxInt8
	case reflect.Int16:
		return math.MaxInt16
	case reflect.Int32:
		return math.MaxInt32
	case reflect.Int64, reflect.Int:
		return math.MaxInt64
	case reflect.Uint8:
		return math.MaxUint8
	case reflect.Uint16:
		return math.MaxUint16
	case reflect.Uint32:
		return math.MaxUint32
	default: // reflect.Uint64, reflect.Uint:
		return math.MaxUint64
	}
}
