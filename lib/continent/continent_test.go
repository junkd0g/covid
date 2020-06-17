package continent

import (
	"testing"

	mcontinent "github.com/junkd0g/covid/lib/model/continent"
)

type requestDataMock struct{}

var requestDataMockFunc func() (mcontinent.Response, error)

func (u requestDataMock) requestContinentData() (mcontinent.Response, error) {
	return requestDataMockFunc()
}

type requestCacheDataMock struct{}

var requestCacheDataMockFunc func() (mcontinent.Response, error)

func (u requestCacheDataMock) getCacheData() (mcontinent.Response, error) {
	return requestCacheDataMockFunc()
}

func TestRegisterUser(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}

	requestDataMockFunc = func() (mcontinent.Response, error) {
		return mcontinent.Response{{Cases: 64}, {Cases: 74}}, nil
	}

	requestCacheDataMockFunc = func() (mcontinent.Response, error) {
		return mcontinent.Response{{Cases: 32}, {Cases: 44}}, nil
	}

	withCashedData, err := GetContinentData()
	if err != nil {
		t.Fatal(err)
	}

	if withCashedData[0].Cases != 32 && withCashedData[1].Cases != 44 {
		t.Fatal("Not getting cached data when both request and cached exist")
	}

	requestCacheDataMockFunc = func() (mcontinent.Response, error) {
		return mcontinent.Response{}, nil
	}

	withNoCashedData, err2 := GetContinentData()
	if err2 != nil {
		t.Fatal(err2)
	}

	if withNoCashedData[0].Cases != 64 && withNoCashedData[1].Cases != 74 {
		t.Fatal("Not getting requested data when cached data are empty")
	}
}
