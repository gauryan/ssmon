function show_dialog(url) {
    $('#myModal .modal-content').load(url, function(e){
        $('#myModal').modal('show');
    });
}

function delete_admin(url) {
    var result = confirm("관리자를 정말로 삭제하시겠습니까?");
    if( result == false ) return;
    location.href = url;
}
