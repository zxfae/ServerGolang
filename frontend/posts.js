export function loadPosts() {
    console.log('Fetching posts...');
    fetch('/api/posts')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            console.log('Posts loaded:', posts);
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            const limitedPosts = posts.slice(0, 12);
            limitedPosts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';
                let truncatedTitle = post.title.length > 60 ? `${post.title.substring(0, 60)}...` : post.title;
                let truncatedDescription = post.description.length > 80 ? `${post.description.substring(0, 80)}...` : post.description;
                postElement.innerHTML = `
                <div class="scrolling">
                    <h2><a class="username">${post.username}</a><span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~</span><span class="white">$/post/</span>${truncatedTitle}</span></h2>
                    <div class="Weigth">
                    <p><a class="username">category</a>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${post.categoryname}</p>
                    </div>
                    <p>${truncatedDescription}</p>
                    <p><span class="time">${post.created}</span></p>
                </div>
                `;
                postContainer.appendChild(postElement);
            });
        })
        .catch(error => console.error('Error loading posts:', error));
}

export function loadPostsByCategory(categoryName) {
    console.log(`Fetching posts for category: ${categoryName}`);
    fetch(`/api/posts/category/${categoryName}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            console.log('Posts loaded:', posts);
            const content = document.getElementById('content');
            content.innerHTML = `
                <div class="category-posts">
                    <h2 class="white-text">Category: ${categoryName}</h2>
                    <div class="posts-container"></div>
                </div>
            `;
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            const limitedPosts = posts.slice(0, 12);
            limitedPosts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';
                let truncatedTitle = post.title.length > 60 ? `${post.title.substring(0, 60)}...` : post.title;
                let truncatedDescription = post.description.length > 80 ? `${post.description.substring(0, 80)}...` : post.description;
                postElement.innerHTML = `
                    <div class="scrolling">
                        <h2><a class="username">${post.username}</a><span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~</span><span class="white">$/post/</span>${truncatedTitle}</span></h2>
                        <div class="Weigth">
                            <p><a class="username">category</a>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${post.categoryname}</p>
                        </div>
                        <p>${truncatedDescription}</p>
                        <p><span class="time">${post.created}</span></p>
                    </div>
                `;
                postContainer.appendChild(postElement);
            });
        })
        .catch(error => console.error('Error loading posts:', error));
}
