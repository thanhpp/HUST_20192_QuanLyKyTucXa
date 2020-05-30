var role = sessionStorage.getItem("role");
var token = sessionStorage.getItem("token");

alert(token);
if (token == null) {
  alert("Vui lòng đăng nhập để truy cập trang này");
  window.location.href = "/app/frontend/pages/dang-nhap.html";
} else {
  var myHeaders = new Headers();
  //myHeaders.append("Content-Type", "text/plain", bearer);

  var requestOptions = {
    method: "GET",
    credentials: "omit",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "text/plain",
    },
    redirect: "follow",
  };

  fetch("http://52.15.50.37:9090/lv1/usrinfo", requestOptions)
    .then((response) => {
      console.log(response.status);
      if (response.status == 401) {
        alert("Vui lòng đăng nhập để truy cập trang này");
        window.location.href = "/app/frontend/pages/dang-nhap.html";
      } else {
        if (role == "1") {
          return;
        } else {
          alert("Bạn không có quyền truy cập trang này");
          window.stop();
        }
      }
    })
    .catch((error) => {
      alert("Không kết nối được tới máy chủ", error);
      window.stop();
    });
}
