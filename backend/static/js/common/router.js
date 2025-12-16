const router = {
    baseUrl: () => { return window.location.origin + "/gin_mcis/system"; },
    push: (path) => { window.location.href = router.baseUrl() + path; }
};