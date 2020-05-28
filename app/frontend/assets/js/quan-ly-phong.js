var roomsList = [
  {
    id: 1,
    room: "101",
    limit: "5",
    empty_bed: "3",
    price: "200000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 2,
    room: "102",
    limit: "5",
    empty_bed: "0",
    price: "300000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 3,
    room: "103",
    limit: "8",
    empty_bed: "6",
    price: "600000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 4,
    room: "104",
    limit: "3",
    empty_bed: "3",
    price: "200000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 5,
    room: "105",
    limit: "5",
    empty_bed: "1",
    price: "100000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 6,
    room: "106",
    limit: "10",
    empty_bed: "5",
    price: "200000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 7,
    room: "107",
    limit: "8",
    empty_bed: "6",
    price: "500000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 8,
    room: "108",
    limit: "5",
    empty_bed: "0",
    price: "400000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
  {
    id: 9,
    room: "109",
    limit: "6",
    empty_bed: "0",
    price: "200000",
    register_date: "1/1/2019",
    update_date: "1/1/2019",
  },
];

var roomsTable = new Tabulator("#room-table", {
  height: 400, // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
  data: roomsList, //assign data to table
  virtualDom: true,
  layout: "fitColumns", //fit columns to width of table (optional)
  columns: [
    //Define Table Columns
    {
      title: "STT",
      field: "id",
      width: 80,
      hozAlign: "center",
      sorter: "number",
    },
    {
      title: "Phòng",
      field: "room",
      width: 120,
      hozAlign: "center",
      sorter: "string",
      hozAlign: "center",
    },
    {
      title: "Giá",
      field: "price",
      hozAlign: "center",
      sorter: "number",
    },
    {
      title: "Giới hạn",
      field: "limit",
      hozAlign: "center",
      sorter: "number",
    },
    {
      title: "Trống",
      field: "empty_bed",
      hozAlign: "center",
      sorter: "number",
    },
    {
      title: "Ngày cập nhật",
      field: "update_date",
      hozAlign: "center",
      sorter: "string",
      visible: false,
    },
    {
      title: "Ngày đăng kí",
      field: "register_date",
      sorter: "string",
      visible: false,
    },
  ],
  rowClick: function (e, row) {
    var id = row.getData().id;
    var room = row.getData().room;
    var limit = row.getData().limit;
    var room = row.getData().room;
    var empty_bed = row.getData().empty_bed;
    var price = row.getData().price;
    var update_date = row.getData().update_date;
    var register_date = row.getData().register_date;
    $("#room-table-modal").modal("show");
    $('input[id="room"]').val(room);
    $('input[id="limit"]').val(limit);
    $('input[id="empty_bed"]').val(empty_bed);
    $('input[id="price"]').val(price);
    $("#update_date").text(update_date);
    $("#register_date").text(register_date);
  },
});

//Define variables for input elements
var fieldEl = document.getElementById("filter_field");
var typeEl = document.getElementById("filter_type");
var valueEl = document.getElementById("filter_value");

//Trigger setFilter function with correct parameters
function updateFilter(){
  var filterVal = fieldEl.options[fieldEl.selectedIndex].value;
  var typeVal = typeEl.options[typeEl.selectedIndex].value;

  var filter = filterVal == "function" ? customFilter : filterVal;

  if(filterVal == "function" ){
    typeEl.disabled = true;
    valueEl.disabled = true;
  }else{
    typeEl.disabled = false;
    valueEl.disabled = false;
  }

  if(filterVal){
    roomsTable.setFilter(filter,typeVal, valueEl.value);
  }
}

//Update filters on value change
document.getElementById("filter_field").addEventListener("change", updateFilter);
document.getElementById("filter_type").addEventListener("change", updateFilter);
document.getElementById("filter_value").addEventListener("keyup", updateFilter);

document.getElementById("download-xlsx").addEventListener("click", function () {
  roomsTable.download("xlsx", "danhsachsinhvienKTX.xlsx", {
    sheetName: "Danh sách",
  });
});

///////////////////////////////////////////////
var studentsAppendingList = [
  {
    id: 1,
    name: "Trần Đình Vũ",
    student_id: "20176126",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 2,
    name: "Trần Đình A",
    student_id: "20176111",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 3,
    name: "Trần Đình B",
    student_id: "20176112",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 4,
    name: "Trần Đình C",
    student_id: "20176113",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 5,
    name: "Trần Đình D",
    student_id: "20176133",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 6,
    name: "Trần Đình E",
    student_id: "20176144",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 7,
    name: "Nguyễn Văn A",
    student_id: "20176134",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 8,
    name: "Nguyễn Văn B",
    student_id: "20176151",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 9,
    name: "Nguyễn Văn C",
    student_id: "20176216",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 10,
    name: "Nguyễn Văn D",
    student_id: "20176335",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 11,
    name: "Nguyễn Văn E",
    student_id: "20176116",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 12,
    name: "Nguyễn Văn F",
    student_id: "20176546",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 13,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 14,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 15,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 16,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 17,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 18,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 19,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 20,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 21,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 22,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
  {
    id: 23,
    name: "Nguyễn Văn G",
    student_id: "20176098",
    register_date: "13/12/1999",
    update_date: "1/1/2020",
    address: "Hà Nội",
    dob: "15/10/1999",
    contact: "0382212381",
  },
];

var studentsPendingTable = new Tabulator("#student-appending-table", {
  height: 400, // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
  data: studentsAppendingList, //assign data to table
  virtualDom: true,
  layout: "fitColumns", //fit columns to width of table (optional)
  columns: [
    //Define Table Columns
    {
      title: "STT",
      field: "id",
      width: 80,
      hozAlign: "center",
      sorter: "number",
    },
    { title: "Họ và tên", field: "name", sorter: "string" },
    { title: "MSSV", field: "student_id", sorter: "number" },
    { title: "Ngày nộp đơn", field: "register_date", sorter: "date" },
    {
      title: "Ngày cập nhật",
      field: "updated_date",
      sorter: "number",
      visible: false,
    },
    { title: "Địa chỉ", field: "address", sorter: "string", visible: false },
    { title: "Ngày sinh", field: "dob", sorter: "date", visible: false },
    {
      title: "Thông tin liên lạc",
      field: "contact",
      sorter: "string",
      visible: false,
    },
  ],
  rowClick: function (e, row) {
    //trigger an alert message when the row is clicked
    alert(row.getData().dob + " Clicked!!!!");
  },
});

var fieldE2 = document.getElementById("filter_field2");
var typeE2 = document.getElementById("filter_type2");
var valueE2 = document.getElementById("filter_value2");

//Trigger setFilter function with correct parameters
function updateFilter2() {
  var filterVal = fieldE2.options[fieldE2.selectedIndex].value;
  var typeVal = typeE2.options[typeE2.selectedIndex].value;

  var filter = filterVal == "function" ? customFilter : filterVal;

  if (filterVal == "function") {
    typeE2.disabled = true;
    valueE2.disabled = true;
  } else {
    typeE2.disabled = false;
    valueE2.disabled = false;
  }

  if (filterVal) {
    studentsPendingTable.setFilter(filter, typeVal, valueE2.value);
  }
}

//Update filters on value change
document.getElementById("filter_field2").addEventListener("change", updateFilter2);
document.getElementById("filter_type2").addEventListener("change", updateFilter2);
document.getElementById("filter_value2").addEventListener("keyup", updateFilter2);

document.getElementById("xuat-excel2").addEventListener("click", function () {
  studentsPendingTable.download("xlsx", "danhsachsinhvienyeucauKTX.xlsx", {
    sheetName: "Danh sách",
  });
});
/////////BUTTON/////////
$("#save-form").on("click", function () {
  //post update info
  alert("Cập nhật thành công");
  $("#room-table-modal").modal("hide");
});
