package user

import (
	"errors"
	"testing"

	"github.com/Go_Testing/Golang_Testing/mockgen/mocks"
	"github.com/golang/mock/gomock"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	//Expect Do to be called once with 123 and "Hello GoMock"
	//as parameters, and return nil from the mocked call
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	//Expect Do to be called once with 123 and "Hello GoMock"
	//as params, and return dummyError from the mocked call
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)

	err := testUser.Use()

	if err != nil {
		t.Fail()
	}

}
