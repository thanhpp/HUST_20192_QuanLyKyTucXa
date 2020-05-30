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

  var tabledata = [
    {id:1, name:"Oli Bob", age:"12", col:"red", dob:""},
    {id:2, name:"Mary May", age:"1", col:"blue", dob:"14/05/1982"},
    {id:3, name:"Christine Lobowski", age:"42", col:"green", dob:"22/05/1982"},
    {id:4, name:"Brendon Philips", age:"125", col:"orange", dob:"01/08/1980"},
    {id:5, name:"Margret Marmajuke", age:"16", col:"yellow", dob:"31/01/1999"},
  ];

  var table = new Tabulator("#fee-history-table", {
    height:205, // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
    data:tabledata, //assign data to table
    layout:"fitColumns", //fit columns to width of table (optional)
    columns:[ //Define Table Columns
      {title:"Name", field:"name", width:150},
      {title:"Age", field:"age", hozAlign:"left", formatter:"progress"},
      {title:"Favourite Color", field:"col"},
      {title:"Date Of Birth", field:"dob", sorter:"date", hozAlign:"center"},
    ],
    rowClick:function(e, row){ //trigger an alert message when the row is clicked
      alert("Row " + row.getData().id + " Clicked!!!!");
    },
 });