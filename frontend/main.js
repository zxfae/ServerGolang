import { checkUserStatus } from './auth.js';
import { handleRoute } from './router.js';

document.addEventListener("DOMContentLoaded", () => {
    checkUserStatus();
    window.addEventListener('hashchange', handleRoute);
});