let token = sessionStorage.getItem("token");
let myHeaders = new Headers();

$("#submit-yeu-cau").on("click", function () {
  let title = $("#title_yeu_cau").val();
  let content = $("#content_yeu_cau").val();
  submitRequest(title, content);
});

function submitRequest(title, content){
  var requestOptions = {
    method: "POST",
    credentials: "omit",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "text/plain",
    },
    body: raw,
    redirect: "follow",
  };
  
  fetch("http://25.43.134.201:8080/lv0/sendreq", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      console.log(result);
      alert("Gửi yêu cầu thành công");
      window.location.reload();
    })
    .catch((error) => {
      console.log("Không kết nối được tới máy chủ", error);
    });
}


  function checkInput() {
    var title = $("#title_yeu_cau").val();
    var content = $("#content_yeu_cau").val();
    if (title != "" && content != "") {
      $("#submit-yeu-cau").removeClass("disabled");
    } else {
      $("#submit-yeu-cau").addClass("disabled");
    }
  }
  setInterval(checkInput, 300);