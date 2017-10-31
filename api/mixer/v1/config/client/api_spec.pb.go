// Code generated by protoc-gen-gogo.
// source: mixer/v1/config/client/api_spec.proto
// DO NOT EDIT!

package istio_mixer_v1_config_client

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import istio_mixer_v1 "istio.io/api/mixer/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HTTPAPISpec defines the canonical configuration for generating
// API-related attributes from HTTP requests based on the method and
// uri templated path matches. It is sufficient for defining the API
// surface of a service for the purposes of API attribute
// generation. It is not intended to represent auth, quota,
// documentation, or other information commonly found in other API
// specifications, e.g. OpenAPI.
//
// Existing standards that define operations (or methods) in terms of
// HTTP methods and paths can be normalized to this format for use in
// Istio. For example, a simple petstore API described by OpenAPIv2
// [here](https://github.com/googleapis/gnostic/blob/master/examples/v2.0/yaml/petstore-simple.yaml)
// can be represented with the following HTTPAPISpec.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: HTTPAPISpec
//     metadata:
//       name: petstore
//       namespace: default
//     spec:
//       attributes:
//         api.service: petstore.swagger.io
//         api.version: 1.0.0
//       patterns:
//       - attributes:
//           api.operation: findPets
//         httpMethod: GET
//         uriTemplate: /api/pets
//       - attributes:
//           api.operation: addPet
//         httpMethod: POST
//         uriTemplate: /api/pets
//       - attributes:
//           api.operation: findPetById
//         httpMethod: GET
//         uriTemplate: /api/pets/{id}
//       - attributes:
//           api.operation: deletePet
//         httpMethod: DELETE
//         uriTemplate: /api/pets/{id}
//
type HTTPAPISpec struct {
	// List of attributes that are generated when *any* of the HTTP
	// patterns match. This list typically includes the "api.service"
	// and "api.version" attributes.
	Attributes *istio_mixer_v1.Attributes `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
	// List of HTTP patterns to match.
	Patterns []*HTTPAPISpecPattern `protobuf:"bytes,2,rep,name=patterns" json:"patterns,omitempty"`
}

func (m *HTTPAPISpec) Reset()                    { *m = HTTPAPISpec{} }
func (*HTTPAPISpec) ProtoMessage()               {}
func (*HTTPAPISpec) Descriptor() ([]byte, []int) { return fileDescriptorApiSpec, []int{0} }

// HTTPAPISpecPattern defines a single pattern to match against
// incoming HTTP requests. The per-pattern list of attributes is
// generated if both the http_method and uri_template match. In
// addition, the top-level list of attributes in the HTTPAPISpec is also
// generated.
//
//     pattern:
//     - attributes
//         api.operation: doFooBar
//       httpMethod: GET
//       uriTemplate: /foo/bar
//
type HTTPAPISpecPattern struct {
	// List of attributes that are generated if the HTTP request matches
	// the specified http_method and uri_template. This typically
	// includes the "api.operation" attribute.
	Attributes *istio_mixer_v1.Attributes `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
	// HTTP request method to match against as defined by
	// [rfc7231](https://tools.ietf.org/html/rfc7231#page-21). For
	// example: GET, HEAD, POST, PUT, DELETE.
	HttpMethod string `protobuf:"bytes,2,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Types that are valid to be assigned to Pattern:
	//	*HTTPAPISpecPattern_UriTemplate
	//	*HTTPAPISpecPattern_Regex
	Pattern isHTTPAPISpecPattern_Pattern `protobuf_oneof:"pattern"`
}

func (m *HTTPAPISpecPattern) Reset()                    { *m = HTTPAPISpecPattern{} }
func (*HTTPAPISpecPattern) ProtoMessage()               {}
func (*HTTPAPISpecPattern) Descriptor() ([]byte, []int) { return fileDescriptorApiSpec, []int{1} }

type isHTTPAPISpecPattern_Pattern interface {
	isHTTPAPISpecPattern_Pattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HTTPAPISpecPattern_UriTemplate struct {
	UriTemplate string `protobuf:"bytes,3,opt,name=uri_template,json=uriTemplate,proto3,oneof"`
}
type HTTPAPISpecPattern_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

func (*HTTPAPISpecPattern_UriTemplate) isHTTPAPISpecPattern_Pattern() {}
func (*HTTPAPISpecPattern_Regex) isHTTPAPISpecPattern_Pattern()       {}

func (m *HTTPAPISpecPattern) GetPattern() isHTTPAPISpecPattern_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HTTPAPISpecPattern) GetUriTemplate() string {
	if x, ok := m.GetPattern().(*HTTPAPISpecPattern_UriTemplate); ok {
		return x.UriTemplate
	}
	return ""
}

func (m *HTTPAPISpecPattern) GetRegex() string {
	if x, ok := m.GetPattern().(*HTTPAPISpecPattern_Regex); ok {
		return x.Regex
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HTTPAPISpecPattern) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HTTPAPISpecPattern_OneofMarshaler, _HTTPAPISpecPattern_OneofUnmarshaler, _HTTPAPISpecPattern_OneofSizer, []interface{}{
		(*HTTPAPISpecPattern_UriTemplate)(nil),
		(*HTTPAPISpecPattern_Regex)(nil),
	}
}

func _HTTPAPISpecPattern_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HTTPAPISpecPattern)
	// pattern
	switch x := m.Pattern.(type) {
	case *HTTPAPISpecPattern_UriTemplate:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.UriTemplate)
	case *HTTPAPISpecPattern_Regex:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Regex)
	case nil:
	default:
		return fmt.Errorf("HTTPAPISpecPattern.Pattern has unexpected type %T", x)
	}
	return nil
}

func _HTTPAPISpecPattern_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HTTPAPISpecPattern)
	switch tag {
	case 3: // pattern.uri_template
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HTTPAPISpecPattern_UriTemplate{x}
		return true, err
	case 4: // pattern.regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HTTPAPISpecPattern_Regex{x}
		return true, err
	default:
		return false, nil
	}
}

func _HTTPAPISpecPattern_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HTTPAPISpecPattern)
	// pattern
	switch x := m.Pattern.(type) {
	case *HTTPAPISpecPattern_UriTemplate:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.UriTemplate)))
		n += len(x.UriTemplate)
	case *HTTPAPISpecPattern_Regex:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Regex)))
		n += len(x.Regex)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// HTTPAPISpecReference defines a reference to an HTTPAPISpec. This is
// typically used for establishing bindings between an HTTPAPISpec and an
// IstioService. For example, the following defines an
// HTTPAPISpecReference for service `foo` in namespace `bar`.
//
//     - name: foo
//       namespace: bar
//
type HTTPAPISpecReference struct {
	// REQUIRED. The short name of the HTTPAPISpec. This is the resource
	// name defined by the metadata name field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional namespace of the HTTPAPISpec. Defaults to the encompassing
	// HTTPAPISpecBinding's metadata namespace field.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *HTTPAPISpecReference) Reset()                    { *m = HTTPAPISpecReference{} }
func (*HTTPAPISpecReference) ProtoMessage()               {}
func (*HTTPAPISpecReference) Descriptor() ([]byte, []int) { return fileDescriptorApiSpec, []int{2} }

// HTTPAPISpecBinding defines the binding between HTTPAPISpecs and one or more
// IstioService. For example, the following establishes a binding
// between the HTTPAPISpec `petstore` and service `foo` in namespace `bar`.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: HTTPAPISpecBinding
//     metadata:
//       name: my-binding
//       namespace: default
//     spec:
//       services:
//       - name: foo
//         namespace: bar
//       api_specs:
//       - name: petstore
//         namespace: default
//
type HTTPAPISpecBinding struct {
	// REQUIRED. One or more services to map the listed HTTPAPISpec onto.
	Services []*IstioService `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	// REQUIRED. One or more HTTPAPISpec references that should be mapped to
	// the specified service(s). The aggregate collection of match
	// conditions defined in the HTTPAPISpecs should not overlap.
	ApiSpecs []*HTTPAPISpecReference `protobuf:"bytes,2,rep,name=api_specs,json=apiSpecs" json:"api_specs,omitempty"`
}

func (m *HTTPAPISpecBinding) Reset()                    { *m = HTTPAPISpecBinding{} }
func (*HTTPAPISpecBinding) ProtoMessage()               {}
func (*HTTPAPISpecBinding) Descriptor() ([]byte, []int) { return fileDescriptorApiSpec, []int{3} }

func init() {
	proto.RegisterType((*HTTPAPISpec)(nil), "istio.mixer.v1.config.client.HTTPAPISpec")
	proto.RegisterType((*HTTPAPISpecPattern)(nil), "istio.mixer.v1.config.client.HTTPAPISpecPattern")
	proto.RegisterType((*HTTPAPISpecReference)(nil), "istio.mixer.v1.config.client.HTTPAPISpecReference")
	proto.RegisterType((*HTTPAPISpecBinding)(nil), "istio.mixer.v1.config.client.HTTPAPISpecBinding")
}
func (m *HTTPAPISpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPAPISpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Attributes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApiSpec(dAtA, i, uint64(m.Attributes.Size()))
		n1, err := m.Attributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Patterns) > 0 {
		for _, msg := range m.Patterns {
			dAtA[i] = 0x12
			i++
			i = encodeVarintApiSpec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HTTPAPISpecPattern) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPAPISpecPattern) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Attributes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApiSpec(dAtA, i, uint64(m.Attributes.Size()))
		n2, err := m.Attributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.HttpMethod) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApiSpec(dAtA, i, uint64(len(m.HttpMethod)))
		i += copy(dAtA[i:], m.HttpMethod)
	}
	if m.Pattern != nil {
		nn3, err := m.Pattern.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *HTTPAPISpecPattern_UriTemplate) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x1a
	i++
	i = encodeVarintApiSpec(dAtA, i, uint64(len(m.UriTemplate)))
	i += copy(dAtA[i:], m.UriTemplate)
	return i, nil
}
func (m *HTTPAPISpecPattern_Regex) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x22
	i++
	i = encodeVarintApiSpec(dAtA, i, uint64(len(m.Regex)))
	i += copy(dAtA[i:], m.Regex)
	return i, nil
}
func (m *HTTPAPISpecReference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPAPISpecReference) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApiSpec(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApiSpec(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	return i, nil
}

func (m *HTTPAPISpecBinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPAPISpecBinding) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, msg := range m.Services {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApiSpec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ApiSpecs) > 0 {
		for _, msg := range m.ApiSpecs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintApiSpec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64ApiSpec(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ApiSpec(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintApiSpec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HTTPAPISpec) Size() (n int) {
	var l int
	_ = l
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovApiSpec(uint64(l))
	}
	if len(m.Patterns) > 0 {
		for _, e := range m.Patterns {
			l = e.Size()
			n += 1 + l + sovApiSpec(uint64(l))
		}
	}
	return n
}

func (m *HTTPAPISpecPattern) Size() (n int) {
	var l int
	_ = l
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovApiSpec(uint64(l))
	}
	l = len(m.HttpMethod)
	if l > 0 {
		n += 1 + l + sovApiSpec(uint64(l))
	}
	if m.Pattern != nil {
		n += m.Pattern.Size()
	}
	return n
}

func (m *HTTPAPISpecPattern_UriTemplate) Size() (n int) {
	var l int
	_ = l
	l = len(m.UriTemplate)
	n += 1 + l + sovApiSpec(uint64(l))
	return n
}
func (m *HTTPAPISpecPattern_Regex) Size() (n int) {
	var l int
	_ = l
	l = len(m.Regex)
	n += 1 + l + sovApiSpec(uint64(l))
	return n
}
func (m *HTTPAPISpecReference) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApiSpec(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovApiSpec(uint64(l))
	}
	return n
}

func (m *HTTPAPISpecBinding) Size() (n int) {
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovApiSpec(uint64(l))
		}
	}
	if len(m.ApiSpecs) > 0 {
		for _, e := range m.ApiSpecs {
			l = e.Size()
			n += 1 + l + sovApiSpec(uint64(l))
		}
	}
	return n
}

func sovApiSpec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApiSpec(x uint64) (n int) {
	return sovApiSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HTTPAPISpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpec{`,
		`Attributes:` + strings.Replace(fmt.Sprintf("%v", this.Attributes), "Attributes", "istio_mixer_v1.Attributes", 1) + `,`,
		`Patterns:` + strings.Replace(fmt.Sprintf("%v", this.Patterns), "HTTPAPISpecPattern", "HTTPAPISpecPattern", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPAPISpecPattern) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpecPattern{`,
		`Attributes:` + strings.Replace(fmt.Sprintf("%v", this.Attributes), "Attributes", "istio_mixer_v1.Attributes", 1) + `,`,
		`HttpMethod:` + fmt.Sprintf("%v", this.HttpMethod) + `,`,
		`Pattern:` + fmt.Sprintf("%v", this.Pattern) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPAPISpecPattern_UriTemplate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpecPattern_UriTemplate{`,
		`UriTemplate:` + fmt.Sprintf("%v", this.UriTemplate) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPAPISpecPattern_Regex) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpecPattern_Regex{`,
		`Regex:` + fmt.Sprintf("%v", this.Regex) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPAPISpecReference) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpecReference{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPAPISpecBinding) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPAPISpecBinding{`,
		`Services:` + strings.Replace(fmt.Sprintf("%v", this.Services), "IstioService", "IstioService", 1) + `,`,
		`ApiSpecs:` + strings.Replace(fmt.Sprintf("%v", this.ApiSpecs), "HTTPAPISpecReference", "HTTPAPISpecReference", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApiSpec(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HTTPAPISpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiSpec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPAPISpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPAPISpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &istio_mixer_v1.Attributes{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patterns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patterns = append(m.Patterns, &HTTPAPISpecPattern{})
			if err := m.Patterns[len(m.Patterns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiSpec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HTTPAPISpecPattern) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiSpec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPAPISpecPattern: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPAPISpecPattern: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &istio_mixer_v1.Attributes{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HttpMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = &HTTPAPISpecPattern_UriTemplate{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = &HTTPAPISpecPattern_Regex{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiSpec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HTTPAPISpecReference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiSpec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPAPISpecReference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPAPISpecReference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiSpec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HTTPAPISpecBinding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiSpec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPAPISpecBinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPAPISpecBinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &IstioService{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiSpecs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiSpecs = append(m.ApiSpecs, &HTTPAPISpecReference{})
			if err := m.ApiSpecs[len(m.ApiSpecs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiSpec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApiSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiSpec
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiSpec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApiSpec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApiSpec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApiSpec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApiSpec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiSpec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/v1/config/client/api_spec.proto", fileDescriptorApiSpec) }

var fileDescriptorApiSpec = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x86, 0x77, 0x93, 0x00, 0xe7, 0x31, 0xd5, 0x2a, 0x42, 0xe6, 0x14, 0x2d, 0xa7, 0x03, 0xa4,
	0x13, 0xc5, 0x3a, 0x39, 0x3a, 0xba, 0x5c, 0x81, 0x2e, 0x12, 0x88, 0xd3, 0xe6, 0xfa, 0x93, 0xe3,
	0x4c, 0x9c, 0x95, 0x62, 0x7b, 0xb5, 0xde, 0x3b, 0xa5, 0xe4, 0x11, 0xe8, 0x78, 0x05, 0x1a, 0x1e,
	0x80, 0x37, 0x48, 0x99, 0x92, 0x12, 0x9b, 0x86, 0x32, 0x8f, 0x80, 0xd6, 0x76, 0x9c, 0x03, 0xc4,
	0x09, 0x89, 0xca, 0xeb, 0x7f, 0xbe, 0x99, 0xfd, 0x77, 0x66, 0xe0, 0x79, 0xaa, 0x2e, 0xd1, 0x84,
	0xab, 0x83, 0x30, 0xce, 0xb3, 0x33, 0x95, 0x84, 0xf1, 0x85, 0xc2, 0xcc, 0x86, 0x91, 0x56, 0x8b,
	0x42, 0x63, 0x2c, 0xb4, 0xc9, 0x6d, 0xce, 0xf6, 0x54, 0x61, 0x55, 0x2e, 0x6a, 0x58, 0xac, 0x0e,
	0x44, 0x03, 0x8b, 0x06, 0xee, 0xef, 0x26, 0x79, 0x92, 0xd7, 0x60, 0xe8, 0x4e, 0x4d, 0x4e, 0xff,
	0x71, 0x57, 0x3a, 0xb2, 0xd6, 0xa8, 0x93, 0xa5, 0xc5, 0xa2, 0x0d, 0x3d, 0xfb, 0xcb, 0xad, 0x05,
	0x9a, 0x95, 0x8a, 0xb1, 0xa1, 0x86, 0x1f, 0x29, 0xf8, 0xd3, 0xf9, 0x7c, 0x76, 0x38, 0x3b, 0x3a,
	0xd6, 0x18, 0xb3, 0x57, 0x00, 0x77, 0x95, 0x02, 0x3a, 0xa0, 0x23, 0x7f, 0xdc, 0x17, 0xbf, 0x39,
	0x3b, 0xec, 0x08, 0xb9, 0x46, 0xb3, 0x37, 0xd0, 0xd3, 0x91, 0xb5, 0x68, 0xb2, 0x22, 0xd8, 0x1a,
	0x6c, 0x8f, 0xfc, 0xf1, 0xbe, 0xd8, 0xf4, 0x26, 0xb1, 0x76, 0xf1, 0xac, 0x49, 0x94, 0x5d, 0x85,
	0xe1, 0x17, 0x0a, 0xec, 0x4f, 0xe0, 0xbf, 0x0c, 0x3e, 0x01, 0xff, 0xdc, 0x5a, 0xbd, 0x48, 0xd1,
	0x9e, 0xe7, 0xa7, 0xc1, 0xd6, 0x80, 0x8e, 0x3c, 0x09, 0x4e, 0x7a, 0x5b, 0x2b, 0xec, 0x29, 0x3c,
	0x5c, 0x1a, 0xb5, 0xb0, 0x98, 0xea, 0x8b, 0xc8, 0x62, 0xb0, 0xed, 0x88, 0x29, 0x91, 0xfe, 0xd2,
	0xa8, 0x79, 0x2b, 0xb2, 0x47, 0x70, 0xcf, 0x60, 0x82, 0x97, 0xc1, 0x4e, 0x1b, 0x6d, 0x7e, 0x27,
	0x1e, 0x3c, 0x68, 0xcd, 0x0f, 0xa7, 0xb0, 0xbb, 0x66, 0x5d, 0xe2, 0x19, 0x1a, 0xcc, 0x62, 0x64,
	0x0c, 0x76, 0xb2, 0x28, 0xc5, 0xda, 0xb6, 0x27, 0xeb, 0x33, 0xdb, 0x03, 0xcf, 0x7d, 0x0b, 0x1d,
	0xc5, 0xd8, 0x5a, 0xba, 0x13, 0x86, 0x9f, 0x7f, 0xed, 0xc2, 0x44, 0x65, 0xa7, 0x2a, 0x4b, 0xd8,
	0x6b, 0xe8, 0xb5, 0x73, 0x74, 0x3d, 0x70, 0xad, 0x7e, 0xb1, 0xb9, 0xd5, 0x47, 0x2e, 0x78, 0xdc,
	0xa4, 0xc8, 0x2e, 0x97, 0xbd, 0x03, 0xef, 0x76, 0x0b, 0x6f, 0x67, 0x36, 0xfe, 0xe7, 0x99, 0x75,
	0xef, 0x92, 0xbd, 0x48, 0x2b, 0xa7, 0x14, 0x93, 0xfd, 0xab, 0x92, 0x93, 0xeb, 0x92, 0x93, 0xaf,
	0x25, 0x27, 0x37, 0x25, 0x27, 0xef, 0x2b, 0x4e, 0x3f, 0x55, 0x9c, 0x5c, 0x55, 0x9c, 0x5e, 0x57,
	0x9c, 0x7e, 0xab, 0x38, 0xfd, 0x51, 0x71, 0x72, 0x53, 0x71, 0xfa, 0xe1, 0x3b, 0x27, 0x27, 0xf7,
	0xeb, 0x45, 0x7c, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0x15, 0x6f, 0x6e, 0x11, 0x26, 0x03, 0x00,
	0x00,
}