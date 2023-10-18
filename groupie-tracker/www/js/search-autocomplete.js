document.addEventListener("DOMContentLoaded", initializeSearchListener);

let timeoutId;

function initializeSearchListener() {
    const searchInput = document.getElementById("search");
    const resultDiv = document.getElementById("results");

    // Attach the event listener for input
    searchInput.addEventListener(
        "input",
        handleInput.bind(null, searchInput, resultDiv)
    );
}

function performSearch(query, resultDiv) {

    const apiUrl = `/search?query=${query}`;

    // Send a GET request to the API
    fetch(apiUrl)
        .then((response) => {
            if (!response.ok) {
                resultDiv.innerHTML = "";
                throw new Error(`Search Error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then((data) => {
            // Clear the resultDiv
            resultDiv.innerHTML = "";

            // Loop over the data (whether it's an array or not)
            data.forEach((item) => {
                // Create an element to display the item
                const itemElement = document.createElement("a");
                itemElement.href = `/band/${item['band_id']}`
                itemElement.title = item['type']
                if (item['type'] === 'band') {
                    itemElement.innerHTML = `<i class="ri-team-line"></i> ${item['text']}`
                } else if  (item['type'] === 'member')  {
                    itemElement.innerHTML = `<i class="ri-user-line"></i> ${item['text']}`
                } else if  (item['type'] === 'location')  {
                    itemElement.innerHTML = `<i class="ri-map-pin-line"></i> ${item['text']}`
                } else if  (item['type'] === 'first_album')  {
                    itemElement.innerHTML = `<i class="ri-disc-line"></i> ${item['text']}`
                } else if  (item['type'] === 'creation_date')  {
                    itemElement.innerHTML = `<i class="ri-calendar-line"></i> ${item['text']}`
                } 
                // Append the item element to the resultDiv
                resultDiv.appendChild(itemElement);
            });
        })
        .catch((error) => {
            console.error( error);
        });
}

function handleInput(inputField, resultDiv, event) {
    const query = event.target.value;

    // Clear any previous timeout
    clearTimeout(timeoutId);

    // Set a new timeout to wait for 300ms after the user stops typing
    timeoutId = setTimeout(() => {
        // Call the performSearch function with the user's query and resultDiv
        performSearch(query, resultDiv);
    }, 700);
}
