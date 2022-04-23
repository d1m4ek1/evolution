const delayedLoading = () => {
  if (document.documentElement.clientWidth <= 960) {
    setTimeout(() => {
      document.querySelectorAll('.delayed-loading').forEach((item = String) => {
        const elem = item;
        const src = elem.getAttribute('delayed-loading');

        elem.src = src;
      });
    }, 500);
  } else {
    document.querySelectorAll('.delayed-loading').forEach((item = String) => {
      const elem = item;
      const src = elem.getAttribute('delayed-loading');

      elem.src = src;
    });
  }
};
delayedLoading();
