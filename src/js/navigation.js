let showInfodate = document.getElementById("showInfodate");
let showInfoRelation = document.getElementById("showInfoRelation");
let showInfolocation = document.getElementById("showInfolocation");

let date = document.getElementById("date");
let relation = document.getElementById("relation");
let locations = document.getElementById("location");

// Default state: date is visible, others are hidden
toggleVisibility(date, [relation, locations]);

function toggleVisibility(showElement, hideElements) {
    showElement.classList.add('visible');
    showElement.classList.remove('hidden');
    hideElements.forEach(element => {
        element.classList.add('hidden');
        element.classList.remove('visible');
    });
}

showInfodate.onclick = function() {
    toggleVisibility(date, [relation, locations]);
}

showInfoRelation.onclick = function() {
    toggleVisibility(relation, [date, locations]);
}

showInfolocation.onclick = function() {
    toggleVisibility(locations, [date, relation]);
}

const navbarLinks = document.querySelectorAll('.navbar-link');

navbarLinks.forEach(link => {
    link.addEventListener('click', function() {
        // Remove the color from all links
        navbarLinks.forEach(nav => nav.style.color = '');
        // Change the color of the clicked link
        this.style.color = '#d39c51';
    });
});