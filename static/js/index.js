$.get("/api/siafile/", function(data){
    for (var i = 0; i < data.length; i++){
        var siafile = data[i];
        var uploadDate = new Date(siafile.uploadedTime);
        $('.siafile-table').append("<tr><td>" + siafile.title + "</td><td>" + siafile.description + "</td><td>" + siafile.tags + "</td><td>" + uploadDate.toLocaleString() + "</tr>");
    }
})
