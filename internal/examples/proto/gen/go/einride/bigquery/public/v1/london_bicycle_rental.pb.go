// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/london_bicycle_rental.proto

package publicv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protobuf schema for the BigQuery public table:
//
//	bigquery-public-data.london_bicycles.cycle_hire
type LondonBicycleRental struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RentalId                    int64                  `protobuf:"varint,1,opt,name=rental_id,json=rentalId,proto3" json:"rental_id,omitempty"`                                                               // INTEGER REQUIRED
	Duration                    *durationpb.Duration   `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`                                                                                // INTEGER NULLABLE
	BikeId                      int64                  `protobuf:"varint,3,opt,name=bike_id,json=bikeId,proto3" json:"bike_id,omitempty"`                                                                     // INTEGER NULLABLE
	EndDate                     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`                                                                   // TIMESTAMP NULLABLE
	EndStationId                int64                  `protobuf:"varint,5,opt,name=end_station_id,json=endStationId,proto3" json:"end_station_id,omitempty"`                                                 // INTEGER NULLABLE
	EndStationName              string                 `protobuf:"bytes,6,opt,name=end_station_name,json=endStationName,proto3" json:"end_station_name,omitempty"`                                            // STRING NULLABLE
	StartDate                   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`                                                             // TIMESTAMP NULLABLE
	StartStationId              int64                  `protobuf:"varint,8,opt,name=start_station_id,json=startStationId,proto3" json:"start_station_id,omitempty"`                                           // INTEGER NULLABLE
	StartStationName            string                 `protobuf:"bytes,9,opt,name=start_station_name,json=startStationName,proto3" json:"start_station_name,omitempty"`                                      // STRING NULLABLE
	EndStationLogicalTerminal   int64                  `protobuf:"varint,10,opt,name=end_station_logical_terminal,json=endStationLogicalTerminal,proto3" json:"end_station_logical_terminal,omitempty"`       // INTEGER NULLABLE
	StartStationLogicalTerminal int64                  `protobuf:"varint,11,opt,name=start_station_logical_terminal,json=startStationLogicalTerminal,proto3" json:"start_station_logical_terminal,omitempty"` // INTEGER NULLABLE
	EndStationPriorityId        int64                  `protobuf:"varint,12,opt,name=end_station_priority_id,json=endStationPriorityId,proto3" json:"end_station_priority_id,omitempty"`                      // INTEGER NULLABLE
}

func (x *LondonBicycleRental) Reset() {
	*x = LondonBicycleRental{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_london_bicycle_rental_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LondonBicycleRental) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LondonBicycleRental) ProtoMessage() {}

func (x *LondonBicycleRental) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_london_bicycle_rental_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LondonBicycleRental.ProtoReflect.Descriptor instead.
func (*LondonBicycleRental) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescGZIP(), []int{0}
}

func (x *LondonBicycleRental) GetRentalId() int64 {
	if x != nil {
		return x.RentalId
	}
	return 0
}

func (x *LondonBicycleRental) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *LondonBicycleRental) GetBikeId() int64 {
	if x != nil {
		return x.BikeId
	}
	return 0
}

func (x *LondonBicycleRental) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *LondonBicycleRental) GetEndStationId() int64 {
	if x != nil {
		return x.EndStationId
	}
	return 0
}

func (x *LondonBicycleRental) GetEndStationName() string {
	if x != nil {
		return x.EndStationName
	}
	return ""
}

func (x *LondonBicycleRental) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *LondonBicycleRental) GetStartStationId() int64 {
	if x != nil {
		return x.StartStationId
	}
	return 0
}

func (x *LondonBicycleRental) GetStartStationName() string {
	if x != nil {
		return x.StartStationName
	}
	return ""
}

func (x *LondonBicycleRental) GetEndStationLogicalTerminal() int64 {
	if x != nil {
		return x.EndStationLogicalTerminal
	}
	return 0
}

func (x *LondonBicycleRental) GetStartStationLogicalTerminal() int64 {
	if x != nil {
		return x.StartStationLogicalTerminal
	}
	return 0
}

func (x *LondonBicycleRental) GetEndStationPriorityId() int64 {
	if x != nil {
		return x.EndStationPriorityId
	}
	return 0
}

var File_einride_bigquery_public_v1_london_bicycle_rental_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x6e,
	0x64, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a, 0x13, 0x4c, 0x6f, 0x6e, 0x64, 0x6f, 0x6e,
	0x42, 0x69, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x69, 0x6b, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x62, 0x69, 0x6b, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x65, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x43, 0x0a, 0x1e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x63,
	0x61, 0x6c, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x17, 0x65, 0x6e,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x65, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x42, 0xaa, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x4c, 0x6f, 0x6e, 0x64, 0x6f, 0x6e, 0x42, 0x69, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x61, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x50, 0xaa, 0x02, 0x1a, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1d, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescData = file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDesc
)

func file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDescData
}

var file_einride_bigquery_public_v1_london_bicycle_rental_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_london_bicycle_rental_proto_goTypes = []interface{}{
	(*LondonBicycleRental)(nil),   // 0: einride.bigquery.public.v1.LondonBicycleRental
	(*durationpb.Duration)(nil),   // 1: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_einride_bigquery_public_v1_london_bicycle_rental_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.LondonBicycleRental.duration:type_name -> google.protobuf.Duration
	2, // 1: einride.bigquery.public.v1.LondonBicycleRental.end_date:type_name -> google.protobuf.Timestamp
	2, // 2: einride.bigquery.public.v1.LondonBicycleRental.start_date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_london_bicycle_rental_proto_init() }
func file_einride_bigquery_public_v1_london_bicycle_rental_proto_init() {
	if File_einride_bigquery_public_v1_london_bicycle_rental_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_london_bicycle_rental_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LondonBicycleRental); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_london_bicycle_rental_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_london_bicycle_rental_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_london_bicycle_rental_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_london_bicycle_rental_proto = out.File
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_rawDesc = nil
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_goTypes = nil
	file_einride_bigquery_public_v1_london_bicycle_rental_proto_depIdxs = nil
}
