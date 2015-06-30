var readFile = function(callback){
    var preview = document.querySelector('img');
    var file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();

    reader.onloadend = function () {
        var encodedFile = reader.result;
        var filename = file.name;

        callback(file.name, encodedFile);
    }

    if (file) {
        reader.readAsText(file);
    }
}

var uploadSiafile = function() {
    // Processes the form and prepares the siafile for upload via json
    var formData = new FormData();

    var formData = $('form').serializeArray().reduce(function(obj, item) {
        obj[item.name] = item.value;
        return obj;
    }, {});

    formData.tags = formData.tags.split(",");

    for (var i = 0; i < formData.tags.length; i++){
        formData.tags[i] = formData.tags[i].trim().toLowerCase();
    };

    formData.siafile = readFile(function(name, fileData){
        formData.filename = name;
        formData.fileData = fileData;
        $.ajax({
            url: '/api/siafile/',
            type: 'POST',
            data: JSON.stringify(formData),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                console.log(msg);
            }
        });
    });
}

var fileDownload = function(siafile){
    var $a = $("<a>", {
        href: 'data:attachment/sia,' + siafile.content,
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
                "<td>" + siafile.tags + "</td>" +
                "<td>" + uploadDate.toLocaleString() + "" +
                "<td id='" + siafile.Id + "'></td></tr>");
        fileDownload(siafile);
    }
};


