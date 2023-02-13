import {Stopwatch} from './stopwatch.js'
import {checkInput} from './checkInput.js'
import {setFormDateTime, setFormTimeRecord, formatManualTimeValues, submitTimeRecord} from './stopwatchForm.js'

let stopwatch = new Stopwatch();

const startBtn = document.querySelector('#start-button');
const stopBtn = document.querySelector('#stop-button');
const pauseBtn = document.querySelector('#pause-button');
const continueBtn = document.querySelector('#continue-button');
const resetBtn = document.querySelector('#reset-button');

startBtn.addEventListener('click', stopwatch.start);
stopBtn.addEventListener('click', () => {
    stopwatch.stop();
    setFormDateTime();
    setFormTimeRecord(stopwatch.dispElapsed);
});
pauseBtn.addEventListener('click', stopwatch.pause);
continueBtn.addEventListener('click', stopwatch.continueTimer);
resetBtn.addEventListener('click', stopwatch.reset);

const manualTimeDiv = document.querySelector('#custom-time-input');
const timeInputElements = manualTimeDiv.querySelectorAll('div > input');

timeInputElements.forEach((item) => {
    let inputBox = document.querySelector(`#${item.id}`)
    inputBox.addEventListener('mouseup', (event) => {event.target.select()});
    inputBox.addEventListener('beforeinput', checkInput);
})

const submitBtn = document.querySelector('#submit-btn');
submitBtn.addEventListener('click', () => {
    let stopwatchTime = document.querySelector('.clock').textContent;
    let manualTimeValues = formatManualTimeValues(timeInputElements);
    submitTimeRecord(stopwatchTime, manualTimeValues);
})

const overrideBtn = document.querySelector('#confirm-override-btn');
const useStopwatchTimeBtn = document.querySelector('#use-stopwatch-time-btn');
const addCustomTimeBtn = document.querySelector('#add-custom-time-btn');
const closeCustomTimeBtn = document.querySelector('#close-custom-time-btn');

overrideBtn.addEventListener('click', () => {
    let manualTimeValues = formatManualTimeValues(timeInputElements);
    setFormTimeRecord(manualTimeValues);
    document.querySelector('#time-record-form').submit();
})

useStopwatchTimeBtn.addEventListener('click', () => {
    document.querySelector('#time-record-form').submit();
})

addCustomTimeBtn.addEventListener('click', () => {
    document.querySelector('#custom-time-container').style.display = "inline-block";
    document.querySelector('#close-custom-time-btn').style.display = "inline-block";
})

closeCustomTimeBtn.addEventListener('click', () => {
    document.querySelector('#custom-time-container').style.display = "none";
    document.querySelector('#close-custom-time-btn').style.display = "none";
})