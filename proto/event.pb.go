// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/event.proto

package bookingpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State int32

const (
	State_unknown          State = 0
	State_unconfirmed      State = 1
	State_confirmed        State = 2
	State_payment_received State = 3
	State_payment_pending  State = 4
	State_planned          State = 5
	State_canceled         State = 6
	State_checked_in       State = 7
	State_checked_out      State = 8
	State_review_pending   State = 9
	State_reviewed         State = 10
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0:  "unknown",
		1:  "unconfirmed",
		2:  "confirmed",
		3:  "payment_received",
		4:  "payment_pending",
		5:  "planned",
		6:  "canceled",
		7:  "checked_in",
		8:  "checked_out",
		9:  "review_pending",
		10: "reviewed",
	}
	State_value = map[string]int32{
		"unknown":          0,
		"unconfirmed":      1,
		"confirmed":        2,
		"payment_received": 3,
		"payment_pending":  4,
		"planned":          5,
		"canceled":         6,
		"checked_in":       7,
		"checked_out":      8,
		"review_pending":   9,
		"reviewed":         10,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_event_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_proto_event_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FromEpochMillis int64  `protobuf:"varint,3,opt,name=from_epoch_millis,json=fromEpochMillis,proto3" json:"from_epoch_millis,omitempty"`
	ToEpochMillis   int64  `protobuf:"varint,4,opt,name=to_epoch_millis,json=toEpochMillis,proto3" json:"to_epoch_millis,omitempty"`
	HotelName       string `protobuf:"bytes,5,opt,name=hotel_name,json=hotelName,proto3" json:"hotel_name,omitempty"`
	HotelId         string `protobuf:"bytes,6,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	FlightId        string `protobuf:"bytes,7,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	AirlineName     string `protobuf:"bytes,8,opt,name=airline_name,json=airlineName,proto3" json:"airline_name,omitempty"`
	BookingState    State  `protobuf:"varint,9,opt,name=booking_state,json=bookingState,proto3,enum=proto.State" json:"booking_state,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Event) GetFromEpochMillis() int64 {
	if x != nil {
		return x.FromEpochMillis
	}
	return 0
}

func (x *Event) GetToEpochMillis() int64 {
	if x != nil {
		return x.ToEpochMillis
	}
	return 0
}

func (x *Event) GetHotelName() string {
	if x != nil {
		return x.HotelName
	}
	return ""
}

func (x *Event) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *Event) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *Event) GetAirlineName() string {
	if x != nil {
		return x.AirlineName
	}
	return ""
}

func (x *Event) GetBookingState() State {
	if x != nil {
		return x.BookingState
	}
	return State_unknown
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x11, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x5f,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x69, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0d, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0xbd,
	0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x12, 0x0b, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x10, 0x05, 0x12, 0x0c, 0x0a,
	0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x09,
	0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x10, 0x0a, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x75, 0x63, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x77, 0x73,
	0x2d, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x2d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_event_proto_goTypes = []interface{}{
	(State)(0),    // 0: proto.State
	(*Event)(nil), // 1: proto.Event
}
var file_proto_event_proto_depIdxs = []int32{
	0, // 0: proto.Event.booking_state:type_name -> proto.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		EnumInfos:         file_proto_event_proto_enumTypes,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}
