import { loadHome } from './home.js';
import { loadRegister } from './register.js';
import { loadLogin, logoutUser } from './login.js';
import { loadPostsByCategory, loadPostsByPosts } from './posts.js';

export function handleRoute() {
    const hash = window.location.hash || '#login';
    const [route, param] = hash.split('/');

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
            if (param) {
                loadPostsByCategory(param);
            } else {
                window.location.hash = '#home';
            }
            break;
        case '#posts':
            if (param) {
                loadPostsByPosts(param);
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