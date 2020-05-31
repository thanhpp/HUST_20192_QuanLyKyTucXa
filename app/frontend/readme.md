# Frontend

## Mục lục

* [Giới thiệu](#giới-thiệu)
* [Ngôn ngữ sử dụng](#ngôn-ngữ-sử-dụng)
* [Các thư viện được sử dụng](#các-thư-viện-được-sử-dụng)
* [Giao diện](#giao-diện)
  * [Đăng nhập](#đăng-nhập)
  * [Sinh viên](#sinh-viên)
    * [Trang chính](#trang-chính)
    * [Trang Thông báo](#trang-thông-báo)
    * [Trang Phòng trọ](#trang-Phòng-trọ)
    * [Trang Cơ sở vật chất](#trang-Cơ-sở-vật-chất)
    * [Trang Theo dõi tiền trọ](#trang-Theo-dõi-tiền-trọ)
    * [Trang Gửi yêu cầu](#trang-Gửi-yêu-cầu)
    * [Trang Yêu cầu đã gửi](#trang-Yêu-cầu-đã-gửi)
    * [Trang Thông tin cán vụ](#trang-Thông-tin-cán-vụ)
  * [Quản lý](#quản-lý)

## Giới thiệu 
- Đây là giao diện cho Project Trang web Ký túc xá cho trường Đại học. 
- Giao diện được làm theo template của [Stisla](https://github.com/stisla/stisla).
- Các trang được chia ra dành cho hai loại người dùng là Sinh viên và Quản lý ký túc xá.
- Người dùng phải đăng nhập để được truy cập vào các trang.

## Ngôn ngữ sử dụng
- HTML
- Javascript

## Các thư viện được sử dụng
- [jQuery](https://jquery.com/)
- [Bootstrap 4](https://getbootstrap.com/)
- [Font Awesome](https://fontawesome.com/)
- [Tabulator](http://tabulator.info/)

## Giao diện
### Giao diện chung
- Màn hình đăng nhập:

![Màn hình đăng nhập](img/dangnhap.png)

- Nút Đăng xuất:

![Đăng xuất](img/dangxuat.png)

  *Sau khi đăng xuất sẽ có hộp thoại thông báo đăng xuất thành công.*

- Thanh Navigation dọc

![Navbar dọc](img/navbar1.png)

  *Thanh Navigation dọc này có thể ấn nút Hamburger trên cùng để thu gọn lại.*
  ![Navbar dọc thu gọn](img/navbar2.png)

- Các trang đều được code responsive với nhiều kích cỡ màn hình, bao gồm trên điện thoại:

  *Giao diện trên màn hình to*

  ![Màn hình to](img/manhinhto.png)

  *Giao diện trên màn hình vừa*

  ![Màn hình vừa](img/manhinhvua.png)

  *Giao diện trên màn hình điện thoại*

  ![Màn hình điện thoại 1](img/manhinhdienthoai1.png) ![Màn hình điện thoại 2](img/manhinhdienthoai2.png)

### Sinh viên
#### Trang chính
- Giao diện:

  *Card Lịch sẽ cho người dùng xem lịch theo tháng và highlight Ngày hiện tại*

  *Chọn "Chi tiết..." sẽ chuyển sang trang [Phòng trọ](#trang-Phòng-trọ)*

![Màn hình trang chính](img/sinhvien_trangchinh.png)

#### Trang Thông báo
- Giao diện:

![Màn hình thông báo](img/sinhvien_thongbao1.png)

  *Chọn "Xem thêm" sẽ hiện toàn bộ thông báo.*

  ![Màn hình Xem thêm](img/sinhvien_thongbao2.png)

#### Trang Phòng trọ
- Giao diện

    *Người dùng có thẻ xem thông tin về phòng trọ và các sinh viên cùng phòng trọ*

 ![Màn hình Phòng trọ](img/sinhvien_phongtro.png)

#### Trang Theo dõi tiền trọ
- Giao diện:

![Màn hình Theo dõi tiền trọ](img/sinhvien_theodoitientro.png)
