var readFile = function(callback){
    var preview = document.querySelector('img');
    var file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();

    reader.onloadend = function () {
        var encodedFile = reader.result;
        var filename = file.name;
        if(filename.indexOf(".sia") != -1){
            callback(file.name, encodedFile);
        } else {
            alert("Error: Not a Siafile!")
        }
    }

    if (file) {
        reader.readAsDataURL(file);
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
              window.location.href = "/";
            }
        });
    });
}
