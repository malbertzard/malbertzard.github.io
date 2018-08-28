window.onload = function() {
  document.getElementById("year").innerHTML = (new Date()).getFullYear();
};

function openInNewTab(url) {
  var win = window.open(url, '_blank');
  win.focus();
}
