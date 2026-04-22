document.addEventListener("DOMContentLoaded", function() {
    const results = document.querySelector('.results-section');
    if (results) {
        results.scrollIntoView({ behavior: 'smooth' });
    }
});

// Kopyalama Fonksiyonu
function copyToClipboard(id) {
    const text = document.getElementById(id).innerText;
    navigator.clipboard.writeText(text).then(() => {
        const btn = event.currentTarget;
        const icon = btn.innerHTML;
        btn.innerHTML = '<i class="fa-solid fa-check" style="color: #48bb78;"></i>';
        setTimeout(() => btn.innerHTML = icon, 1500);
    });
}