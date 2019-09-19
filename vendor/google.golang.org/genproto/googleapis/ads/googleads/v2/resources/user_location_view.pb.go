// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/user_location_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A user location view.
//
// User Location View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at the actual physical location of
// the user by targeted or not targeted location. If other segment fields are
// used, you may get more than one row per country.
type UserLocationView struct {
	// The resource name of the user location view.
	// UserLocation view resource names have the form:
	//
	// `customers/{customer_id}/userLocationViews/{country_criterion_id}~{targeting_location}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Criterion Id for the country.
	CountryCriterionId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=country_criterion_id,json=countryCriterionId,proto3" json:"country_criterion_id,omitempty"`
	// Indicates whether location was targeted or not.
	TargetingLocation    *wrappers.BoolValue `protobuf:"bytes,3,opt,name=targeting_location,json=targetingLocation,proto3" json:"targeting_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserLocationView) Reset()         { *m = UserLocationView{} }
func (m *UserLocationView) String() string { return proto.CompactTextString(m) }
func (*UserLocationView) ProtoMessage()    {}
func (*UserLocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f59642a3431baa2, []int{0}
}

func (m *UserLocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLocationView.Unmarshal(m, b)
}
func (m *UserLocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLocationView.Marshal(b, m, deterministic)
}
func (m *UserLocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLocationView.Merge(m, src)
}
func (m *UserLocationView) XXX_Size() int {
	return xxx_messageInfo_UserLocationView.Size(m)
}
func (m *UserLocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLocationView.DiscardUnknown(m)
}

var xxx_messageInfo_UserLocationView proto.InternalMessageInfo

func (m *UserLocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserLocationView) GetCountryCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CountryCriterionId
	}
	return nil
}

func (m *UserLocationView) GetTargetingLocation() *wrappers.BoolValue {
	if m != nil {
		return m.TargetingLocation
	}
	return nil
}

func init() {
	proto.RegisterType((*UserLocationView)(nil), "google.ads.googleads.v2.resources.UserLocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/user_location_view.proto", fileDescriptor_0f59642a3431baa2)
}

var fileDescriptor_0f59642a3431baa2 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x49, 0x0a, 0x17, 0x6e, 0xee, 0x15, 0x34, 0x28, 0x94, 0x2a, 0xd2, 0x2a, 0x85, 0xae,
	0x26, 0x10, 0xc5, 0xc5, 0xb8, 0x4a, 0x5d, 0x94, 0x8a, 0x4a, 0x29, 0x98, 0x85, 0x04, 0xc2, 0x34,
	0x19, 0x87, 0x81, 0x74, 0xbe, 0x30, 0x33, 0x69, 0x71, 0xeb, 0xa3, 0xb8, 0xf4, 0x51, 0xdc, 0xf9,
	0x1a, 0x3e, 0x85, 0xe4, 0xcf, 0xcc, 0x42, 0x41, 0x77, 0x87, 0x99, 0xf3, 0x3b, 0xe7, 0x9b, 0xf9,
	0x3c, 0xcc, 0x00, 0x58, 0x41, 0x03, 0x92, 0xab, 0xa0, 0x95, 0xb5, 0xda, 0x84, 0x81, 0xa4, 0x0a,
	0x2a, 0x99, 0x51, 0x15, 0x54, 0x8a, 0xca, 0xb4, 0x80, 0x8c, 0x68, 0x0e, 0x22, 0xdd, 0x70, 0xba,
	0x45, 0xa5, 0x04, 0x0d, 0xfe, 0xa8, 0x05, 0x10, 0xc9, 0x15, 0xb2, 0x2c, 0xda, 0x84, 0xc8, 0xb2,
	0x83, 0xe3, 0x2e, 0xbe, 0x01, 0x56, 0xd5, 0x63, 0xb0, 0x95, 0xa4, 0x2c, 0xa9, 0x54, 0x6d, 0xc4,
	0xe0, 0xc8, 0xd4, 0x97, 0x3c, 0x20, 0x42, 0x80, 0x6e, 0x4a, 0xba, 0xdb, 0x93, 0x77, 0xc7, 0xdb,
	0xbd, 0x57, 0x54, 0xde, 0x74, 0xe5, 0x31, 0xa7, 0x5b, 0xff, 0xd4, 0xdb, 0x31, 0xf9, 0xa9, 0x20,
	0x6b, 0xda, 0x77, 0x86, 0xce, 0xe4, 0xef, 0xf2, 0xbf, 0x39, 0xbc, 0x23, 0x6b, 0xea, 0xdf, 0x7a,
	0xfb, 0x19, 0x54, 0x42, 0xcb, 0xa7, 0x34, 0x93, 0x5c, 0x53, 0x59, 0x8f, 0xce, 0xf3, 0xbe, 0x3b,
	0x74, 0x26, 0xff, 0xc2, 0xc3, 0x6e, 0x5c, 0x64, 0xc6, 0x42, 0x73, 0xa1, 0x2f, 0xce, 0x63, 0x52,
	0x54, 0x74, 0xe9, 0x77, 0xe0, 0x95, 0xe1, 0xe6, 0xb9, 0x3f, 0xf7, 0x7c, 0x4d, 0x24, 0xa3, 0x9a,
	0x0b, 0x66, 0xbf, 0xa2, 0xdf, 0x6b, 0xc2, 0x06, 0xdf, 0xc2, 0xa6, 0x00, 0x45, 0x9b, 0xb5, 0x67,
	0x29, 0xf3, 0x84, 0xe9, 0xb3, 0xeb, 0x8d, 0x33, 0x58, 0xa3, 0x5f, 0xff, 0x6e, 0x7a, 0xf0, 0xf5,
	0xe9, 0x8b, 0xba, 0x60, 0xe1, 0x3c, 0x5c, 0x77, 0x2c, 0x83, 0x82, 0x08, 0x86, 0x40, 0xb2, 0x80,
	0x51, 0xd1, 0xd4, 0x9b, 0x1d, 0x96, 0x5c, 0xfd, 0xb0, 0xd2, 0x4b, 0xab, 0x5e, 0xdc, 0xde, 0x2c,
	0x8a, 0x5e, 0xdd, 0xd1, 0xac, 0x8d, 0x8c, 0x72, 0x85, 0x5a, 0x59, 0xab, 0x38, 0x44, 0x4b, 0xe3,
	0x7c, 0x33, 0x9e, 0x24, 0xca, 0x55, 0x62, 0x3d, 0x49, 0x1c, 0x26, 0xd6, 0xf3, 0xe1, 0x8e, 0xdb,
	0x0b, 0x8c, 0xa3, 0x5c, 0x61, 0x6c, 0x5d, 0x18, 0xc7, 0x21, 0xc6, 0xd6, 0xb7, 0xfa, 0xd3, 0x0c,
	0x7b, 0xf6, 0x19, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x36, 0x49, 0xd0, 0x7e, 0x02, 0x00, 0x00,
}