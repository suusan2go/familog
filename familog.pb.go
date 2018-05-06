// Code generated by protoc-gen-go. DO NOT EDIT.
// source: familog.proto

/*
Package familog is a generated protocol buffer package.

It is generated from these files:
	familog.proto

It has these top-level messages:
	User
	Diary
	DiaryEntry
	StartDiaryRequest
	StartDiaryResponse
	SubscribeDiaryRequest
	SubscribeDiaryResponse
	AllDiariesRequest
	AllDiariesResponse
	FindDiaryEntriesRequest
	FindDiaryEntriesResponse
	ReadDiaryEntryRequest
	ReadDiaryEntryResponse
	DiaryEntryForm
	PublishDiaryEntryRequest
	PublishDiaryEntryResponse
	UpdateDiaryEntryRequest
	UpdateDiaryEntryResponse
	ShowMyProfileRequest
	ShowMyProfileResponse
	FindProfileByUserIdRequest
	FindProfileByUserIdResponse
	UpdateProfileRequest
	UpdateProfileResponse
*/
package familog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type Diary struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
}

func (m *Diary) Reset()                    { *m = Diary{} }
func (m *Diary) String() string            { return proto.CompactTextString(m) }
func (*Diary) ProtoMessage()               {}
func (*Diary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Diary) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Diary) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type DiaryEntry struct {
	Id            int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Body          string                     `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	Emoji         string                     `protobuf:"bytes,3,opt,name=emoji" json:"emoji,omitempty"`
	User          *User                      `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	CoverImageUrl string                     `protobuf:"bytes,5,opt,name=cover_image_url,json=coverImageUrl" json:"cover_image_url,omitempty"`
	ImageUrls     []string                   `protobuf:"bytes,6,rep,name=image_urls,json=imageUrls" json:"image_urls,omitempty"`
	PublishedAt   *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=published_at,json=publishedAt" json:"published_at,omitempty"`
}

func (m *DiaryEntry) Reset()                    { *m = DiaryEntry{} }
func (m *DiaryEntry) String() string            { return proto.CompactTextString(m) }
func (*DiaryEntry) ProtoMessage()               {}
func (*DiaryEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DiaryEntry) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DiaryEntry) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *DiaryEntry) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func (m *DiaryEntry) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *DiaryEntry) GetCoverImageUrl() string {
	if m != nil {
		return m.CoverImageUrl
	}
	return ""
}

func (m *DiaryEntry) GetImageUrls() []string {
	if m != nil {
		return m.ImageUrls
	}
	return nil
}

func (m *DiaryEntry) GetPublishedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.PublishedAt
	}
	return nil
}

type StartDiaryRequest struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
}

func (m *StartDiaryRequest) Reset()                    { *m = StartDiaryRequest{} }
func (m *StartDiaryRequest) String() string            { return proto.CompactTextString(m) }
func (*StartDiaryRequest) ProtoMessage()               {}
func (*StartDiaryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StartDiaryRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type StartDiaryResponse struct {
	Diary *Diary `protobuf:"bytes,1,opt,name=diary" json:"diary,omitempty"`
}

func (m *StartDiaryResponse) Reset()                    { *m = StartDiaryResponse{} }
func (m *StartDiaryResponse) String() string            { return proto.CompactTextString(m) }
func (*StartDiaryResponse) ProtoMessage()               {}
func (*StartDiaryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StartDiaryResponse) GetDiary() *Diary {
	if m != nil {
		return m.Diary
	}
	return nil
}

type SubscribeDiaryRequest struct {
	SecretCode string `protobuf:"bytes,1,opt,name=secret_code,json=secretCode" json:"secret_code,omitempty"`
}

func (m *SubscribeDiaryRequest) Reset()                    { *m = SubscribeDiaryRequest{} }
func (m *SubscribeDiaryRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeDiaryRequest) ProtoMessage()               {}
func (*SubscribeDiaryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SubscribeDiaryRequest) GetSecretCode() string {
	if m != nil {
		return m.SecretCode
	}
	return ""
}

type SubscribeDiaryResponse struct {
}

func (m *SubscribeDiaryResponse) Reset()                    { *m = SubscribeDiaryResponse{} }
func (m *SubscribeDiaryResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeDiaryResponse) ProtoMessage()               {}
func (*SubscribeDiaryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type AllDiariesRequest struct {
}

func (m *AllDiariesRequest) Reset()                    { *m = AllDiariesRequest{} }
func (m *AllDiariesRequest) String() string            { return proto.CompactTextString(m) }
func (*AllDiariesRequest) ProtoMessage()               {}
func (*AllDiariesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AllDiariesResponse struct {
	Diaries []*Diary `protobuf:"bytes,1,rep,name=diaries" json:"diaries,omitempty"`
}

func (m *AllDiariesResponse) Reset()                    { *m = AllDiariesResponse{} }
func (m *AllDiariesResponse) String() string            { return proto.CompactTextString(m) }
func (*AllDiariesResponse) ProtoMessage()               {}
func (*AllDiariesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllDiariesResponse) GetDiaries() []*Diary {
	if m != nil {
		return m.Diaries
	}
	return nil
}

type FindDiaryEntriesRequest struct {
	DiaryId string `protobuf:"bytes,1,opt,name=diary_id,json=diaryId" json:"diary_id,omitempty"`
}

func (m *FindDiaryEntriesRequest) Reset()                    { *m = FindDiaryEntriesRequest{} }
func (m *FindDiaryEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindDiaryEntriesRequest) ProtoMessage()               {}
func (*FindDiaryEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FindDiaryEntriesRequest) GetDiaryId() string {
	if m != nil {
		return m.DiaryId
	}
	return ""
}

type FindDiaryEntriesResponse struct {
	DiaryEntries []*DiaryEntry `protobuf:"bytes,1,rep,name=diary_entries,json=diaryEntries" json:"diary_entries,omitempty"`
}

func (m *FindDiaryEntriesResponse) Reset()                    { *m = FindDiaryEntriesResponse{} }
func (m *FindDiaryEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindDiaryEntriesResponse) ProtoMessage()               {}
func (*FindDiaryEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *FindDiaryEntriesResponse) GetDiaryEntries() []*DiaryEntry {
	if m != nil {
		return m.DiaryEntries
	}
	return nil
}

type ReadDiaryEntryRequest struct {
	DiaryEntryId string `protobuf:"bytes,1,opt,name=diary_entry_id,json=diaryEntryId" json:"diary_entry_id,omitempty"`
}

func (m *ReadDiaryEntryRequest) Reset()                    { *m = ReadDiaryEntryRequest{} }
func (m *ReadDiaryEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadDiaryEntryRequest) ProtoMessage()               {}
func (*ReadDiaryEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReadDiaryEntryRequest) GetDiaryEntryId() string {
	if m != nil {
		return m.DiaryEntryId
	}
	return ""
}

type ReadDiaryEntryResponse struct {
	DiaryEntries []*DiaryEntry `protobuf:"bytes,1,rep,name=diary_entries,json=diaryEntries" json:"diary_entries,omitempty"`
}

func (m *ReadDiaryEntryResponse) Reset()                    { *m = ReadDiaryEntryResponse{} }
func (m *ReadDiaryEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadDiaryEntryResponse) ProtoMessage()               {}
func (*ReadDiaryEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReadDiaryEntryResponse) GetDiaryEntries() []*DiaryEntry {
	if m != nil {
		return m.DiaryEntries
	}
	return nil
}

type DiaryEntryForm struct {
	Body           string   `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Emoji          string   `protobuf:"bytes,2,opt,name=emoji" json:"emoji,omitempty"`
	CoverImageBlob string   `protobuf:"bytes,3,opt,name=cover_image_blob,json=coverImageBlob" json:"cover_image_blob,omitempty"`
	ImageBlobs     []string `protobuf:"bytes,4,rep,name=image_blobs,json=imageBlobs" json:"image_blobs,omitempty"`
}

func (m *DiaryEntryForm) Reset()                    { *m = DiaryEntryForm{} }
func (m *DiaryEntryForm) String() string            { return proto.CompactTextString(m) }
func (*DiaryEntryForm) ProtoMessage()               {}
func (*DiaryEntryForm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DiaryEntryForm) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *DiaryEntryForm) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func (m *DiaryEntryForm) GetCoverImageBlob() string {
	if m != nil {
		return m.CoverImageBlob
	}
	return ""
}

func (m *DiaryEntryForm) GetImageBlobs() []string {
	if m != nil {
		return m.ImageBlobs
	}
	return nil
}

type PublishDiaryEntryRequest struct {
	DiaryEntryForm *DiaryEntryForm `protobuf:"bytes,1,opt,name=diary_entry_form,json=diaryEntryForm" json:"diary_entry_form,omitempty"`
}

func (m *PublishDiaryEntryRequest) Reset()                    { *m = PublishDiaryEntryRequest{} }
func (m *PublishDiaryEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishDiaryEntryRequest) ProtoMessage()               {}
func (*PublishDiaryEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PublishDiaryEntryRequest) GetDiaryEntryForm() *DiaryEntryForm {
	if m != nil {
		return m.DiaryEntryForm
	}
	return nil
}

type PublishDiaryEntryResponse struct {
	DiaryEntry *DiaryEntry `protobuf:"bytes,1,opt,name=diary_entry,json=diaryEntry" json:"diary_entry,omitempty"`
}

func (m *PublishDiaryEntryResponse) Reset()                    { *m = PublishDiaryEntryResponse{} }
func (m *PublishDiaryEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishDiaryEntryResponse) ProtoMessage()               {}
func (*PublishDiaryEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *PublishDiaryEntryResponse) GetDiaryEntry() *DiaryEntry {
	if m != nil {
		return m.DiaryEntry
	}
	return nil
}

type UpdateDiaryEntryRequest struct {
	DiaryEntryId   string          `protobuf:"bytes,1,opt,name=diary_entry_id,json=diaryEntryId" json:"diary_entry_id,omitempty"`
	DiaryEntryForm *DiaryEntryForm `protobuf:"bytes,2,opt,name=diary_entry_form,json=diaryEntryForm" json:"diary_entry_form,omitempty"`
}

func (m *UpdateDiaryEntryRequest) Reset()                    { *m = UpdateDiaryEntryRequest{} }
func (m *UpdateDiaryEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDiaryEntryRequest) ProtoMessage()               {}
func (*UpdateDiaryEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *UpdateDiaryEntryRequest) GetDiaryEntryId() string {
	if m != nil {
		return m.DiaryEntryId
	}
	return ""
}

func (m *UpdateDiaryEntryRequest) GetDiaryEntryForm() *DiaryEntryForm {
	if m != nil {
		return m.DiaryEntryForm
	}
	return nil
}

type UpdateDiaryEntryResponse struct {
	DiaryEntry *DiaryEntry `protobuf:"bytes,1,opt,name=diary_entry,json=diaryEntry" json:"diary_entry,omitempty"`
}

func (m *UpdateDiaryEntryResponse) Reset()                    { *m = UpdateDiaryEntryResponse{} }
func (m *UpdateDiaryEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateDiaryEntryResponse) ProtoMessage()               {}
func (*UpdateDiaryEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *UpdateDiaryEntryResponse) GetDiaryEntry() *DiaryEntry {
	if m != nil {
		return m.DiaryEntry
	}
	return nil
}

type ShowMyProfileRequest struct {
}

func (m *ShowMyProfileRequest) Reset()                    { *m = ShowMyProfileRequest{} }
func (m *ShowMyProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowMyProfileRequest) ProtoMessage()               {}
func (*ShowMyProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type ShowMyProfileResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ShowMyProfileResponse) Reset()                    { *m = ShowMyProfileResponse{} }
func (m *ShowMyProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowMyProfileResponse) ProtoMessage()               {}
func (*ShowMyProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ShowMyProfileResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type FindProfileByUserIdRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *FindProfileByUserIdRequest) Reset()                    { *m = FindProfileByUserIdRequest{} }
func (m *FindProfileByUserIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FindProfileByUserIdRequest) ProtoMessage()               {}
func (*FindProfileByUserIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *FindProfileByUserIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FindProfileByUserIdResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *FindProfileByUserIdResponse) Reset()                    { *m = FindProfileByUserIdResponse{} }
func (m *FindProfileByUserIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FindProfileByUserIdResponse) ProtoMessage()               {}
func (*FindProfileByUserIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *FindProfileByUserIdResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateProfileRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	AvatarBlob string `protobuf:"bytes,2,opt,name=avatar_blob,json=avatarBlob" json:"avatar_blob,omitempty"`
}

func (m *UpdateProfileRequest) Reset()                    { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()               {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *UpdateProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateProfileRequest) GetAvatarBlob() string {
	if m != nil {
		return m.AvatarBlob
	}
	return ""
}

type UpdateProfileResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateProfileResponse) Reset()                    { *m = UpdateProfileResponse{} }
func (m *UpdateProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateProfileResponse) ProtoMessage()               {}
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *UpdateProfileResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "familog.User")
	proto.RegisterType((*Diary)(nil), "familog.Diary")
	proto.RegisterType((*DiaryEntry)(nil), "familog.DiaryEntry")
	proto.RegisterType((*StartDiaryRequest)(nil), "familog.StartDiaryRequest")
	proto.RegisterType((*StartDiaryResponse)(nil), "familog.StartDiaryResponse")
	proto.RegisterType((*SubscribeDiaryRequest)(nil), "familog.SubscribeDiaryRequest")
	proto.RegisterType((*SubscribeDiaryResponse)(nil), "familog.SubscribeDiaryResponse")
	proto.RegisterType((*AllDiariesRequest)(nil), "familog.AllDiariesRequest")
	proto.RegisterType((*AllDiariesResponse)(nil), "familog.AllDiariesResponse")
	proto.RegisterType((*FindDiaryEntriesRequest)(nil), "familog.FindDiaryEntriesRequest")
	proto.RegisterType((*FindDiaryEntriesResponse)(nil), "familog.FindDiaryEntriesResponse")
	proto.RegisterType((*ReadDiaryEntryRequest)(nil), "familog.ReadDiaryEntryRequest")
	proto.RegisterType((*ReadDiaryEntryResponse)(nil), "familog.ReadDiaryEntryResponse")
	proto.RegisterType((*DiaryEntryForm)(nil), "familog.DiaryEntryForm")
	proto.RegisterType((*PublishDiaryEntryRequest)(nil), "familog.PublishDiaryEntryRequest")
	proto.RegisterType((*PublishDiaryEntryResponse)(nil), "familog.PublishDiaryEntryResponse")
	proto.RegisterType((*UpdateDiaryEntryRequest)(nil), "familog.UpdateDiaryEntryRequest")
	proto.RegisterType((*UpdateDiaryEntryResponse)(nil), "familog.UpdateDiaryEntryResponse")
	proto.RegisterType((*ShowMyProfileRequest)(nil), "familog.ShowMyProfileRequest")
	proto.RegisterType((*ShowMyProfileResponse)(nil), "familog.ShowMyProfileResponse")
	proto.RegisterType((*FindProfileByUserIdRequest)(nil), "familog.FindProfileByUserIdRequest")
	proto.RegisterType((*FindProfileByUserIdResponse)(nil), "familog.FindProfileByUserIdResponse")
	proto.RegisterType((*UpdateProfileRequest)(nil), "familog.UpdateProfileRequest")
	proto.RegisterType((*UpdateProfileResponse)(nil), "familog.UpdateProfileResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Familog service

type FamilogClient interface {
	// 日記を始める
	StartDiary(ctx context.Context, in *StartDiaryRequest, opts ...grpc.CallOption) (*StartDiaryResponse, error)
	// 日記を購読する
	SubscribeDiary(ctx context.Context, in *SubscribeDiaryRequest, opts ...grpc.CallOption) (*SubscribeDiaryResponse, error)
	// 日記帳を検索する
	AllDiaries(ctx context.Context, in *AllDiariesRequest, opts ...grpc.CallOption) (*AllDiariesResponse, error)
	// 日記を検索する
	FindDiaryEntries(ctx context.Context, in *FindDiaryEntriesRequest, opts ...grpc.CallOption) (*FindDiaryEntriesResponse, error)
	// 日記をみる
	ReadDiaryEntry(ctx context.Context, in *ReadDiaryEntryRequest, opts ...grpc.CallOption) (*ReadDiaryEntryResponse, error)
	// 日記を投稿する
	PublishDiaryEntry(ctx context.Context, in *PublishDiaryEntryRequest, opts ...grpc.CallOption) (*PublishDiaryEntryResponse, error)
	// 日記を更新する
	UpdateDiaryEntry(ctx context.Context, in *UpdateDiaryEntryRequest, opts ...grpc.CallOption) (*UpdateDiaryEntryResponse, error)
	// 自分のプロフィールを表示する
	ShowMyProfile(ctx context.Context, in *ShowMyProfileRequest, opts ...grpc.CallOption) (*ShowMyProfileResponse, error)
	// 他のプロフィールを表示する
	FindProfileByUserId(ctx context.Context, in *FindProfileByUserIdRequest, opts ...grpc.CallOption) (*FindProfileByUserIdResponse, error)
	// プロフィールを更新する
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
}

type familogClient struct {
	cc *grpc.ClientConn
}

func NewFamilogClient(cc *grpc.ClientConn) FamilogClient {
	return &familogClient{cc}
}

func (c *familogClient) StartDiary(ctx context.Context, in *StartDiaryRequest, opts ...grpc.CallOption) (*StartDiaryResponse, error) {
	out := new(StartDiaryResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/StartDiary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) SubscribeDiary(ctx context.Context, in *SubscribeDiaryRequest, opts ...grpc.CallOption) (*SubscribeDiaryResponse, error) {
	out := new(SubscribeDiaryResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/SubscribeDiary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) AllDiaries(ctx context.Context, in *AllDiariesRequest, opts ...grpc.CallOption) (*AllDiariesResponse, error) {
	out := new(AllDiariesResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/AllDiaries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) FindDiaryEntries(ctx context.Context, in *FindDiaryEntriesRequest, opts ...grpc.CallOption) (*FindDiaryEntriesResponse, error) {
	out := new(FindDiaryEntriesResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/FindDiaryEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) ReadDiaryEntry(ctx context.Context, in *ReadDiaryEntryRequest, opts ...grpc.CallOption) (*ReadDiaryEntryResponse, error) {
	out := new(ReadDiaryEntryResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/ReadDiaryEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) PublishDiaryEntry(ctx context.Context, in *PublishDiaryEntryRequest, opts ...grpc.CallOption) (*PublishDiaryEntryResponse, error) {
	out := new(PublishDiaryEntryResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/PublishDiaryEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) UpdateDiaryEntry(ctx context.Context, in *UpdateDiaryEntryRequest, opts ...grpc.CallOption) (*UpdateDiaryEntryResponse, error) {
	out := new(UpdateDiaryEntryResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/UpdateDiaryEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) ShowMyProfile(ctx context.Context, in *ShowMyProfileRequest, opts ...grpc.CallOption) (*ShowMyProfileResponse, error) {
	out := new(ShowMyProfileResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/ShowMyProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) FindProfileByUserId(ctx context.Context, in *FindProfileByUserIdRequest, opts ...grpc.CallOption) (*FindProfileByUserIdResponse, error) {
	out := new(FindProfileByUserIdResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/FindProfileByUserId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familogClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	out := new(UpdateProfileResponse)
	err := grpc.Invoke(ctx, "/familog.Familog/UpdateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Familog service

type FamilogServer interface {
	// 日記を始める
	StartDiary(context.Context, *StartDiaryRequest) (*StartDiaryResponse, error)
	// 日記を購読する
	SubscribeDiary(context.Context, *SubscribeDiaryRequest) (*SubscribeDiaryResponse, error)
	// 日記帳を検索する
	AllDiaries(context.Context, *AllDiariesRequest) (*AllDiariesResponse, error)
	// 日記を検索する
	FindDiaryEntries(context.Context, *FindDiaryEntriesRequest) (*FindDiaryEntriesResponse, error)
	// 日記をみる
	ReadDiaryEntry(context.Context, *ReadDiaryEntryRequest) (*ReadDiaryEntryResponse, error)
	// 日記を投稿する
	PublishDiaryEntry(context.Context, *PublishDiaryEntryRequest) (*PublishDiaryEntryResponse, error)
	// 日記を更新する
	UpdateDiaryEntry(context.Context, *UpdateDiaryEntryRequest) (*UpdateDiaryEntryResponse, error)
	// 自分のプロフィールを表示する
	ShowMyProfile(context.Context, *ShowMyProfileRequest) (*ShowMyProfileResponse, error)
	// 他のプロフィールを表示する
	FindProfileByUserId(context.Context, *FindProfileByUserIdRequest) (*FindProfileByUserIdResponse, error)
	// プロフィールを更新する
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
}

func RegisterFamilogServer(s *grpc.Server, srv FamilogServer) {
	s.RegisterService(&_Familog_serviceDesc, srv)
}

func _Familog_StartDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDiaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).StartDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/StartDiary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).StartDiary(ctx, req.(*StartDiaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_SubscribeDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeDiaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).SubscribeDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/SubscribeDiary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).SubscribeDiary(ctx, req.(*SubscribeDiaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_AllDiaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllDiariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).AllDiaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/AllDiaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).AllDiaries(ctx, req.(*AllDiariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_FindDiaryEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDiaryEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).FindDiaryEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/FindDiaryEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).FindDiaryEntries(ctx, req.(*FindDiaryEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_ReadDiaryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDiaryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).ReadDiaryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/ReadDiaryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).ReadDiaryEntry(ctx, req.(*ReadDiaryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_PublishDiaryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishDiaryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).PublishDiaryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/PublishDiaryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).PublishDiaryEntry(ctx, req.(*PublishDiaryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_UpdateDiaryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiaryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).UpdateDiaryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/UpdateDiaryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).UpdateDiaryEntry(ctx, req.(*UpdateDiaryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_ShowMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowMyProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).ShowMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/ShowMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).ShowMyProfile(ctx, req.(*ShowMyProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_FindProfileByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProfileByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).FindProfileByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/FindProfileByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).FindProfileByUserId(ctx, req.(*FindProfileByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Familog_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilogServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/familog.Familog/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilogServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Familog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "familog.Familog",
	HandlerType: (*FamilogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartDiary",
			Handler:    _Familog_StartDiary_Handler,
		},
		{
			MethodName: "SubscribeDiary",
			Handler:    _Familog_SubscribeDiary_Handler,
		},
		{
			MethodName: "AllDiaries",
			Handler:    _Familog_AllDiaries_Handler,
		},
		{
			MethodName: "FindDiaryEntries",
			Handler:    _Familog_FindDiaryEntries_Handler,
		},
		{
			MethodName: "ReadDiaryEntry",
			Handler:    _Familog_ReadDiaryEntry_Handler,
		},
		{
			MethodName: "PublishDiaryEntry",
			Handler:    _Familog_PublishDiaryEntry_Handler,
		},
		{
			MethodName: "UpdateDiaryEntry",
			Handler:    _Familog_UpdateDiaryEntry_Handler,
		},
		{
			MethodName: "ShowMyProfile",
			Handler:    _Familog_ShowMyProfile_Handler,
		},
		{
			MethodName: "FindProfileByUserId",
			Handler:    _Familog_FindProfileByUserId_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Familog_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "familog.proto",
}

func init() { proto.RegisterFile("familog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0xd3, 0xa4, 0xa1, 0x27, 0x9b, 0xd0, 0x4e, 0xff, 0xbc, 0x5e, 0xb6, 0xcd, 0x0e, 0x15,
	0x0a, 0x48, 0x24, 0x52, 0x59, 0xa4, 0xd5, 0x4a, 0x8b, 0x68, 0x81, 0xae, 0x2a, 0x84, 0x54, 0xd2,
	0xf6, 0x02, 0x04, 0x8a, 0xec, 0x78, 0x9a, 0x1a, 0xd9, 0x99, 0x30, 0x1e, 0x2f, 0xca, 0x2d, 0x12,
	0x0f, 0xca, 0x13, 0xf0, 0x0a, 0x68, 0x7e, 0x6c, 0x8f, 0x93, 0x49, 0xb5, 0xd0, 0xbb, 0xcc, 0xf9,
	0x9b, 0xef, 0x3b, 0x3e, 0xe7, 0x9b, 0x40, 0xe7, 0xce, 0x4f, 0xa2, 0x98, 0x4e, 0x07, 0x73, 0x46,
	0x39, 0x45, 0x2d, 0x7d, 0xf4, 0x3e, 0xe3, 0xf7, 0x11, 0x0b, 0xe7, 0x3e, 0xe3, 0x8b, 0xe1, 0x94,
	0xd2, 0x69, 0x4c, 0x86, 0x32, 0x22, 0xc8, 0xee, 0x86, 0x3c, 0x4a, 0x48, 0xca, 0xfd, 0x64, 0xae,
	0x92, 0xf0, 0x25, 0x34, 0x6e, 0x53, 0xc2, 0x50, 0x17, 0xea, 0x51, 0xe8, 0x3a, 0x3d, 0xa7, 0xdf,
	0x1c, 0xd5, 0xa3, 0x10, 0x21, 0x68, 0xcc, 0xfc, 0x84, 0xb8, 0xf5, 0x9e, 0xd3, 0xdf, 0x1a, 0xc9,
	0xdf, 0xe8, 0x39, 0x80, 0xff, 0xce, 0xe7, 0x3e, 0x1b, 0x67, 0x2c, 0x76, 0x37, 0xa4, 0x67, 0x4b,
	0x59, 0x6e, 0x59, 0x8c, 0x3f, 0x87, 0xe6, 0xb7, 0x91, 0xcf, 0x16, 0x2b, 0xb5, 0xf6, 0xa0, 0x79,
	0x13, 0xf1, 0x38, 0x2f, 0xa6, 0x0e, 0xf8, 0x1f, 0x07, 0x40, 0xc6, 0x7f, 0x37, 0xe3, 0x96, 0x24,
	0x04, 0x8d, 0x80, 0x86, 0x8b, 0x1c, 0x80, 0xf8, 0x2d, 0x0a, 0x91, 0x84, 0xfe, 0x16, 0xe9, 0xbb,
	0xd5, 0x01, 0xbd, 0x80, 0x46, 0x96, 0x12, 0xe6, 0x36, 0x7a, 0x4e, 0xbf, 0x7d, 0xda, 0x19, 0xe4,
	0x5d, 0x11, 0xbc, 0x46, 0xd2, 0x85, 0x3e, 0x81, 0x0f, 0x27, 0xf4, 0x1d, 0x61, 0xe3, 0x28, 0xf1,
	0xa7, 0x44, 0xc2, 0x6f, 0xca, 0x12, 0x1d, 0x69, 0xbe, 0x14, 0xd6, 0x5b, 0x16, 0x0b, 0x86, 0x45,
	0x44, 0xea, 0x6e, 0xf6, 0x36, 0x04, 0xc3, 0x48, 0x7b, 0x53, 0xf4, 0x06, 0x9e, 0xcc, 0xb3, 0x20,
	0x8e, 0xd2, 0x7b, 0x12, 0x8e, 0x7d, 0xee, 0xb6, 0xe4, 0x8d, 0xde, 0x40, 0x35, 0x79, 0x90, 0x37,
	0x79, 0x70, 0x93, 0x37, 0x79, 0xd4, 0x2e, 0xe2, 0xcf, 0x38, 0xfe, 0x14, 0x76, 0xae, 0xb9, 0xcf,
	0xb8, 0x64, 0x3d, 0x22, 0xbf, 0x67, 0x24, 0xe5, 0x82, 0x13, 0x97, 0xcd, 0x71, 0x14, 0x27, 0x79,
	0xc0, 0xaf, 0x01, 0x99, 0xa1, 0xe9, 0x9c, 0xce, 0x52, 0x82, 0x4e, 0xa0, 0x19, 0x0a, 0x83, 0x8c,
	0x6d, 0x9f, 0x76, 0x0b, 0xaa, 0x2a, 0x4c, 0x39, 0xf1, 0x2b, 0xd8, 0xbf, 0xce, 0x82, 0x74, 0xc2,
	0xa2, 0x80, 0x54, 0xae, 0x3a, 0x86, 0x76, 0x4a, 0x26, 0x8c, 0xf0, 0xf1, 0x84, 0x86, 0xf9, 0x85,
	0xa0, 0x4c, 0xdf, 0xd0, 0x90, 0x60, 0x17, 0x0e, 0x96, 0x33, 0xd5, 0xcd, 0x78, 0x17, 0x76, 0xce,
	0xe2, 0x58, 0xd8, 0x22, 0x92, 0xea, 0x7a, 0xf8, 0x2b, 0x40, 0xa6, 0x51, 0x83, 0xec, 0x43, 0x2b,
	0x54, 0x26, 0xd7, 0xe9, 0x6d, 0x58, 0x60, 0xe6, 0x6e, 0xfc, 0x12, 0x0e, 0x2f, 0xa2, 0x59, 0x58,
	0x0c, 0x41, 0x59, 0x1a, 0x3d, 0x85, 0x0f, 0x24, 0x99, 0xb1, 0x9e, 0x89, 0x2d, 0x95, 0xb5, 0xb8,
	0x0c, 0xf1, 0x0d, 0xb8, 0xab, 0x59, 0xfa, 0xee, 0x57, 0xd0, 0x51, 0x69, 0x44, 0x39, 0x34, 0x82,
	0xdd, 0x2a, 0x02, 0x39, 0x70, 0xa3, 0x27, 0xa1, 0x51, 0x01, 0xbf, 0x81, 0xfd, 0x11, 0xf1, 0x43,
	0xc3, 0xaf, 0x91, 0x9c, 0x40, 0xb7, 0x2c, 0x69, 0xe0, 0x29, 0xd3, 0x05, 0xa8, 0x11, 0x1c, 0x2c,
	0xa7, 0x3f, 0x1a, 0xd2, 0x5f, 0x0e, 0x74, 0x4b, 0xe7, 0x05, 0x65, 0x49, 0xb1, 0x14, 0x8e, 0x6d,
	0x29, 0xea, 0xe6, 0x52, 0xf4, 0x61, 0xdb, 0x9c, 0xf8, 0x20, 0xa6, 0x81, 0xde, 0x9a, 0x6e, 0x39,
	0xf2, 0xe7, 0x31, 0x0d, 0xc4, 0x54, 0x94, 0x31, 0xa9, 0xdb, 0x90, 0x43, 0xaf, 0xd6, 0x40, 0xf8,
	0x53, 0xfc, 0x2b, 0xb8, 0x57, 0x6a, 0x8a, 0x57, 0xbb, 0x73, 0x06, 0xdb, 0x66, 0x77, 0xee, 0x28,
	0x4b, 0xf4, 0x70, 0x1e, 0x5a, 0x08, 0x0a, 0x0e, 0xa3, 0x6e, 0x58, 0x39, 0xe3, 0x1f, 0xe1, 0xa9,
	0xa5, 0xbc, 0xee, 0xde, 0x4b, 0x68, 0x1b, 0xf5, 0x75, 0x69, 0x6b, 0xef, 0xa0, 0x2c, 0x8b, 0xff,
	0x74, 0xe0, 0xf0, 0x76, 0x1e, 0xfa, 0x9c, 0xfc, 0xcf, 0xef, 0x69, 0xe5, 0x55, 0xff, 0x6f, 0xbc,
	0xae, 0xc0, 0x5d, 0xc5, 0xf0, 0x28, 0x5a, 0x07, 0xb0, 0x77, 0x7d, 0x4f, 0xff, 0xf8, 0x61, 0x71,
	0xc5, 0xe8, 0x5d, 0x14, 0x93, 0x7c, 0x0f, 0x5f, 0xc3, 0xfe, 0x92, 0x5d, 0x5f, 0x93, 0x2b, 0xa3,
	0xb3, 0x56, 0x19, 0xf1, 0x97, 0xe0, 0x89, 0x6d, 0xd2, 0x99, 0xe7, 0x0b, 0xe1, 0xba, 0x0c, 0xf3,
	0x66, 0x1d, 0x42, 0x4b, 0x44, 0x95, 0x5d, 0xda, 0xcc, 0xa4, 0x1f, 0x7f, 0x0d, 0xcf, 0xac, 0x69,
	0xef, 0x7f, 0xf1, 0xf7, 0xb0, 0xa7, 0xda, 0x53, 0x25, 0x53, 0x3c, 0x3c, 0x8e, 0xf1, 0xf0, 0x1c,
	0x43, 0x5b, 0x3f, 0x3c, 0x72, 0x8e, 0xd5, 0xa0, 0xeb, 0xb7, 0x48, 0xcc, 0xa8, 0xe8, 0xc0, 0x52,
	0xb1, 0xf7, 0x06, 0x72, 0xfa, 0xf7, 0x26, 0xb4, 0x2e, 0x94, 0x19, 0xbd, 0x05, 0x28, 0x65, 0x17,
	0x79, 0x45, 0xf8, 0x8a, 0x6c, 0x7b, 0xcf, 0xac, 0x3e, 0xad, 0x96, 0x35, 0x74, 0x0d, 0xdd, 0xaa,
	0x92, 0xa2, 0xa3, 0x32, 0xc1, 0x26, 0xce, 0xde, 0xf1, 0x5a, 0x7f, 0x51, 0xf4, 0x2d, 0x40, 0xa9,
	0xb7, 0x06, 0xba, 0x15, 0x65, 0x36, 0xd0, 0xad, 0x0a, 0x34, 0xae, 0xa1, 0x9f, 0x60, 0x7b, 0x59,
	0x42, 0x51, 0xaf, 0x48, 0x59, 0xa3, 0xc9, 0xde, 0x8b, 0x07, 0x22, 0x4c, 0xe2, 0x55, 0x21, 0x34,
	0x88, 0x5b, 0x05, 0xd6, 0x20, 0x6e, 0x57, 0x50, 0x5c, 0x43, 0xbf, 0xc0, 0xce, 0x8a, 0x44, 0xa0,
	0x12, 0xce, 0x3a, 0x75, 0xf2, 0xf0, 0x43, 0x21, 0x66, 0x37, 0x96, 0x17, 0xd5, 0xe8, 0xc6, 0x1a,
	0x1d, 0x31, 0xba, 0xb1, 0x6e, 0xcb, 0x71, 0x0d, 0x5d, 0x41, 0xa7, 0xb2, 0x99, 0xe8, 0x79, 0xf9,
	0x95, 0x2d, 0x9b, 0xec, 0x1d, 0xad, 0x73, 0x17, 0x15, 0x03, 0xd8, 0xb5, 0x2c, 0x1e, 0xfa, 0xb8,
	0xf2, 0x6d, 0xec, 0xdb, 0xec, 0x9d, 0x3c, 0x1c, 0x64, 0xa2, 0xae, 0x6c, 0x93, 0x81, 0xda, 0xb6,
	0xb2, 0x06, 0x6a, 0xeb, 0x12, 0xe2, 0xda, 0xf9, 0xd1, 0xcf, 0x1f, 0x4d, 0x23, 0x7e, 0x9f, 0x05,
	0x83, 0x09, 0x4d, 0x86, 0x69, 0x96, 0xa5, 0xfe, 0xec, 0x74, 0x4a, 0x87, 0x3a, 0x2f, 0xd8, 0x94,
	0x7f, 0x9d, 0xbe, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xe8, 0xa9, 0xc3, 0xd2, 0x0a, 0x00,
	0x00,
}
