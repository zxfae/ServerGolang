import { loadHome } from './home.js';
import { loadRegister } from './register.js';
import { loadLogin, logoutUser } from './login.js';
import { loadPostsByCategory } from './posts.js';

export function handleRoute() {
    const hash = window.location.hash || '#login';
    const [route, category] = hash.split('/');

    switch (route) {
        case '#home':
            if (window.isUserAuthenticated) {
                loadHome();
            } else {
                window.location.hash = '#login';
            }
            break;
        case '#register':
            loadRegister();
            break;
        case '#login':
            loadLogin();
            break;
        case '#logout':
            logoutUser();
            break;
        case '#category':
            if (category) {
                loadPostsByCategory(category);
            } else {
                window.location.hash = '#home';
            }
            break;
        default:
            window.location.hash = '#home';
            break;
    }
}

export function navigate(event, route) {
    event.preventDefault();
    window.location.hash = route;
}
