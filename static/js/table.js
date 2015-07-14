$.get("/api/siafile/", function(siafiles){
    makeTable(siafiles);
});

var fileDownload = function(siafile){
    var $a = $("<a>", {
        href: siafile.fileData,
        target: '_blank',
        download: siafile.filename,
        text: siafile.filename
    });

    $("#" + siafile.Id).append($a);
};

var makeTable = function(siafiles){
    $('.siafile-table').empty();
    for (var i = 0; i < siafiles.length; i++){
        var siafile = siafiles[i];
        var uploadDate = new Date(siafile.uploadedTime);
        $('.siafile-table').append(
                "<tr>" + 
                "<td>" + siafile.title + "</td>" + 
                "<td>" + siafile.description + "</td>" +
                "<td>" + siafile.tags.join(', ') + "</td>" +
                "<td>" + uploadDate.toLocaleString() + "</td>" +
                "<td id='" + siafile.Id + "'></td></tr>"
                );
        fileDownload(siafile);
    }
};
