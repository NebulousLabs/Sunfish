var dataUrlEncode = function(str, callback) {
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    callback(str); 

}

var asciiEncode = function(str, callback) {
    str = str.replace(/\+/g, '-').replace(/\//g, '_');
    callback(str);
}

var readFile = function(callback){
    var preview = document.querySelector('img');
    var file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();

    reader.onloadend = function () {
        var filename = file.name;
        if(filename.indexOf(".sia") != -1){
            asciiEncode(btoa(reader.result), function(ascii) {
                callback(file.name, ascii);
            });
        } else {
            alert("Error: Not a Siafile!")
        }
    }

    if (file) {
        reader.readAsBinaryString(file);
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

    formData.siafile = readFile(function(name, ascii){
        formData.filename = name;
        formData.ascii = ascii;
        $.ajax({
            url: '/api/siafile/',
            type: 'POST',
            data: JSON.stringify(formData),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                window.location.href = "/";
            }
        });
    });
}

var fileDownload = function(siafile){
    dataUrlEncode(siafile.ascii, function(dataUrl) {
        var $a = $("<a>", {
            href: "data:application/octet-stream;base64," + dataUrl,
            target: '_blank',
            download: siafile.filename,
            text: siafile.filename
        });

        $("#" + siafile.Id).append($a);
    });

    var $a = $("<a>", {
        onclick: "window.prompt('Copy Ascii','" + siafile.ascii + "')",
        text: "Copy ASCII"
    });

    $("#ascii-" + siafile.Id).append($a);
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
                "<td id='" + siafile.Id + "'></td>" +
                "<td id='ascii-" + siafile.Id + "'></td></tr>"
                );
        fileDownload(siafile);
    }
};
