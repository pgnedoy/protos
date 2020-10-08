# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [hello-proto/v2/hello_api.proto](#hello-proto/v2/hello_api.proto)
    - [CreateUserRequest](#helloproto.v2.CreateUserRequest)
    - [CreateUserResponse](#helloproto.v2.CreateUserResponse)
    - [HelloRequest](#helloproto.v2.HelloRequest)
    - [HelloResponse](#helloproto.v2.HelloResponse)
  
    - [HelloAPI](#helloproto.v2.HelloAPI)
  
- [hello-proto/v2/types.proto](#hello-proto/v2/types.proto)
    - [User](#helloproto.v2.User)
  
    - [AuthType](#helloproto.v2.AuthType)
    - [DeletionReason](#helloproto.v2.DeletionReason)
    - [GenderType](#helloproto.v2.GenderType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hello-proto/v2/hello_api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hello-proto/v2/hello_api.proto



<a name="helloproto.v2.CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth_id | [string](#string) |  |  |
| auth_type | [AuthType](#helloproto.v2.AuthType) |  |  |
| name | [string](#string) |  |  |
| age | [int32](#int32) |  |  |
| country | [string](#string) |  |  |






<a name="helloproto.v2.CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#helloproto.v2.User) |  |  |






<a name="helloproto.v2.HelloRequest"></a>

### HelloRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="helloproto.v2.HelloResponse"></a>

### HelloResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| greeting | [string](#string) |  |  |
| greet_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 

 

 


<a name="helloproto.v2.HelloAPI"></a>

### HelloAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#helloproto.v2.CreateUserRequest) | [CreateUserResponse](#helloproto.v2.CreateUserResponse) |  |
| Hello | [HelloRequest](#helloproto.v2.HelloRequest) | [HelloResponse](#helloproto.v2.HelloResponse) |  |

 



<a name="hello-proto/v2/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hello-proto/v2/types.proto



<a name="helloproto.v2.User"></a>

### User
This is a leading comment for a message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the uuid of user |
| name | [string](#string) |  |  |
| age | [int32](#int32) |  |  |
| wink_id | [string](#string) |  | unique value for each user |
| country | [string](#string) |  |  |
| gender | [GenderType](#helloproto.v2.GenderType) |  |  |
| about | [string](#string) |  |  |
| auth_type | [AuthType](#helloproto.v2.AuthType) |  |  |
| birthday | [string](#string) |  | birthday of this user |
| deletion_reason | [DeletionReason](#helloproto.v2.DeletionReason) |  |  |
| create_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| update_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| delete_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| snap_name | [string](#string) |  | snapchat username |





 


<a name="helloproto.v2.AuthType"></a>

### AuthType


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTH_TYPE_INVALID | 0 |  |
| AUTH_TYPE_PHONE | 1 |  |
| AUTH_TYPE_SNAP | 2 |  |



<a name="helloproto.v2.DeletionReason"></a>

### DeletionReason


| Name | Number | Description |
| ---- | ------ | ----------- |
| DELETION_REASON_INVALID | 0 |  |
| DELETION_REASON_SELF_DELETION | 1 |  |
| DELETION_REASON_BOT_BEHAVIOUR | 2 |  |
| DELETION_REASON_NUDITY | 3 |  |
| DELETION_REASON_SEXUAL_ACTIVITY | 4 |  |
| DELETION_REASON_ADULT_TOYS | 5 |  |
| DELETION_REASON_UNDERWEAR | 6 |  |
| DELETION_REASON_PHYSICAL_VIOLENCE | 7 |  |
| DELETION_REASON_WEAPON_VIOLENCE | 8 |  |
| DELETION_REASON_WEAPONS | 9 |  |
| DELETION_REASON_SELF_INJURY | 10 |  |



<a name="helloproto.v2.GenderType"></a>

### GenderType


| Name | Number | Description |
| ---- | ------ | ----------- |
| GENDER_TYPE_INVALID | 0 |  |
| GENDER_TYPE_MALE | 1 |  |
| GENDER_TYPE_FEMALE | 2 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

