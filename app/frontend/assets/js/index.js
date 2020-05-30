let token = localStorage.getItem("token");
console.log(token);
let user_name;
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

fetch("http://52.15.50.37:9090/lv0/usrinfo", requestOptions)
  .then((response) => response.json())
  .then((result) => {
    console.log(result);
  })
  .catch((error) => console.log("Không kết nối được tới máy chủ", error));
