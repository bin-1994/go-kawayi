package utils

// 加密key
// U2FsdGVkX18Wv2TVLmvrnKZ3mGRb0XdlDxprEoKQEOU=
const auth_key = "U2FsdGVkX18Wv2TVLmvrnKZ3mGRb0XdlDxprEoKQEOU="

//func UserMd5 (value string) string {
//	encrypt := Sha1(value) + auth_key
//	return Md5(encrypt)
//}
//
//func Sha1 (value string) string {
//	sha1 := sha1.New()
//	sha1.Write([]byte(value))
//	return hex.EncodeToString(sha1.Sum([]byte("")))
//}
//
//func Md5(value string) string {
//	md5 := md5.New()
//	md5.Write([]byte(value))
//	return hex.EncodeToString(md5.Sum([]byte("")))
//}
