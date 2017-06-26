// Code generated by protoc-gen-go. DO NOT EDIT.
// source: database.proto

/*
Package database is a generated protocol buffer package.

It is generated from these files:
	database.proto

It has these top-level messages:
	AddGameRequest
	AddGameResponse
*/
package database

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type ErrorCode int32

const (
	ErrorCode_SUCCESS   ErrorCode = 0
	ErrorCode_NOT_FOUND ErrorCode = 1
)

var ErrorCode_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var ErrorCode_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Beware! Changing these will require a re-install of the system.
type TeamCode int32

const (
	TeamCode_Buffalo      TeamCode = 0
	TeamCode_Cleveland    TeamCode = 1
	TeamCode_NewOrleans   TeamCode = 2
	TeamCode_NewEngland   TeamCode = 3
	TeamCode_Detroit      TeamCode = 4
	TeamCode_GreenBay     TeamCode = 5
	TeamCode_Seattle      TeamCode = 6
	TeamCode_Baltimore    TeamCode = 7
	TeamCode_Miami        TeamCode = 8
	TeamCode_Minnesota    TeamCode = 9
	TeamCode_Cincinnati   TeamCode = 10
	TeamCode_Philadelphia TeamCode = 11
	TeamCode_Pittsburgh   TeamCode = 12
	TeamCode_Chicago      TeamCode = 13
	TeamCode_Indianapolis TeamCode = 14
	TeamCode_NYGiants     TeamCode = 15
	TeamCode_Jacksonville TeamCode = 16
	TeamCode_StLouis      TeamCode = 17
	TeamCode_KansasCity   TeamCode = 18
	TeamCode_Tennessee    TeamCode = 19
	TeamCode_Carolina     TeamCode = 20
	TeamCode_Arizona      TeamCode = 21
	TeamCode_Denver       TeamCode = 22
	TeamCode_Dallas       TeamCode = 23
	TeamCode_Houston      TeamCode = 24
	TeamCode_SanFrancisco TeamCode = 25
	TeamCode_SanDiego     TeamCode = 26
	TeamCode_Oakland      TeamCode = 27
	TeamCode_NYJets       TeamCode = 28
	TeamCode_Washington   TeamCode = 29
	TeamCode_TampaBay     TeamCode = 30
	TeamCode_Atlanta      TeamCode = 31
)

var TeamCode_name = map[int32]string{
	0:  "Buffalo",
	1:  "Cleveland",
	2:  "NewOrleans",
	3:  "NewEngland",
	4:  "Detroit",
	5:  "GreenBay",
	6:  "Seattle",
	7:  "Baltimore",
	8:  "Miami",
	9:  "Minnesota",
	10: "Cincinnati",
	11: "Philadelphia",
	12: "Pittsburgh",
	13: "Chicago",
	14: "Indianapolis",
	15: "NYGiants",
	16: "Jacksonville",
	17: "StLouis",
	18: "KansasCity",
	19: "Tennessee",
	20: "Carolina",
	21: "Arizona",
	22: "Denver",
	23: "Dallas",
	24: "Houston",
	25: "SanFrancisco",
	26: "SanDiego",
	27: "Oakland",
	28: "NYJets",
	29: "Washington",
	30: "TampaBay",
	31: "Atlanta",
}
var TeamCode_value = map[string]int32{
	"Buffalo":      0,
	"Cleveland":    1,
	"NewOrleans":   2,
	"NewEngland":   3,
	"Detroit":      4,
	"GreenBay":     5,
	"Seattle":      6,
	"Baltimore":    7,
	"Miami":        8,
	"Minnesota":    9,
	"Cincinnati":   10,
	"Philadelphia": 11,
	"Pittsburgh":   12,
	"Chicago":      13,
	"Indianapolis": 14,
	"NYGiants":     15,
	"Jacksonville": 16,
	"StLouis":      17,
	"KansasCity":   18,
	"Tennessee":    19,
	"Carolina":     20,
	"Arizona":      21,
	"Denver":       22,
	"Dallas":       23,
	"Houston":      24,
	"SanFrancisco": 25,
	"SanDiego":     26,
	"Oakland":      27,
	"NYJets":       28,
	"Washington":   29,
	"TampaBay":     30,
	"Atlanta":      31,
}

func (x TeamCode) String() string {
	return proto.EnumName(TeamCode_name, int32(x))
}
func (TeamCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AddGameRequest struct {
	Week     int32    `protobuf:"varint,1,opt,name=week" json:"week,omitempty"`
	HomeTeam TeamCode `protobuf:"varint,2,opt,name=home_team,json=homeTeam,enum=TeamCode" json:"home_team,omitempty"`
	AwayTeam TeamCode `protobuf:"varint,3,opt,name=away_team,json=awayTeam,enum=TeamCode" json:"away_team,omitempty"`
}

func (m *AddGameRequest) Reset()                    { *m = AddGameRequest{} }
func (m *AddGameRequest) String() string            { return proto.CompactTextString(m) }
func (*AddGameRequest) ProtoMessage()               {}
func (*AddGameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddGameRequest) GetWeek() int32 {
	if m != nil {
		return m.Week
	}
	return 0
}

func (m *AddGameRequest) GetHomeTeam() TeamCode {
	if m != nil {
		return m.HomeTeam
	}
	return TeamCode_Buffalo
}

func (m *AddGameRequest) GetAwayTeam() TeamCode {
	if m != nil {
		return m.AwayTeam
	}
	return TeamCode_Buffalo
}

type AddGameResponse struct {
	Error ErrorCode `protobuf:"varint,1,opt,name=error,enum=ErrorCode" json:"error,omitempty"`
	Uuid  string    `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *AddGameResponse) Reset()                    { *m = AddGameResponse{} }
func (m *AddGameResponse) String() string            { return proto.CompactTextString(m) }
func (*AddGameResponse) ProtoMessage()               {}
func (*AddGameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddGameResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_SUCCESS
}

func (m *AddGameResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*AddGameRequest)(nil), "AddGameRequest")
	proto.RegisterType((*AddGameResponse)(nil), "AddGameResponse")
	proto.RegisterEnum("ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("TeamCode", TeamCode_name, TeamCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DatabaseService service

type DatabaseServiceClient interface {
	AddGame(ctx context.Context, in *AddGameRequest, opts ...client.CallOption) (*AddGameResponse, error)
}

type databaseServiceClient struct {
	c           client.Client
	serviceName string
}

func NewDatabaseServiceClient(serviceName string, c client.Client) DatabaseServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "databaseservice"
	}
	return &databaseServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *databaseServiceClient) AddGame(ctx context.Context, in *AddGameRequest, opts ...client.CallOption) (*AddGameResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DatabaseService.AddGame", in)
	out := new(AddGameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DatabaseService service

type DatabaseServiceHandler interface {
	AddGame(context.Context, *AddGameRequest, *AddGameResponse) error
}

func RegisterDatabaseServiceHandler(s server.Server, hdlr DatabaseServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DatabaseService{hdlr}, opts...))
}

type DatabaseService struct {
	DatabaseServiceHandler
}

func (h *DatabaseService) AddGame(ctx context.Context, in *AddGameRequest, out *AddGameResponse) error {
	return h.DatabaseServiceHandler.AddGame(ctx, in, out)
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xdd, 0x4e, 0x1b, 0x3f,
	0x10, 0xc5, 0xff, 0x01, 0x02, 0x59, 0x03, 0xc9, 0xfc, 0xdd, 0x2f, 0x4a, 0xbf, 0x10, 0x17, 0x2d,
	0x42, 0x55, 0x2e, 0xe8, 0x03, 0x54, 0xb0, 0x81, 0xb4, 0xb4, 0x24, 0x28, 0x09, 0xaa, 0xb8, 0x42,
	0x43, 0x76, 0x48, 0x46, 0x78, 0xed, 0xd4, 0xf6, 0x06, 0xd1, 0x27, 0xed, 0xe3, 0x54, 0x63, 0x28,
	0x52, 0xdb, 0x3b, 0x8f, 0xce, 0x6f, 0xce, 0x39, 0x5a, 0xcd, 0xaa, 0x66, 0x81, 0x11, 0x2f, 0x31,
	0x50, 0x7b, 0xe6, 0x5d, 0x74, 0xdb, 0x51, 0x35, 0xf7, 0x8b, 0xa2, 0x8b, 0x25, 0x0d, 0xe8, 0x7b,
	0x45, 0x21, 0x6a, 0xad, 0x96, 0x6e, 0x88, 0xae, 0x37, 0x6a, 0x5b, 0xb5, 0x9d, 0xfa, 0x20, 0xbd,
	0xf5, 0x5b, 0x95, 0x4d, 0x5d, 0x49, 0x17, 0x91, 0xb0, 0xdc, 0x58, 0xd8, 0xaa, 0xed, 0x34, 0xf7,
	0xb2, 0xf6, 0x88, 0xb0, 0xcc, 0x5d, 0x41, 0x83, 0x86, 0x68, 0x32, 0x09, 0x87, 0x37, 0x78, 0x7b,
	0xc7, 0x2d, 0xfe, 0xc3, 0x89, 0x26, 0xd3, 0x76, 0x57, 0xb5, 0x1e, 0x52, 0xc3, 0xcc, 0xd9, 0x40,
	0x7a, 0x4b, 0xd5, 0xc9, 0x7b, 0xe7, 0x53, 0x6e, 0x73, 0x4f, 0xb5, 0x0f, 0x65, 0x4a, 0x7b, 0x77,
	0x82, 0x14, 0xab, 0x2a, 0x2e, 0x52, 0x7e, 0x36, 0x48, 0xef, 0xdd, 0x77, 0x2a, 0x7b, 0xe0, 0xf4,
	0xaa, 0x5a, 0x19, 0x9e, 0xe5, 0xf9, 0xe1, 0x70, 0x08, 0xff, 0xe9, 0x75, 0x95, 0xf5, 0xfa, 0xa3,
	0x8b, 0xa3, 0xfe, 0x59, 0xaf, 0x03, 0xb5, 0xdd, 0x9f, 0x8b, 0xaa, 0xf1, 0xbb, 0x88, 0x80, 0x07,
	0xd5, 0xd5, 0x15, 0x1a, 0x77, 0x07, 0xe6, 0x86, 0xe6, 0x64, 0xd0, 0x16, 0x50, 0xd3, 0x4d, 0xa5,
	0x7a, 0x74, 0xd3, 0xf7, 0x86, 0xd0, 0x06, 0x58, 0xb8, 0x9f, 0x0f, 0xed, 0x24, 0xe9, 0x8b, 0xb2,
	0xdb, 0xa1, 0xe8, 0x1d, 0x47, 0x58, 0xd2, 0x6b, 0xaa, 0xd1, 0xf5, 0x44, 0xf6, 0x00, 0x6f, 0xa1,
	0x9e, 0xf2, 0x09, 0x63, 0x34, 0x04, 0xcb, 0x62, 0x7b, 0x80, 0x26, 0x72, 0xe9, 0x3c, 0xc1, 0x8a,
	0xce, 0x54, 0xfd, 0x84, 0xb1, 0x64, 0x68, 0x88, 0x72, 0xc2, 0xd6, 0x52, 0x70, 0x11, 0x21, 0x93,
	0x80, 0x9c, 0xed, 0x98, 0xad, 0xc5, 0xc8, 0xa0, 0x34, 0xa8, 0xb5, 0xd3, 0x29, 0x1b, 0x2c, 0xc8,
	0xcc, 0xa6, 0x8c, 0xb0, 0x2a, 0xc4, 0x29, 0xc7, 0x18, 0x2e, 0x2b, 0x3f, 0x99, 0xc2, 0x9a, 0xe4,
	0xe4, 0x53, 0x1e, 0xe3, 0xc4, 0xc1, 0xba, 0xe0, 0x9f, 0x6d, 0xc1, 0x68, 0x71, 0xe6, 0x0c, 0x07,
	0x68, 0x4a, 0xa9, 0xde, 0x79, 0x97, 0xd1, 0xc6, 0x00, 0x2d, 0xd1, 0x8f, 0x71, 0x7c, 0x1d, 0x9c,
	0x9d, 0xb3, 0x31, 0x04, 0x90, 0x6a, 0xc6, 0xaf, 0xae, 0xe2, 0x00, 0xff, 0x8b, 0xf7, 0x17, 0xb4,
	0x01, 0x43, 0xce, 0xf1, 0x16, 0xb4, 0x94, 0x1b, 0x91, 0x94, 0x0b, 0x44, 0xf0, 0x48, 0xbc, 0x72,
	0xf4, 0xce, 0xb0, 0x45, 0x78, 0x2c, 0x9b, 0xfb, 0x9e, 0x7f, 0x38, 0x8b, 0xf0, 0x44, 0x2b, 0xb5,
	0xdc, 0x21, 0x3b, 0x27, 0x0f, 0x4f, 0xd3, 0x1b, 0x8d, 0xc1, 0x00, 0xcf, 0x04, 0xfa, 0xe4, 0xaa,
	0x10, 0x9d, 0x85, 0x0d, 0x49, 0x1f, 0xa2, 0x3d, 0xf2, 0x68, 0xc7, 0x1c, 0xc6, 0x0e, 0x9e, 0x8b,
	0xe3, 0x10, 0x6d, 0x87, 0x69, 0xe2, 0x60, 0x53, 0xe0, 0x3e, 0x5e, 0xa7, 0x4f, 0xfb, 0x42, 0x5c,
	0x7a, 0xe7, 0xc7, 0x14, 0x03, 0xbc, 0x94, 0x5e, 0xdf, 0x30, 0x4c, 0xd9, 0x4e, 0xc4, 0xe8, 0x95,
	0xac, 0x8d, 0xb0, 0x9c, 0xa1, 0x7c, 0xe9, 0xd7, 0xa9, 0x48, 0x34, 0x68, 0x23, 0xc2, 0x9b, 0xbd,
	0x8f, 0xaa, 0xd5, 0xb9, 0x3f, 0xea, 0x21, 0xf9, 0x39, 0x8f, 0x49, 0xbf, 0x57, 0x2b, 0xf7, 0xf7,
	0xa5, 0x5b, 0xed, 0x3f, 0xef, 0x7b, 0x13, 0xda, 0x7f, 0x9d, 0xde, 0xe5, 0x72, 0xfa, 0x15, 0x3e,
	0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x19, 0x9c, 0x95, 0x61, 0x1c, 0x03, 0x00, 0x00,
}
