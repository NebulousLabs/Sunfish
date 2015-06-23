var fileDownload = function(siafile){
    var $a = $("<a>", {
        href: 'data:attachment/sia,' + siafile.content,
        target: '_blank',
        download: siafile.filename,
        text: siafile.filename
    });

    $("#" + siafile.title).append($a);
};

$.get("/api/siafile/", function(data){
    for (var i = 0; i < data.length; i++){
        var siafile = data[i];
        var uploadDate = new Date(siafile.uploadedTime);
        $('.siafile-table').append("<tr><td>" + siafile.title + "</td><td>" + siafile.description + "</td><td>" + siafile.tags + "</td><td>" + uploadDate.toLocaleString() + "<td id='" + siafile.title + "'></td></tr>");
        fileDownload(siafile);

    }
})
