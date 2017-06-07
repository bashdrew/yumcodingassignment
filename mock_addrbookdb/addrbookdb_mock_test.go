package mock_addrbookdb_test

import (
	"fmt"
	"testing"

	pb "bashdrew/yumcodingassignment/addrbookdb"
	pbmock "bashdrew/yumcodingassignment/mock_addrbookdb"

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

func TestGetPersonDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookDBlient := pbmock.NewMockAddrBookDBClient(ctrl)
	req := &pb.PersonRequestDB{Id: int64(0)}
	mockAddrBookDBlient.EXPECT().ReadPersonDB(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReplyDB{Firstname: "Andrew"}, nil)
	testReadPersonDB(t, mockAddrBookDBlient)
}

func testReadPersonDB(t *testing.T, client pb.AddrBookDBClient) {
	r, err := client.ReadPersonDB(context.Background(), &pb.PersonRequestDB{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Firstname != "Andrew" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestCreatePersonDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookDBlient := pbmock.NewMockAddrBookDBClient(ctrl)
	req := &pb.PersonRequestDB{Id: int64(0)}
	mockAddrBookDBlient.EXPECT().CreatePersonDB(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReplyDB{Lastname: "Amargo"}, nil)
	testCreatePersonDB(t, mockAddrBookDBlient)
}

func testCreatePersonDB(t *testing.T, client pb.AddrBookDBClient) {
	r, err := client.CreatePersonDB(context.Background(), &pb.PersonRequestDB{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Lastname != "Amargo" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestDeletePersonDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookDBlient := pbmock.NewMockAddrBookDBClient(ctrl)
	req := &pb.PersonRequestDB{Id: int64(0)}
	mockAddrBookDBlient.EXPECT().DeletePersonDB(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReplyDB{Email: "aamargo@gmail.com"}, nil)
	testDeletePersonDB(t, mockAddrBookDBlient)
}

func testDeletePersonDB(t *testing.T, client pb.AddrBookDBClient) {
	r, err := client.DeletePersonDB(context.Background(), &pb.PersonRequestDB{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Email != "aamargo@gmail.com" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}

func TestUpdatePersonDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAddrBookDBClient := pbmock.NewMockAddrBookDBClient(ctrl)
	req := &pb.PersonRequestDB{Id: int64(0)}
	mockAddrBookDBClient.EXPECT().UpdatePersonDB(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.PersonReplyDB{Phoneno: "123-123-1234"}, nil)
	testUpdatePersonDB(t, mockAddrBookDBClient)
}

func testUpdatePersonDB(t *testing.T, client pb.AddrBookDBClient) {
	r, err := client.UpdatePersonDB(context.Background(), &pb.PersonRequestDB{Id: int64(0)})
	fmt.Println(r)
	if err != nil || r.Phoneno != "123-123-1234" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Firstname)
}
