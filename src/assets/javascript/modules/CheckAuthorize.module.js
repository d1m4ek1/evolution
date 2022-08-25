import GetCookie from "../getCookie";
import OpenConnectWebSocket from "../websockets";

const MODULE_CHECK_AUTHORIZE_USER = () => {
  const token = GetCookie("token");
  const userId = GetCookie("userId");

  if (
    token !== undefined &&
    userId !== undefined &&
    token !== "" &&
    userId !== ""
  ) {
    fetch(`/api/check_authorization?token=${token}&userId=${userId}`, {
      method: "GET",
    }).then((response) => {
      response.json().then((data) => {
        if (data !== null) {
          if (data.isVerify) {
            sessionStorage.setItem("verify", data.isVerify);
            OpenConnectWebSocket();
          }
        }
      });
    });
  } else {
    sessionStorage.setItem("verify", false);
  }
};

export default MODULE_CHECK_AUTHORIZE_USER;
