function checkChar(char, regExp) {
    return regExp.test(char);  
}

function checkNumRange(inputVal, min, max) {
    let inputNum = Number(inputVal);
    return (inputNum >= Number(min) && inputNum <= Number(max));
}

function checkInput(event) {
    let inputVal = event.data;
    let inputType = event.inputType;
    const isNum = checkChar(inputVal, /\d/);

    if (isNum) {
        event.preventDefault();
        
        let currentInputBoxVal = String(event.target.value);

        if (event.target.selectionStart === 0) {
            event.target.value = "0" + String(inputVal);
        } else if (event.target.selectionStart === 1) {
            event.target.value = currentInputBoxVal.substr(0,1) + String(inputVal);
        } else if (event.target.selectionStart === 2) {    
            if (currentInputBoxVal.startsWith("0")) {
                // move current number to left and add new value
                event.target.value = currentInputBoxVal.substr(1) + String(inputVal);
            } else {
                return;
            }      
        }

        let inRange = checkNumRange(event.target.value, event.target.min, event.target.max);
        
        if (!inRange) {
            let errorMsg = document.querySelector('#errorMsg');
            errorMsg.innerText = `Please enter a number from ${event.target.min} to ${event.target.max}`;
            errorMsg.style.display = 'inline-block';
        } else {
            errorMsg.style.display = 'none';
        }
        
    } else if (inputType.includes('delete')) {
        // allow deletion
        return;
    } else {
        event.preventDefault();
    }

}

export {checkChar, checkNumRange, checkInput}