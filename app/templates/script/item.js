
$(document).ready(function () {
    $('#product-img').click(function () {
        $('#popup-img').attr('src', $(this).attr('src'));
        $('#image-popup').fadeIn();
    });

    // × ボタンまたは背景をクリックしたらポップアップを閉じる
    $('#close-popup, #image-popup').click(function () {
        $('#image-popup').fadeOut();
    });

    // ポップアップ内の画像をクリックしたらポップアップを閉じないようにする
    $('#popup-img').click(function (event) {
        event.stopPropagation();
    });
});
