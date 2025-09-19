package public_func

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

// 定义地区数据结构
type AreaData struct {
	AreaData []Province `json:"area_data"`
}

type Province struct {
	ProvinceCode string `json:"province_code"`
	Province     string `json:"province"`
	CityList     []City `json:"city_list"`
}

type City struct {
	CityCode     string     `json:"city_code"`
	City         string     `json:"city"`
	DistrictList []District `json:"district_list"`
}

type District struct {
	DistrictCode string `json:"district_code"`
	District     string `json:"district"`
}

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

/*生成测试手机号, 第一个bool为是否进行base64编码，第二个bool为是否白名单手机号*/
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

// 根据前17位计算校验码
func calculateCheckCode(id17 string) string {
	// 权重系数
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码对应关系
	checkCodes := "10X98765432"

	// 计算加权和
	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(id17[i]-'0') * weights[i]
	}

	// 取模并获取校验码
	mod := sum % 11
	return string(checkCodes[mod])
}

// 生成指定年龄的出生日期
func generateBirthDate(age int) string {
	// 获取当前年份
	now := time.Now()
	birthYear := now.Year() - age

	// 随机生成月份和日期
	birthMonth := rand.Intn(12) + 1
	birthDay := rand.Intn(28) + 1 // 简单处理，避免处理每月天数差异

	// 格式化为YYYYMMDD
	return fmt.Sprintf("%04d%02d%02d", birthYear, birthMonth, birthDay)
}

// 生成顺序码，根据性别确定奇偶性
func generateOrderCode(gender int) string {
	var orderCode int
	switch gender {
	case 1: // 男性
		// 男性，生成奇数
		orderCode = rand.Intn(500)*2 + 1
	case 2: // 女性
		// 女性，生成偶数
		orderCode = rand.Intn(500) * 2
	default: // 随机
		orderCode = rand.Intn(1000)
	}
	return fmt.Sprintf("%03d", orderCode)
}

// 查找area_code.json文件
func findAreaCodeFile() (string, error) {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 尝试几种可能的路径
	possiblePaths := []string{
		filepath.Join(wd, "baiju_go_script", "public_func", "data", "area_code.json"),
		filepath.Join(wd, "data", "area_code.json"),
		filepath.Join(wd, "..", "baiju_go_script", "public_func", "data", "area_code.json"),
		filepath.Join(wd, "..", "data", "area_code.json"),
		"baiju_go_script/public_func/data/area_code.json",
		"data/area_code.json",
		"area_code.json",
	}

	// 如果是Windows系统，也尝试反斜杠
	if runtime.GOOS == "windows" {
		possiblePaths = append(possiblePaths, filepath.Join(wd, "baiju_go_script", "public_func", "data", "area_code.json"))
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("无法找到area_code.json文件")
}

/*生成随机测试身份证*/
func GetCertificateNum(isBase64 bool) string {
	// 查找并加载地区数据
	areaDataPath, err := findAreaCodeFile()
	if err != nil {
		// 如果查找文件失败，返回默认值
		certificateNum := "130407202208114156"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	data, err := os.ReadFile(areaDataPath)
	if err != nil {
		// 如果读取文件失败，返回默认值
		certificateNum := "130407202208114151"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	var area AreaData
	err = json.Unmarshal(data, &area)
	if err != nil {
		// 如果解析JSON失败，返回默认值
		certificateNum := "130407202208114152"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	// 随机选择一个地区代码
	province := area.AreaData[rand.Intn(len(area.AreaData))]
	city := province.CityList[rand.Intn(len(province.CityList))]
	district := city.DistrictList[rand.Intn(len(city.DistrictList))]

	// 随机生成年龄（18-60岁）
	age := rand.Intn(43) + 18

	// 随机生成性别 (0表示随机)
	gender := rand.Intn(3) // 0, 1, 或 2

	// 构造身份证号码
	// 1-6位：地区代码
	areaCode := district.DistrictCode
	// 7-14位：出生日期
	birthDate := generateBirthDate(age)
	// 15-17位：顺序码
	orderCode := generateOrderCode(gender)

	// 前17位
	id17 := areaCode + birthDate + orderCode

	// 18位：校验码
	checkCode := calculateCheckCode(id17)

	// 完整身份证号码
	certificateNum := id17 + checkCode

	if isBase64 {
		return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
	}
	return certificateNum
}

// 根据指定性别和年龄生成身份证号码 (isBase64-是否base64加密,  gender 0-不限制 1-男 2-女, age 指定身份证年龄)
func GetCertificateNumByGenderAndAge(isBase64 bool, gender int, age int) string {
	// 查找并加载地区数据
	areaDataPath, err := findAreaCodeFile()
	if err != nil {
		// 如果查找文件失败，返回默认值
		certificateNum := "130407202208114156"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	data, err := os.ReadFile(areaDataPath)
	if err != nil {
		// 如果读取文件失败，返回默认值
		certificateNum := "130407202208114151"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	var area AreaData
	err = json.Unmarshal(data, &area)
	if err != nil {
		// 如果解析JSON失败，返回默认值
		certificateNum := "130407202208114152"
		if isBase64 {
			return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
		}
		return certificateNum
	}

	// 随机选择一个地区代码
	province := area.AreaData[rand.Intn(len(area.AreaData))]
	city := province.CityList[rand.Intn(len(province.CityList))]
	district := city.DistrictList[rand.Intn(len(city.DistrictList))]

	// 构造身份证号码
	// 1-6位：地区代码
	areaCode := district.DistrictCode
	// 7-14位：出生日期
	birthDate := generateBirthDate(age)
	// 15-17位：顺序码
	orderCode := generateOrderCode(gender)

	// 前17位
	id17 := areaCode + birthDate + orderCode

	// 18位：校验码
	checkCode := calculateCheckCode(id17)

	// 完整身份证号码
	certificateNum := id17 + checkCode

	if isBase64 {
		return string(base64.StdEncoding.EncodeToString([]byte(certificateNum)))
	}
	return certificateNum
}

// 随机生成省市区地址
type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}

func GetAddress() Address {

	// 查找并加载地区数据
	areaDataPath, _ := findAreaCodeFile()
	data, _ := os.ReadFile(areaDataPath)

	var area AreaData
	err := json.Unmarshal(data, &area)
	if err != nil {
		return Address{} // 返回空的Address结构体
	}

	// 随机选择一个地区代码
	province := area.AreaData[rand.Intn(len(area.AreaData))]
	city := province.CityList[rand.Intn(len(province.CityList))]
	district := city.DistrictList[rand.Intn(len(city.DistrictList))]

	return Address{
		Province: province.Province,
		City:     city.City,
		District: district.District,
	}
}
func PublicDataCase() {
	fmt.Println("This is a public data function.")
	// GetUUID()
	// address := GetAddress()
	// fmt.Println(address.Province)

}
