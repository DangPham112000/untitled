const options = {
    enableHighAccuracy: false,
    timeout: 15000,
    maximumAge: 120000,
};

var tsArr = [],
    geoTS,
    oldPosTS;
console.log("Begin");
console.log("maximumAge: ", options.maximumAge);

function success(newPos) {
    const crd = newPos.coords;
    if (oldPosTS !== newPos.timestamp) {
        oldPosTS = newPos.timestamp;
        console.log("-------------New location TS--------------------");
    } else {
    }
    console.log(
        newPos.timestamp,
        `${crd.accuracy} ${crd.latitude} ${crd.longitude}`
    );
    console.log("Age: ", geoTS - newPos.timestamp);
}

function error(err) {
    console.warn(`(${err.code}): ${err.message}`);
}

function btnOnClick() {
    geoTS = new Date().getTime();
    console.log(geoTS, "geolocate...");
    navigator.geolocation.getCurrentPosition(success, error, options);
}
