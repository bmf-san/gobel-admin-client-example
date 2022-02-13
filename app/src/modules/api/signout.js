import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const postSignout = async () => {
  try {
    const res = await apiClient
      .post(
        `/private/signout`,
        {},
        {
          headers: {
            Authorization: "Bearer " + storage.getAccessToken()
          }
        }
      )
      .then(res => {
        storage.removeAccessToken();
        storage.removeRefreshToken();
        storage.setIsSignin(false);
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const apiSignout = {
  postSignout
};

export default apiSignout;
