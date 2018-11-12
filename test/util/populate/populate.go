package populate

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Walker struct {
	mogrify func(v reflect.Value)
}

var dummyPrivateKey *rsa.PrivateKey

func init() {
	var err error
	dummyPrivateKey, err = x509.ParsePKCS1PrivateKey([]byte{
		0x30, 0x25, 0x02, 0x01, 0x00, 0x02, 0x03, 0x00, 0xa2, 0x5f, 0x02, 0x03,
		0x01, 0x00, 0x01, 0x02, 0x02, 0x43, 0xf1, 0x02, 0x02, 0x00, 0xc5, 0x02,
		0x02, 0x00, 0xd3, 0x02, 0x02, 0x00, 0x91, 0x02, 0x02, 0x00, 0xad, 0x02,
		0x02, 0x00, 0xb7,
	})
	if err != nil {
		panic(err)
	}
}

// Walk is a recursive struct value population function. Given a pointer to an arbitrarily complex value v, it fills
// in the complete structure of that value, setting each string with the path taken to reach it.
//
// This function has the following caveats:
//  - Signed integers are set to int(1)
//  - Unsigned integers are set to uint(1)
//  - Floating point numbers are set to float(1.0)
//  - Booleans are set to True
//  - Arrays and slices are allocated 1 element
//  - Maps are allocated 1 element
//  - Only map[string][string] types are supported
//  - strings are set to the value of the path taken to reach the string
func Walk(v interface{}, mogrify func(v reflect.Value)) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		panic("argument is not a pointer to a value")
	}
	Walker{mogrify: mogrify}.walk(val, "")
}

// walk fills in the complete structure of a complex value v using path as the root of the labelling.
func (w Walker) walk(v reflect.Value, path string) {
	if !v.IsValid() {
		return
	}

	// special cases
	switch v.Interface().(type) {
	case []byte:
		v.SetBytes([]byte(path))
		return
	case *rsa.PrivateKey:
		// use a dummy value because the zero value cannot be marshalled
		v.Set(reflect.ValueOf(dummyPrivateKey))
		return
	}

	w.mogrify(v)

	switch v.Kind() {
	case reflect.Interface:
		w.walk(v.Elem(), path)
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		w.walk(v.Elem(), path)
	case reflect.Struct:
		// do not go on with the recursion if it isn't one of the core openshift-azure types
		if !strings.HasPrefix(v.Type().PkgPath(), "github.com/openshift/openshift-azure/") ||
			strings.HasPrefix(v.Type().PkgPath(), "github.com/openshift/openshift-azure/vendor/") {
			return
		}
		for i := 0; i < v.NumField(); i++ {
			// do not walk AADIdentityProvider.Kind to prevent breaking AADIdentityProvider unmarshall
			if v.Type().Field(i).Name == "Kind" {
				continue
			}
			field := v.Field(i)
			newpath := extendPath(path, v.Type().Field(i).Name, v.Kind())
			w.walk(field, newpath)
		}
	case reflect.Array, reflect.Slice:
		// if the array/slice has length 0 allocate a new slice of length 1
		if v.Len() == 0 {
			v.Set(reflect.MakeSlice(v.Type(), 1, 1))
		}
		for i := 0; i < v.Len(); i++ {
			field := v.Index(i)
			newpath := extendPath(path, strconv.Itoa(i), v.Kind())
			w.walk(field, newpath)
		}
	case reflect.Map:
		// only map[string]string types are supported
		if v.Type().Key().Kind() != reflect.String || v.Type().Elem().Kind() != reflect.String {
			return
		}
		v.Set(reflect.MakeMap(v.Type()))
		v.SetMapIndex(reflect.ValueOf(path+".key"), reflect.ValueOf(path+".val"))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(1)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.SetUint(1)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(1.0)
	case reflect.Bool:
		v.SetBool(true)
	case reflect.String:
		v.SetString(path)
	default:
		panic("unimplemented: " + v.Kind().String())
	}
}

// extendPath takes a path and a proposed extension to that path and returns a new path based on the kind of value for which
// the new path is being constructed
func extendPath(path, extension string, kind reflect.Kind) string {
	if path == "" {
		return extension
	}
	switch kind {
	case reflect.Struct:
		return fmt.Sprintf("%s.%s", path, extension)
	case reflect.Slice, reflect.Array:
		return fmt.Sprintf("%s[%s]", path, extension)
	}
	return ""
}
