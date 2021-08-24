# serde_rust_go
in rs/
```sh
cargo run
msg bytes: [123, 34, 116, 111, 107, 101, 110, 95, 116, 121, 112, 101, 34, 58, 34, 69, 84, 72, 34, 44, 34, 112, 114, 105, 99, 101, 34, 58, 49, 48, 48, 48, 48, 44, 34, 116, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 48, 125]
msg in hex: "7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d"
```
in go_serde/
```sh
go run main.go
msg bytes: [123 34 84 111 107 101 110 84 121 112 101 34 58 34 69 84 72 34 44 34 80 114 105 99 101 34 58 49 48 48 48 48 44 34 84 105 109 101 83 116 97 109 112 34 58 48 125]
msg in hex:  0x7b22546f6b656e54797065223a22455448222c225072696365223a31303030302c2254696d655374616d70223a307d
```

## 对比


```sh
[123 34 116 111 107 101 110 95 116 121 112 101 34 58 34 69 84 72 34 44 34 112 114 105 99 101 34 58 49 48 48 48 48 44 34 116 105 109 101 115 116 97 109 112 34 58 48 125]        rust
[123 34 116 111 107 101 110 95 116 121 112 101 34 58 34 69 84 72 34 44 34 112 114 105 99 101 34 58 49 48 48 48 48 44 34 116 105 109 101 95 115 116 97 109 112 34 58 48 125]     go

7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d
7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d655f7374616d70223a307d
```

## 最新版本
```sh
#rust
msg bytes: [123, 34, 116, 111, 107, 101, 110, 95, 116, 121, 112, 101, 34, 58, 34, 69, 84, 72, 34, 44, 34, 112, 114, 105, 99, 101, 34, 58, 49, 48, 48, 48, 48, 44, 34, 116, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 48, 125]
msg in hex: "7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d"

#go
msg bytes: [123 34 116 111 107 101 110 95 116 121 112 101 34 58 34 69 84 72 34 44 34 112 114 105 99 101 34 58 49 48 48 48 48 44 34 116 105 109 101 115 116 97 109 112 34 58 48 125]
msg in hex: 0x7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d


#compare
7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d
7b22746f6b656e5f74797065223a22455448222c227072696365223a31303030302c2274696d657374616d70223a307d
```

## 问题
字段的命名不统一导致的