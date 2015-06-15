function uploadSiafile() {
    // $.post()
  var formData = new FormData();
  console.log($('#siafile').file);
  console.log(formData.get('siafile'));

    var formData = $('form').serializeArray().reduce(function(obj, item) {
            obj[item.name] = item.value;
                return obj;
    }, {});
    console.log(formData);
}
