// Code generated by gocode.Generate; DO NOT EDIT.

package api

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
	_ "cuelang.org/go/pkg"
)

var cuegenvalUser = cuegenMake("User", &User{})

// CUEValidate validates x.
func (x *User) CUEValidate() error {
	return cuegenCodec.Validate(cuegenvalUser, x)
}

// CUEComplete completes x.
func (x *User) CUEComplete() error {
	return cuegenCodec.Complete(cuegenvalUser, x)
}

var cuegenvalUserPermissions = cuegenMake("UserPermissions", &UserPermissions{})

// CUEValidate validates x.
func (x *UserPermissions) CUEValidate() error {
	return cuegenCodec.Validate(cuegenvalUserPermissions, x)
}

// CUEComplete completes x.
func (x *UserPermissions) CUEComplete() error {
	return cuegenCodec.Complete(cuegenvalUserPermissions, x)
}

var cuegenvalPermission = cuegenMake("Permission", &Permission{})

// CUEValidate validates x.
func (x *Permission) CUEValidate() error {
	return cuegenCodec.Validate(cuegenvalPermission, x)
}

// CUEComplete completes x.
func (x *Permission) CUEComplete() error {
	return cuegenCodec.Complete(cuegenvalPermission, x)
}

var cuegenvalResourcesSelector = cuegenMake("ResourcesSelector", nil)

// CUEValidate validates x.
func (x ResourcesSelector) CUEValidate() error {
	return cuegenCodec.Validate(cuegenvalResourcesSelector, x)
}

// CUEComplete completes x.
func (x ResourcesSelector) CUEComplete() error {
	return cuegenCodec.Complete(cuegenvalResourcesSelector, x)
}

var cuegenvalResourceAction = cuegenMake("ResourceAction", nil)

// CUEValidate validates x.
func (x ResourceAction) CUEValidate() error {
	return cuegenCodec.Validate(cuegenvalResourceAction, x)
}

// CUEComplete completes x.
func (x ResourceAction) CUEComplete() error {
	return cuegenCodec.Complete(cuegenvalResourceAction, x)
}

var cuegenCodec, cuegenInstance_, cuegenValue = func() (*gocodec.Codec, *cue.Instance, cue.Value) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0], instances[0].Value()
}()

// Deprecated: cue.Instance is deprecated. Use cuegenValue instead.
var cuegenInstance = cuegenInstance_

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenValue.FieldByName(name, true)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 561 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff|R\xcdj\xdb@\x10\u07b1]\xe8\x0ei\u07e0\xb0\xac!\xc4`+\x90C\x0f\xbe4\x86\xa4\xd0CKH\xe9\u0248\xb2\x91\xa6\xceR\xfd\xb1+5-A\xb4M\u04feL_\xad\xcf\x10\x95\x95d[VB|\xb1\xf6\x9b\x99o\xe6\xfbf\x9eU\xbf\a0\xa8\xfe0\xa8~0\xf6\xf2\xfb\x10`O'6WI@'*W\x0e\x86!\x8c\xce\xd34\x87\x01\x83\u0459\xca/a\x8f\xc1\x93\xd7:\"\v\xd5-c\xecE\xf5k\x00\xf0|\xe9\a\x05y\x9ft\xd4V\xde2\xa8n\x18;\xa8~\x0e\x01\x9en\xf1\x1b\x06\x03\x18\xbdS19\xa2Q\r\"c\xecn\xf8\xcf\r\x02\x00\xfb+\x9d_\x16\x17^\x90\u0187\xa1IM\xa2\xcdaP\u040c\xbe\xaa8\x8bh\xa62}\xa82\r\x00\xbcHbe\uc94a\xe0n\xf87S\xc1g\xb5\"\xa12\x8d\xa8\xe3,5\xb9\x9067:YY\x89\xf8\xc1\x92\x99\x8b\xb1\xfb\xc3q\xf3\xb8F\xae\u00f9hr\xc4\xf1*=xs2A\x9e\xa8\x98\u05a8\xf5\xde\xea\xe4\xbcH\xc8\x1e\x1cM\u013e+\xe1M\x04yY\x978)\x13\xe4\x14+\x1d\xedp\x9d:d\x82e\xdd\xf9\x8cL\xac\xad\xd5ib\xe7B4ct0\xdc~\xbbp\xfd\x1bo1<'\x9b\x16& \xfb\x9e\"\n\xf2\xd4)\xb9\x87a\x9f\xb5VXX2\x1f{2]^-5\xeb&/=\xcf\xeb4\xf5\xeb\xd4\x0e\xddt\xe9o_N\u0638;\xf55r\xb3\x9e\xe8\xa1\xe9j\xb6\r:A\xae\xa2(\xbd\xb2\xaf\u06be\xeb\xc8\"\xc87\xbd\x17u\xc6t\xe9\xef\x06'\xc8CJ4=Z{Rg<P[\xe2\xfd\xd9\xd6\xe6\xe0isc\x8ds\x8b\xce\u0274\xabw>\n!\xa4\xba\bB\x89\xbc=\x15!\xc3\u013d\xda\x1b\x901\x1d\xb7\x87\xeb\x85\xf4E\"/\x91/v\xf6\xd2\xdfTK\xbf\xd9U\xdd\xda\xd3!\xf2\u078a\\V\xd7g\x19\x7f\x9b\xd9V\x85\x1b\xa1uu.\x96\u0490\n\xe5T\xc8+\xa3s\x92>r^N\xc5c\xf5\xb3\xa3]\x06W\xd2:\xed\xf8:<\xbe\x93T\u2bb5s!z\x9b\xc0q?\xa3\xb5\x99\xb1\xff\x01\x00\x00\xff\xff\xb0\u01860{\x04\x00\x00")
