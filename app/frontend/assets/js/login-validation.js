let token = sessionStorage.getItem('token');
if (token != null){
    //redirect to page
}
else{
    alert("Vui lòng đăng nhập để truy cập trang!");
    window.location.href = "/app/frontend/pages/dang-nhap.html";
}