document.addEventListener('DOMContentLoaded', cyclePlaceholders)

let intervalId; // Variable to store the interval ID
let currentIndex = 0;
const placeholderTexts = ["Search...", "50 Cent...", "Post Malone..."];
async function cyclePlaceholders() {
    const searchInput = document.getElementById('search');
    if (searchInput === null) {
        console.error("searchInput is null");
        return;
    }
    clearInterval(intervalId); // Clear any existing interval

    for (let i = searchInput.placeholder.length - 1; i >= 0; i--) {
        searchInput.placeholder = searchInput.placeholder.slice(0, -1);
        await new Promise(r => setTimeout(r, 110));
    }
    const newText = placeholderTexts[currentIndex];
    currentIndex = (currentIndex + 1) % placeholderTexts.length;

    for (let i = 0; i < newText.length; i++) {
        searchInput.placeholder += newText[i];
        await new Promise(r => setTimeout(r, 110));
    }

    intervalId = setInterval(cyclePlaceholders, 2000); // Schedule the next run after 4 seconds
}

function showOptions(e) {
    const filterId = e.getAttribute("data-items-id");
    const filterItems = document.getElementById(filterId)
    if (filterItems.style.display === "block") {
        filterItems.style.display = "none"; // If it's currently "block", set it to "none"
    } else {
        filterItems.style.display = "block"; // If it's not "block" (e.g., "none" or ""), set it to "block"
    }
}

function updateTheme(e) {
    const themeValue = e.getAttribute("data-theme-value");
    document.cookie = `theme=${themeValue}; path=/; SameSite=Strict`;
    location.reload();
}