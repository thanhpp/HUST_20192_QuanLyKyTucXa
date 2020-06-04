# Backend

## MỤC LỤC

### [I. Thiết kế hệ thống](#thiết-kế-hế-thống)

### [II. Các thư viện ngoài sử dụng](#các-thư-viện-ngoài-sử-dụng)

### [III. SQL Tables](#sql-tables)

### [IV. Danh sách các API đang cung cấp](#danh-sách-các-api-đang-cung-cấp)

#### [POSTMAN](#POSTMAN)

### [V. Thiết kế các API](#thiết-kế-các-api)

#### [1. Level 0](#1-level-0)

#### [2. Level 1](#2-level-1)

## Thiết kế hệ thống

![BackEndDesign](img/DormAppBackendDesign.png)

## Các thư viện ngoài sử dụng

1.[Gin-Gonic](https://github.com/gin-gonic):

- Một trong các framework nổi tiếng của Go.

- Có tốc độ cao và nhiều middlewares hỗ trợ.

2.[gORM](https://gorm.io/)

- Thư viện hỗ trợ ORM(Object-relational mapping cho Go)

3.[Viper](https://github.com/spf13/viper)

- Thư viện hỗ trợ đọc và cài đặt cấc biến môi trường cho Go.

4.[Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt?tab=doc)

- Thư viện dùng cho việc mã hóa, hash trong Go.

5.[Jwt-go](https://github.com/dgrijalva/jwt-go)

- Thư viện dùng cho việc tạo và xác thực JWT trong Go.

## [SQL TABLES](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/blob/master/app/backend/SQL.md)

## Danh sách các API đang cung cấp

### [POSTMAN](https://www.getpostman.com/collections/8894497461d3adc2ec1e)

## Thiết kế các API:
- Chia người sử dụng thành 3 dạng:
    - user (Khi chưa đăng nhập)
    - level 1 (người quản lý)
    - level 0 (sinh viên)

- Thông tin về phân quyền, xác thực được lưu trong token trả về và được kiểm tra mỗi khi người dùng gọi request. 

- Refresh token được dùng để lấy thêm access token mới.
### 1. Level 0

- MSSV được đính kèm vào jwt để định danh sinh viên khi thực hiện request.
- Sinh viên có thể lấy các thông tin liên quan đến mình.
    - Lấy thông tin bản thân. (Dự kiến kết nối với hệ thống khác)
    - Lấy thông tin về phòng trọ:
        - Danh sách bạn cùng phòng
        - Lịch sử đóng tiền trọ
        - Danh sách cơ sở vật chất có trong phòng.
    - Lấy thông tin về thông báo chung.
- Vấn đề gửi yêu cầu:
    - Sinh viên có thể lấy toàn bộ danh sách yêu cầu mình đã tạo
    - Sinh viên có thể tạo yêu cầu mới(tuy nhiên không được trùng về trạng thái và title)
- Sửa đổi thông tin cá nhân.

### 2. Level 1

- Người quản lý có khả năng lấy các thông tin:
    - Sinh viên
    - Phòng trọ
    - CSVC của các phòng
    - Lấy toàn bộ các thông báo chung.

- Người quản lý có quyền thêm và sửa đổi CSDL:
    - Đổi phòng trọ cho sinh viên
    - Tính tiền trọ cho tháng mới của toàn bộ sinh viên.
    - Trả lời yêu cầu của sinh viên.
    - Thay đổi trạng thái trả tiền trọ của sinh viên.
    - Tạo CSVC mới.
    - Thêm CSVC cho các phòng.
    - Chỉnh sửa số lượng CSVC.