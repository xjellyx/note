package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " AA bb CC dd FF 7哈哈 "
	fmt.Println("前缀", hasPrefix(str, "AAb"))
	fmt.Println("后缀", hasSuffix(str, "dFd"))
	fmt.Println("字符串包含关系", contains(str, "bbC"))
	fmt.Println("字符串交集关系", containsAny(str, "BbC"))
	fmt.Println("索引", index(str, "哈"))
	fmt.Println("索引", lastIndex(str, "d"))
	fmt.Println("字符串替换", replace(str, "bbCC", "lloo", -1))
	fmt.Println("统计字符串", count(str, "A"))
	fmt.Println("重复输出字符", repeat(str, 3))
	fmt.Println("大写转小写", toLower(str))
	fmt.Println("小写转大写", toUpper(str))
	fmt.Println("删除开头和结尾的空白符", trimSpace(str))
	fmt.Println("删除指定字符", trim(str, "cut"))
	fmt.Println("利用空白符作为分隔符, 返回slice", fields(str))
	fmt.Println("分割", split(str, " "))

	arr := []string{"a", "b", "v", "d"}
	fmt.Println("拼接slice到字符串", join(arr, ""))
	// 其他方法与这些分类相似，原理一样
}

// hasPrefix 前缀
func hasPrefix(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// hasSuffix  后缀
func hasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// contains 字符串包含关系
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// containsAny 字符串交集关系
func containsAny(str, chars string) bool {
	return strings.ContainsAny(str, chars)
}

// index 索引, 获取字符出现的第一个位置值,-1表示字符串s不包含字符串str
func index(str, s string) int {
	return strings.Index(str, s)
}

// 获取某字符出现最后的那个位置
func lastIndex(str, s string) int {
	return strings.LastIndex(str, s)
}

// replace 字符串替换,将str中前n个字符串old替换为new，并返回新字符串，n=-1，替换所有
func replace(str, old, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

// count Count用于计算字符串str在字符串s中出现的非重叠次数
func count(str, s string) int {
	return strings.Count(str, s)
}

// repeat Repeat用于重复count次字符串s，并返回一个新的字符串
func repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// toLower 转小写
func toLower(s string) string {
	return strings.ToLower(s)
}

// toUpper 转大写
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// trimSpace 删除开头和结尾的空白符
func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

// trim 剔除指定字符
func trim(str, s string) string {
	return strings.Trim(str, s)
}

// fields 利用空白符作为分隔符, 返回slice
func fields(s string) []string {
	return strings.Fields(s)
}

// split 分割str
func split(s, sep string) []string {
	return strings.Split(s, sep)
}

// join 拼接slice到字符串 Join用来将元素类型为string的slice使用分割符号拼接起来
func join(arr []string, sep string) string {
	return strings.Join(arr, sep)
}
