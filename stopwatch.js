var dispElapsed;
var msElapsed = 0;
var createClock;

function padTime(ms) {
    if (ms) {
        let paddedTime = String(ms).padStart(2, '0');
        return paddedTime;
    } else {
        return "00";
    }		
}

function displayTime(ms) {
    let hr = Math.floor(ms/3600000);
    let min = Math.floor(ms/60000) % 60;
    let sec = Math.floor(ms/1000) % 60;
    let milliSec = Math.floor(ms/10) % 100;
    let hrPad = padTime(hr)
    let minPad = padTime(min)
    let secPad = padTime(sec)
    let milliSecPad = padTime(milliSec)

    return hrPad + ":" + minPad + ":" + secPad + "." + milliSecPad
}

function startTimer(startTime) {
    msElapsed = Date.now() - startTime;
    dispElapsed = displayTime(msElapsed);
    document.querySelector('.clock').textContent = dispElapsed;
}

function setTimeRecord(timeValue) {
    let timeRecordInput = document.querySelector("input[id='time-record']");
    timeRecordInput.value = timeValue;
}

function hideAllElements() {
    document.querySelector('#start-button').style.display = "none";
    document.querySelector('#pause-button').style.display = "none";
    document.querySelector('#continue-button').style.display = "none";
    document.querySelector('#stop-button').style.display = "none";			
    document.querySelector('#reset-button').style.display = "none";			
    document.querySelector('#reset-button').style.display = "none";			
    document.querySelector('#form-div').style.display = "none";
    document.querySelector('#warning-msg').style.display = "none";
    document.querySelector('#custom-time-container').style.display = "none";
}

function start(msElapsed = 0) {
    let startTime = Date.now() - msElapsed;
    createClock = setInterval(startTimer, 10, startTime);
    hideAllElements()
    document.querySelector('#pause-button').style.display = "inline-block";
    document.querySelector('#stop-button').style.display = "inline-block";
}

function stop() {
    clearInterval(createClock);
    setFormDateTime();
    hideAllElements();
    document.querySelector('#reset-button').style.display = "inline-block";
    document.querySelector('#form-div').style.display = "inline-block";
    setTimeRecord(dispElapsed);
}

function reset() {
    msElapsed = 0
    document.querySelector('.clock').textContent = "00:00:00.00";
    hideAllElements();
    document.querySelector('#start-button').style.display = "inline-block";
}

function pause() {
    stop();
    hideAllElements();
    document.querySelector('#form-div').style.display = "inline-block";
    document.querySelector('#continue-button').style.display = "inline-block";
    document.querySelector('#stop-button').style.display = "inline-block";
}

function continueTimer() {
    start(msElapsed);
    hideAllElements();
    document.querySelector('#pause-button').style.display = "inline-block";
    document.querySelector('#stop-button').style.display = "inline-block";
}

function setFormDateTime() {
    let datetimeNow = new Date().toISOString().slice(0,19);
    let [dateStr, timeStr] = datetimeNow.split("T");
    
    let formDate = document.querySelector("input[id='form-date']");
    let formTime = document.querySelector("input[id='form-time']");
    formDate.value = dateStr;
    formTime.value = timeStr;
}

export {start, stop, reset, pause, continueTimer, startTimer, padTime, displayTime, dispElapsed, msElapsed, createClock};