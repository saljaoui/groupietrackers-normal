document.querySelector('#search-btn').onclick = () => {
    let searchForm = document.querySelector('.search-artists'); // Corrected variable name
    searchForm.classList.toggle('active'); // Corrected classList method
};

let text = document.getElementById('text');
window.addEventListener('scroll', () => {
    let value = window.scrollY;
text.style.marginTop = value * 2.5 + 'px';
console.log(value)
if (value >= 265) {
    text.style.display = "none" 
} else {
    text.style.display = "block"
}
});

let loadMoreBtn = document.querySelector('#load-more');
let loadLessBtn = document.querySelector('#Load-less');
let load = document.getElementById('load-more');
let load2 = document.getElementById('Load-less');
let currentItem = 24;
loadMoreBtn.onclick = () => {
    let boxes = [...document.querySelectorAll('.cards .card')];
    for (let i = currentItem; i >= currentItem - 24 && i < boxes.length; i++){
        boxes[i].style.display = 'inline-block'
    }
    currentItem += 24;
    console.log(currentItem)

    if (currentItem >= 50) {
        load.style.display = 'none'
    }
    if (currentItem >= 28) {
        load2.style.display = 'flex'
    }

};
loadLessBtn.onclick = () => {
    let boxes = [...document.querySelectorAll('.cards .card')];
    currentItem -= 24;
    for (let i = currentItem; i >= currentItem - 24 && i < boxes.length; i++){
        boxes[i].style.display = 'none'
    }
    if (currentItem <= 50) {
        load.style.display = 'flex'
    }
    console.log(currentItem)
    if (currentItem <= 24) {
        load2.style.display = 'none'
    }
};
