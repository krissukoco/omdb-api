// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: omdb.proto

package omdb

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

type GetMovieByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieByIDRequest) Reset() {
	*x = GetMovieByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByIDRequest) ProtoMessage() {}

func (x *GetMovieByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByIDRequest.ProtoReflect.Descriptor instead.
func (*GetMovieByIDRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMovieByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year      string   `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Rated     string   `protobuf:"bytes,4,opt,name=rated,proto3" json:"rated,omitempty"`
	Genre     string   `protobuf:"bytes,5,opt,name=genre,proto3" json:"genre,omitempty"`
	Plot      string   `protobuf:"bytes,6,opt,name=plot,proto3" json:"plot,omitempty"`
	Director  string   `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Actors    []string `protobuf:"bytes,8,rep,name=actors,proto3" json:"actors,omitempty"`
	Language  string   `protobuf:"bytes,9,opt,name=language,proto3" json:"language,omitempty"`
	Country   string   `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	Type      string   `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	PosterUrl string   `protobuf:"bytes,12,opt,name=poster_url,json=posterUrl,proto3" json:"poster_url,omitempty"`
}

func (x *GetMovieByIDResponse) Reset() {
	*x = GetMovieByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByIDResponse) ProtoMessage() {}

func (x *GetMovieByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByIDResponse.ProtoReflect.Descriptor instead.
func (*GetMovieByIDResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{1}
}

func (x *GetMovieByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMovieByIDResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMovieByIDResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMovieByIDResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetMovieByIDResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetMovieByIDResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetMovieByIDResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetMovieByIDResponse) GetActors() []string {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *GetMovieByIDResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetMovieByIDResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetMovieByIDResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMovieByIDResponse) GetPosterUrl() string {
	if x != nil {
		return x.PosterUrl
	}
	return ""
}

type SearchMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Page  uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchMoviesRequest) Reset() {
	*x = SearchMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoviesRequest) ProtoMessage() {}

func (x *SearchMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoviesRequest.ProtoReflect.Descriptor instead.
func (*SearchMoviesRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{2}
}

func (x *SearchMoviesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchMoviesRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SearchMoviesRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type SearchMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies       []*MovieResult `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	TotalResults uint64         `protobuf:"varint,2,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
}

func (x *SearchMoviesResponse) Reset() {
	*x = SearchMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoviesResponse) ProtoMessage() {}

func (x *SearchMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoviesResponse.ProtoReflect.Descriptor instead.
func (*SearchMoviesResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{3}
}

func (x *SearchMoviesResponse) GetMovies() []*MovieResult {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *SearchMoviesResponse) GetTotalResults() uint64 {
	if x != nil {
		return x.TotalResults
	}
	return 0
}

type MovieResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year      string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	PosterUrl string `protobuf:"bytes,5,opt,name=poster_url,json=posterUrl,proto3" json:"poster_url,omitempty"`
}

func (x *MovieResult) Reset() {
	*x = MovieResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResult) ProtoMessage() {}

func (x *MovieResult) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResult.ProtoReflect.Descriptor instead.
func (*MovieResult) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{4}
}

func (x *MovieResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MovieResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieResult) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieResult) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieResult) GetPosterUrl() string {
	if x != nil {
		return x.PosterUrl
	}
	return ""
}

var File_omdb_proto protoreflect.FileDescriptor

var file_omdb_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6f, 0x6d,
	0x64, 0x62, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xad, 0x02, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x53, 0x0a, 0x13, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x66,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x7a, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x4f, 0x4d, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x19, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x6f, 0x6d,
	0x64, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x6f, 0x6d, 0x64, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omdb_proto_rawDescOnce sync.Once
	file_omdb_proto_rawDescData = file_omdb_proto_rawDesc
)

func file_omdb_proto_rawDescGZIP() []byte {
	file_omdb_proto_rawDescOnce.Do(func() {
		file_omdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_omdb_proto_rawDescData)
	})
	return file_omdb_proto_rawDescData
}

var file_omdb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_omdb_proto_goTypes = []interface{}{
	(*GetMovieByIDRequest)(nil),  // 0: omdb.GetMovieByIDRequest
	(*GetMovieByIDResponse)(nil), // 1: omdb.GetMovieByIDResponse
	(*SearchMoviesRequest)(nil),  // 2: omdb.SearchMoviesRequest
	(*SearchMoviesResponse)(nil), // 3: omdb.SearchMoviesResponse
	(*MovieResult)(nil),          // 4: omdb.MovieResult
}
var file_omdb_proto_depIdxs = []int32{
	4, // 0: omdb.SearchMoviesResponse.movies:type_name -> omdb.MovieResult
	0, // 1: omdb.OMDBService.GetMovieByID:input_type -> omdb.GetMovieByIDRequest
	2, // 2: omdb.OMDBService.SearchMovies:input_type -> omdb.SearchMoviesRequest
	1, // 3: omdb.OMDBService.GetMovieByID:output_type -> omdb.GetMovieByIDResponse
	3, // 4: omdb.OMDBService.SearchMovies:output_type -> omdb.SearchMoviesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_omdb_proto_init() }
func file_omdb_proto_init() {
	if File_omdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByIDRequest); i {
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
		file_omdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByIDResponse); i {
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
		file_omdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoviesRequest); i {
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
		file_omdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoviesResponse); i {
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
		file_omdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResult); i {
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
			RawDescriptor: file_omdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omdb_proto_goTypes,
		DependencyIndexes: file_omdb_proto_depIdxs,
		MessageInfos:      file_omdb_proto_msgTypes,
	}.Build()
	File_omdb_proto = out.File
	file_omdb_proto_rawDesc = nil
	file_omdb_proto_goTypes = nil
	file_omdb_proto_depIdxs = nil
}
