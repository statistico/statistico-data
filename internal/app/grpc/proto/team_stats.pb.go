// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/grpc/proto/team_stats.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type TeamStatsResponse struct {
	HomeTeam             *TeamStats `protobuf:"bytes,1,opt,name=home_team,json=homeTeam,proto3" json:"home_team,omitempty"`
	AwayTeam             *TeamStats `protobuf:"bytes,2,opt,name=away_team,json=awayTeam,proto3" json:"away_team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TeamStatsResponse) Reset()         { *m = TeamStatsResponse{} }
func (m *TeamStatsResponse) String() string { return proto.CompactTextString(m) }
func (*TeamStatsResponse) ProtoMessage()    {}
func (*TeamStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5fe853e65352bc, []int{0}
}

func (m *TeamStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamStatsResponse.Unmarshal(m, b)
}
func (m *TeamStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamStatsResponse.Marshal(b, m, deterministic)
}
func (m *TeamStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamStatsResponse.Merge(m, src)
}
func (m *TeamStatsResponse) XXX_Size() int {
	return xxx_messageInfo_TeamStatsResponse.Size(m)
}
func (m *TeamStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TeamStatsResponse proto.InternalMessageInfo

func (m *TeamStatsResponse) GetHomeTeam() *TeamStats {
	if m != nil {
		return m.HomeTeam
	}
	return nil
}

func (m *TeamStatsResponse) GetAwayTeam() *TeamStats {
	if m != nil {
		return m.AwayTeam
	}
	return nil
}

type TeamStats struct {
	TeamId               uint64                `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ShotsTotal           *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=shots_total,json=shotsTotal,proto3" json:"shots_total,omitempty"`
	ShotsOnGoal          *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=shots_on_goal,json=shotsOnGoal,proto3" json:"shots_on_goal,omitempty"`
	ShotsOffGoal         *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=shots_off_goal,json=shotsOffGoal,proto3" json:"shots_off_goal,omitempty"`
	ShotsBlocked         *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=shots_blocked,json=shotsBlocked,proto3" json:"shots_blocked,omitempty"`
	ShotsInsideBox       *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=shots_inside_box,json=shotsInsideBox,proto3" json:"shots_inside_box,omitempty"`
	ShotsOutsideBox      *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=shots_outside_box,json=shotsOutsideBox,proto3" json:"shots_outside_box,omitempty"`
	PassesTotal          *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=passes_total,json=passesTotal,proto3" json:"passes_total,omitempty"`
	PassesAccuracy       *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=passes_accuracy,json=passesAccuracy,proto3" json:"passes_accuracy,omitempty"`
	PassesPercentage     *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=passes_percentage,json=passesPercentage,proto3" json:"passes_percentage,omitempty"`
	AttacksTotal         *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=attacks_total,json=attacksTotal,proto3" json:"attacks_total,omitempty"`
	AttacksDangerous     *wrappers.UInt32Value `protobuf:"bytes,12,opt,name=attacks_dangerous,json=attacksDangerous,proto3" json:"attacks_dangerous,omitempty"`
	Fouls                *wrappers.UInt32Value `protobuf:"bytes,13,opt,name=fouls,proto3" json:"fouls,omitempty"`
	Corners              *wrappers.UInt32Value `protobuf:"bytes,14,opt,name=corners,proto3" json:"corners,omitempty"`
	Offsides             *wrappers.UInt32Value `protobuf:"bytes,15,opt,name=offsides,proto3" json:"offsides,omitempty"`
	Possession           *wrappers.UInt32Value `protobuf:"bytes,16,opt,name=possession,proto3" json:"possession,omitempty"`
	YellowCards          *wrappers.UInt32Value `protobuf:"bytes,17,opt,name=yellow_cards,json=yellowCards,proto3" json:"yellow_cards,omitempty"`
	RedCards             *wrappers.UInt32Value `protobuf:"bytes,18,opt,name=red_cards,json=redCards,proto3" json:"red_cards,omitempty"`
	Saves                *wrappers.UInt32Value `protobuf:"bytes,19,opt,name=saves,proto3" json:"saves,omitempty"`
	Substitutions        *wrappers.UInt32Value `protobuf:"bytes,20,opt,name=substitutions,proto3" json:"substitutions,omitempty"`
	GoalKicks            *wrappers.UInt32Value `protobuf:"bytes,21,opt,name=goal_kicks,json=goalKicks,proto3" json:"goal_kicks,omitempty"`
	GoalAttempts         *wrappers.UInt32Value `protobuf:"bytes,22,opt,name=goal_attempts,json=goalAttempts,proto3" json:"goal_attempts,omitempty"`
	FreeKicks            *wrappers.UInt32Value `protobuf:"bytes,23,opt,name=free_kicks,json=freeKicks,proto3" json:"free_kicks,omitempty"`
	ThrowIns             *wrappers.UInt32Value `protobuf:"bytes,24,opt,name=throw_ins,json=throwIns,proto3" json:"throw_ins,omitempty"`
	HomeXg               *wrappers.UInt32Value `protobuf:"bytes,25,opt,name=home_xg,json=homeXg,proto3" json:"home_xg,omitempty"`
	AwayXg               *wrappers.UInt32Value `protobuf:"bytes,26,opt,name=away_xg,json=awayXg,proto3" json:"away_xg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TeamStats) Reset()         { *m = TeamStats{} }
func (m *TeamStats) String() string { return proto.CompactTextString(m) }
func (*TeamStats) ProtoMessage()    {}
func (*TeamStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5fe853e65352bc, []int{1}
}

func (m *TeamStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamStats.Unmarshal(m, b)
}
func (m *TeamStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamStats.Marshal(b, m, deterministic)
}
func (m *TeamStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamStats.Merge(m, src)
}
func (m *TeamStats) XXX_Size() int {
	return xxx_messageInfo_TeamStats.Size(m)
}
func (m *TeamStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamStats.DiscardUnknown(m)
}

var xxx_messageInfo_TeamStats proto.InternalMessageInfo

func (m *TeamStats) GetTeamId() uint64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *TeamStats) GetShotsTotal() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsTotal
	}
	return nil
}

func (m *TeamStats) GetShotsOnGoal() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsOnGoal
	}
	return nil
}

func (m *TeamStats) GetShotsOffGoal() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsOffGoal
	}
	return nil
}

func (m *TeamStats) GetShotsBlocked() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsBlocked
	}
	return nil
}

func (m *TeamStats) GetShotsInsideBox() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsInsideBox
	}
	return nil
}

func (m *TeamStats) GetShotsOutsideBox() *wrappers.UInt32Value {
	if m != nil {
		return m.ShotsOutsideBox
	}
	return nil
}

func (m *TeamStats) GetPassesTotal() *wrappers.UInt32Value {
	if m != nil {
		return m.PassesTotal
	}
	return nil
}

func (m *TeamStats) GetPassesAccuracy() *wrappers.UInt32Value {
	if m != nil {
		return m.PassesAccuracy
	}
	return nil
}

func (m *TeamStats) GetPassesPercentage() *wrappers.UInt32Value {
	if m != nil {
		return m.PassesPercentage
	}
	return nil
}

func (m *TeamStats) GetAttacksTotal() *wrappers.UInt32Value {
	if m != nil {
		return m.AttacksTotal
	}
	return nil
}

func (m *TeamStats) GetAttacksDangerous() *wrappers.UInt32Value {
	if m != nil {
		return m.AttacksDangerous
	}
	return nil
}

func (m *TeamStats) GetFouls() *wrappers.UInt32Value {
	if m != nil {
		return m.Fouls
	}
	return nil
}

func (m *TeamStats) GetCorners() *wrappers.UInt32Value {
	if m != nil {
		return m.Corners
	}
	return nil
}

func (m *TeamStats) GetOffsides() *wrappers.UInt32Value {
	if m != nil {
		return m.Offsides
	}
	return nil
}

func (m *TeamStats) GetPossession() *wrappers.UInt32Value {
	if m != nil {
		return m.Possession
	}
	return nil
}

func (m *TeamStats) GetYellowCards() *wrappers.UInt32Value {
	if m != nil {
		return m.YellowCards
	}
	return nil
}

func (m *TeamStats) GetRedCards() *wrappers.UInt32Value {
	if m != nil {
		return m.RedCards
	}
	return nil
}

func (m *TeamStats) GetSaves() *wrappers.UInt32Value {
	if m != nil {
		return m.Saves
	}
	return nil
}

func (m *TeamStats) GetSubstitutions() *wrappers.UInt32Value {
	if m != nil {
		return m.Substitutions
	}
	return nil
}

func (m *TeamStats) GetGoalKicks() *wrappers.UInt32Value {
	if m != nil {
		return m.GoalKicks
	}
	return nil
}

func (m *TeamStats) GetGoalAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.GoalAttempts
	}
	return nil
}

func (m *TeamStats) GetFreeKicks() *wrappers.UInt32Value {
	if m != nil {
		return m.FreeKicks
	}
	return nil
}

func (m *TeamStats) GetThrowIns() *wrappers.UInt32Value {
	if m != nil {
		return m.ThrowIns
	}
	return nil
}

func (m *TeamStats) GetHomeXg() *wrappers.UInt32Value {
	if m != nil {
		return m.HomeXg
	}
	return nil
}

func (m *TeamStats) GetAwayXg() *wrappers.UInt32Value {
	if m != nil {
		return m.AwayXg
	}
	return nil
}

func init() {
	proto.RegisterType((*TeamStatsResponse)(nil), "proto.TeamStatsResponse")
	proto.RegisterType((*TeamStats)(nil), "proto.TeamStats")
}

func init() {
	proto.RegisterFile("internal/app/grpc/proto/team_stats.proto", fileDescriptor_ff5fe853e65352bc)
}

var fileDescriptor_ff5fe853e65352bc = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0x5f, 0x4f, 0x13, 0x4d,
	0x14, 0x87, 0x5f, 0x5e, 0x69, 0x4b, 0x0f, 0x05, 0xda, 0x51, 0x60, 0x6c, 0x8c, 0x31, 0x5c, 0x18,
	0x6e, 0x6c, 0x93, 0x12, 0x8d, 0x44, 0x8d, 0x52, 0x15, 0x6c, 0xb8, 0xd0, 0x14, 0x34, 0xdc, 0x6d,
	0xa6, 0xbb, 0x67, 0x97, 0x4d, 0xb7, 0x3b, 0xcb, 0xcc, 0x2c, 0x2d, 0x1f, 0xd6, 0xef, 0x62, 0xe6,
	0xcf, 0x2e, 0x31, 0x86, 0x64, 0xae, 0xc8, 0x2e, 0xcf, 0xf3, 0x3b, 0x73, 0x66, 0xa6, 0x7b, 0xe0,
	0x30, 0xcd, 0x15, 0x8a, 0x9c, 0x65, 0x43, 0x56, 0x14, 0xc3, 0x44, 0x14, 0xe1, 0xb0, 0x10, 0x5c,
	0xf1, 0xa1, 0x42, 0xb6, 0x08, 0xa4, 0x62, 0x4a, 0x0e, 0xcc, 0x0b, 0xd2, 0x30, 0x7f, 0xfa, 0xcf,
	0x13, 0xce, 0x93, 0x0c, 0x2d, 0x35, 0x2b, 0xe3, 0xe1, 0x52, 0xb0, 0xa2, 0x40, 0xe1, 0xb0, 0xfe,
	0xcb, 0x87, 0x02, 0x05, 0xde, 0x94, 0x28, 0xab, 0xb8, 0x83, 0x1b, 0xe8, 0x5d, 0x22, 0x5b, 0x5c,
	0xe8, 0x0a, 0x53, 0x94, 0x05, 0xcf, 0x25, 0x92, 0x57, 0xd0, 0xbe, 0xe6, 0x0b, 0x0c, 0x74, 0x71,
	0xba, 0xf6, 0x62, 0xed, 0x70, 0x73, 0xd4, 0xb5, 0xfc, 0xe0, 0x1e, 0xde, 0xd0, 0x88, 0x7e, 0xd4,
	0x38, 0x5b, 0xb2, 0x3b, 0x8b, 0xff, 0xff, 0x10, 0xae, 0x11, 0xfd, 0x78, 0xf0, 0xbb, 0x03, 0xed,
	0xfa, 0x3d, 0xd9, 0x87, 0x96, 0xe9, 0x31, 0x8d, 0x4c, 0xa5, 0xf5, 0x69, 0x53, 0x3f, 0x4e, 0x22,
	0xf2, 0x01, 0x36, 0xe5, 0x35, 0x57, 0x32, 0x50, 0x5c, 0xb1, 0xcc, 0xe5, 0x3e, 0x1b, 0xd8, 0xbe,
	0x07, 0x55, 0xdf, 0x83, 0x9f, 0x93, 0x5c, 0x1d, 0x8d, 0x7e, 0xb1, 0xac, 0xc4, 0x29, 0x18, 0xe1,
	0x52, 0xf3, 0xe4, 0x13, 0x6c, 0x59, 0x9d, 0xe7, 0x41, 0xc2, 0x59, 0x46, 0x1f, 0x79, 0x04, 0xd8,
	0x8a, 0xdf, 0xf3, 0x33, 0xce, 0x32, 0x32, 0x86, 0x6d, 0x97, 0x10, 0xc7, 0x36, 0x62, 0xdd, 0x23,
	0xa2, 0x63, 0x23, 0xe2, 0xd8, 0x64, 0x9c, 0x54, 0xab, 0x98, 0x65, 0x3c, 0x9c, 0x63, 0x44, 0x1b,
	0xde, 0x11, 0x63, 0x6b, 0x90, 0x53, 0xe8, 0xda, 0x88, 0x34, 0x97, 0x69, 0x84, 0xc1, 0x8c, 0xaf,
	0x68, 0xd3, 0x23, 0xc5, 0x2e, 0x7e, 0x62, 0xa4, 0x31, 0x5f, 0x91, 0x6f, 0xd0, 0x73, 0xed, 0x94,
	0xaa, 0x0e, 0x6a, 0x79, 0x04, 0xed, 0xd8, 0x8e, 0xac, 0xa5, 0x93, 0x3e, 0x42, 0xa7, 0x60, 0x52,
	0x62, 0x75, 0x34, 0x1b, 0x3e, 0x3b, 0x6b, 0x0d, 0x7b, 0x36, 0x5f, 0x61, 0xc7, 0x05, 0xb0, 0x30,
	0x2c, 0x05, 0x0b, 0xef, 0x68, 0xdb, 0xa7, 0x23, 0x2b, 0x9d, 0x38, 0x87, 0x4c, 0xa0, 0xe7, 0x62,
	0x0a, 0x14, 0x21, 0xe6, 0x8a, 0x25, 0x48, 0xc1, 0x23, 0xa8, 0x6b, 0xb5, 0x1f, 0xb5, 0xa5, 0xcf,
	0x89, 0x29, 0xc5, 0xc2, 0x79, 0xd5, 0xd3, 0xa6, 0xcf, 0x39, 0x39, 0xc5, 0x36, 0x35, 0x81, 0x5e,
	0x15, 0x11, 0xb1, 0x3c, 0x41, 0xc1, 0x4b, 0x49, 0x3b, 0x3e, 0xab, 0x71, 0xda, 0x97, 0xca, 0x22,
	0x23, 0x68, 0xc4, 0xbc, 0xcc, 0x24, 0xdd, 0xf2, 0xd0, 0x2d, 0x4a, 0xde, 0x40, 0x2b, 0xe4, 0x22,
	0x47, 0x21, 0xe9, 0xb6, 0x87, 0x55, 0xc1, 0xe4, 0x2d, 0x6c, 0xf0, 0x38, 0xd6, 0x47, 0x2b, 0xe9,
	0x8e, 0x87, 0x58, 0xd3, 0xe4, 0x3d, 0x40, 0xc1, 0xf5, 0x3e, 0xca, 0x94, 0xe7, 0xb4, 0xeb, 0xf3,
	0xfb, 0xbc, 0xe7, 0xf5, 0x25, 0xba, 0xc3, 0x2c, 0xe3, 0xcb, 0x20, 0x64, 0x22, 0x92, 0xb4, 0xe7,
	0x73, 0x89, 0xac, 0xf1, 0x59, 0x0b, 0xe4, 0x18, 0xda, 0x02, 0x23, 0x67, 0x13, 0x9f, 0x95, 0x0b,
	0x8c, 0xac, 0x3a, 0x82, 0x86, 0x64, 0xb7, 0x28, 0xe9, 0x63, 0x9f, 0xfd, 0x35, 0x28, 0x19, 0xc3,
	0x96, 0x2c, 0x67, 0x52, 0xa5, 0xaa, 0x54, 0x29, 0xcf, 0x25, 0x7d, 0xe2, 0xe1, 0xfe, 0xad, 0x90,
	0x77, 0x00, 0xfa, 0x3b, 0x12, 0xcc, 0xd3, 0x70, 0x2e, 0xe9, 0xae, 0x47, 0x40, 0x5b, 0xf3, 0xe7,
	0x1a, 0xd7, 0x57, 0xd4, 0xc8, 0x4c, 0x29, 0x5c, 0x14, 0x4a, 0xd2, 0x3d, 0x9f, 0x2b, 0xaa, 0x95,
	0x13, 0x67, 0xe8, 0xfa, 0xb1, 0x40, 0x74, 0xf5, 0xf7, 0x7d, 0xea, 0x6b, 0xde, 0xd6, 0x3f, 0x86,
	0xb6, 0xba, 0x16, 0x7c, 0xa9, 0xbf, 0x43, 0x94, 0xfa, 0xec, 0xb7, 0xc1, 0x27, 0xb9, 0x24, 0xaf,
	0xa1, 0x65, 0xe6, 0xc9, 0x2a, 0xa1, 0x4f, 0x3d, 0xc4, 0xa6, 0x86, 0xaf, 0x12, 0xad, 0x99, 0xb9,
	0xb2, 0x4a, 0x68, 0xdf, 0x47, 0xd3, 0xf0, 0x55, 0x32, 0x0a, 0xa0, 0x5b, 0x8f, 0x97, 0x0b, 0x14,
	0xb7, 0x69, 0x88, 0xe4, 0x1c, 0xf6, 0xce, 0x50, 0xd5, 0xaf, 0x4f, 0xb9, 0x38, 0x4d, 0x57, 0xaa,
	0x14, 0x48, 0x76, 0xdd, 0xa4, 0x72, 0xcf, 0x53, 0x3b, 0x1e, 0xfb, 0xf4, 0x9f, 0x01, 0xe6, 0x86,
	0xe3, 0xc1, 0x7f, 0xb3, 0xa6, 0xf9, 0xd7, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x16,
	0x6c, 0x18, 0xb5, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeamStatsServiceClient is the client API for TeamStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamStatsServiceClient interface {
	GetTeamStatsForFixture(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*TeamStatsResponse, error)
}

type teamStatsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTeamStatsServiceClient(cc *grpc.ClientConn) TeamStatsServiceClient {
	return &teamStatsServiceClient{cc}
}

func (c *teamStatsServiceClient) GetTeamStatsForFixture(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*TeamStatsResponse, error) {
	out := new(TeamStatsResponse)
	err := c.cc.Invoke(ctx, "/proto.TeamStatsService/GetTeamStatsForFixture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamStatsServiceServer is the server API for TeamStatsService service.
type TeamStatsServiceServer interface {
	GetTeamStatsForFixture(context.Context, *FixtureRequest) (*TeamStatsResponse, error)
}

// UnimplementedTeamStatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTeamStatsServiceServer struct {
}

func (*UnimplementedTeamStatsServiceServer) GetTeamStatsForFixture(ctx context.Context, req *FixtureRequest) (*TeamStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamStatsForFixture not implemented")
}

func RegisterTeamStatsServiceServer(s *grpc.Server, srv TeamStatsServiceServer) {
	s.RegisterService(&_TeamStatsService_serviceDesc, srv)
}

func _TeamStatsService_GetTeamStatsForFixture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixtureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamStatsServiceServer).GetTeamStatsForFixture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeamStatsService/GetTeamStatsForFixture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamStatsServiceServer).GetTeamStatsForFixture(ctx, req.(*FixtureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeamStatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TeamStatsService",
	HandlerType: (*TeamStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamStatsForFixture",
			Handler:    _TeamStatsService_GetTeamStatsForFixture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/grpc/proto/team_stats.proto",
}
