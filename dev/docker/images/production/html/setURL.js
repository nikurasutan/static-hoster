let url = document.getElementById('siteURL')
let postURL = `${location.protocol}//${location.host}/api/update`
url.innerText = postURL
url.addEventListener('click', function() {
    navigator.clipboard.writeText(postURL)
})