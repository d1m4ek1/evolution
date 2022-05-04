let stickyHeaderOpened = false;

const stickyHeader = () => {
  if (document.documentElement.clientWidth <= 960) {
    let x1 = null;
    let y1 = null;

    const handleTouchStart = (e) => {
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

    const handleTouchMove = (e) => {
      if (!x1 || !y1) {
        return false;
      }

      const x2 = e.touches[0].clientX;
      const y2 = e.touches[0].clientY;

      const xDiff = x2 - x1;
      const yDiff = y2 - y1;

      if (Math.abs(xDiff) > Math.abs(yDiff)) {
        // right
        if (xDiff > 0) {
          document.querySelector('.header_sticky').style.transform = 'translateX(0px)';
          document
            .querySelector('.main')
            .classList.remove('main_squeeze_before_remove');
          document
            .querySelector('.main')
            .classList.add('main_squeeze_before_add');
          document
            .querySelector('.main__body_content')
            .classList.remove('main_unclench');
          document
            .querySelector('.main__body_content')
            .classList.add('main_squeeze');
          document.body.style.overflow = 'hidden';
          stickyHeaderOpened = true;
        } else if (stickyHeaderOpened) {
          document.querySelector('.header_sticky').style.transform = 'translateX(-200px)';
          document
            .querySelector('.main')
            .classList.remove('main_squeeze_before_add');
          document
            .querySelector('.main')
            .classList.add('main_squeeze_before_remove');
          document
            .querySelector('.main__body_content')
            .classList.remove('main_squeeze');
          document
            .querySelector('.main__body_content')
            .classList.add('main_unclench');
          document.body.style.overflow = null;
          setTimeout(() => {
            document
              .querySelector('.main')
              .classList.remove('main_squeeze_before_remove');
          }, 490);
          stickyHeaderOpened = false;
        }
      }
      return true;
    };

    const handleTouchEnd = () => {
      x1 = null;
      y1 = null;
    };

    document.addEventListener('touchstart', handleTouchStart, false);
    document.addEventListener('touchmove', handleTouchMove, false);
    document.addEventListener('touchend', handleTouchEnd, false);
  }
};
stickyHeader();
