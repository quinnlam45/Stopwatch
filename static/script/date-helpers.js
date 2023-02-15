function addDays(date, days) {			
    let currentDate = new Date(Number(date));
    let noOfDays = Number(days);
    return new Date(currentDate.setDate(currentDate.getDate() + noOfDays));
}

function getOffsetCurrentDate() {
    return addDays(new Date(Date.now()), 1)
}

function convertToISODate(dateObj) {
    const dateStr = dateObj.toISOString().substr(0, 10);
    return dateStr;
}

function dateStrToISODateFormat(dateStr) {
    let [dd, mm, yyyy] = dateStr.split("/");
    return `${yyyy}-${mm}-${dd}`
}

export {addDays, getOffsetCurrentDate, convertToISODate, dateStrToISODateFormat}