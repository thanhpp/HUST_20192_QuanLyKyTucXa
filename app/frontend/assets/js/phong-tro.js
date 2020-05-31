var token = localStorage.getItem("token");
var myHeaders = new Headers();
//myHeaders.append("Content-Type", "text/plain", bearer);

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
    $("#current-number").val(result.room.occupied);
    $("#max-number").val(result.room.roomMax);
    $("#due-date").val(result.room.duedDate);
    $("#fee-status").val(result.room.feeStatus);
  })
  .catch((error) => {
    console.log("Không kết nối được tới máy chủ", error);
    alert("Không kết nối được tới máy chủ");
  });

var friendsRequestOptions = {
  method: "GET",
  credentials: "omit",
  headers: {
    Authorization: "Bearer " + token,
    "Content-Type": "text/plain",
  },
  redirect: "follow",
};

let tabledata;

fetch("http://52.15.50.37:9090/lv0/friends", friendsRequestOptions)
  .then((response) => {
    response.json();
  })
  .then((result) => {
    tabledata = result;
    for (var i = 0; i < Object.keys(tabledata).length; i++) {
      tabledata[i].stt = i + 1;
    }
    handleTable();
  })
  .catch((error) => {
    console.log("Không kết nối được tới máy chủ", error);
  });

// let tabledata = [
//   {
//     name: "Tran Dinh Vu",
//     dob: "13/12/1999",
//     contact: "09191919191",
//     address: "Hà Nội",
//   },
//   {
//     name: "Tran Dinh A",
//     dob: "13/12/1999",
//     contact: "0382212381",
//     address: "Hạ Long",
//   },
//   {
//     name: "Tran Dinh B",
//     dob: "13/12/1999",
//     contact: "0382212381",
//     address: "Hưng Yên",
//   },
//   {
//     name: "Tran Dinh C",
//     dob: "13/12/1999",
//     contact: "09191919191",
//     address: "New York",
//   },
//   {
//     name: "Tran Dinh D",
//     dob: "13/12/1999",
//     contact: "09191919191",
//     address: "Bác Ninh",
//   },
  
// ];
// for (var i = 0; i < Object.keys(tabledata).length; i++) {
//   tabledata[i].stt = i + 1;
// }

// handleTable();

function handleTable() {
  let table = new Tabulator("#roomate-table", {
    height: "100%", // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
    data: tabledata, //assign data to table
    layout: "fitDataStretch", //fit columns to width of table (optional)
    columns: [
      //Define Table Columns
      {
        title: "STT",
        field: "stt",
        width: 70,
        hozAlign: "center",
        headerSort: false,
      },
      { title: "Họ và tên", field: "name", headerSort: false, width: 200 },
      { title: "Ngày sinh", field: "dob", headerSort: false },
      {
        title: "SĐT",
        field: "contact",
        headerSort: false,
      },
      { title: "Địa chỉ", field: "address", headerSort: false },
    ],
  });
}
