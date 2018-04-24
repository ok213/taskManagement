
function rework(){
    var content = document.getElementById('rowbtn');
    content.style.display = "none";
    content = document.getElementById('rework');
    content.style.display = "block";
}

function addUser(obj){
    obj.style.display = "none";
    var content = document.getElementById('addUser');
    content.style.display = "block";
}

function addDept(obj){
    obj.style.display = "none";
    var content = document.getElementById('addDept');
    content.style.display = "block";
}

function addRole(obj){
    obj.style.display = "none";
    var content = document.getElementById('addRole');
    content.style.display = "block";
}

function addStatus(obj){
    obj.style.display = "none";
    var content = document.getElementById('addStatus');
    content.style.display = "block";
}

function checkUser(link) {
    var content = document.getElementById('contentMain');
    // var loading = document.getElementById('loading');
    // content.innerHTML = loading.innerHTML;
    var req = getXMLHttpRequest();
    req.onreadystatechange = function () {
        if (req.readyState != 4) return;
        if (req.status != 200) {
            alert("Error.... " + req.statusText);
            return;
        }
        content.innerHTML = req.responseText;
    }
    req.open("GET", link, true);
    req.send(null);
}

// получить объект XMLHttpRequest
function getXMLHttpRequest() {
    if (window.XMLHttpRequest) {
        try {
            return new XMLHttpRequest();
        } catch (error) { }
    } else if (window.ActiveXObject) {
        try {
            return new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e) { }
        try {
            return new ActiveXObject("Microsoft.XMLHTTP");
        } catch (e) { }
    }
    return null;
}
