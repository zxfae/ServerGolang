export function loadCategories() {
    console.log('Fetching categories...');
    fetch('/api/categories')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(categories => {
            console.log('Categories loaded:', categories);
            const categoriesContainer = document.querySelector('.categories-container');
            categoriesContainer.innerHTML = '';
            const limitedCategories = categories.slice(0, 12);
            limitedCategories.forEach(category => {
                const categoryElement = document.createElement('div');
                categoryElement.className = 'category';
                categoryElement.innerHTML = `
                    <h2><a href="#category/${category.name}" class="category-link">${category.name}</a></h2>
                    <p>${category.description}</p>
                `;
                categoriesContainer.appendChild(categoryElement);
            });
        })
        .catch(error => console.error('Error loading categories:', error));
}
