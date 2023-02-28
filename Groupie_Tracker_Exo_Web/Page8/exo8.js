const descriptions = document.querySelectorAll('.card-description');

descriptions.forEach(description => {
  const fullText = description.innerHTML;
  const maxLength = 100;

  if (fullText.length > maxLength) {
    const shortText = fullText.substring(0, maxLength);
    const trimmedText = shortText.trim();
    const newText = trimmedText + '...';
    description.innerHTML = newText;
  }
});