var token = localStorage.getItem("token");
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
  .then((response) => {
    response.json();
  })
  .then((result) => {
    $("#student-name").val(result.student_info.name);
    $("#student-id").val(result.student_info.studentid);
    $("#student-birth").val(result.student_info.dob);
    $("#student-address").val(result.student_info.address);
    $("#student-contact").val(result.student_info.contact);
    $("#student-room").val(result.student_info.room);
  })
  .catch((error) => {
    console.log("Không kết nối được tới máy chủ", error);
    alert("Không kết nối được tới máy chủ");
  });

  var roomRequestOptions = {
    method: "GET",
    credentials: "omit",
    headers: {
      Authorization: "Bearer " + token,
      "Content-Type": "text/plain",
    },
    redirect: "follow",
  };
  
  fetch("http://52.15.50.37:9090/lv0/roominfo", roomRequestOptions)
    .then((response) => {
      response.json();
    })
    .then((result) => {
      $("#student-room").val(result.room.roomID);
      $("#room-fee").val(result.room.roomPrice);
      $("#due-date").val(result.room.duedDate);
      $("#fee-status").val(result.room.feeStatus);
  
    })
    .catch((error) => {
      console.log("Không kết nối được tới máy chủ", error);
      alert("Không kết nối được tới máy chủ");
    });
