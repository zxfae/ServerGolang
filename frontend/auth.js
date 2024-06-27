import { handleRoute } from './router.js';

export function checkUserStatus() {
    fetch('/api/checkAuth')
        .then(response => {
            if (response.ok) {
                window.isUserAuthenticated = true;
                if (!window.location.hash || window.location.hash === '#login' || window.location.hash === '#register') {
                    window.location.hash = '#home';
                }
            } else {
                window.isUserAuthenticated = false;
                if (!window.location.hash || window.location.hash === '#home') {
                    window.location.hash = '#login';
                }
            }
            handleRoute();
        })
        .catch(error => {
            console.error('Error checking user status:', error);
            window.isUserAuthenticated = false;
            handleRoute();
        });
}
