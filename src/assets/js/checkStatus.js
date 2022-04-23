const checkStatus = () => {
  const cookies = document.cookie.split(';');
  cookies.forEach((item) => {
    if (item.includes('token=')) {
      fetch('/check_status').catch((err) => console.error(err));
      setInterval(() => {
        fetch('/check_status').catch((err) => console.error(err));
      }, 10000);
    }
  });
};
checkStatus();
