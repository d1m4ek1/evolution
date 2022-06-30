import GetCookie from '../getCookie';
import OpenConnectWebSocket from '../websockets';

const MODULE_CHECK_AUTHORIZE_USER = () => {
  const token: string | undefined = GetCookie('token');
  const userId: string | undefined = GetCookie('userId');

  if (token !== undefined && userId !== undefined
      && token !== "" && userId !== "") {
    fetch(`/api/check_authorization?token=${token}&userId=${userId}`, {
      method: 'GET',
    })
      .then((response: Response) => {
        response.json().then((data) => {
          if (data !== null) {
            if (data.isVerify) {
              OpenConnectWebSocket();
            }
          }
        });
      });
  }
};

export default MODULE_CHECK_AUTHORIZE_USER;
