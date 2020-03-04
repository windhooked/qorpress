// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package grafeas

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/api/option"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"

	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockGrafeasServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	grafeaspb.GrafeasServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockGrafeasServer) GetOccurrence(ctx context.Context, req *grafeaspb.GetOccurrenceRequest) (*grafeaspb.Occurrence, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Occurrence), nil
}

func (s *mockGrafeasServer) ListOccurrences(ctx context.Context, req *grafeaspb.ListOccurrencesRequest) (*grafeaspb.ListOccurrencesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.ListOccurrencesResponse), nil
}

func (s *mockGrafeasServer) DeleteOccurrence(ctx context.Context, req *grafeaspb.DeleteOccurrenceRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockGrafeasServer) CreateOccurrence(ctx context.Context, req *grafeaspb.CreateOccurrenceRequest) (*grafeaspb.Occurrence, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Occurrence), nil
}

func (s *mockGrafeasServer) BatchCreateOccurrences(ctx context.Context, req *grafeaspb.BatchCreateOccurrencesRequest) (*grafeaspb.BatchCreateOccurrencesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.BatchCreateOccurrencesResponse), nil
}

func (s *mockGrafeasServer) UpdateOccurrence(ctx context.Context, req *grafeaspb.UpdateOccurrenceRequest) (*grafeaspb.Occurrence, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Occurrence), nil
}

func (s *mockGrafeasServer) GetOccurrenceNote(ctx context.Context, req *grafeaspb.GetOccurrenceNoteRequest) (*grafeaspb.Note, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Note), nil
}

func (s *mockGrafeasServer) GetNote(ctx context.Context, req *grafeaspb.GetNoteRequest) (*grafeaspb.Note, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Note), nil
}

func (s *mockGrafeasServer) ListNotes(ctx context.Context, req *grafeaspb.ListNotesRequest) (*grafeaspb.ListNotesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.ListNotesResponse), nil
}

func (s *mockGrafeasServer) DeleteNote(ctx context.Context, req *grafeaspb.DeleteNoteRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockGrafeasServer) CreateNote(ctx context.Context, req *grafeaspb.CreateNoteRequest) (*grafeaspb.Note, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Note), nil
}

func (s *mockGrafeasServer) BatchCreateNotes(ctx context.Context, req *grafeaspb.BatchCreateNotesRequest) (*grafeaspb.BatchCreateNotesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.BatchCreateNotesResponse), nil
}

func (s *mockGrafeasServer) UpdateNote(ctx context.Context, req *grafeaspb.UpdateNoteRequest) (*grafeaspb.Note, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.Note), nil
}

func (s *mockGrafeasServer) ListNoteOccurrences(ctx context.Context, req *grafeaspb.ListNoteOccurrencesRequest) (*grafeaspb.ListNoteOccurrencesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*grafeaspb.ListNoteOccurrencesResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockGrafeas mockGrafeasServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	grafeaspb.RegisterGrafeasServer(serv, &mockGrafeas)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestGrafeasGetOccurrence(t *testing.T) {
	var name2 string = "name2-1052831874"
	var resourceUri string = "resourceUri-384040517"
	var noteName string = "noteName1780787896"
	var remediation string = "remediation779381797"
	var expectedResponse = &grafeaspb.Occurrence{
		Name:        name2,
		ResourceUri: resourceUri,
		NoteName:    noteName,
		Remediation: remediation,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.GetOccurrenceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOccurrence(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasGetOccurrenceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.GetOccurrenceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOccurrence(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasListOccurrences(t *testing.T) {
	var nextPageToken string = ""
	var occurrencesElement *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var occurrences = []*grafeaspb.Occurrence{occurrencesElement}
	var expectedResponse = &grafeaspb.ListOccurrencesResponse{
		NextPageToken: nextPageToken,
		Occurrences:   occurrences,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var request = &grafeaspb.ListOccurrencesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListOccurrences(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Occurrences[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasListOccurrencesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var request = &grafeaspb.ListOccurrencesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListOccurrences(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasDeleteOccurrence(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.DeleteOccurrenceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteOccurrence(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestGrafeasDeleteOccurrenceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.DeleteOccurrenceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteOccurrence(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestGrafeasCreateOccurrence(t *testing.T) {
	var name string = "name3373707"
	var resourceUri string = "resourceUri-384040517"
	var noteName string = "noteName1780787896"
	var remediation string = "remediation779381797"
	var expectedResponse = &grafeaspb.Occurrence{
		Name:        name,
		ResourceUri: resourceUri,
		NoteName:    noteName,
		Remediation: remediation,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var occurrence *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var request = &grafeaspb.CreateOccurrenceRequest{
		Parent:     formattedParent,
		Occurrence: occurrence,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateOccurrence(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasCreateOccurrenceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var occurrence *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var request = &grafeaspb.CreateOccurrenceRequest{
		Parent:     formattedParent,
		Occurrence: occurrence,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateOccurrence(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasBatchCreateOccurrences(t *testing.T) {
	var expectedResponse *grafeaspb.BatchCreateOccurrencesResponse = &grafeaspb.BatchCreateOccurrencesResponse{}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var occurrences []*grafeaspb.Occurrence = nil
	var request = &grafeaspb.BatchCreateOccurrencesRequest{
		Parent:      formattedParent,
		Occurrences: occurrences,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateOccurrences(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasBatchCreateOccurrencesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var occurrences []*grafeaspb.Occurrence = nil
	var request = &grafeaspb.BatchCreateOccurrencesRequest{
		Parent:      formattedParent,
		Occurrences: occurrences,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateOccurrences(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasUpdateOccurrence(t *testing.T) {
	var name2 string = "name2-1052831874"
	var resourceUri string = "resourceUri-384040517"
	var noteName string = "noteName1780787896"
	var remediation string = "remediation779381797"
	var expectedResponse = &grafeaspb.Occurrence{
		Name:        name2,
		ResourceUri: resourceUri,
		NoteName:    noteName,
		Remediation: remediation,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var occurrence *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var request = &grafeaspb.UpdateOccurrenceRequest{
		Name:       formattedName,
		Occurrence: occurrence,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateOccurrence(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasUpdateOccurrenceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var occurrence *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var request = &grafeaspb.UpdateOccurrenceRequest{
		Name:       formattedName,
		Occurrence: occurrence,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateOccurrence(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasGetOccurrenceNote(t *testing.T) {
	var name2 string = "name2-1052831874"
	var shortDescription string = "shortDescription-235369287"
	var longDescription string = "longDescription-1747792199"
	var expectedResponse = &grafeaspb.Note{
		Name:             name2,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.GetOccurrenceNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOccurrenceNote(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasGetOccurrenceNoteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/occurrences/%s", "[PROJECT]", "[OCCURRENCE]")
	var request = &grafeaspb.GetOccurrenceNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOccurrenceNote(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasGetNote(t *testing.T) {
	var name2 string = "name2-1052831874"
	var shortDescription string = "shortDescription-235369287"
	var longDescription string = "longDescription-1747792199"
	var expectedResponse = &grafeaspb.Note{
		Name:             name2,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.GetNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetNote(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasGetNoteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.GetNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetNote(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasListNotes(t *testing.T) {
	var nextPageToken string = ""
	var notesElement *grafeaspb.Note = &grafeaspb.Note{}
	var notes = []*grafeaspb.Note{notesElement}
	var expectedResponse = &grafeaspb.ListNotesResponse{
		NextPageToken: nextPageToken,
		Notes:         notes,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var request = &grafeaspb.ListNotesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListNotes(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Notes[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasListNotesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var request = &grafeaspb.ListNotesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListNotes(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasDeleteNote(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.DeleteNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteNote(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestGrafeasDeleteNoteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.DeleteNoteRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteNote(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestGrafeasCreateNote(t *testing.T) {
	var name string = "name3373707"
	var shortDescription string = "shortDescription-235369287"
	var longDescription string = "longDescription-1747792199"
	var expectedResponse = &grafeaspb.Note{
		Name:             name,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var noteId string = "noteId2129224840"
	var note *grafeaspb.Note = &grafeaspb.Note{}
	var request = &grafeaspb.CreateNoteRequest{
		Parent: formattedParent,
		NoteId: noteId,
		Note:   note,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateNote(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasCreateNoteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var noteId string = "noteId2129224840"
	var note *grafeaspb.Note = &grafeaspb.Note{}
	var request = &grafeaspb.CreateNoteRequest{
		Parent: formattedParent,
		NoteId: noteId,
		Note:   note,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateNote(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasBatchCreateNotes(t *testing.T) {
	var expectedResponse *grafeaspb.BatchCreateNotesResponse = &grafeaspb.BatchCreateNotesResponse{}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var notes map[string]*grafeaspb.Note = nil
	var request = &grafeaspb.BatchCreateNotesRequest{
		Parent: formattedParent,
		Notes:  notes,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateNotes(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasBatchCreateNotesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s", "[PROJECT]")
	var notes map[string]*grafeaspb.Note = nil
	var request = &grafeaspb.BatchCreateNotesRequest{
		Parent: formattedParent,
		Notes:  notes,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateNotes(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasUpdateNote(t *testing.T) {
	var name2 string = "name2-1052831874"
	var shortDescription string = "shortDescription-235369287"
	var longDescription string = "longDescription-1747792199"
	var expectedResponse = &grafeaspb.Note{
		Name:             name2,
		ShortDescription: shortDescription,
		LongDescription:  longDescription,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var note *grafeaspb.Note = &grafeaspb.Note{}
	var request = &grafeaspb.UpdateNoteRequest{
		Name: formattedName,
		Note: note,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateNote(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasUpdateNoteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var note *grafeaspb.Note = &grafeaspb.Note{}
	var request = &grafeaspb.UpdateNoteRequest{
		Name: formattedName,
		Note: note,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateNote(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestGrafeasListNoteOccurrences(t *testing.T) {
	var nextPageToken string = ""
	var occurrencesElement *grafeaspb.Occurrence = &grafeaspb.Occurrence{}
	var occurrences = []*grafeaspb.Occurrence{occurrencesElement}
	var expectedResponse = &grafeaspb.ListNoteOccurrencesResponse{
		NextPageToken: nextPageToken,
		Occurrences:   occurrences,
	}

	mockGrafeas.err = nil
	mockGrafeas.reqs = nil

	mockGrafeas.resps = append(mockGrafeas.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.ListNoteOccurrencesRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListNoteOccurrences(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockGrafeas.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Occurrences[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestGrafeasListNoteOccurrencesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockGrafeas.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/notes/%s", "[PROJECT]", "[NOTE]")
	var request = &grafeaspb.ListNoteOccurrencesRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListNoteOccurrences(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
