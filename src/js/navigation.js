
let showInfodate = document.getElementById("showInfodate");
let showInfoRelation = document.getElementById("showInfoRelation");
let showInfolocation = document.getElementById("showInfolocation");

let about = document.getElementById("date");
let relation = document.getElementById("relation");
let locations = document.getElementById("location");

// Default state: About is visible, others are hidden
about.classList.add('visible');
relation.classList.add('hidden');
locations.classList.add('hidden')

showInfodate.onclick = function() {
    about.classList.add('visible');
    about.classList.remove('hidden');

    relation.classList.add('hidden');
    relation.classList.remove('visible');
    locations.classList.add('hidden');
    locations.classList.remove('visible');
}

showInfoRelation.onclick = function() {
    relation.classList.add('visible');
    relation.classList.remove('hidden');

    about.classList.remove('info')
    about.classList.add('hidden');
    about.classList.remove('visible');

    locations.classList.add('hidden');
    locations.classList.remove('visible');
}



showInfolocation.onclick = function() {
    locations.classList.add('visible');
    locations.classList.remove('hidden');

    about.classList.remove('info')
    about.classList.add('hidden');
    about.classList.remove('visible');

    relation.classList.add('hidden');
    relation.classList.remove('visible');
}