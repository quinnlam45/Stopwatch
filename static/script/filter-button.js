export class FilterButton {
    constructor(filterVar, buttonId) {
        this.filterVar = filterVar;
        this.buttonId = buttonId;
        this.toggleFlag = false;
        this.buttonObj = document.getElementById(buttonId);
        this.closeButton = ' <i class="bi bi-x" style="display: inline-block;"></i>';
    }
    
    changeToggleFlag = () => {
        if (this.toggleFlag) {
            this.toggleFlag = false;
        } else {
            this.toggleFlag = true;
        }
    }

    toggleStyle = () => {
        if (this.toggleFlag) {
            this.buttonObj.className = "btn btn-secondary btn-sm";
        } else {
            this.buttonObj.className = "btn btn-outline-secondary btn-sm";
        }
    }

    toggleClose = () => {
        if (this.toggleFlag) {
            let html = this.buttonObj.innerHTML;
            this.buttonObj.innerHTML = html + this.closeButton;
        } else {
            // remove icon by using inner text to reset
            this.buttonObj.innerHTML = this.buttonObj.innerText;
        }
    }

    changeToggleStyle = () => {
        this.changeToggleFlag();
        this.toggleStyle();
        this.toggleClose();
    }

}