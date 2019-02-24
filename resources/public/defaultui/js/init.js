requirejs.config({
    "paths": {
        "text":"https://cdn.staticfile.org/require-text/2.0.12/text.min",
        "vue":"https://cdn.jsdelivr.net/npm/vue@2.5.21/dist/vue",
        "ELEMENT":"https://unpkg.com/element-ui/lib/index",
        "html-top":"/public/defaultui/block/top.html",
        "main":"/public/defaultui/js/main"
    },
});
define(function (require) {
    var htmltop=require("text!html-top");
    document.getElementById("top").innerHTML=htmltop;

    require(["main"],function(main){
        main()
    });
})