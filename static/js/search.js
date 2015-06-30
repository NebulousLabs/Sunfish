var searchSiafiles = function(){
    var searchQuery = $('#query').val();
    console.log(searchQuery);
    $.ajax({
        type: 'get',
        url: '/api/siafile/search/',
        data: 'tags=' + searchQuery,
        success: function(siafiles){
            makeTable(siafiles);
        },
    });
};
