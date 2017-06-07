package mock_addrbookrestapi_test

import (
	"fmt"
	"testing"

	pb "bashdrew/yumcodingassignment/addrbookrestapi"
	pbmock "bashdrew/yumcodingassignment/mock_addrbookrestapi"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestGetPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookRestAPIClient := pbmock.NewMockAddrBookRestAPIClient(ctrl)
	req := &pb.PersonRequest{Id: int64(0)}
	mockAddrBookRestAPIClient.EXPECT().GetPerson(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReply{Firstname: "Andrew"}, nil)
	testGetPerson(t, mockAddrBookRestAPIClient)
}

func testGetPerson(t *testing.T, client pb.AddrBookRestAPIClient) {
	r, err := client.GetPerson(context.Background(), &pb.PersonRequest{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Firstname != "Andrew" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestPostPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookRestAPIClient := pbmock.NewMockAddrBookRestAPIClient(ctrl)
	req := &pb.PersonRequest{Id: int64(0)}
	mockAddrBookRestAPIClient.EXPECT().PostPerson(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReply{Lastname: "Amargo"}, nil)
	testPostPerson(t, mockAddrBookRestAPIClient)
}

func testPostPerson(t *testing.T, client pb.AddrBookRestAPIClient) {
	r, err := client.PostPerson(context.Background(), &pb.PersonRequest{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Lastname != "Amargo" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestDeletePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookRestAPIClient := pbmock.NewMockAddrBookRestAPIClient(ctrl)
	req := &pb.PersonRequest{Id: int64(0)}
	mockAddrBookRestAPIClient.EXPECT().DeletePerson(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReply{Email: "aamargo@gmail.com"}, nil)
	testDeletePerson(t, mockAddrBookRestAPIClient)
}

func testDeletePerson(t *testing.T, client pb.AddrBookRestAPIClient) {
	r, err := client.DeletePerson(context.Background(), &pb.PersonRequest{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Email != "aamargo@gmail.com" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestPutPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookRestAPIClient := pbmock.NewMockAddrBookRestAPIClient(ctrl)
	req := &pb.PersonRequest{Id: int64(0)}
	mockAddrBookRestAPIClient.EXPECT().PutPerson(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReply{Phoneno: "123-123-1234"}, nil)
	testPutPerson(t, mockAddrBookRestAPIClient)
}

func testPutPerson(t *testing.T, client pb.AddrBookRestAPIClient) {
	r, err := client.PutPerson(context.Background(), &pb.PersonRequest{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Phoneno != "123-123-1234" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}
