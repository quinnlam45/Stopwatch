export class Stopwatch {
    constructor() {
        this.dispElapsed;
        this.msElapsed = 0;
        this.createClock;
    }

    padTime = (ms) => {
        if (ms) {
            let paddedTime = String(ms).padStart(2, '0');
            return paddedTime;
        } else {
            return "00";
        }		
    }

    displayTime = (ms) => {
        let hr = Math.floor(ms/3600000);
        let min = Math.floor(ms/60000) % 60;
        let sec = Math.floor(ms/1000) % 60;
        let milliSec = Math.floor(ms/10) % 100;
        let hrPad = this.padTime(hr)
        let minPad = this.padTime(min)
        let secPad = this.padTime(sec)
        let milliSecPad = this.padTime(milliSec)

        return hrPad + ":" + minPad + ":" + secPad + "." + milliSecPad
    }

    startTimer = (startTime) => {
        this.msElapsed = Date.now() - startTime;
        this.dispElapsed = this.displayTime(this.msElapsed);
        document.querySelector('.clock').textContent = this.dispElapsed;
    }

    hideAllElements = () => {
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

    start = () => {
        let startTime = Date.now() - this.msElapsed;
        this.createClock = setInterval(this.startTimer, 10, startTime);
        this.hideAllElements()
        document.querySelector('#pause-button').style.display = "inline-block";
        document.querySelector('#stop-button').style.display = "inline-block";
    }

    stop = () => {
        clearInterval(this.createClock);
        this.hideAllElements();
        document.querySelector('#reset-button').style.display = "inline-block";
        document.querySelector('#form-div').style.display = "inline-block";
    }

    reset = () => {
        this.msElapsed = 0
        document.querySelector('.clock').textContent = "00:00:00.00";
        this.hideAllElements();
        document.querySelector('#start-button').style.display = "inline-block";
    }

    pause = () => {
        this.stop();
        this.hideAllElements();
        document.querySelector('#form-div').style.display = "inline-block";
        document.querySelector('#continue-button').style.display = "inline-block";
        document.querySelector('#stop-button').style.display = "inline-block";
    }

    continueTimer = () => {
        this.start(msElapsed);
        this.hideAllElements();
        document.querySelector('#pause-button').style.display = "inline-block";
        document.querySelector('#stop-button').style.display = "inline-block";
    }

}