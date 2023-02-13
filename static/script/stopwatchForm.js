function setFormDateTime() {
    let datetimeNow = new Date().toISOString().slice(0,19);
    let [dateStr, timeStr] = datetimeNow.split("T");
    
    let formDate = document.querySelector("input[id='form-date']");
    let formTime = document.querySelector("input[id='form-time']");
    formDate.value = dateStr;
    formTime.value = timeStr;
}

function setFormTimeRecord(timeValue) {
    let timeRecordInput = document.querySelector("input[id='time-record']");
    timeRecordInput.value = timeValue;
}

function getManualTimeValues(manualTimeInputElements) {
    let timeArray = [];
    
    manualTimeInputElements.forEach((item) => {
        timeArray.push(item.value);
    })
    
    return timeArray;
}

function formatManualTimeValues(manualTimeInputElements) {
    let [hh, mm, ss, ms] = getManualTimeValues(manualTimeInputElements);
    return `${hh}:${mm}:${ss}.${ms}`;
}

function submitTimeRecord(stopwatchTime, manualTimeValues) {
    if (stopwatchTime != "00:00:00.00" && manualTimeValues != "00:00:00.00") {
        document.querySelector('#warning-msg').style.display = "inline-block";
    } else if (stopwatchTime === "00:00:00.00" && manualTimeValues != "00:00:00.00") {
        setFormTimeRecord(manualTimeValues);
        document.querySelector('#time-record-form').submit();
    } else {
        document.querySelector('#time-record-form').submit();
    }
}

export {setFormDateTime, setFormTimeRecord, formatManualTimeValues, submitTimeRecord}