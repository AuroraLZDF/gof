package utils

import (
	"strings"
	"fmt"
	"encoding/json"
	"strconv"
	"net/smtp"
	"time"
	"reflect"
	"errors"
	"path/filepath"
	"os"
	"log"
	"crypto/md5"
	"encoding/hex"
	"github.com/auroraLZDF/goconf"
)

/**
 * 获取执行可执行文件的目录
 */
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * 判断 文件/文件夹 是否存在
 */
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

/**
 * 区分目录和文件
 */
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

/**
 * md5 加密
 */
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	result := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果

	return result
}

/**
 * 获取变量类型
 */
func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

/**
 * map to json
 */
func MapToJson(obj map[string]interface{}) string {
	jsonBytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return ""
	}

	return string(jsonBytes)
}

/**
 * json to map
 */
func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println("json.Unmarshal failed:", err)
	}

	return mapResult
}

/**
 * struct to json
 */
func StructToJson(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}

	return string(jsonBytes)
}

/**
 * int to string
 */
func IntToString(v int) string {
	return strconv.Itoa(v)
}

/**
 * string to int
 */
func StringToInt(str string) int {
	_int, _ := strconv.Atoi(str)
	return _int
}

/**
 * float64 to int
 */
func FloatToInt(v float64) int {
	return int(v)
}

/**
 * struct to map
 */
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/**
 * 发送邮件
 */
func SendMail(to, subject, body, mailType string, conf *goconf.Config) error {
	user := conf.GetValue("mail", "username")
	password := conf.GetValue("mail", "password")
	host := conf.GetValue("mail", "hostname")
	port := conf.GetValue("mail", "port")

	auth := smtp.PlainAuth("", user, password, host)

	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/html" + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host+":"+port, auth, user, sendTo, msg)

	return err
}

/**
 * 去除空格
 */
func TrimS(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

/**
 * 声明 errors 类型
 */
func Error(str string) error {
	err := errors.New(str)
	return err
}

/**
 * 格式化日期
 */
func Date(format string) string {
	return time.Now().Format(format)
}
