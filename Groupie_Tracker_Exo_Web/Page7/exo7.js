const carouselItems = document.querySelector('.carousel-items');
const prevButton = document.querySelector('.carousel-prev');
const nextButton = document.querySelector('.carousel-next');
const itemWidth = carouselItems.querySelector('.carousel-item').getBoundingClientRect().width;

let scrollPosition = 0;

prevButton.addEventListener('click', () => {
  scrollPosition -= itemWidth;
  carouselItems.scroll({
    left: scrollPosition,
    behavior: 'smooth'
  });
});

nextButton.addEventListener('click', () => {
  scrollPosition += itemWidth;
  carouselItems.scroll({
    left: scrollPosition,
    behavior: 'smooth'
  });
});