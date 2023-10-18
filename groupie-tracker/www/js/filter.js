// show/hide filter view
function toggleFilters() {
    const filtersSection = document.getElementById("filter");
    filtersSection.classList.toggle("active");
}

function updateCreationFilter() {
    const min = document.getElementById("creationMin");
    const max = document.getElementById("creationMax");
    const display = document.getElementById("creation");

    const set = () => {
        const minValue = parseInt(min.value, 10);
        const maxValue = parseInt(max.value, 10);
        if (minValue <= maxValue) {
            display.textContent = `Select Creation Date: ${minValue} - ${maxValue}`;
        } else if (minValue > maxValue) {
            min.value = max.value;
            set();
        }
    };
    set();
    min.addEventListener("input", set);
    max.addEventListener("input", set);
}
function updateAlbumFilter() {
    const min = document.getElementById("albumMin");
    const max = document.getElementById("albumMax");
    const display = document.getElementById("album");

    const set = () => {
        const minValue = parseInt(min.value, 10);
        const maxValue = parseInt(max.value, 10);
        if (minValue <= maxValue) {
            display.textContent = `Select First Album Date: ${minValue} - ${maxValue}`;
        } else if (minValue > maxValue) {
            min.value = max.value;
            set();
        }
    };
    set();
    min.addEventListener("input", set);
    max.addEventListener("input", set);
}

document.addEventListener("DOMContentLoaded", () => {
    updateCreationFilter()
    updateAlbumFilter()
})
