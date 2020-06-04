# Quản lý ký túc xá

## Thông tin đề tài 

**Học phần** : Các phương pháp phát triển phần mềm nhanh.

**Tên đề tài** : Làm website quản lý ký túc xá.

**Thành viên** : 
- Trần Đình Vũ
- Phan Phú Thành

**Ngôn ngữ sử dụng** :
- [Golang](https://golang.org/)
- Javascript.

**Tài liệu khác** :
- [Google Drive folder](https://drive.google.com/open?id=1mdFRZLyr1xLBjlqofsPvLP6bdbBcQS_B)

  -  [Sprint Planning](https://docs.google.com/spreadsheets/d/1EoV6SjYCSxJXXKl3eWgFJAMr5WD1BUkhHbn7Jj3Wa5A/edit?usp=sharing)
  -  [Q&A](https://docs.google.com/spreadsheets/d/1qhD5vAsDTaZFDzS6ogCaKI8Yk6_AQ2E665HbKgXV0JI/edit?usp=sharing)

## [Tài liệu về phát triển phần mềm](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu)

### 1. [Xác định yêu cầu phần mềm](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/XacDinhYeuCauPhanMem)

### 2. [Phân tích và thiết kế](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/PhanTich%26ThietKe)

#### 2.1. [Biểu đồ trình tự](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/PhanTich%26ThietKe/Bieu%20do%20trinh%20tu)

#### 2.2. [Biểu đồ hoạt động](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/PhanTich%26ThietKe/Bieu%20do%20trinh%20tu)

#### 2.3. [Biểu đồ lớp](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/PhanTich%26ThietKe/Bieu%20do%20lop)

#### 2.4. [Biểu đồ E-R](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/TaiLieu/PhanTich%26ThietKe/Bieu%20do%20ER)

## [Chương trình](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/app)

### [1. Backend](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/app/backend)

### [2. Frontend](https://github.com/ThanhPP/HUST_20192_QuanLyKyTucXa/tree/master/app/frontend)

## Triển khai hệ thống

### 1. Backend:
- Chương trình chính
  - Cách 1: Chạy file binary(Thích hợp với linux/64 bit)
  - Cách 2: Cài đặt Go compiler và chạy lệnh
    ```bash
    # Lưu ý: để biến môi trường GOMODULE111=on
    go build -v
    ``` 
    Sau đó chạy chương trình đã được biên dịch.

- Cài đặt các CSDL liên quan:
  - MySQL
  - Redis

- Thay đổi các config trong folder config
  - Các tham số để khởi tạo CSDL

### 2. Frontend:
- Sử dụng phần mềm VSCode và Extension Live Server
  - Tải và cài đặt phần mềm VSCode: https://code.visualstudio.com/
  - Tìm và cài đặt extension Live Server: https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer
  - Sau khi cài đặt extension, mở file HTML bạn muốn rồi ấn vào nút Go Live ở góc trái dưới màn hình
### 3. Hamachi:
- Đăng ký tài khoản tại: https://www.vpn.net/
- Tải và cài đặt Hamachi trên máy tính tại : https://www.vpn.net/
- Đăng nhập tài khoản Hamachi.
- Tạo 1 mạng riêng ảo(dạng: Mesh) - Gói free cho 5 máy tính 1 mạng.
- Kết nối tới mạng đó và lấy Virtual IP của máy.
- Chạy chương trình backend trên Virtual IP và cổng tự lựa chọn.
