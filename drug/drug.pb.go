// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: drug.proto

package drug

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

type Drug struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt      string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt      string `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Detail         string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail,omitempty"`
	DrugType       string `protobuf:"bytes,6,opt,name=drug_type,json=drugType,proto3" json:"drug_type,omitempty"`
	IsPrescription string `protobuf:"bytes,7,opt,name=is_prescription,json=isPrescription,proto3" json:"is_prescription,omitempty"`
	InsDrugs       string `protobuf:"bytes,8,opt,name=ins_drugs,json=insDrugs,proto3" json:"ins_drugs,omitempty"`
	Dosage         string `protobuf:"bytes,9,opt,name=dosage,proto3" json:"dosage,omitempty"`
	Taboo          string `protobuf:"bytes,10,opt,name=taboo,proto3" json:"taboo,omitempty"`
}

func (x *Drug) Reset() {
	*x = Drug{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Drug) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drug) ProtoMessage() {}

func (x *Drug) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drug.ProtoReflect.Descriptor instead.
func (*Drug) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{0}
}

func (x *Drug) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Drug) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Drug) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Drug) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Drug) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Drug) GetDrugType() string {
	if x != nil {
		return x.DrugType
	}
	return ""
}

func (x *Drug) GetIsPrescription() string {
	if x != nil {
		return x.IsPrescription
	}
	return ""
}

func (x *Drug) GetInsDrugs() string {
	if x != nil {
		return x.InsDrugs
	}
	return ""
}

func (x *Drug) GetDosage() string {
	if x != nil {
		return x.Dosage
	}
	return ""
}

func (x *Drug) GetTaboo() string {
	if x != nil {
		return x.Taboo
	}
	return ""
}

type CreateDrugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Detail         string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail,omitempty"`
	DrugType       string `protobuf:"bytes,6,opt,name=drug_type,json=drugType,proto3" json:"drug_type,omitempty"`
	IsPrescription string `protobuf:"bytes,7,opt,name=is_prescription,json=isPrescription,proto3" json:"is_prescription,omitempty"`
	InsDrugs       string `protobuf:"bytes,8,opt,name=ins_drugs,json=insDrugs,proto3" json:"ins_drugs,omitempty"`
	Dosage         string `protobuf:"bytes,9,opt,name=dosage,proto3" json:"dosage,omitempty"`
	Taboo          string `protobuf:"bytes,10,opt,name=taboo,proto3" json:"taboo,omitempty"`
}

func (x *CreateDrugRequest) Reset() {
	*x = CreateDrugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDrugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDrugRequest) ProtoMessage() {}

func (x *CreateDrugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDrugRequest.ProtoReflect.Descriptor instead.
func (*CreateDrugRequest) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDrugRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDrugRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *CreateDrugRequest) GetDrugType() string {
	if x != nil {
		return x.DrugType
	}
	return ""
}

func (x *CreateDrugRequest) GetIsPrescription() string {
	if x != nil {
		return x.IsPrescription
	}
	return ""
}

func (x *CreateDrugRequest) GetInsDrugs() string {
	if x != nil {
		return x.InsDrugs
	}
	return ""
}

func (x *CreateDrugRequest) GetDosage() string {
	if x != nil {
		return x.Dosage
	}
	return ""
}

func (x *CreateDrugRequest) GetTaboo() string {
	if x != nil {
		return x.Taboo
	}
	return ""
}

type CreateDrugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drug []*Drug `protobuf:"bytes,1,rep,name=Drug,proto3" json:"Drug,omitempty"`
}

func (x *CreateDrugResponse) Reset() {
	*x = CreateDrugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDrugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDrugResponse) ProtoMessage() {}

func (x *CreateDrugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDrugResponse.ProtoReflect.Descriptor instead.
func (*CreateDrugResponse) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDrugResponse) GetDrug() []*Drug {
	if x != nil {
		return x.Drug
	}
	return nil
}

type GetDrugDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDrugDetailRequest) Reset() {
	*x = GetDrugDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugDetailRequest) ProtoMessage() {}

func (x *GetDrugDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugDetailRequest.ProtoReflect.Descriptor instead.
func (*GetDrugDetailRequest) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{3}
}

func (x *GetDrugDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDrugDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drug *Drug `protobuf:"bytes,1,opt,name=Drug,proto3" json:"Drug,omitempty"`
}

func (x *GetDrugDetailResponse) Reset() {
	*x = GetDrugDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugDetailResponse) ProtoMessage() {}

func (x *GetDrugDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugDetailResponse.ProtoReflect.Descriptor instead.
func (*GetDrugDetailResponse) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{4}
}

func (x *GetDrugDetailResponse) GetDrug() *Drug {
	if x != nil {
		return x.Drug
	}
	return nil
}

type GetDrugListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DrugType       string `protobuf:"bytes,2,opt,name=drug_type,json=drugType,proto3" json:"drug_type,omitempty"`
	IsPrescription string `protobuf:"bytes,3,opt,name=is_prescription,json=isPrescription,proto3" json:"is_prescription,omitempty"`
	InsDrugs       string `protobuf:"bytes,4,opt,name=ins_drugs,json=insDrugs,proto3" json:"ins_drugs,omitempty"`
	Dosage         string `protobuf:"bytes,5,opt,name=dosage,proto3" json:"dosage,omitempty"`
	Taboo          string `protobuf:"bytes,6,opt,name=taboo,proto3" json:"taboo,omitempty"`
	Page           string `protobuf:"bytes,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize       string `protobuf:"bytes,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetDrugListRequest) Reset() {
	*x = GetDrugListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugListRequest) ProtoMessage() {}

func (x *GetDrugListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugListRequest.ProtoReflect.Descriptor instead.
func (*GetDrugListRequest) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{5}
}

func (x *GetDrugListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDrugListRequest) GetDrugType() string {
	if x != nil {
		return x.DrugType
	}
	return ""
}

func (x *GetDrugListRequest) GetIsPrescription() string {
	if x != nil {
		return x.IsPrescription
	}
	return ""
}

func (x *GetDrugListRequest) GetInsDrugs() string {
	if x != nil {
		return x.InsDrugs
	}
	return ""
}

func (x *GetDrugListRequest) GetDosage() string {
	if x != nil {
		return x.Dosage
	}
	return ""
}

func (x *GetDrugListRequest) GetTaboo() string {
	if x != nil {
		return x.Taboo
	}
	return ""
}

func (x *GetDrugListRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetDrugListRequest) GetPageSize() string {
	if x != nil {
		return x.PageSize
	}
	return ""
}

type GetDrugListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drug  []*Drug `protobuf:"bytes,1,rep,name=Drug,proto3" json:"Drug,omitempty"`
	Total string  `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetDrugListResponse) Reset() {
	*x = GetDrugListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugListResponse) ProtoMessage() {}

func (x *GetDrugListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugListResponse.ProtoReflect.Descriptor instead.
func (*GetDrugListResponse) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{6}
}

func (x *GetDrugListResponse) GetDrug() []*Drug {
	if x != nil {
		return x.Drug
	}
	return nil
}

func (x *GetDrugListResponse) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

type UpdateDrugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Detail         string `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail,omitempty"`
	DrugType       string `protobuf:"bytes,6,opt,name=drug_type,json=drugType,proto3" json:"drug_type,omitempty"`
	IsPrescription string `protobuf:"bytes,7,opt,name=is_prescription,json=isPrescription,proto3" json:"is_prescription,omitempty"`
	InsDrugs       string `protobuf:"bytes,8,opt,name=ins_drugs,json=insDrugs,proto3" json:"ins_drugs,omitempty"`
	Dosage         string `protobuf:"bytes,9,opt,name=dosage,proto3" json:"dosage,omitempty"`
	Taboo          string `protobuf:"bytes,10,opt,name=taboo,proto3" json:"taboo,omitempty"`
}

func (x *UpdateDrugRequest) Reset() {
	*x = UpdateDrugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDrugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDrugRequest) ProtoMessage() {}

func (x *UpdateDrugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDrugRequest.ProtoReflect.Descriptor instead.
func (*UpdateDrugRequest) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDrugRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDrugRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDrugRequest) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *UpdateDrugRequest) GetDrugType() string {
	if x != nil {
		return x.DrugType
	}
	return ""
}

func (x *UpdateDrugRequest) GetIsPrescription() string {
	if x != nil {
		return x.IsPrescription
	}
	return ""
}

func (x *UpdateDrugRequest) GetInsDrugs() string {
	if x != nil {
		return x.InsDrugs
	}
	return ""
}

func (x *UpdateDrugRequest) GetDosage() string {
	if x != nil {
		return x.Dosage
	}
	return ""
}

func (x *UpdateDrugRequest) GetTaboo() string {
	if x != nil {
		return x.Taboo
	}
	return ""
}

type DeleteDrugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDrugRequest) Reset() {
	*x = DeleteDrugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDrugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDrugRequest) ProtoMessage() {}

func (x *DeleteDrugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDrugRequest.ProtoReflect.Descriptor instead.
func (*DeleteDrugRequest) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDrugRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drug_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_drug_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_drug_proto_rawDescGZIP(), []int{9}
}

var File_drug_proto protoreflect.FileDescriptor

var file_drug_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x72,
	0x75, 0x67, 0x22, 0x8f, 0x02, 0x0a, 0x04, 0x44, 0x72, 0x75, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x75, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x75, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x72, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x73,
	0x5f, 0x64, 0x72, 0x75, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e,
	0x73, 0x44, 0x72, 0x75, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6f, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x75, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x75, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x73,
	0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6e, 0x73, 0x5f, 0x64, 0x72, 0x75, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x44, 0x72, 0x75, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x44, 0x72, 0x75, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x72,
	0x75, 0x67, 0x2e, 0x44, 0x72, 0x75, 0x67, 0x52, 0x04, 0x44, 0x72, 0x75, 0x67, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x44, 0x72, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64,
	0x72, 0x75, 0x67, 0x2e, 0x44, 0x72, 0x75, 0x67, 0x52, 0x04, 0x44, 0x72, 0x75, 0x67, 0x22, 0xea,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x75,
	0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72,
	0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x5f, 0x64, 0x72, 0x75, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x44, 0x72, 0x75, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x44, 0x72, 0x75, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x44, 0x72, 0x75, 0x67, 0x52, 0x04, 0x44, 0x72,
	0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xe0, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72,
	0x75, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x72, 0x75, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x5f, 0x64, 0x72, 0x75, 0x67, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x44, 0x72, 0x75, 0x67, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6f, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xc3, 0x02, 0x0a, 0x0a, 0x44, 0x72,
	0x75, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x12, 0x17, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x75,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x44, 0x72, 0x75, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x72, 0x75,
	0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x75, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64,
	0x72, 0x75, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x75, 0x67, 0x12, 0x17, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b,
	0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x12, 0x17, 0x2e, 0x64, 0x72, 0x75, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x64, 0x72, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_drug_proto_rawDescOnce sync.Once
	file_drug_proto_rawDescData = file_drug_proto_rawDesc
)

func file_drug_proto_rawDescGZIP() []byte {
	file_drug_proto_rawDescOnce.Do(func() {
		file_drug_proto_rawDescData = protoimpl.X.CompressGZIP(file_drug_proto_rawDescData)
	})
	return file_drug_proto_rawDescData
}

var file_drug_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_drug_proto_goTypes = []interface{}{
	(*Drug)(nil),                  // 0: drug.Drug
	(*CreateDrugRequest)(nil),     // 1: drug.CreateDrugRequest
	(*CreateDrugResponse)(nil),    // 2: drug.CreateDrugResponse
	(*GetDrugDetailRequest)(nil),  // 3: drug.GetDrugDetailRequest
	(*GetDrugDetailResponse)(nil), // 4: drug.GetDrugDetailResponse
	(*GetDrugListRequest)(nil),    // 5: drug.GetDrugListRequest
	(*GetDrugListResponse)(nil),   // 6: drug.GetDrugListResponse
	(*UpdateDrugRequest)(nil),     // 7: drug.UpdateDrugRequest
	(*DeleteDrugRequest)(nil),     // 8: drug.DeleteDrugRequest
	(*Empty)(nil),                 // 9: drug.Empty
}
var file_drug_proto_depIdxs = []int32{
	0, // 0: drug.CreateDrugResponse.Drug:type_name -> drug.Drug
	0, // 1: drug.GetDrugDetailResponse.Drug:type_name -> drug.Drug
	0, // 2: drug.GetDrugListResponse.Drug:type_name -> drug.Drug
	1, // 3: drug.DrugServer.CreateDrug:input_type -> drug.CreateDrugRequest
	3, // 4: drug.DrugServer.GetDrugDetail:input_type -> drug.GetDrugDetailRequest
	5, // 5: drug.DrugServer.GetDrugList:input_type -> drug.GetDrugListRequest
	7, // 6: drug.DrugServer.UpdateDrug:input_type -> drug.UpdateDrugRequest
	8, // 7: drug.DrugServer.DeleteDrug:input_type -> drug.DeleteDrugRequest
	2, // 8: drug.DrugServer.CreateDrug:output_type -> drug.CreateDrugResponse
	4, // 9: drug.DrugServer.GetDrugDetail:output_type -> drug.GetDrugDetailResponse
	6, // 10: drug.DrugServer.GetDrugList:output_type -> drug.GetDrugListResponse
	9, // 11: drug.DrugServer.UpdateDrug:output_type -> drug.Empty
	9, // 12: drug.DrugServer.DeleteDrug:output_type -> drug.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_drug_proto_init() }
func file_drug_proto_init() {
	if File_drug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Drug); i {
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
		file_drug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDrugRequest); i {
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
		file_drug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDrugResponse); i {
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
		file_drug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugDetailRequest); i {
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
		file_drug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugDetailResponse); i {
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
		file_drug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugListRequest); i {
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
		file_drug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugListResponse); i {
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
		file_drug_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDrugRequest); i {
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
		file_drug_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDrugRequest); i {
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
		file_drug_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_drug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drug_proto_goTypes,
		DependencyIndexes: file_drug_proto_depIdxs,
		MessageInfos:      file_drug_proto_msgTypes,
	}.Build()
	File_drug_proto = out.File
	file_drug_proto_rawDesc = nil
	file_drug_proto_goTypes = nil
	file_drug_proto_depIdxs = nil
}
