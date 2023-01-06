// Code generated by protoc-gen-struct-transformer, version: <dev>. DO NOT EDIT.
// source file: team.proto
// source package: team

package transform

import (
	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/utils/transformhelpers"
)

// "google.protobuf.Timestamp": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.TimestampRules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetListHubRequest": target: "", Omitted: true, OneofDecl: ""
// "team.GetListTeamResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.BoolRules": target: "", Omitted: true, OneofDecl: ""
// "validate.EnumRules": target: "", Omitted: true, OneofDecl: ""
// "validate.DurationRules": target: "", Omitted: true, OneofDecl: ""
// "team.GetTeamByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "example.HelloRequest": target: "", Omitted: true, OneofDecl: ""
// "google.api.CustomHttpPattern": target: "", Omitted: true, OneofDecl: ""
// "hub.CreateHubRequest": target: "", Omitted: true, OneofDecl: ""
// "hub.UpdateHubResponse": target: "", Omitted: true, OneofDecl: ""
// "user.GetUserByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofOptions": target: "", Omitted: true, OneofDecl: ""
// "google.api.Http": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.StringRules": target: "", Omitted: true, OneofDecl: ""
// "hub.UpdateHubRequest": target: "", Omitted: true, OneofDecl: ""
// "user.GetListUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "nrpc.Error": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "team.GetListTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "user.GetListUserRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Duration": target: "", Omitted: true, OneofDecl: ""
// "nrpc.HeartBeat": target: "", Omitted: true, OneofDecl: ""
// "team.GetTeamByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "team.UpdateTeamResponse": target: "", Omitted: true, OneofDecl: ""
// "user.GetUserByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorSet": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "search.SearchTeamHubResponse": target: "", Omitted: true, OneofDecl: ""
// "user.CreateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "validate.FloatRules": target: "", Omitted: true, OneofDecl: ""
// "team.CreateTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.CreateHubResponse": target: "", Omitted: true, OneofDecl: ""
// "validate.BytesRules": target: "", Omitted: true, OneofDecl: ""
// "team.UpdateTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "example.HelloReply": target: "", Omitted: true, OneofDecl: ""
// "validate.Int64Rules": target: "", Omitted: true, OneofDecl: ""
// "nrpc.Void": target: "", Omitted: true, OneofDecl: ""
// "nrpc.NoReply": target: "", Omitted: true, OneofDecl: ""
// "validate.MessageRules": target: "", Omitted: true, OneofDecl: ""
// "validate.RepeatedRules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetHubByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.MapRules": target: "", Omitted: true, OneofDecl: ""
// "team.Team": target: "Team", Omitted: false, OneofDecl: ""
// "team.CreateTeamResponse": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ExtensionRangeOptions": target: "", Omitted: true, OneofDecl: ""
// "nrpc.NoRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodOptions": target: "", Omitted: true, OneofDecl: ""
// "google.api.HttpRule": target: "", Omitted: true, OneofDecl: ""
// "validate.Int32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetListHubResponse": target: "", Omitted: true, OneofDecl: ""
// "user.User": target: "User", Omitted: false, OneofDecl: ""
// "google.protobuf.FileDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.DescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.GeneratedCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "validate.DoubleRules": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.Hub": target: "Hub", Omitted: false, OneofDecl: ""
// "search.SearchTeamHubRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "user.CreateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetHubByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.UninterpretedOption": target: "", Omitted: true, OneofDecl: ""
// "validate.FieldRules": target: "", Omitted: true, OneofDecl: ""
// "validate.AnyRules": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MessageOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.SourceCodeInfo": target: "", Omitted: true, OneofDecl: ""
// message "GetTeamByIDRequest" has no option "transformer.go_struct", skipped...
// message "GetTeamByIDResponse" has no option "transformer.go_struct", skipped...

// Target struct fields:
// Field: "HubID", Type: "pgtype.Text", isPointer: false
// Field: "LocationID", Type: "pgtype.Text", isPointer: false
// Field: "CreatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "UpdatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "DeletedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "ID", Type: "pgtype.Text", isPointer: false
// Field: "Name", Type: "pgtype.Text", isPointer: false
// Field: "Type", Type: "pgtype.Text", isPointer: false

// ===============================
// fdp.Name: "id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Id", gname: "ID"

// ===============================
// fdp.Name: "name", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Name", gname: "Name"

// ===============================
// fdp.Name: "type", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Type", gname: "Type"

// ===============================
// fdp.Name: "location_id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "LocationId", gname: "LocationID"

// ===============================
// fdp.Name: "hub_id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "HubId", gname: "HubID"
// message "CreateTeamRequest" has no option "transformer.go_struct", skipped...
// message "CreateTeamResponse" has no option "transformer.go_struct", skipped...
// message "GetListTeamRequest" has no option "transformer.go_struct", skipped...
// message "GetListTeamResponse" has no option "transformer.go_struct", skipped...
// message "UpdateTeamRequest" has no option "transformer.go_struct", skipped...
// message "UpdateTeamResponse" has no option "transformer.go_struct", skipped...
func PbToTeamPtr(src *pb.Team, opts ...TransformParam) *models.Team {
	if src == nil {
		return nil
	}

	d := PbToTeam(*src, opts...)
	return &d
}

func PbToTeamPtrList(src []*pb.Team, opts ...TransformParam) []*models.Team {
	resp := make([]*models.Team, len(src))

	for i, s := range src {
		resp[i] = PbToTeamPtr(s, opts...)
	}

	return resp
}

func PbToTeamPtrVal(src *pb.Team, opts ...TransformParam) models.Team {
	if src == nil {
		return models.Team{}
	}

	return PbToTeam(*src, opts...)
}

func PbToTeamPtrValList(src []*pb.Team, opts ...TransformParam) []models.Team {
	resp := make([]models.Team, len(src))

	for i, s := range src {
		resp[i] = PbToTeam(*s)
	}

	return resp
}

// PbToTeamList is DEPRECATED. Use PbToTeamPtrValList instead.
func PbToTeamList(src []*pb.Team, opts ...TransformParam) []models.Team {
	return PbToTeamPtrValList(src)
}

func PbToTeam(src pb.Team, opts ...TransformParam) models.Team {
	s := models.Team{
		ID:         transformhelpers.StringToPgtypeText(src.Id),
		Name:       transformhelpers.StringToPgtypeText(src.Name),
		Type:       transformhelpers.StringToPgtypeText(src.Type),
		LocationID: transformhelpers.StringToPgtypeText(src.LocationId),
		HubID:      transformhelpers.StringToPgtypeText(src.HubId),
	}

	applyOptions(opts...)

	return s
}

func PbToTeamValPtr(src pb.Team, opts ...TransformParam) *models.Team {
	d := PbToTeam(src, opts...)
	return &d
}

func PbToTeamValList(src []pb.Team, opts ...TransformParam) []models.Team {
	resp := make([]models.Team, len(src))

	for i, s := range src {
		resp[i] = PbToTeam(s, opts...)
	}

	return resp
}

func TeamToPbPtr(src *models.Team, opts ...TransformParam) *pb.Team {
	if src == nil {
		return nil
	}

	d := TeamToPb(*src, opts...)
	return &d
}

func TeamToPbPtrList(src []*models.Team, opts ...TransformParam) []*pb.Team {
	resp := make([]*pb.Team, len(src))

	for i, s := range src {
		resp[i] = TeamToPbPtr(s, opts...)
	}

	return resp
}

func TeamToPbPtrVal(src *models.Team, opts ...TransformParam) pb.Team {
	if src == nil {
		return pb.Team{}
	}

	return TeamToPb(*src, opts...)
}

func TeamToPbValPtrList(src []models.Team, opts ...TransformParam) []*pb.Team {
	resp := make([]*pb.Team, len(src))

	for i, s := range src {
		g := TeamToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// TeamToPbList is DEPRECATED. Use TeamToPbValPtrList instead.
func TeamToPbList(src []models.Team, opts ...TransformParam) []*pb.Team {
	return TeamToPbValPtrList(src)
}

func TeamToPb(src models.Team, opts ...TransformParam) pb.Team {
	s := pb.Team{
		Id:         transformhelpers.PgtypeTextToString(src.ID),
		Name:       transformhelpers.PgtypeTextToString(src.Name),
		Type:       transformhelpers.PgtypeTextToString(src.Type),
		LocationId: transformhelpers.PgtypeTextToString(src.LocationID),
		HubId:      transformhelpers.PgtypeTextToString(src.HubID),
	}

	applyOptions(opts...)

	return s
}

func TeamToPbValPtr(src models.Team, opts ...TransformParam) *pb.Team {
	d := TeamToPb(src, opts...)
	return &d
}

func TeamToPbValList(src []models.Team, opts ...TransformParam) []pb.Team {
	resp := make([]pb.Team, len(src))

	for i, s := range src {
		resp[i] = TeamToPb(s, opts...)
	}

	return resp
}
