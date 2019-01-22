# struct2map

golang 的 struct 转为 map，可以只转结构体的部份字段，常用于 API Server 输出个部份字段给前端，示例代码见 test 文件

## 说明：

```
Convert(s, fields, tagName)
s 要转换的结构体
fields 要转换的字段名（以标签名为准）
tagName 标签名
```

```
ConvertSlice(s, fields, tagName)
s 要转换的结构体 Slice
fields 要转换的字段名（以标签名为准）
tagName 标签名
```
