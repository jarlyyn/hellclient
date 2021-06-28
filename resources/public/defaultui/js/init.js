requirejs.config({
    "paths": {
        "text":"/public/defaultui/js/text.min",
        "vue":"/public/defaultui/js/vue",
        "ELEMENT":"/public/defaultui/theme-chalk/index",
        "html-top":"/public/defaultui/block/top.html",
        "html-triggers":"/public/defaultui/block/triggers.html",
        "html-creategameform":"/public/defaultui/block/creategameform.html",
        "html-createscriptform":"/public/defaultui/block/createscriptform.html",
        "html-createtimerform":"/public/defaultui/block/createtimerform.html",
        "html-updatetimerform":"/public/defaultui/block/updatetimerform.html",
        "html-alllines":"/public/defaultui/block/alllines.html",
        "html-notopened":"/public/defaultui/block/notopened.html",
        "html-scriptlist":"/public/defaultui/block/scriptlist.html",
        "html-timerlist":"/public/defaultui/block/timerlist.html",
        "html-script":"/public/defaultui/block/script.html",
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
    var htmlalllines=require("text!html-notopened");
    document.getElementById("notopened").innerHTML=htmlalllines;
    var htmlscript=require("text!html-script");
    document.getElementById("script").innerHTML=htmlscript;
    var htmlcreatescriptform=require("text!html-createscriptform");
    document.getElementById("createscriptform").innerHTML=htmlcreatescriptform;
    var htmlscriptlist=require("text!html-scriptlist");
    document.getElementById("scriptlist").innerHTML=htmlscriptlist;
    var htmltimerlist=require("text!html-timerlist");
    document.getElementById("timerlist").innerHTML=htmltimerlist;
    var htmlcreatetimerform=require("text!html-createtimerform");
    document.getElementById("createtimerform").innerHTML=htmlcreatetimerform;
    var htmlupdatetimerform=require("text!html-updatetimerform");
    document.getElementById("updatetimerform").innerHTML=htmlupdatetimerform;

    
    var htmlalllines=require("text!html-alllines");

    require(["main"],function(main){
    });
})