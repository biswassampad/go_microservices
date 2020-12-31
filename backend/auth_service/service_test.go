package main

import (
	"context"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/biswassampads/blog_applications/global"
	"github.com/biswassampads/blog_applications/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("example"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{ID: primitive.NewObjectID(), Email: "test@gmail.com", Username: "Carl", Password: string(pw)})
	server := authServer{}
	_, err := server.Login(context.Background(), &proto.LoginRequest{Login: "test@gmail.com", Password: "example"})
	if err != nil {
		t.Error("1: An error was returned: ", err.Error())
	}
	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "something", Password: "someting"})
	if err == nil {
		t.Error("2: Error is null")
	}

	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "Carl", Password: "example"})
	if err != nil {
		t.Error("3: An error was occured: ", err.Error())
	}
}
