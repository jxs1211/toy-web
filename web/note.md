所有int数据类型都有已出的

不要为难自己，看不懂使用英文翻译deepL

## string的总结

len计算的是字节数，不是字符数,计算字符数使用utf8.RuneCountInString

如果字符串中出现了非ASCII码，就用uft8库计算长度

![image-20220627085809605](C:\Users\xjshen\AppData\Roaming\Typora\typora-user-images\image-20220627085809605.png)

不支持修改

切片

不要使用在子切片中修改元素，会影响原值

切片不支持随机填加、删除，需要自己实现

append操作可能触发扩容，会生成新切片，并将原数组的元素全部拷贝到新数组

fmt

保留2为小数

检查序列化和反序列化是否正确：将[]byte输出为16进制

```go
func TestAny(t *testing.T) {
	// c := utf8.RuneCountInString("沈先捷")
	// fmt.Println(c)
	str := "沈先捷"
	b := []byte(str)
	fmt.Printf(" => bytes(hex): [% x]\n", b)
	b2, _ := json.Marshal(str)
	fmt.Printf("Marshal string => bytes(hex): [% x]\n", b2)
	var str2 interface{}
	_ = json.Unmarshal(b2, &str2)
	fmt.Printf("Unmarshal bytes => string bytes(hex): [% x]\n", str2)
	fmt.Printf(" => : %s\n", str2)
	fmt.Printf("is equal: %v\n", str == str2.(string)[1:]) // unmarhshal process is ok
	fmt.Printf("is equal: %v\n", str == str2.(string)[1:]) // something during unmarshal process
}
```

