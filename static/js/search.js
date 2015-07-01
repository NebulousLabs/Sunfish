$("#query").keyup(function(event){
    if(event.keyCode == 13){
        searchSiafiles();
    };
});

var searchSiafiles = function(){
    var searchQuery = $('#query').val().toLowerCase();
    $.ajax({
        type: 'get',
        url: '/api/siafile/search/',
        data: 'tags=' + searchQuery,
        success: function(siafiles){
            makeTable(siafiles);
        },
    });
};
