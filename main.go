package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var input_filename string
var result_arry []string

func main() {
	data_arry, err := initialize()

	if err != nil {
		fmt.Println(err)
	}

	// filter the data_arry
	data_arry = data_arry[:len(data_arry)-1]
	//

	for i := 0; i < len(data_arry); i++ {
		formatted_range_str, _ := translate_ip_and_mask_to_range(data_arry[i])
		fmt.Println(formatted_range_str)
	}
}

func translate_ip_and_mask_to_range(value string) (string, error) {
	var ipconfig IpConfig
	// Split the string on the slash
	ip_and_mask_str, _ := split_on_slash(value)

	// if there is no more than one element
	// in the array I should throw something

	ip_int64, _ := convert_tradIp_to_int64(ip_and_mask_str[0])
	mask_int64, _ := convert_tradIp_to_int64(ip_and_mask_str[1])

	ipconfig.default_ip = ip_int64
	ipconfig.subnet_mask = mask_int64

	mask_full := int64(4294967295)

	ipconfig.subnet_range = int64(mask_full ^ mask_int64)
	ipconfig.broadcast_ip = int64(ipconfig.default_ip | ipconfig.subnet_range)
	ipconfig.range_start = ipconfig.default_ip + 1
	ipconfig.range_end = ipconfig.broadcast_ip - 1

	range_start, _ := convert_int64_to_tradIp(ipconfig.range_start)
	range_end, _ := convert_int64_to_tradIp(ipconfig.range_end)

	var sb strings.Builder
	sb.WriteString(range_start)
	sb.WriteString("-")
	sb.WriteString(range_end)
	sb.WriteString(";")
	formatted_range_str := sb.String()

	return formatted_range_str, nil
}

func convert_int64_to_tradIp(value int64) (string, error) {

	ip_seg_int64, _ := split_int64_to_fourbytes(value)

	var sb strings.Builder
	for i := 0; i < len(ip_seg_int64); i++ {
		sb.WriteString(strconv.FormatInt(ip_seg_int64[i], 10))
		if i != len(ip_seg_int64)-1 {
			sb.WriteString(".")
		}
	}

	ipaddress := sb.String()
	return ipaddress, nil
}

func convert_tradIp_to_int64(value string) (int64, error) {
	ip_segments_str, _ := split_on_dot(value)

	// if there are not 4 segments
	// I should throw something

	for i := 0; i < 4; i++ {
		ip_segments_str[i], _ = convert_strNum_to_strBinary(ip_segments_str[i])
		if len(ip_segments_str[i]) < 8 {
			delta := 8 - len(ip_segments_str[i])
			var sb strings.Builder
			for i := 0; i < delta; i++ {
				sb.WriteString("0")
			}
			sb.WriteString(ip_segments_str[i])
			ip_segments_str[i] = sb.String()
		}
	}

	var sb strings.Builder
	sb.WriteString("0b")
	for i := 0; i < len(ip_segments_str); i++ {
		sb.WriteString(ip_segments_str[i])
	}

	ip_binary_str := sb.String()

	ip_int, _ := strconv.ParseInt(ip_binary_str, 0, 64)

	return ip_int, nil
}

func split_int64_to_fourbytes(value int64) ([]int64, error) {
	seg1_mask := int64(4278190080)
	seg2_mask := int64(16711680)
	seg3_mask := int64(65280)
	seg4_mask := int64(255)

	var ip_seg_int64 []int64

	ip_seg_int64 = append(ip_seg_int64, (value&seg1_mask)>>24)
	ip_seg_int64 = append(ip_seg_int64, (value&seg2_mask)>>16)
	ip_seg_int64 = append(ip_seg_int64, (value&seg3_mask)>>8)
	ip_seg_int64 = append(ip_seg_int64, value&seg4_mask)

	return ip_seg_int64, nil
}

func split_on_slash(value string) ([]string, error) {
	result := strings.Split(value, "/")
	if len(result) == 1 {
		result = strings.Split(value, "\\")
	}
	return result, nil
}

func split_on_dot(value string) ([]string, error) {
	result := strings.Split(value, ".")
	if len(result) == 1 {

	}
	return result, nil
}

type IpConfig struct {
	default_ip   int64
	broadcast_ip int64
	subnet_mask  int64
	subnet_range int64
	range_start  int64
	range_end    int64
}

func convert_strNum_to_strBinary(value string) (string, error) {
	int_num, _ := strconv.Atoi(value)
	int64_num := int64(int_num)
	str_bin := strconv.FormatInt(int64_num, 2)
	return str_bin, nil
}

func filter(value string) (string, error) {

	return "", nil
}

func initialize() ([]string, error) {
	if len(os.Args) < 2 {
		fmt.Println("Less than two arguments, exiting")

		return nil, errors.New("Correct Usage: go run [program].go [filename]")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	data_str := string(data)
	data_str_arry := strings.Split(data_str, "\n")
	for i := 0; i < len(data_str_arry); i++ {
		fmt.Println(data_str_arry[i])
	}
	return data_str_arry, err
}
