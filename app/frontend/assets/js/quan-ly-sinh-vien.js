var studentsList = [
  {id:1, name:"Trần Đình Vũ", student_id: "20176126", room:"201 B6", fee_status:"Đã thanh toán"},
  {id:2, name:"Trần Đình A", student_id: "20176111", room:"201 B6", fee_status:"Đã thanh toán"},
  {id:3, name:"Trần Đình B", student_id: "20176112", room:"201 B6", fee_status:"Chưa thanh toán"},
  {id:4, name:"Trần Đình C", student_id: "20176113", room:"203 B6", fee_status:"Đã thanh toán"},
  {id:5, name:"Trần Đình D", student_id: "20176133", room:"203 B6", fee_status:"Đã thanh toán"},
  {id:6, name:"Trần Đình E", student_id: "20176144", room:"202 B6", fee_status:"Đã thanh toán"},
  {id:7, name:"Nguyễn Văn A", student_id: "20176134", room:"202 B6", fee_status:"Chưa thanh toán"},
  {id:8, name:"Nguyễn Văn B", student_id: "20176151", room:"201 B6", fee_status:"Đã thanh toán"},
  {id:9, name:"Nguyễn Văn C", student_id: "20176216", room:"202 B6", fee_status:"Chưa thanh toán"},
  {id:10, name:"Nguyễn Văn D", student_id: "20176335", room:"203 B6", fee_status:"Chưa thanh toán"},
  {id:11, name:"Nguyễn Văn E", student_id: "20176116", room:"204 B6", fee_status:"Đã thanh toán"},
  {id:12, name:"Nguyễn Văn F", student_id: "20176546", room:"204 B6", fee_status:"Đã thanh toán"},
  {id:13, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:14, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:15, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:16, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:17, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:18, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:19, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:20, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:21, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:22, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  {id:23, name:"Nguyễn Văn G", student_id: "20176098", room:"204 B6", fee_status:"Chưa thanh toán"},
  
];

var studentsAppendingList = [
  {id:1, name:"Trần Đình Vũ", student_id: "20176126", register_date:"13/12/1999",},
  {id:2, name:"Trần Đình A", student_id: "20176111", register_date:"13/12/1999",},
  {id:3, name:"Trần Đình B", student_id: "20176112", register_date:"13/12/1999",},
  {id:4, name:"Trần Đình C", student_id: "20176113", register_date:"13/12/1999",},
  {id:5, name:"Trần Đình D", student_id: "20176133", register_date:"13/12/1999",},
  {id:6, name:"Trần Đình E", student_id: "20176144", register_date:"13/12/1999",},
  {id:7, name:"Nguyễn Văn A", student_id: "20176134", register_date:"13/12/1999",},
  {id:8, name:"Nguyễn Văn B", student_id: "20176151", register_date:"13/12/1999",},
  {id:9, name:"Nguyễn Văn C", student_id: "20176216", register_date:"13/12/1999",},
  {id:10, name:"Nguyễn Văn D", student_id: "20176335", register_date:"13/12/1999",},
  {id:11, name:"Nguyễn Văn E", student_id: "20176116", register_date:"13/12/1999",},
  {id:12, name:"Nguyễn Văn F", student_id: "20176546", register_date:"13/12/1999",},
  {id:13, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:14, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:15, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:16, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:17, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:18, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:19, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:20, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:21, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:22, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  {id:23, name:"Nguyễn Văn G", student_id: "20176098", register_date:"13/12/1999",},
  
];

//create Tabulator on DOM element with id "example-table"
var studentsTable = new Tabulator("#student-table", {
  height:400 , // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
  data:studentsList,//assign data to table
  virtualDom: true, 
  layout:"fitColumns", //fit columns to width of table (optional)
  columns:[ //Define Table Columns
    {title:"STT", field:"id", width:80, hozAlign:"center", sorter:"number"},
    {title:"Họ và tên", field:"name", sorter:"string"},
    {title:"MSSV", field:"student_id", sorter:"number"},
    {title:"Phòng", field:"room", sorter:"string"},
    {title:"Tình trạng nộp phí", field:"fee_status", sorter:"string"},
    
  ],
  rowClick:function(e, row){ //trigger an alert message when the row is clicked
    alert("Row " + row.getData().id + " Clicked!!!!");
  },
});

var studentsPendingTable = new Tabulator("#student-appending-table", {
  height:400 , // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
  data:studentsAppendingList,//assign data to table
  virtualDom: true, 
  layout:"fitColumns", //fit columns to width of table (optional)
  columns:[ //Define Table Columns
    {title:"STT", field:"id", width:80, hozAlign:"center", sorter:"number"},
    {title:"Họ và tên", field:"name", sorter:"string"},
    {title:"MSSV", field:"student_id", sorter:"number"},
    {title:"Ngày nộp đơn", field:"register_date", sorter:"date"},
  ],
  rowClick:function(e, row){ //trigger an alert message when the row is clicked
    alert("Row " + row.getData().id + " Clicked!!!!");
  },
});

//Define variables for input elements
var fieldEl = document.getElementById("filter-field");
var typeEl = document.getElementById("filter-type");
var valueEl = document.getElementById("filter-value");

//Custom filter example
function customFilter(data){
    return data.car && data.rating < 3;
}

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
    studentsTable.setFilter(filter,typeVal, valueEl.value);
  }
}

//Update filters on value change
document.getElementById("filter-field").addEventListener("change", updateFilter);
document.getElementById("filter-type").addEventListener("change", updateFilter);
document.getElementById("filter-value").addEventListener("keyup", updateFilter);

document.getElementById("download-xlsx").addEventListener("click", function(){
  studentsTable.download("xlsx", "danhsachsinhvienKTX.xlsx", {sheetName:"Danh sách"});
});

var fieldE2 = document.getElementById("filter-field2");
var typeE2 = document.getElementById("filter-type2");
var valueE2 = document.getElementById("filter-value2");

//Custom filter example
function customFilter(data){
    return data.car && data.rating < 3;
}

//Trigger setFilter function with correct parameters
function updateFilter(){
  var filterVal = fieldE2.options[fieldE2.selectedIndex].value;
  var typeVal = typeE2.options[typeE2.selectedIndex].value;

  var filter = filterVal == "function" ? customFilter : filterVal;

  if(filterVal == "function" ){
    typeE2.disabled = true;
    valueE2.disabled = true;
  }else{
    typeE2.disabled = false;
    valueE2.disabled = false;
  }

  if(filterVal){
    studentsPendingTable.setFilter(filter,typeVal, valueE2.value);
  }
}

//Update filters on value change
document.getElementById("filter-field2").addEventListener("change", updateFilter);
document.getElementById("filter-type2").addEventListener("change", updateFilter);
document.getElementById("filter-value2").addEventListener("keyup", updateFilter);

document.getElementById("xuat-excel2").addEventListener("click", function(){
  studentsPendingTable.download("xlsx", "danhsachsinhvienyeucauKTX.xlsx", {sheetName:"Danh sách"});
});
