const API_BASE_URL = 'http://localhost:8000';
const CSRF_COOKIE_NAME = 'csrftoken';
const CSRF_SAFE_METHODS = new Set(['GET', 'HEAD', 'OPTIONS']);

function getCookie(url, name) {
    return new Promise((resolve) => {
        chrome.cookies.get({ url, name }, (cookie) => {
            resolve(cookie ? cookie.value : null);
        });
    });
}

async function apiFetch(path, options = {}) {
    const method = (options.method || 'GET').toUpperCase();
    const headers = new Headers(options.headers || {});

    if (!CSRF_SAFE_METHODS.has(method)) {
        const csrfToken = await getCookie(API_BASE_URL, CSRF_COOKIE_NAME);
        if (csrfToken) {
            headers.set('X-CSRFToken', csrfToken);
        }
    }

    const response = await fetch(`${API_BASE_URL}${path}`, {
        ...options,
        method,
        credentials: 'include',
        headers,
    });

    const contentType = response.headers.get('content-type') || '';
    const responseBody = contentType.includes('application/json')
        ? await response.json()
        : await response.text();

    if (!response.ok) {
        throw new Error(
            typeof responseBody === 'string'
                ? responseBody || response.statusText
                : responseBody.error || response.statusText
        );
    }

    return responseBody;
}

chrome.runtime.onMessage.addListener((request, _sender, sendResponse) => {
    if (request.action === 'API_FETCH') {
        apiFetch(request.path, request.options)
            .then((data) => sendResponse({ ok: true, data }))
            .catch((error) => sendResponse({ ok: false, error: error.message }));
        return true;
    }
});
