const registry = require("./registry.json");

function getFonts() {
    window.queryLocalFonts().then((availableFonts) => {
        var fonts = JSON.stringify(availableFonts.map((font) => font.fullName));
        console.log(fonts);
        document.querySelector("#amount").innerHTML = fonts.length;
    });
}

console.log(registry);
