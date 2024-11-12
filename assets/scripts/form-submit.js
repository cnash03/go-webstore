document.addEventListener('DOMContentLoaded', function() {
    var form = document.getElementById('purchase-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent the default form submission
        console.log('Form submitted'); // Log to console for debugging
        
        // Perform any client-side validation here if needed
        
        // If all validation passes, submit the form
        form.submit();
    });
});