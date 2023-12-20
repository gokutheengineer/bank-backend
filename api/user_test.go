package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mockdb "github.com/gokutheengineer/bank-backend/db/mock"
	db "github.com/gokutheengineer/bank-backend/db/sqlc"
	util "github.com/gokutheengineer/bank-backend/util"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/require"
)

type eqUserMatcher struct {
	arg      db.User
	password string
}

func (e eqUserMatcher) Matches(x interface{}) bool {
	// Check if inputArgs assignable to db.CreateUserParams
	inputArgs, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	// Check password
	if err := util.VerifyPasswordBcrypt(e.password, e.arg.PasswordHashed); err != nil {
		return false
	}

	// check rest of the args fullname and username
	return (e.arg.Fullname == inputArgs.Fullname) && (e.arg.Username == inputArgs.Username)
}

func (e eqUserMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.arg)
}

func EqUser(arg db.User, password string) gomock.Matcher {
	return eqUserMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username,
				"password": password,
				"fullname": user.Fullname,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), EqUser(user, password)). //specific matcher
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireResponseBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "Duplicate Username",
			body: gin.H{
				"username": user.Username,
				"password": password,
				"fullname": user.Fullname,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pgconn.PgError{Code: "23505"}) // https://www.postgresql.org/docs/11/errcodes-appendix.html
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "Invalid Username",
			body: gin.H{
				"username": "aaa",
				"password": password,
				"fullname": user.Fullname,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0). //test won't run due to invalid username
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Password",
			body: gin.H{
				"username": user.Username,
				"password": password[:2], //invalid short password
				"fullname": user.Fullname,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0). //test won't run due to invalid password
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Error",
			body: gin.H{
				"username": user.Username,
				"password": password,
				"fullname": user.Fullname,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, fmt.Errorf("some error"))
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func requireResponseBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var _user db.User
	err = json.Unmarshal(data, &_user)
	require.NoError(t, err)
	require.Equal(t, user.Username, _user.Username)
	require.Equal(t, user.Fullname, _user.Fullname)
	require.Empty(t, _user.PasswordHashed, "password hashed should not be returned")

}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomPassword()
	hashedPassword, err := util.HashPasswordBcrypt(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		PasswordHashed: hashedPassword,
		Fullname:       util.RandomFullname(),
	}

	return
}
