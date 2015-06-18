var readFile = function(){
    var preview = document.querySelector('img');
    var file    = document.querySelector('input[type=file]').files[0];
    var reader  = new FileReader();

    reader.onloadend = function () {
        console.log(reader.result);
        return reader.result;
    }

    if (file) {
        reader.readAsDataURL(file);
    }
}

var uploadSiafile = function() {
    // $.post()
    var formData = new FormData();

    var formData = $('form').serializeArray().reduce(function(obj, item) {
        obj[item.name] = item.value;
        return obj;
    }, {});
    formData.siafilepath = '';
    formData.hash = '';
    formData.tags = formData.tags.split(",");
    // TODO convert to callback and send base64 file in json
    formData.siafile = readFile();

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
}
