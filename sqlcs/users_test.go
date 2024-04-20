package sqlcs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"sqlcs.sqlc.dev/app/util"
)

func CreateRandomUser(t *testing.T) User{
	hashedpass,_ := util.Hashedpassword(util.RandomString(5))
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		Hashedpassword:hashedpass,
		FullName:util.RandomString(5)+" "+ util.RandomString(5),
		Email:util.RandomString(5)+"@gmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	return user
}

func TestCreateUser(t *testing.T){
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	_,err := testQueries.GetUser(context.Background(),user1.Username) 
	require.NoError(t,err)
}