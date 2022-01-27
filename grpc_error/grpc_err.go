package main

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// sample server to returning error and using status
func server() error {

	st := status.New(codes.InvalidArgument, "One or more fields are invalid")

	f1 := &errdetails.BadRequest_FieldViolation{
		Field:       "Email",
		Description: "Invalid email format",
	}

	f2 := &errdetails.BadRequest_FieldViolation{
		Field:       "Password",
		Description: "Must be at least 10 characters",
	}

	f3 := &errdetails.PreconditionFailure_Violation{
		Type:        "USER",
		Subject:     "no users created",
		Description: "No users have been created",
	}
	st, _ = st.WithDetails(f1)
	st, _ = st.WithDetails(f2)
	st, _ = st.WithDetails(f3)

	return st.Err()
}

// This should be done in UI/UX on client side
func client(err error) {
	st, _ := status.FromError(err)
	p := st.Proto()
	data, _ := protojson.Marshal(p)
	fmt.Println(string(data))
}

func main() {

	err := server()
	client(err)

}
