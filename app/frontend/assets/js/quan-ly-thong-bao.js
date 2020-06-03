$( document ).ready(function() {

  function checkInput() {
    var header_area = $("#header").val();
    var content_area = $("#noi-dung-thong-bao").val();
    if (header_area != "" && content_area != "") {
      $("#submit-button").removeClass("disabled");
    } else {
      $("#submit-button").addClass("disabled");
    }
  }
  setInterval(checkInput, 300);


  var studentRequestList = [
    {
      id:"1",
      subject: "Tình hình điều hoà bị hỏng",
      name: "Trần Đình Vũ",
      student_id: "20176126",
      room: "201",
      sent_date: "13/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đã xử lý",
    },
    {
      id:"2",
      subject: "Vòi nước có vấn đề ở phòng vệ sinh tầng 2",
      name: "Lê Hồng Quang",
      student_id: "20176001",
      room: "201",
      sent_date: "15/02/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đã xử lý",
    },
    {
      id:"3",
      subject: "Báo cáo cái ghế bị hỏng",
      name: "Mai Thị Mây",
      student_id: "20176126",
      room: "301",
      sent_date: "01/04/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Chưa xử lý",
    },
    {
      id:"4",
      subject: "Gãy cầu thang ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "101",
      sent_date: "13/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đang xử lý",
    },
    {
      id:"5",
      subject: "Gãy cầu thang ở tầng 4",
      name: "Nguyễn Trần Đình",
      student_id: "2017600",
      room: "101",
      sent_date: "13/04/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Chưa xử lý",
    },
    {
      id:"6",
      subject: "Gãy cầu thang ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "502",
      sent_date: "20/05/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Chưa xử lý",
    },
    {
      id:"7",
      subject: "Gãy cầu thang ở tầng 6",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "201",
      sent_date: "03/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đã xử lý",
    },
    {
      id:"8",
      subject: "Gãy cầu thang ở tầng 7",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "503",
      sent_date: "02/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đang xử lý",
    },
    {
      id:"9",
      subject: "Nứt tường ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "201",
      sent_date: "02/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Chưa xử lý",
    },
    {
      id:"10",
      subject: "Gãy cầu thang ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "202",
      sent_date: "16/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đã xử lý",
    },
    {
      id:"11",
      subject: "Gãy cầu thang ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "202",
      sent_date: "13/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Chưa xử lý",
    },
    {
      id:"12",
      subject: "Gãy cầu thang ở tầng 5",
      name: "Nguyễn Trần Đình",
      student_id: "20176126",
      room: "203",
      sent_date: "13/01/2020",
      content:"Oh noooooooooooooooooooooooooooooooooo",
      status: "Đang xử lý",
    },
  ];
var studentsTable = new Tabulator("#student-parking-table", {
  height: 400, // set height of table (in CSS or here), this enables the Virtual DOM and improves render speed dramatically (can be any valid css height value)
  data: studentRequestList, //assign data to table
  virtualDom: true,
  layout: "fitColumns", //fit columns to width of table (optional)
  initialSort:[
    {column:"sent_date", dir:"desc"}, //sort by this first
],
  columns: [
    //Define Table Columns
    { title: "Mã", field: "id",width:"80", hozAlign: "center", sorter: "string",},
    { title: "Tiêu đề",field: "subject",sorter: "string" },
    { title: "Nội dung",field: "content",sorter: "string",visible: false },
    { title: "Họ và tên", field: "name", sorter: "string", visible: false },
    { title: "MSSV", width:"130",field: "student_id", sorter: "number" },
    { title: "Phòng", width:"100", field: "room", sorter: "string" },
    {
      title: "Ngày gửi",
      field: "sent_date",
      width:"200",
      sorter: "date",
      sorterParams: {
        format: "DD-MM-YYYY",
        alignEmptyValues: "top",
      },
    },
    { title: "Tình trạng", width:"150",field: "status", sorter: "string" },
  ],
  rowClick: function (e, row) {
    var id = row.getData().id;
    var subject = row.getData().subject;
    var content = row.getData().content;
    var student_id = row.getData().student_id;
    var room = row.getData().room;
    var sent_date = row.getData().sent_date;
    var status = row.getData().status;
    var name = row.getData().name;
    $("#student-request-modal").modal("show");
    $("#subject").text(subject);
    $('#id').text(id);
    $("#sent-date").text(sent_date);
    $('#student_name').text(name);
    $('#student-id').text(student_id);
    $('#room').text(room);
    $('#content').text(content);
    $('#status').val(status);
  },
});

var fieldEl = document.getElementById("filter-field");
var typeEl = document.getElementById("filter-type");
var valueEl = document.getElementById("filter-value");

//Custom filter example
function customFilter(data) {
  return data.car && data.rating < 3;
}

//Trigger setFilter function with correct parameters
function updateFilter() {
  var filterVal = fieldEl.options[fieldEl.selectedIndex].value;
  var typeVal = typeEl.options[typeEl.selectedIndex].value;

  var filter = filterVal == "function" ? customFilter : filterVal;

  if (filterVal == "function") {
    typeEl.disabled = true;
    valueEl.disabled = true;
  } else {
    typeEl.disabled = false;
    valueEl.disabled = false;
  }

  if (filterVal) {
    studentsTable.setFilter(filter, typeVal, valueEl.value);
  }
}


//Update filters on value change
document
  .getElementById("filter-field")
  .addEventListener("change", updateFilter);
document.getElementById("filter-type").addEventListener("change", updateFilter);
document.getElementById("filter-value").addEventListener("keyup", updateFilter);

$("#add-student").on("click", function () {
  $("#add-student-table-modal").modal("show");
});

$("#change-student-button").click(function () {
  //gui request;
});

$("#add-student-button").click(function () {
  //gui-request;
});
})
