let stickyHeaderOpened = false;

interface iHeaderElements {
  headerSticky: HTMLElement
  main: HTMLElement
  mainBodyContent: HTMLElement
}

const StickyHeader = (): void => {
  if (document.documentElement.clientWidth <= 960) {
    let x1: number = null;
    let y1: number = null;

    const handleTouchStart = (e): void => {
      if (
        e.target.offsetParent !== null
        && e.target.offsetParent.classList[0] === 'swiper-slide'
      ) {
        return;
      }

      const firstTouch = e.touches[0];

      x1 = firstTouch.clientX;
      y1 = firstTouch.clientY;
    };

    const handleTouchMove = (e): boolean => {
      if (!x1 || !y1) {
        return false;
      }

      const x2: number = e.touches[0].clientX;
      const y2: number = e.touches[0].clientY;

      const xDiff: number = x2 - x1;
      const yDiff: number = y2 - y1;

      if (Math.abs(xDiff) > Math.abs(yDiff)) {
        // right
        if (xDiff > 0) {

          const headerElements: iHeaderElements = {
            headerSticky: document.querySelector('.header_sticky'),
            main: document.querySelector('.main'),
            mainBodyContent: document.querySelector('.main__body_content')
          }
          headerElements.headerSticky.style.transform = 'translateX(0px)';

          headerElements.main.classList.remove('main_squeeze_before_remove');
          headerElements.main.classList.add('main_squeeze_before_add');

          headerElements.mainBodyContent.classList.remove('main_unclench');
          headerElements.mainBodyContent.classList.add('main_squeeze');

          document.body.style.overflow = 'hidden';
          stickyHeaderOpened = true;

        } else if (stickyHeaderOpened) {
          
          const headerElements: iHeaderElements = {
            headerSticky: document.querySelector('.header_sticky'),
            main: document.querySelector('.main'),
            mainBodyContent: document.querySelector('.main__body_content')
          }
          headerElements.headerSticky.style.transform = 'translateX(-200px)';

          headerElements.main.classList.remove('main_squeeze_before_add');
          headerElements.main.classList.add('main_squeeze_before_remove');

          headerElements.mainBodyContent.classList.remove('main_squeeze');
          headerElements.mainBodyContent.classList.add('main_unclench');
          document.body.style.overflow = null;

          setTimeout(() => {
            headerElements.main.classList.remove('main_squeeze_before_remove');
          }, 490);
          stickyHeaderOpened = false;
        }
      }
      return true;
    };

    const handleTouchEnd = (): void => {
      x1 = null;
      y1 = null;
    };

    document.addEventListener('touchstart', handleTouchStart, false);
    document.addEventListener('touchmove', handleTouchMove, false);
    document.addEventListener('touchend', handleTouchEnd, false);
  }
};
export default StickyHeader;
