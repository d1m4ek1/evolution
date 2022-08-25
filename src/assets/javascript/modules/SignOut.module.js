const MODULE_SIGN_OUT = () => {
  fetch('/api/signout_account', {
    method: 'POST',
  })
    .then((response) => {
      if (response.ok) {
        document.cookie = 'token=; path=/; max-age=-1;';
        document.cookie = 'userId=; path=/; max-age=-1;';
        window.location.href = '/';
      }
    })
    .catch((err) => console.error(err));
};

export default MODULE_SIGN_OUT;
