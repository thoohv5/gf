// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user/message.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_InfoReq_Code = regexp.MustCompile(`^[a-z0-9]{5,30}$`)

func (this *InfoReq) Validate() error {
	if !_regex_InfoReq_Code.MatchString(this.Code) {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9]{5,30}$"`, this.Code))
	}
	return nil
}
func (this *InfoResp) Validate() error {
	return nil
}
