package public_func

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

/* 生成UUID v4 (随机)*/
func GetUUID() string {
	id := uuid.New()
	return id.String()
}

/* 生成IP (随机)*/
func GetRandomIPv4InChina() string {
	// 中国常用IP段
	chinaIPSets := []struct {
		baseIP string
		mask   int
	}{
		{"223.160.0.0", 16}, // 中国移动
		{"117.160.0.0", 16}, // 中国联通
		{"1.180.0.0", 16},   // 中国电信
		{"114.114.0.0", 16}, // 中国电信
	}

	// 随机选择一个IP段
	selected := chinaIPSets[rand.Intn(len(chinaIPSets))]

	// 生成该网段内的随机IP
	ip := net.ParseIP(selected.baseIP).To4()
	for i := 0; i < 4; i++ {
		if (selected.mask / 8) <= i {
			ip[i] = byte(rand.Intn(256))
		}
	}

	return ip.String()
}

/*生成测试姓名*/
func GetName(isBase64 bool) string {
	name := faker.ChineseName()
	if isBase64 {
		return string(base64.StdEncoding.EncodeToString([]byte(name)))
	}
	return name
}

/*生成测试手机号*/
func GetPhone(isBase64, isWhiteList bool) string {
	var phone string
	// 中国手机号前缀
	if isWhiteList {
		phone = "15555555555"
	} else {
		prefixes := []string{
			"130", "131", "132", "133", "134", "135", "136", "137", "138", "139",
			"145", "147", "149",
			"150", "151", "152", "153", "155", "156", "157", "158", "159",
			"166",
			"170", "171", "173", "175", "176", "177", "178",
			"180", "181", "182", "183", "184", "185", "186", "187", "188", "189",
			"191", "198", "199",
		}
		// 随机选择前缀
		prefix := prefixes[rand.Intn(len(prefixes))]
		// 生成8位随机数字
		suffix := fmt.Sprintf("%08d", rand.Intn(100000000))
		phone = prefix + suffix
	}

	if isBase64 {
		return string(base64.StdEncoding.EncodeToString([]byte(phone)))
	}
	return phone
}

/*生成测试身份证*/
func GetCertificateNum(isBase64 bool) string {
	certificateNum := "130407202208114156"
	if isBase64 {
		return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
	}
	return certificateNum
}
func PublicDataCase() {
	fmt.Println("This is a public data function.")
	// GetUUID()
}
