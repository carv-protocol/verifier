document.getElementById('submit').addEventListener('click', function (event) {
    event.preventDefault();

    const rpcUrlInput = document.getElementById('rpcUrl');
    let rpcUrl = rpcUrlInput.value
    if (!rpcUrl) {
        rpcUrl = rpcUrlInput.placeholder
    }

    const privateKey = document.getElementById('privateKey').value

    console.log( "rpcUrl, privateKey: ", rpcUrl, privateKey);
    let formData = JSON.stringify({
        rpcUrl: rpcUrl,
        privateKey: privateKey
    });

    var requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: formData,
        redirect: 'follow'
    };
    fetch("http://localhost:8080/verifier/run", requestOptions)
    .then(response => response.json())
    .then(data => {
        alert(data.msg)
    })
    .catch(error => console.log('error', error));
})

async function checkActive() {
    try {
        const response = await fetch('verifier/active',{
            method: 'GET',
        });
        if (!response.ok) {
            throw new Error(`Failed to check active status: ${response.statusText}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        throw error;  // Re-throw for handling in ipcMain
    }
}

var checkActiveTimerId = setInterval( async function() {
    try {
        const data = await checkActive();
        console.log('Active:', JSON.stringify(data));

        var activeDiv = document.getElementById('isActive');
        if (data.success) {
            activeDiv.innerText = 'Active';
            return
        }
        activeDiv.innerText = 'InActive';
    } catch (error) {
        console.error(error);
    }
}, 1000);
