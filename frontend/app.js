const postKeyField = document.getElementById('post-key')
const postValueField = document.getElementById('post-value')
const postOutput = document.getElementById('output-post')
const postButton = document.getElementById('post-key-value-button')
const postForm = document.getElementById('post-form')

const getKeyField = document.getElementById('get-key')
const getOutput = document.getElementById('output')
const getForm = document.getElementById('get-form')

async function postKeyValue() {
    let key =  postKeyField.value;
    let value = postValueField.value;
    let url = "https://go-api.johnlangs.dev/v1/" + key;

    console.log(key)
    console.log(value)
    console.log(url)

    await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: value
    });

    postOutput.textContent = "Posted!";
}

async function getKey() {
    let key =  postKeyField.value;
    url = "https://go-api.johnlangs.dev/v1/" + key;

    console.log(key)
    console.log(url)

    let res = await fetch(url, {
        method: "GET",
    });
    res = await res.text();

    getOutput.textContent = res;
}
