document.addEventListener('DOMContentLoaded', function() {
    // Theme switcher
    const themeToggle = document.getElementById('theme-toggle');
    if (themeToggle) {
        themeToggle.addEventListener('click', function() {
            document.body.classList.toggle('dark-mode');
            const isDark = document.body.classList.contains('dark-mode');
            localStorage.setItem('darkMode', isDark);
        });

        // Load saved theme preference
        const savedTheme = localStorage.getItem('darkMode');
        if (savedTheme === 'true') {
            document.body.classList.add('dark-mode');
        }
    }

    // Real-time server status checker
    const statusCheck = document.getElementById('server-status');
    if (statusCheck) {
        function checkServerStatus() {
            fetch('/ping')
                .then(response => response.json())
                .then(data => {
                    statusCheck.textContent = '🟢 Server Online';
                    statusCheck.classList.remove('offline');
                    statusCheck.classList.add('online');
                })
                .catch(() => {
                    statusCheck.textContent = '🔴 Server Offline';
                    statusCheck.classList.remove('online');
                    statusCheck.classList.add('offline');
                });
        }

        // Check status every 30 seconds
        checkServerStatus();
        setInterval(checkServerStatus, 30000);
    }

    // Add animation to page transitions
    document.querySelectorAll('a').forEach(link => {
        link.addEventListener('click', function(e) {
            if (this.hostname === window.location.hostname) {
                e.preventDefault();
                document.body.classList.add('fade-out');
                setTimeout(() => {
                    window.location = this.href;
                }, 300);
            }
        });
    });
});