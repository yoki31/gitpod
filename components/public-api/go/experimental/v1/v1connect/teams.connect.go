// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/experimental/v1/teams.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TeamsServiceName is the fully-qualified name of the TeamsService service.
	TeamsServiceName = "gitpod.experimental.v1.TeamsService"
)

// TeamsServiceClient is a client for the gitpod.experimental.v1.TeamsService service.
type TeamsServiceClient interface {
	// CreateTeam creates a new Team.
	CreateTeam(context.Context, *connect_go.Request[v1.CreateTeamRequest]) (*connect_go.Response[v1.CreateTeamResponse], error)
	// GetTeam retrieves a single Team.
	GetTeam(context.Context, *connect_go.Request[v1.GetTeamRequest]) (*connect_go.Response[v1.GetTeamResponse], error)
	// ListTeams lists the caller has access to.
	ListTeams(context.Context, *connect_go.Request[v1.ListTeamsRequest]) (*connect_go.Response[v1.ListTeamsResponse], error)
	// DeleteTeam deletes the specified team.
	DeleteTeam(context.Context, *connect_go.Request[v1.DeleteTeamRequest]) (*connect_go.Response[v1.DeleteTeamResponse], error)
	// GetTeamInvitation retrieves the invitation for a Team.
	GetTeamInvitation(context.Context, *connect_go.Request[v1.GetTeamInvitationRequest]) (*connect_go.Response[v1.GetTeamInvitationResponse], error)
	// JoinTeam makes the caller a TeamMember of the Team.
	JoinTeam(context.Context, *connect_go.Request[v1.JoinTeamRequest]) (*connect_go.Response[v1.JoinTeamResponse], error)
	// ResetTeamInvitation resets the invitation_id for a Team.
	ResetTeamInvitation(context.Context, *connect_go.Request[v1.ResetTeamInvitationRequest]) (*connect_go.Response[v1.ResetTeamInvitationResponse], error)
	// ListTeamMembers lists the members of a Team.
	ListTeamMembers(context.Context, *connect_go.Request[v1.ListTeamMembersRequest]) (*connect_go.Response[v1.ListTeamMembersResponse], error)
	// UpdateTeamMember updates team membership properties.
	UpdateTeamMember(context.Context, *connect_go.Request[v1.UpdateTeamMemberRequest]) (*connect_go.Response[v1.UpdateTeamMemberResponse], error)
	// DeleteTeamMember removes a TeamMember from the Team.
	DeleteTeamMember(context.Context, *connect_go.Request[v1.DeleteTeamMemberRequest]) (*connect_go.Response[v1.DeleteTeamMemberResponse], error)
}

// NewTeamsServiceClient constructs a client for the gitpod.experimental.v1.TeamsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTeamsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TeamsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &teamsServiceClient{
		createTeam: connect_go.NewClient[v1.CreateTeamRequest, v1.CreateTeamResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/CreateTeam",
			opts...,
		),
		getTeam: connect_go.NewClient[v1.GetTeamRequest, v1.GetTeamResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/GetTeam",
			opts...,
		),
		listTeams: connect_go.NewClient[v1.ListTeamsRequest, v1.ListTeamsResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/ListTeams",
			opts...,
		),
		deleteTeam: connect_go.NewClient[v1.DeleteTeamRequest, v1.DeleteTeamResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/DeleteTeam",
			opts...,
		),
		getTeamInvitation: connect_go.NewClient[v1.GetTeamInvitationRequest, v1.GetTeamInvitationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/GetTeamInvitation",
			opts...,
		),
		joinTeam: connect_go.NewClient[v1.JoinTeamRequest, v1.JoinTeamResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/JoinTeam",
			opts...,
		),
		resetTeamInvitation: connect_go.NewClient[v1.ResetTeamInvitationRequest, v1.ResetTeamInvitationResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/ResetTeamInvitation",
			opts...,
		),
		listTeamMembers: connect_go.NewClient[v1.ListTeamMembersRequest, v1.ListTeamMembersResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/ListTeamMembers",
			opts...,
		),
		updateTeamMember: connect_go.NewClient[v1.UpdateTeamMemberRequest, v1.UpdateTeamMemberResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/UpdateTeamMember",
			opts...,
		),
		deleteTeamMember: connect_go.NewClient[v1.DeleteTeamMemberRequest, v1.DeleteTeamMemberResponse](
			httpClient,
			baseURL+"/gitpod.experimental.v1.TeamsService/DeleteTeamMember",
			opts...,
		),
	}
}

// teamsServiceClient implements TeamsServiceClient.
type teamsServiceClient struct {
	createTeam          *connect_go.Client[v1.CreateTeamRequest, v1.CreateTeamResponse]
	getTeam             *connect_go.Client[v1.GetTeamRequest, v1.GetTeamResponse]
	listTeams           *connect_go.Client[v1.ListTeamsRequest, v1.ListTeamsResponse]
	deleteTeam          *connect_go.Client[v1.DeleteTeamRequest, v1.DeleteTeamResponse]
	getTeamInvitation   *connect_go.Client[v1.GetTeamInvitationRequest, v1.GetTeamInvitationResponse]
	joinTeam            *connect_go.Client[v1.JoinTeamRequest, v1.JoinTeamResponse]
	resetTeamInvitation *connect_go.Client[v1.ResetTeamInvitationRequest, v1.ResetTeamInvitationResponse]
	listTeamMembers     *connect_go.Client[v1.ListTeamMembersRequest, v1.ListTeamMembersResponse]
	updateTeamMember    *connect_go.Client[v1.UpdateTeamMemberRequest, v1.UpdateTeamMemberResponse]
	deleteTeamMember    *connect_go.Client[v1.DeleteTeamMemberRequest, v1.DeleteTeamMemberResponse]
}

// CreateTeam calls gitpod.experimental.v1.TeamsService.CreateTeam.
func (c *teamsServiceClient) CreateTeam(ctx context.Context, req *connect_go.Request[v1.CreateTeamRequest]) (*connect_go.Response[v1.CreateTeamResponse], error) {
	return c.createTeam.CallUnary(ctx, req)
}

// GetTeam calls gitpod.experimental.v1.TeamsService.GetTeam.
func (c *teamsServiceClient) GetTeam(ctx context.Context, req *connect_go.Request[v1.GetTeamRequest]) (*connect_go.Response[v1.GetTeamResponse], error) {
	return c.getTeam.CallUnary(ctx, req)
}

// ListTeams calls gitpod.experimental.v1.TeamsService.ListTeams.
func (c *teamsServiceClient) ListTeams(ctx context.Context, req *connect_go.Request[v1.ListTeamsRequest]) (*connect_go.Response[v1.ListTeamsResponse], error) {
	return c.listTeams.CallUnary(ctx, req)
}

// DeleteTeam calls gitpod.experimental.v1.TeamsService.DeleteTeam.
func (c *teamsServiceClient) DeleteTeam(ctx context.Context, req *connect_go.Request[v1.DeleteTeamRequest]) (*connect_go.Response[v1.DeleteTeamResponse], error) {
	return c.deleteTeam.CallUnary(ctx, req)
}

// GetTeamInvitation calls gitpod.experimental.v1.TeamsService.GetTeamInvitation.
func (c *teamsServiceClient) GetTeamInvitation(ctx context.Context, req *connect_go.Request[v1.GetTeamInvitationRequest]) (*connect_go.Response[v1.GetTeamInvitationResponse], error) {
	return c.getTeamInvitation.CallUnary(ctx, req)
}

// JoinTeam calls gitpod.experimental.v1.TeamsService.JoinTeam.
func (c *teamsServiceClient) JoinTeam(ctx context.Context, req *connect_go.Request[v1.JoinTeamRequest]) (*connect_go.Response[v1.JoinTeamResponse], error) {
	return c.joinTeam.CallUnary(ctx, req)
}

// ResetTeamInvitation calls gitpod.experimental.v1.TeamsService.ResetTeamInvitation.
func (c *teamsServiceClient) ResetTeamInvitation(ctx context.Context, req *connect_go.Request[v1.ResetTeamInvitationRequest]) (*connect_go.Response[v1.ResetTeamInvitationResponse], error) {
	return c.resetTeamInvitation.CallUnary(ctx, req)
}

// ListTeamMembers calls gitpod.experimental.v1.TeamsService.ListTeamMembers.
func (c *teamsServiceClient) ListTeamMembers(ctx context.Context, req *connect_go.Request[v1.ListTeamMembersRequest]) (*connect_go.Response[v1.ListTeamMembersResponse], error) {
	return c.listTeamMembers.CallUnary(ctx, req)
}

// UpdateTeamMember calls gitpod.experimental.v1.TeamsService.UpdateTeamMember.
func (c *teamsServiceClient) UpdateTeamMember(ctx context.Context, req *connect_go.Request[v1.UpdateTeamMemberRequest]) (*connect_go.Response[v1.UpdateTeamMemberResponse], error) {
	return c.updateTeamMember.CallUnary(ctx, req)
}

// DeleteTeamMember calls gitpod.experimental.v1.TeamsService.DeleteTeamMember.
func (c *teamsServiceClient) DeleteTeamMember(ctx context.Context, req *connect_go.Request[v1.DeleteTeamMemberRequest]) (*connect_go.Response[v1.DeleteTeamMemberResponse], error) {
	return c.deleteTeamMember.CallUnary(ctx, req)
}

// TeamsServiceHandler is an implementation of the gitpod.experimental.v1.TeamsService service.
type TeamsServiceHandler interface {
	// CreateTeam creates a new Team.
	CreateTeam(context.Context, *connect_go.Request[v1.CreateTeamRequest]) (*connect_go.Response[v1.CreateTeamResponse], error)
	// GetTeam retrieves a single Team.
	GetTeam(context.Context, *connect_go.Request[v1.GetTeamRequest]) (*connect_go.Response[v1.GetTeamResponse], error)
	// ListTeams lists the caller has access to.
	ListTeams(context.Context, *connect_go.Request[v1.ListTeamsRequest]) (*connect_go.Response[v1.ListTeamsResponse], error)
	// DeleteTeam deletes the specified team.
	DeleteTeam(context.Context, *connect_go.Request[v1.DeleteTeamRequest]) (*connect_go.Response[v1.DeleteTeamResponse], error)
	// GetTeamInvitation retrieves the invitation for a Team.
	GetTeamInvitation(context.Context, *connect_go.Request[v1.GetTeamInvitationRequest]) (*connect_go.Response[v1.GetTeamInvitationResponse], error)
	// JoinTeam makes the caller a TeamMember of the Team.
	JoinTeam(context.Context, *connect_go.Request[v1.JoinTeamRequest]) (*connect_go.Response[v1.JoinTeamResponse], error)
	// ResetTeamInvitation resets the invitation_id for a Team.
	ResetTeamInvitation(context.Context, *connect_go.Request[v1.ResetTeamInvitationRequest]) (*connect_go.Response[v1.ResetTeamInvitationResponse], error)
	// ListTeamMembers lists the members of a Team.
	ListTeamMembers(context.Context, *connect_go.Request[v1.ListTeamMembersRequest]) (*connect_go.Response[v1.ListTeamMembersResponse], error)
	// UpdateTeamMember updates team membership properties.
	UpdateTeamMember(context.Context, *connect_go.Request[v1.UpdateTeamMemberRequest]) (*connect_go.Response[v1.UpdateTeamMemberResponse], error)
	// DeleteTeamMember removes a TeamMember from the Team.
	DeleteTeamMember(context.Context, *connect_go.Request[v1.DeleteTeamMemberRequest]) (*connect_go.Response[v1.DeleteTeamMemberResponse], error)
}

// NewTeamsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTeamsServiceHandler(svc TeamsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.experimental.v1.TeamsService/CreateTeam", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/CreateTeam",
		svc.CreateTeam,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/GetTeam", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/GetTeam",
		svc.GetTeam,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/ListTeams", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/ListTeams",
		svc.ListTeams,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/DeleteTeam", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/DeleteTeam",
		svc.DeleteTeam,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/GetTeamInvitation", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/GetTeamInvitation",
		svc.GetTeamInvitation,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/JoinTeam", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/JoinTeam",
		svc.JoinTeam,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/ResetTeamInvitation", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/ResetTeamInvitation",
		svc.ResetTeamInvitation,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/ListTeamMembers", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/ListTeamMembers",
		svc.ListTeamMembers,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/UpdateTeamMember", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/UpdateTeamMember",
		svc.UpdateTeamMember,
		opts...,
	))
	mux.Handle("/gitpod.experimental.v1.TeamsService/DeleteTeamMember", connect_go.NewUnaryHandler(
		"/gitpod.experimental.v1.TeamsService/DeleteTeamMember",
		svc.DeleteTeamMember,
		opts...,
	))
	return "/gitpod.experimental.v1.TeamsService/", mux
}

// UnimplementedTeamsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTeamsServiceHandler struct{}

func (UnimplementedTeamsServiceHandler) CreateTeam(context.Context, *connect_go.Request[v1.CreateTeamRequest]) (*connect_go.Response[v1.CreateTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.CreateTeam is not implemented"))
}

func (UnimplementedTeamsServiceHandler) GetTeam(context.Context, *connect_go.Request[v1.GetTeamRequest]) (*connect_go.Response[v1.GetTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.GetTeam is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ListTeams(context.Context, *connect_go.Request[v1.ListTeamsRequest]) (*connect_go.Response[v1.ListTeamsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.ListTeams is not implemented"))
}

func (UnimplementedTeamsServiceHandler) DeleteTeam(context.Context, *connect_go.Request[v1.DeleteTeamRequest]) (*connect_go.Response[v1.DeleteTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.DeleteTeam is not implemented"))
}

func (UnimplementedTeamsServiceHandler) GetTeamInvitation(context.Context, *connect_go.Request[v1.GetTeamInvitationRequest]) (*connect_go.Response[v1.GetTeamInvitationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.GetTeamInvitation is not implemented"))
}

func (UnimplementedTeamsServiceHandler) JoinTeam(context.Context, *connect_go.Request[v1.JoinTeamRequest]) (*connect_go.Response[v1.JoinTeamResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.JoinTeam is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ResetTeamInvitation(context.Context, *connect_go.Request[v1.ResetTeamInvitationRequest]) (*connect_go.Response[v1.ResetTeamInvitationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.ResetTeamInvitation is not implemented"))
}

func (UnimplementedTeamsServiceHandler) ListTeamMembers(context.Context, *connect_go.Request[v1.ListTeamMembersRequest]) (*connect_go.Response[v1.ListTeamMembersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.ListTeamMembers is not implemented"))
}

func (UnimplementedTeamsServiceHandler) UpdateTeamMember(context.Context, *connect_go.Request[v1.UpdateTeamMemberRequest]) (*connect_go.Response[v1.UpdateTeamMemberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.UpdateTeamMember is not implemented"))
}

func (UnimplementedTeamsServiceHandler) DeleteTeamMember(context.Context, *connect_go.Request[v1.DeleteTeamMemberRequest]) (*connect_go.Response[v1.DeleteTeamMemberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.experimental.v1.TeamsService.DeleteTeamMember is not implemented"))
}
