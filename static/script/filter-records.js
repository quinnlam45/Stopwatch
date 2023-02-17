import {addDays, getOffsetCurrentDate, dateStrToISODateFormat} from './date-helpers.js'

function filterDate(days) {
    let table = document.getElementById("records-table");
    let tr = table.getElementsByTagName("tr");

    const endRange = getOffsetCurrentDate();
    const startRange = addDays(new Date(Date.now()), days);

    for (let i = 0; i < tr.length; i++) {
        let td = tr[i].getElementsByTagName("td")[0];
        if (td) {
            let tableTxt = td.textContent || td.innerText;
            //convert text to date
            let recordDate = new Date(dateStrToISODateFormat(tableTxt));

            if (recordDate >= startRange && recordDate <= endRange) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }    
}

function renderRecordsTable(data) {
    let tableHtml = "";
    
    for (let row of data) {
        tableHtml += 
        `<tr>
            <td>${row.date}</td>
            <td>${row.time}</td>
            <td>${row.distance} ${row.distanceUnit}</td>
            <td>${row.by}</td>
        </tr>`;
    }

    document.getElementById("records-table").innerHTML = tableHtml;
}

function filterTableByText(targetText) {
    // Declare variables
    let table = document.getElementById("records-table");
    let tr = table.getElementsByTagName("tr");
  
    // Loop through all table rows, and hide those who don't match the search query
    for (let i = 0; i < tr.length; i++) {
        let td = tr[i].getElementsByTagName("td")[3];
        if (td) {
            let tableTxt = td.textContent || td.innerText;

            if (tableTxt === targetText) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

export {renderRecordsTable, filterTableByText, filterDate}