requirejs.config({
    "paths": {
        "text":"https://cdn.staticfile.org/require-text/2.0.12/text.min",
        "vue":"https://cdn.jsdelivr.net/npm/vue@2.5.21/dist/vue",
        "ELEMENT":"https://unpkg.com/element-ui/lib/index",
        "html-top":"/public/defaultui/block/top.html",
        "html-triggers":"/public/defaultui/block/triggers.html",
        "html-creategameform":"/public/defaultui/block/creategameform.html",
        "html-alllines":"/public/defaultui/block/alllines.html",
        "main":"/public/defaultui/js/main"
    },
});
define(function (require) {
    var htmltop=require("text!html-top");
    document.getElementById("top").innerHTML=htmltop;
    var htmltriggers=require("text!html-triggers");
    document.getElementById("triggers").innerHTML=htmltriggers;
    var htmlcreategameform=require("text!html-creategameform");
    document.getElementById("creategameform").innerHTML=htmlcreategameform;
    var htmlalllines=require("text!html-alllines");
    document.getElementById("alllines").innerHTML=htmlalllines;

    require(["main"],function(main){
        main()
    });
})