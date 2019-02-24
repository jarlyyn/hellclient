requirejs.config({
    "paths": {
        "vue":"https://cdn.jsdelivr.net/npm/vue@2.5.21/dist/vue",
        "ELEMENT":"https://unpkg.com/element-ui/lib/index",
    },
});
define(function (require) {
    require("/public/defaultui/js/websocket.js");
    require("/public/defaultui/js/vm.js");
    require("/public/defaultui/js/buttons.js");
    require("/public/defaultui/js/handlers.js");

})