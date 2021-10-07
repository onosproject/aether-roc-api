// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package utils

import (
	"fmt"
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
	splitCaps    = regexp.MustCompile(`[A-Z][^A-Z]*`)
	splitNumbers = regexp.MustCompile(`[0-9][^A-Z]*`)
)

var log = logging.GetLogger("main")

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

// NewGnmiSetRequest -- new set request including updates and deletes
func NewGnmiSetRequest(updates []*gnmi.Update, deletes []*gnmi.Path,
	ext100Name *string, ext101Version *string, ext102Type *string,
) (*gnmi.SetRequest, error) {
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

	return gnmiSet, nil
}

// ExtractExtension100 - the name of the change will be returned as extension 100
func ExtractExtension100(gnmiResponse *gnmi.SetResponse) *string {
	for _, ext := range gnmiResponse.Extension {
		switch extTyped := ext.Ext.(type) {
		case *gnmi_ext.Extension_RegisteredExt:
			if extTyped.RegisteredExt.Id == 100 {
				changeName := string(extTyped.RegisteredExt.Msg)
				return &changeName
			}
		}
	}

	return nil
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
	modelType := strings.Title(oapiParts[1]) // Change to title case
	modelVersion := oapiParts[2][1:]         // Remove the "v" at the start of "v1.0.0"

	extensions := make([]*gnmi_ext.Extension, 0)
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
	case "[]string":
		valueStrArr := value.([]string)
		llVals := make([]*gnmi.TypedValue, 0)
		for _, str := range valueStrArr {
			llVal := gnmi.TypedValue{
				Value: &gnmi.TypedValue_StringVal{StringVal: str},
			}
			llVals = append(llVals, &llVal)
		}
		update.Val.Value = &gnmi.TypedValue_LeaflistVal{
			LeaflistVal: &gnmi.ScalarArray{
				Element: llVals,
			},
		}
	case "*uint8", "*uint16", "*uint32", "*uint64":
		update.Val.Value = &gnmi.TypedValue_UintVal{UintVal: reflect.Indirect(reflectValue).Uint()}
	case "*int8", "*int16", "*int32", "*int64":
		update.Val.Value = &gnmi.TypedValue_IntVal{IntVal: reflect.Indirect(reflectValue).Int()}
	case "*bool":
		update.Val.Value = &gnmi.TypedValue_BoolVal{BoolVal: reflect.Indirect(reflectValue).Bool()}
	default:
		switch reflectValue.Kind().String() {
		case "int64":
			update.Val.Value = &gnmi.TypedValue_IntVal{IntVal: reflect.Indirect(reflectValue).Int()}
		case "ptr":
			// It might be an enum
			enumListMethod := reflectValue.Elem().MethodByName("ΛMap")
			if !enumListMethod.IsZero() {
				returnMap := enumListMethod.Call(nil)
				if len(returnMap) != 1 {
					return nil, fmt.Errorf("error reading enum values")
				}
				enumMap := returnMap[0].Interface().(map[string]map[int64]ygot.EnumDefinition)
				enumValues, ok := enumMap[reflectValue.Type().Elem().Name()]
				if !ok {
					return nil, fmt.Errorf("could not find enum %s", reflectValue.Type().Elem().Name())
				}
				enumValue, ok := enumValues[reflect.Indirect(reflectValue).Int()]
				if !ok {
					return nil, fmt.Errorf("unexpected enum value %d", reflect.Indirect(reflectValue).Int())
				}
				update.Val.Value = &gnmi.TypedValue_StringVal{StringVal: enumValue.Name}
			} else {
				return nil, err
			}
		default:
			n := reflectValue.Type().String()
			k := reflectValue.Kind()
			return nil, fmt.Errorf("unhandled type %s %v", n, k)
		}
	}

	return update, nil
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
		return nil, fmt.Errorf("error calling ΛListKeyMap %s", methodReturn[1].Interface().(error).Error())
	}
	return methodReturn[0].Interface().(map[string]interface{}), nil
}

// ExtractGnmiEnumMap - extract an enum value from YGOT
func ExtractGnmiEnumMap(gnmiElement interface{}, path string, oaiValue interface{}) (string, *ygot.EnumDefinition, error) {
	submatchall := splitCaps.FindAllString(path, -1)
	enumPath := fmt.Sprintf("/%s", strings.ToLower(strings.Join(submatchall, "/")))
	value := reflect.ValueOf(gnmiElement)
	keysMethod := value.MethodByName("ΛEnumTypeMap")
	if !keysMethod.IsZero() {
		methodReturn := keysMethod.Call(make([]reflect.Value, 0))
		if len(methodReturn) != 1 {
			return "", nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
		}
		yangEnumTypeMapIf := methodReturn[0].Interface()
		yangEnumTypeMap, ok := yangEnumTypeMapIf.(map[string][]reflect.Type)
		if !ok {
			return "", nil, fmt.Errorf("unable to cast to a map")
		}
		enums, ok := yangEnumTypeMap[enumPath]
		if !ok {
			return "", nil, fmt.Errorf("could not find Enum %s in device enum map", enumPath)
		}
		for _, e := range enums {
			eVal := reflect.Zero(e)
			lambdaMap := eVal.MethodByName("ΛMap")
			if !lambdaMap.IsZero() {
				lambdaMapReturn := lambdaMap.Call(make([]reflect.Value, 0))
				if len(lambdaMapReturn) != 1 {
					return "", nil, fmt.Errorf("expecting 2 values back from method ΛListKeyMap")
				}
				lambdaMapIf := lambdaMapReturn[0].Interface()
				yangEnumTypeMap, ok := lambdaMapIf.(map[string]map[int64]ygot.EnumDefinition)
				if !ok {
					return "", nil, fmt.Errorf("unable to cast to a map")
				}
				mapDefs, ok := yangEnumTypeMap[e.Name()]
				if !ok {
					return "", nil, fmt.Errorf("enum %s not present", e.Name())
				}
				oaiVal := reflect.ValueOf(oaiValue).Int()
				def, ok := mapDefs[oaiVal]
				if !ok {
					return "", nil, fmt.Errorf("value %d in enum %s not present", oaiValue, e.Name())
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
	submatchall := splitCaps.FindAllString(path, -1)
	for i, sm := range submatchall {
		numbers := splitNumbers.FindAllString(sm, -1)
		if len(numbers) == 1 {
			replace := fmt.Sprintf("%s_%s", sm[:len(sm)-len(numbers[0])], strings.ToTitle(numbers[0]))
			submatchall[i] = replace
			log.Infof("Numbers %s\n", replace)
		}
	}
	return submatchall
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
		field = value.FieldByName(pathParts[0])
		if (!field.IsValid() && len(pathParts) > 1) || (field.IsValid() && len(pathParts) > 1 && !checkValue(pathParts[1:], field)) {
			// Try again with more parts
			field = value.FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1]))
			skipPathParts++
			if (!field.IsValid() && len(pathParts) > 2) || (field.IsValid() && len(pathParts) > 2 && !checkValue(pathParts[2:], field)) {
				// Try again with more parts
				field = value.FieldByName(fmt.Sprintf("%s%s%s", pathParts[0], pathParts[1], pathParts[2]))
				skipPathParts++
				if !field.IsValid() {
					return nil, fmt.Errorf("error getting fieldname %v on %v", pathParts, element)
				}
			}
		}
		skipPathParts++
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
		p := reflect.ValueOf(params[0])
		field = value.MapIndex(p)
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
			values[i] = res.Interface().(string)
		}
		return &value, nil
	default:
		return nil, fmt.Errorf("unhandled %v", value.Kind())
	}

	return recurseFindMp(field.Interface(), pathParts[skipPathParts:], params[skipParams:])
}

// CreateModelPluginObject - iterate through model plugin model structure to build object
func CreateModelPluginObject(modelPluginPtr interface{}, path string, params ...string) (interface{}, error) {
	submatchall := splitPath(path)
	return recurseCreateMp(modelPluginPtr, submatchall, params)
}

func recurseCreateMp(mpObjectPtr interface{}, pathParts []string, params []string) (interface{}, error) {
	skipPathParts := 0
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
	structField, ok := mpType.Elem().FieldByName(pathParts[0])
	if (!ok && len(pathParts) > 1) || (ok && len(pathParts) > 1 && !checkBranch(pathParts[1:], structField)) {
		structField, ok = mpType.Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1]))
		if (!ok && len(pathParts) > 2) || (ok && len(pathParts) > 2 && !checkBranch(pathParts[2:], structField)) {
			structField, ok = mpType.Elem().FieldByName(fmt.Sprintf("%s%s%s", pathParts[0], pathParts[1], pathParts[2]))
			if (!ok && len(pathParts) > 3) || (ok && len(pathParts) > 3 && !checkBranch(pathParts[3:], structField)) {
				structField, ok = mpType.Elem().FieldByName(fmt.Sprintf("%s%s%s%s", pathParts[0], pathParts[1], pathParts[2], pathParts[3]))
				if !ok {
					return nil, fmt.Errorf("unable to get field %s", pathParts[0])
				}
			}
		}
	} else if !ok {
		return nil, fmt.Errorf("cannot find child %s", pathParts[0])
	}
	switch structField.Type.Kind() {
	case reflect.Ptr:
		secondField := false
		thirdField := false
		fourthField := false
		fieldValue := reflect.ValueOf(mpObjectPtr).Elem().FieldByName(pathParts[0])
		if !fieldValue.IsValid() && len(pathParts) > 1 {
			fieldValue = reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1]))
			if !fieldValue.IsValid() && len(pathParts) > 2 {
				fieldValue = reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s%s", pathParts[0], pathParts[1], pathParts[2]))
				if !fieldValue.IsValid() && len(pathParts) > 3 {
					fieldValue = reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s%s%s", pathParts[0], pathParts[1], pathParts[2], pathParts[3]))
					if !fieldValue.IsValid() {
						return nil, fmt.Errorf("unexpected field name %s", pathParts[0])
					}
					fourthField = true
					skipPathParts++
				}
				thirdField = true
				skipPathParts++
			}
			secondField = true
			skipPathParts++
		}
		if fieldValue.IsNil() {
			fieldValue := reflect.New(structField.Type.Elem())
			if !secondField && !thirdField && !fourthField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(pathParts[0]).Set(fieldValue)
			} else if secondField && !thirdField && !fourthField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1])).Set(fieldValue)
			} else if thirdField && !fourthField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s%s", pathParts[0], pathParts[1], pathParts[2])).Set(fieldValue)
			} else if fourthField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s%s%s", pathParts[0], pathParts[1], pathParts[2], pathParts[3])).Set(fieldValue)
			} else {
				return nil, fmt.Errorf("unexpected logic error with multiple fields %v", pathParts)
			}
		}
		skipPathParts++
		return recurseCreateMp(fieldValue.Interface(), pathParts[skipPathParts:], params[skipParam:])
	case reflect.Map:
		theMap := reflect.ValueOf(mpObjectPtr).Elem().FieldByName(pathParts[0])
		secondField := false
		if (!theMap.IsValid() && len(pathParts) > 1) || (theMap.IsValid() && len(pathParts) > 1 && !checkValue(pathParts[1:], theMap)) {
			theMap = reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1]))
			if !theMap.IsValid() {
				return nil, fmt.Errorf("unexpected field name %s", pathParts[0])
			}
			skipPathParts++
			secondField = true
		}
		if theMap.IsNil() {
			theMap := reflect.MakeMap(structField.Type)
			if !secondField {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(pathParts[0]).Set(theMap)
			} else {
				reflect.ValueOf(mpObjectPtr).Elem().FieldByName(fmt.Sprintf("%s%s", pathParts[0], pathParts[1])).Set(theMap)
			}
		}
		valueType := reflect.TypeOf(theMap.Interface()).Elem().Elem()
		skipParam++
		skipPathParts++
		existingValue := reflect.Zero(valueType)
		keyType := reflect.TypeOf(theMap.Interface()).Key()
		for _, existingkey := range theMap.MapKeys() {
			existingKeyStr := fmt.Sprintf("%v", existingkey.Interface())
			if keyType.Kind() == reflect.Struct {
				// Strip off bracket at start and end for struct keys
				existingKeyStr = existingKeyStr[1 : len(existingKeyStr)-1]
			}
			if strings.Contains(existingKeyStr, params[0]) {
				existingValue = theMap.MapIndex(existingkey)
			}
		}
		if existingValue.IsZero() {
			key := reflect.New(keyType)
			if keyType.Kind() == reflect.Struct {
				keyValueParts := strings.Split(params[0], " ")
				if len(keyValueParts) != keyType.NumField() {
					return nil, fmt.Errorf("unexpected key structure. Expected %d space separated parts, got %d. Value %s",
						keyType.NumField(), len(keyValueParts), params[0])
				}
				for i := 0; i < keyType.NumField(); i++ {
					if err := setReflectValue(keyType.Field(i).Type, key.Elem().Field(i), keyValueParts[i]); err != nil {
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
		}
		return recurseCreateMp(existingValue.Interface(), pathParts[skipPathParts:], params[skipParam:])
	case reflect.Slice:
		newSlice := reflect.MakeSlice(structField.Type, 0, 0)
		reflect.ValueOf(mpObjectPtr).Elem().FieldByName(pathParts[0]).Set(newSlice)
		valueType := reflect.TypeOf(newSlice.Interface()).Elem()
		skipPathParts++
		for range params {
			sliceEntry := reflect.New(valueType)
			sliceValue, err := recurseCreateMp(sliceEntry.Interface(), pathParts[skipPathParts:], params[skipParam:])
			if err != nil {
				return nil, err
			}
			skipParam++
			newSlice = reflect.Append(newSlice, reflect.ValueOf(sliceValue).Elem())
		}
		return newSlice.Interface(), nil
	case reflect.Int64: // For enums
		newValue := reflect.New(structField.Type)
		skipPathParts++
		return recurseCreateMp(newValue.Interface(), pathParts[skipPathParts:], params[skipParam:])
	default:
		return nil, fmt.Errorf("recurseCreateMp unhandled %v %v", structField.Type.Kind(), mpObjectPtr)
	}
}

// checkBranch - check that the struct field has children named pathParts[0], or pathParts[0]pathParts[1]
func checkBranch(pathParts []string, structField reflect.StructField) bool {
	if structField.Type.Kind() == reflect.Ptr {
		var ppStr = ""
		for _, pp := range pathParts {
			ppStr = ppStr + pp
			_, ok := structField.Type.Elem().FieldByName(ppStr)
			if ok {
				return true
			}
		}
		return false
	}
	return true
}

func checkValue(pathParts []string, value reflect.Value) bool {
	if value.Kind() == reflect.Ptr {
		var ppStr = ""
		for _, pp := range pathParts {
			ppStr = ppStr + pp
			_, ok := value.Type().Elem().FieldByName(ppStr)
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
		uintVal, err := strconv.Atoi(theValue)
		if err != nil {
			return err
		}
		theStruct.SetUint(uint64(uintVal))
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
				enumValues, ok := enumMap[theType.Elem().Name()]
				if !ok {
					return fmt.Errorf("could not find enum %s", theType.Elem().Name())
				}
				for k, v := range enumValues {
					if strings.EqualFold(v.Name, theValue) {
						theStruct.SetInt(k)
						break
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
