import GetCookie from '../getCookie';
import OpenConnectWebSocket from '../websockets';

const MODULE_CHECK_AUTHORIZE_USER = async () => {
  const response = await fetch(`/api/check_authorization`, {
    method: 'GET',
  });
  const jsonResponse = await response.json();

  sessionStorage.setItem('verify', jsonResponse.isVerify);

  if (jsonResponse.isVerify) {
    OpenConnectWebSocket();
  }

  return jsonResponse.isVerify;
};

export default MODULE_CHECK_AUTHORIZE_USER;
