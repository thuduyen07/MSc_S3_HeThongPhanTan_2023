# Distributed Systems

This is the repo of Distributed Systems course

## Lecture 01
[How to Write Go Code](https://go.dev/doc/code#:~:text=Go%20programs%20are%20organized%20into,files%20within%20the%20same%20package.)

## Lab01

## Lecture 02

## Lab02
[rpc](https://pkg.go.dev/net/rpc)
[lab02 submit link](https://docs.google.com/forms/d/e/1FAIpQLSfCGAMuYYLqXx6jTMc3Fmx5dyTD5aJIQPshe8XjIMaf3YvZlw/viewform)
## Lecture 03

## Lab03

## Lecture 04

Các server phân biệt với nhau bằng IP (và port)

Trong key-value store, thêm phần replication vào. Mỗi lần user set giá trị vào bảng key-value khi put dữ liệu lên, ta cần copy dữ liệu qua các server còn lại (giả sử là 5 server)

Các test cases:
- TC_01: Put/get
- TC_02: Put/get when primary is dead
- TC_03: put and get in slow network 

hàm ping để kiểm tra primary server còn sống

tất cả các server tự gửi cho nhau bằng hàm check để lấy một cặp ip-version, chờ vài chục giây -> chọn version cao nhất và ip thấp nhất làm primary mới nếu primảy hiện tại chớt

các máy phải lưu primary_ip, có biến isPrimary


Deadline: 16h00 18/08/2023

Link: https://docs.google.com/forms/d/e/1FAIpQLSdc2hOqxAeM4oY5JvTxe8kG6YGj67zF27OcAZXZI236IFFkew/viewform

dành riêng một IP để client connect đến, nếu primary chớt thì thằng mới vẫn listen thằng IP đó, và một lúc chỉ có một ông listen để tránh conflict

consistency: tính nhất quán. Ví dụ: 
- set x=1 thì get x phải bằng 1 ==> read-your-write consistency 
- eventual consistency
- client set x=1 --get-x--> set x=3 --get-x  ==> sequentual consistency

xử lí tình huống mạng chậm -> set x=1 đến sau set x=3
- gửi thêm timeStamp 
- sử dụng NTP (một loại time server) để triển khai đồng hồ dùng chung
- [Lamport Clock](https://en.wikipedia.org/wiki/Lamport_timestamp)

