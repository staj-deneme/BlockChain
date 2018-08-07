$( document ).ready(function() {
    $("#register-btn").on("click",function() {
        $(".loginbox").addClass("flipped");
        $(".registerbox").removeClass("flipped");
    });
    $("#login-btn").on("click",function() {
        $(".registerbox").addClass("flipped");
        $(".loginbox").removeClass("flipped");
    });
});
