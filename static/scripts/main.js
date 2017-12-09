(function ($) {
    "use strict";

    $(document).ready(function(){



        screenSizeChecker();

        $(window).resize(screenSizeChecker);

    });

    function screenSizeChecker() {
        var articles = $(".card");
        console.log(($(window).width() + 15));

        var width = ($(window).width() + 15);

        $(articles).each(function (index) {
            $(this).removeClass("featured");
            $(this).css("cssText", "margin-right:");
            if (width < 1200) {
                if (width > 768) {
                    if (index == 2 || index == 7 || index == 10) {
                        $(this).addClass("featured");
                    }

                    if (index == 4 || index == 6 || index == 9) {
                        $(this).css("cssText", "margin-right: 0px !important;");
                    }

                    if (index == 3 || index == 5 || index == 8) {
                        $(this).css("cssText", "margin-right: 22px !important;");
                    }
                }
            } else {
                if (index == 3 || index == 10) {
                    $(this).addClass("featured");
                }

                if (index == 3 || index == 6 || index == 9) {
                    $(this).css("cssText", "margin-right: 0px !important;");
                }

                if (index == 5 || index == 8) {
                    $(this).css("cssText", "margin-right: 22px !important;");
                }
            }
        });
    }

}(jQuery));
