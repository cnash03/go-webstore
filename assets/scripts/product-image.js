document.addEventListener("DOMContentLoaded", function() {
    const productSelect = document.getElementById("product");
    const selectedImage = document.getElementById("selected-product-image");

    // Set a default generic image when the page loads
    const defaultImage = "/assets/images/genericCar.jpg";
    selectedImage.src = defaultImage;
    selectedImage.style.display = "block"; // Show the default image initially

    productSelect.addEventListener("change", function() {
        const selectedOption = productSelect.options[productSelect.selectedIndex];
        const imagePath = selectedOption.getAttribute("data-image");

        if (imagePath) {
            selectedImage.src = "/assets/images/" + imagePath; // Update to selected product image
            selectedImage.style.display = "block"; // Ensure the image is visible
        } else {
            selectedImage.src = defaultImage; // Reset to the generic image if no product selected
        }
    });
});
