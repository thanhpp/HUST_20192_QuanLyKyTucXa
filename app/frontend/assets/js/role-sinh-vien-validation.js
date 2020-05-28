let role = localStorage.getItem('role');

if(role === ""){

}else{
  alert("Bạn không có quyền truy cập trang này!");
  window.stop();
}