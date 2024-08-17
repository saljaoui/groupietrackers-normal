// navigation.js

document.addEventListener('DOMContentLoaded', () => {
    // Get all navigation buttons and article sections
    const navLinks = document.querySelectorAll('.navbar-link');
    const articles = document.querySelectorAll('.contact-artists');

    // Function to handle tab clicks
    function handleNavClick(event) {
        // Remove 'active' class from all nav links and articles
        navLinks.forEach(link => link.classList.remove('active'));
        articles.forEach(article => article.classList.remove('active'));

        // Add 'active' class to the clicked nav link and the corresponding article
        const clickedLink = event.target;
        const articleId = clickedLink.id;
        const targetArticle = document.getElementById(articleId);

        clickedLink.classList.add('active');
        targetArticle.classList.add('active');
    }

    // Attach click event listeners to all nav links
    navLinks.forEach(link => {
        link.addEventListener('click', handleNavClick);
    });

    // Initialize the first tab as active
    if (navLinks.length > 0) {
        navLinks[0].click(); // Simulate a click on the first nav link to show its content by default
    }
});
